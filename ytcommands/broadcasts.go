/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterBroadcastFormat(params *ytservice.Params,table *ytservice.Table) error {

	// register parts

	table.RegisterPart("id",[]ytservice.FieldSpec{
		ytservice.FieldSpec{ "broadcast","Id",ytservice.FIELD_STRING },
	})

	table.RegisterPart("snippet",[]ytservice.FieldSpec{
		ytservice.FieldSpec{ "title","Snippet/Title",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "description","Snippet/Description",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "channel","Snippet/ChannelId",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "publishedAt","Snippet/PublishedAt",ytservice.FIELD_DATETIME },
		ytservice.FieldSpec{ "scheduledStartTime","Snippet/ScheduledStartTime",ytservice.FIELD_DATETIME },
		ytservice.FieldSpec{ "scheduledEndTime","Snippet/ScheduledEndTime",ytservice.FIELD_DATETIME },
		ytservice.FieldSpec{ "actualStartTime","Snippet/ActualStartTime",ytservice.FIELD_DATETIME },
		ytservice.FieldSpec{ "actualEndTime","Snippet/ActualEndTime",ytservice.FIELD_DATETIME },
		ytservice.FieldSpec{ "isLiveBroadcast","Snippet/IsDefaultBroadcast",ytservice.FIELD_BOOLEAN },
		ytservice.FieldSpec{ "chat","Snippet/LiveChatId",ytservice.FIELD_STRING },
	})

	table.RegisterPart("status",[]ytservice.FieldSpec{
		ytservice.FieldSpec{ "status","Status/LifeCycleStatus",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "privacyStatus","Status/PrivacyStatus",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "recordingStatus","Status/RecordingStatus",ytservice.FIELD_STRING },
	})

	// set default columns
	table.SetColumns([]string{ "broadcast","title","description","status","chat" })

	// success
	return nil
}


////////////////////////////////////////////////////////////////////////////////
// List Broadcasts

func ListBroadcasts(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call
	call := service.API.LiveBroadcasts.List(strings.Join(table.Parts(),","))

	// set filter parameters
	if service.ServiceAccount && params.IsValidChannel() {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).OnBehalfOfContentOwnerChannel(*params.Channel)
	}
	call = call.Mine(true)

	// Perform search, and return results
	return service.DoBroadcastsList(call,table,params.MaxResults)
}

////////////////////////////////////////////////////////////////////////////////
// Delete Broadcast

func DeleteBroadcast(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call
	call := service.API.LiveBroadcasts.Delete("XX")

	// set filter parameters
	if service.ServiceAccount && params.IsValidChannel() {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).OnBehalfOfContentOwnerChannel(*params.Channel)
	}

	// Perform search, and return results
	err := call.Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse,err)
	}

	// success
	return nil
}


