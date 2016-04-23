/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
    "errors"

    "github.com/djthorpe/ytapi/ytapi"
    "github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterBroadcastCommands() []ytapi.Command {
    return []ytapi.Command{
        ytapi.Command{
            Name: "ListBroadcasts",
            Description: "List broadcasts",
            Optional: []*ytapi.Flag{ &ytapi.FlagContentOwner,&ytapi.FlagChannel,&ytapi.FlagBroadcastStatus,&ytapi.FlagMaxResults },
            Setup: RegisterBroadcastFormat,
            Execute: ListBroadcasts,
        },
        ytapi.Command{
            Name: "DeleteBroadcast",
            Description: "Delete broadcast",
            Optional: []*ytapi.Flag{ &ytapi.FlagContentOwner,&ytapi.FlagChannel },
            Required: []*ytapi.Flag{ &ytapi.FlagVideo },
            Execute: DeleteBroadcast,
        },
        ytapi.Command{
            Name: "NewBroadcast",
            Description: "Create a new broadcast",
            Optional: []*ytapi.Flag{ &ytapi.FlagContentOwner,&ytapi.FlagChannel,&ytapi.FlagDescription },
            Required: []*ytapi.Flag{ &ytapi.FlagTitle },
            Execute: InsertBroadcast,
        },
    }
}

func RegisterBroadcastFormat(values *ytapi.Values, table *ytservice.Table) error {

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

func ListBroadcasts(service *ytservice.Service, values *ytapi.Values, table *ytservice.Table) error {
    // set parameters
    maxresults := values.GetUint(&ytapi.FlagMaxResults)
    contentowner := values.GetString(&ytapi.FlagContentOwner)
    channel := values.GetString(&ytapi.FlagChannel)
    status := values.GetString(&ytapi.FlagBroadcastStatus)
    parts := "id" //strings.Join(table.Parts(), ",")

    // create call and set parameters
	call := service.API.LiveBroadcasts.List(parts)
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
    if status != "" {
        call = call.BroadcastStatus(status)
    }

    // Perform search, and return results
	return service.DoBroadcastsList(call, table, int64(maxresults))
}

////////////////////////////////////////////////////////////////////////////////
// Delete Broadcast

func DeleteBroadcast(service *ytservice.Service, values *ytapi.Values, table *ytservice.Table) error {
/*
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
*/
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Insert Broadcast

func InsertBroadcast(service *ytservice.Service, values *ytapi.Values, table *ytservice.Table) error {
    // success
    return nil
}

