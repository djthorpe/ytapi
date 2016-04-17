/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"github.com/djthorpe/ytapi/ytservice"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterStreamFormat(params *ytservice.Params, table *ytservice.Table) error {

	// register parts

	table.RegisterPart("id", []ytservice.FieldSpec{
		ytservice.FieldSpec{"stream", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytservice.FieldSpec{
		ytservice.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"isDefaultStream", "Snippet/IsDefaultStream", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("cdn", []ytservice.FieldSpec{
		ytservice.FieldSpec{"format", "Cdn/Format", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"ingestionType", "Cdn/IngestionType", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"streamkey", "Cdn/IngestionInfo/StreamName", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"ingestionAddress", "Cdn/IngestionInfo/IngestionAddress", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"backupIngestionAddress", "Cdn/IngestionInfo/BackupIngestionAddress", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("status", []ytservice.FieldSpec{
		ytservice.FieldSpec{"streamStatus", "Status/StreamStatus", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"healthStatus", "Status/HealthStatus/Status", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"lastUpdateTime", "Status/HealthStatus/LastUpdateTimeSeconds", ytservice.FIELD_NUMBER},
	})

	table.RegisterPart("contentDetails", []ytservice.FieldSpec{
		ytservice.FieldSpec{"closedCaptionsIngestionUrl", "ContentDetails/ClosedCaptionsIngestionUrl", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"isReusable", "ContentDetails/IsReusable", ytservice.FIELD_BOOLEAN},
	})

	// set default columns
	table.SetColumns([]string{"streamkey", "title", "format", "isDefaultStream"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// LiveStreams.List

func ListStreams(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call
	call := service.API.LiveStreams.List(strings.Join(table.Parts(), ","))

	// set filter parameters
	if service.ServiceAccount && params.IsValidChannel() {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).OnBehalfOfContentOwnerChannel(*params.Channel)
	}
	call = call.Mine(true)

	// Perform operation, and return results
	return service.DoStreamsList(call, table, params.MaxResults)
}


////////////////////////////////////////////////////////////////////////////////
// LiveStreams.Delete

func DeleteStream(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// Get stream
	if params.IsValidStream() == false {
		return ytservice.NewError(ytservice.ErrorBadParameter,nil)
	}

	// create call
	call := service.API.LiveStreams.Delete(*params.Stream)

	// set filter parameters
	if service.ServiceAccount && params.IsValidChannel() {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).OnBehalfOfContentOwnerChannel(*params.Channel)
	}

	// Perform delete
	err := call.Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}

	// success
	return nil
}
