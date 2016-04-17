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

func RegisterBroadcastFormat(params *ytservice.Params, table *ytservice.Table) error {

	// register parts

	table.RegisterPart("id", []ytservice.FieldSpec{
		ytservice.FieldSpec{"broadcast", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytservice.FieldSpec{
		ytservice.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"scheduledStartTime", "Snippet/ScheduledStartTime", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"scheduledEndTime", "Snippet/ScheduledEndTime", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"actualStartTime", "Snippet/ActualStartTime", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"actualEndTime", "Snippet/ActualEndTime", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"isLiveBroadcast", "Snippet/IsDefaultBroadcast", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"chat", "Snippet/LiveChatId", ytservice.FIELD_STRING},
	})

	table.RegisterPart("status", []ytservice.FieldSpec{
		ytservice.FieldSpec{"status", "Status/LifeCycleStatus", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"privacyStatus", "Status/PrivacyStatus", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"recordingStatus", "Status/RecordingStatus", ytservice.FIELD_STRING},
	})

	table.RegisterPart("contentDetails", []ytservice.FieldSpec{
		ytservice.FieldSpec{"stream", "ContentDetails/BoundStreamId", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"enableMonitorStream", "ContentDetails/MonitorStream/EnableMonitorStream", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"broadcastStreamDelayMs", "ContentDetails/MonitorStream/BroadcastStreamDelayMs", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"embedHtml", "ContentDetails/MonitorStream/EmbedHtml", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"enableEmbed", "ContentDetails/EnableEmbed", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"enableDvr", "ContentDetails/EnableDvr", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"enableContentEncryption", "ContentDetails/EnableContentEncryption", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"startWithSlate", "ContentDetails/StartWithSlate", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"recordFromStart", "ContentDetails/RecordFromStart", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"enableClosedCaptions", "ContentDetails/EnableClosedCaptions", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"closedCaptionsType", "ContentDetails/ClosedCaptionsType", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"enableLowLatency", "ContentDetails/EnableLowLatency", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("statistics", []ytservice.FieldSpec{
		ytservice.FieldSpec{"chatcount", "Statistics/TotalChatCount", ytservice.FIELD_NUMBER},
	})

	// set default columns
	table.SetColumns([]string{"broadcast", "title", "description", "status", "chat"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Broadcasts

func ListBroadcasts(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call
	call := service.API.LiveBroadcasts.List(strings.Join(table.Parts(), ","))

	// set filter parameters
	if service.ServiceAccount && params.IsValidChannel() {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).OnBehalfOfContentOwnerChannel(*params.Channel)
	}
	if params.IsEmptyBroadcastStatus() == false {
		call = call.BroadcastStatus(*params.BroadcastStatus)
	} else {
		call = call.Mine(true)
	}

	// Perform search, and return results
	return service.DoBroadcastsList(call, table, params.MaxResults)
}

////////////////////////////////////////////////////////////////////////////////
// Delete Broadcast

func DeleteBroadcast(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// Get video
	if params.IsValidVideo() == false {
		return ytservice.NewError(ytservice.ErrorBadParameter,nil)
	}

	// create call
	call := service.API.LiveBroadcasts.Delete(*params.Video)

	// set filter parameters
	if service.ServiceAccount && params.IsValidChannel() {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).OnBehalfOfContentOwnerChannel(*params.Channel)
	}

	// Perform search, and return results
	err := call.Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}

	// success
	return nil
}
