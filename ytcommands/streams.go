package ytcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
	"strings"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterStreamCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListStreams",
			Description: "List streams",
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults},
			Setup:       RegisterStreamFormat,
			Execute:     ListStreams,
		},
		&ytapi.Command{
			Name:        "DeleteStream",
			Description: "Delete stream",
			Required:    []*ytapi.Flag{&ytapi.FlagStream},
			Execute:     DeleteStream,
		},
		&ytapi.Command{
			Name:        "NewStream",
			Description: "Create live stream",
			Required:    []*ytapi.Flag{&ytapi.FlagTitle, &ytapi.FlagStreamResolution},
			Optional:    []*ytapi.Flag{&ytapi.FlagDescription, &ytapi.FlagStreamType, &ytapi.FlagStreamReusable},
			Setup:       RegisterStreamFormat,
			Execute:     InsertStream,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterStreamFormat(values *ytapi.Values, table *ytapi.Table) error {

	// register parts

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "stream", Path: "Id", Type: ytapi.FLAG_STREAM},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "default", Path: "Snippet/IsDefaultStream", Type: ytapi.FLAG_BOOL},
	})

	table.RegisterPart("cdn", []*ytapi.Flag{
		&ytapi.Flag{Name: "format", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "type", Path: "Cdn/IngestionType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "key", Path: "Cdn/IngestionInfo/StreamName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "ingestionAddress", Path: "Cdn/IngestionInfo/IngestionAddress", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "backupIngestionAddress", Path: "Cdn/IngestionInfo/BackupIngestionAddress", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "resolution", Path: "Cdn/Resolution", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "frameRate", Path: "Cdn/FrameRate", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "streamStatus", Path: "Status/StreamStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "healthStatus", Path: "Status/HealthStatus/Status", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "lastUpdateTime", Path: "Status/HealthStatus/LastUpdateTimeSeconds", Type: ytapi.FLAG_UINT},
	})

	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "closedCaptionsIngestionUrl", Path: "ContentDetails/ClosedCaptionsIngestionUrl", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "isReusable", Path: "ContentDetails/IsReusable", Type: ytapi.FLAG_BOOL},
	})

	// set default columns
	table.SetColumns([]string{"key", "title", "type", "default", "stream"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// LiveStreams.List

func ListStreams(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// Set the call parameters
	call := service.API.LiveStreams.List(strings.Join(table.Parts(false), ","))
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		if values.IsSet(&ytapi.FlagChannel) {
			call = call.OnBehalfOfContentOwnerChannel(values.GetString(&ytapi.FlagChannel))
		}
	}
	call = call.Mine(true)

	// Perform search, and return results
	return ytapi.DoStreamsList(call, table, int64(maxresults), service.CallOptions()...)
}

////////////////////////////////////////////////////////////////////////////////
// LiveStreams.Delete

func DeleteStream(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	stream := values.GetString(&ytapi.FlagStream)

	// Create call, set parameters
	call := service.API.LiveStreams.Delete(stream)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
		if channel == "" {
			return errors.New("Invalid channel parameter")
		} else {
			call = call.OnBehalfOfContentOwnerChannel(channel)
		}
	} else if channel != "" {
		return errors.New("Invalid channel parameter")
	}

	// Perform search, and return results
	err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// LiveStreams.Insert

func InsertStream(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)

	// Create call, set parameters
	call := service.API.LiveStreams.Insert("snippet,cdn,contentDetails", &youtube.LiveStream{
		Snippet: &youtube.LiveStreamSnippet{
			Title:       values.GetString(&ytapi.FlagTitle),
			Description: values.GetString(&ytapi.FlagDescription),
		},
		Cdn: &youtube.CdnSettings{
			Format:        values.GetString(&ytapi.FlagStreamResolution),
			IngestionType: values.GetString(&ytapi.FlagStreamType),
			ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
				"format":        &ytapi.FlagStreamResolution,
				"ingestionType": &ytapi.FlagStreamType,
			}),
		},
		ContentDetails: &youtube.LiveStreamContentDetails{
			IsReusable: values.GetBool(&ytapi.FlagStreamReusable),
			ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
				"IsReusable": &ytapi.FlagStreamReusable,
			}),
		},
	})

	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
		if channel == "" {
			return errors.New("Invalid channel parameter")
		} else {
			call = call.OnBehalfOfContentOwnerChannel(channel)
		}
	} else if channel != "" {
		return errors.New("Invalid channel parameter")
	}

	// Insert broadcast and get response
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Retrieve stream details
	call2 := service.API.LiveStreams.List(strings.Join(table.Parts(false), ",")).Id(response.Id)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
		if channel == "" {
			return errors.New("Invalid channel parameter")
		} else {
			call2 = call2.OnBehalfOfContentOwnerChannel(channel)
		}
	} else if channel != "" {
		return errors.New("Invalid channel parameter")
	}

	// Perform LiveStreams.list and return results
	return ytapi.DoStreamsList(call2, table, 1, service.CallOptions()...)
}
