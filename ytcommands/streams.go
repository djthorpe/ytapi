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
	channelCache map[ytapi.Channel]bool // tells us we have cached keys for this channel
	streamKeyCache map[ytapi.Stream]string // maps stream keys to id's
)

////////////////////////////////////////////////////////////////////////////////
// Generate cache of stream key => stream id for a content-owner/channel pair

func CacheStreamKeys(service *ytservice.Service) error {
	if streamKeyCache != nil {
		// cached
		return nil
	}

	streamKeyCache = make(map[ytapi.Stream]string,0)
	channelCache = make(map[ytapi.Channel]bool,0)

	// TODO!!!!!

	return nil
}

func StreamLookup(service *ytservice.Service,value string) (string,error) {
	if err := CacheStreamKeys(service); err != nil {
		return "",err
	}
	return "",errors.New(fmt.Sprint("Stream key not found: ",value))
}

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterStreamCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListStreams",
			Description: "List streams",
			Optional:    []*ytapi.Flag{ &ytapi.FlagMaxResults },
			Setup:       RegisterStreamFormat,
			Execute:     ListStreams,
		},
		ytapi.Command{
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

	table.RegisterPart("id", []ytapi.FieldSpec{
		ytapi.FieldSpec{"stream", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytapi.FieldSpec{
		ytapi.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"isDefaultStream", "Snippet/IsDefaultStream", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("cdn", []ytapi.FieldSpec{
		ytapi.FieldSpec{"format", "Cdn/Format", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"ingestionType", "Cdn/IngestionType", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"key", "Cdn/IngestionInfo/StreamName", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"ingestionAddress", "Cdn/IngestionInfo/IngestionAddress", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"backupIngestionAddress", "Cdn/IngestionInfo/BackupIngestionAddress", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("status", []ytapi.FieldSpec{
		ytapi.FieldSpec{"streamStatus", "Status/StreamStatus", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"healthStatus", "Status/HealthStatus/Status", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"lastUpdateTime", "Status/HealthStatus/LastUpdateTimeSeconds", ytservice.FIELD_NUMBER},
	})

	table.RegisterPart("contentDetails", []ytapi.FieldSpec{
		ytapi.FieldSpec{"closedCaptionsIngestionUrl", "ContentDetails/ClosedCaptionsIngestionUrl", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"isReusable", "ContentDetails/IsReusable", ytservice.FIELD_BOOLEAN},
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
	stream,err := StreamLookup(service,values.GetString(&ytapi.FlagStream))
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
