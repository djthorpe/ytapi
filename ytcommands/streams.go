/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"errors"
	"fmt"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

var (
	channelCache   map[ytapi.Channel]bool  // tells us we have cached keys for this channel
	streamKeyCache map[ytapi.Stream]string // maps stream keys to id's
)

////////////////////////////////////////////////////////////////////////////////
// Generate cache of stream key => stream id for a content-owner/channel pair

func CacheStreamKeys(service *ytservice.Service) error {
	if streamKeyCache != nil {
		// cached
		return nil
	}

	streamKeyCache = make(map[ytapi.Stream]string, 0)
	channelCache = make(map[ytapi.Channel]bool, 0)

	// TODO!!!!!

	return nil
}

func StreamLookup(service *ytservice.Service, value string) (string, error) {
	if err := CacheStreamKeys(service); err != nil {
		return "", err
	}
	return "", errors.New(fmt.Sprint("Stream key not found: ", value))
}

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
		&ytapi.Flag{Name: "title", Path: "Snippet/Title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Path: "Snippet/Description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "publishedAt", Path: "Snippet/PublishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "isDefaultStream", Path: "Snippet/IsDefaultStream", Type: ytapi.FLAG_BOOL},
	})

	table.RegisterPart("cdn", []*ytapi.Flag{
		&ytapi.Flag{Name: "format", Path: "Cdn/Format", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "ingestionType", Path: "Cdn/IngestionType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "key", Path: "Cdn/IngestionInfo/StreamName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "ingestionAddress", Path: "Cdn/IngestionInfo/IngestionAddress", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "backupIngestionAddress", Path: "Cdn/IngestionInfo/BackupIngestionAddress", Type: ytapi.FLAG_STRING},
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
	table.SetColumns([]string{"key", "title", "format", "isDefaultStream"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// LiveStreams.List

func ListStreams(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	parts := "id,snippet,cdn,status" //strings.Join(table.Parts(), ",")

	// create call and set parameters
	call := service.API.LiveStreams.List(parts)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
		if channel == "" {
			return errors.New("Invalid channel parameter")
		} else {
			call = call.OnBehalfOfContentOwnerChannel(channel)
		}
	} else if channel != "" {
		return errors.New("Invalid channel parameter")
	} else {
		call = call.Mine(true)
	}

	// Perform search, and return results
	return ytapi.DoStreamsList(call, table, int64(maxresults))
}

////////////////////////////////////////////////////////////////////////////////
// LiveStreams.Delete

func DeleteStream(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	stream, err := StreamLookup(service, values.GetString(&ytapi.FlagStream))
	if err != nil {
		return err
	}

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
	err = call.Do()
	if err != nil {
		return err
	}

	// success
	return nil
}
