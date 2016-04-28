/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"errors"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register broadcast commands

func RegisterBroadcastCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListBroadcasts",
			Description: "List broadcasts",
			Optional:    []*ytapi.Flag{ &ytapi.FlagBroadcastStatus, &ytapi.FlagMaxResults},
			Setup:       RegisterBroadcastFormat,
			Execute:     ListBroadcasts,
		},
		ytapi.Command{
			Name:        "DeleteBroadcast",
			Description: "Delete broadcast",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Execute:     DeleteBroadcast,
		},
		ytapi.Command{
			Name:        "TransitionBroadcast",
			Description: "Transition broadcast to another state",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagBroadcastTransition},
			Execute:     TransitionBroadcast,
		},
		ytapi.Command{
			Name:        "BindBroadcast",
			Description: "Bind or unbind broadcast to stream",
			Optional:    []*ytapi.Flag{&ytapi.FlagContentOwner, &ytapi.FlagStream},
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Execute:     TransitionBroadcast,
		},
		ytapi.Command{
			Name:        "NewBroadcast",
			Description: "Create a new broadcast",
			Optional: []*ytapi.Flag{
				&ytapi.FlagDescription, &ytapi.FlagEndTime, &ytapi.FlagDvr,
				&ytapi.FlagContentEncryption, &ytapi.FlagEmbed,
				&ytapi.FlagRecordFromStart, &ytapi.FlagStartWithSlate,
				&ytapi.FlagClosedCaptions, &ytapi.FlagMonitorStream,
				&ytapi.FlagBroadcastDelay, &ytapi.FlagLowLatency,
			},
			Required: []*ytapi.Flag{
				&ytapi.FlagTitle, &ytapi.FlagStartTime, &ytapi.FlagPrivacyStatus,
			},
			Execute: InsertBroadcast,
		},
	}
}

func RegisterBroadcastFormat(values *ytapi.Values, table *ytapi.Table) error {

	// register parts
	table.RegisterPart("id", []ytapi.FieldSpec{
		ytapi.FieldSpec{"broadcast", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytapi.FieldSpec{
		ytapi.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"scheduledStartTime", "Snippet/ScheduledStartTime", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"scheduledEndTime", "Snippet/ScheduledEndTime", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"actualStartTime", "Snippet/ActualStartTime", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"actualEndTime", "Snippet/ActualEndTime", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"isLiveBroadcast", "Snippet/IsDefaultBroadcast", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"chat", "Snippet/LiveChatId", ytservice.FIELD_STRING},
	})

	table.RegisterPart("status", []ytapi.FieldSpec{
		ytapi.FieldSpec{"status", "Status/LifeCycleStatus", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"privacyStatus", "Status/PrivacyStatus", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"recordingStatus", "Status/RecordingStatus", ytservice.FIELD_STRING},
	})

	table.RegisterPart("contentDetails", []ytapi.FieldSpec{
		ytapi.FieldSpec{"stream", "ContentDetails/BoundStreamId", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"enableMonitorStream", "ContentDetails/MonitorStream/EnableMonitorStream", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"broadcastStreamDelayMs", "ContentDetails/MonitorStream/BroadcastStreamDelayMs", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"embedHtml", "ContentDetails/MonitorStream/EmbedHtml", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"enableEmbed", "ContentDetails/EnableEmbed", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"enableDvr", "ContentDetails/EnableDvr", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"enableContentEncryption", "ContentDetails/EnableContentEncryption", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"startWithSlate", "ContentDetails/StartWithSlate", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"recordFromStart", "ContentDetails/RecordFromStart", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"enableClosedCaptions", "ContentDetails/EnableClosedCaptions", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"closedCaptionsType", "ContentDetails/ClosedCaptionsType", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"enableLowLatency", "ContentDetails/EnableLowLatency", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("statistics", []ytapi.FieldSpec{
		ytapi.FieldSpec{"chatcount", "Statistics/TotalChatCount", ytservice.FIELD_NUMBER},
	})

	// set default columns
	table.SetColumns([]string{"broadcast", "title", "description", "status", "chat"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Broadcasts

func ListBroadcasts(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	status := values.GetString(&ytapi.FlagBroadcastStatus)
	parts := "id,snippet,status" //strings.Join(table.Parts(), ",")

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
	} else if status != "" {
		call = call.BroadcastStatus(status)
	} else {
		call = call.Mine(true)
	}

	// Perform search, and return results
	return ytapi.DoBroadcastsList(call, table, int64(maxresults))
}

////////////////////////////////////////////////////////////////////////////////
// Delete Broadcast

func DeleteBroadcast(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	video := values.GetString(&ytapi.FlagVideo)

	// Create call, set parameters
	call := service.API.LiveBroadcasts.Delete(video)
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
	err := call.Do()
	if err != nil {
		return err
	}

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Transition Broadcast

func TransitionBroadcast(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	video := values.GetString(&ytapi.FlagVideo)
	transition := values.GetString(&ytapi.FlagBroadcastTransition)

	// Create call, set parameters
	call := service.API.LiveBroadcasts.Transition(transition, video, "id,snippet,status")
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
	_, err := call.Do()
	if err != nil {
		return err
	}

	// TODO: retrieve broadcast again and print out values

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Insert Broadcast

func InsertBroadcast(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)

	// Create call, set parameters
	call := service.API.LiveBroadcasts.Insert("id,snippet,status,contentDetails", &youtube.LiveBroadcast{
		Snippet: &youtube.LiveBroadcastSnippet{
			Title:              values.GetString(&ytapi.FlagTitle),
			Description:        values.GetString(&ytapi.FlagDescription),
			ScheduledStartTime: values.GetTimeInISOFormat(&ytapi.FlagStartTime),
			ScheduledEndTime:   values.GetTimeInISOFormat(&ytapi.FlagEndTime),
		},
		Status: &youtube.LiveBroadcastStatus{
			PrivacyStatus: values.GetString(&ytapi.FlagPrivacyStatus),
		},
		ContentDetails: &youtube.LiveBroadcastContentDetails{
			EnableDvr:               values.GetBool(&ytapi.FlagDvr),
			EnableContentEncryption: values.GetBool(&ytapi.FlagContentEncryption),
			EnableEmbed:             values.GetBool(&ytapi.FlagEmbed),
			EnableLowLatency:        values.GetBool(&ytapi.FlagLowLatency),
			RecordFromStart:         values.GetBool(&ytapi.FlagRecordFromStart),
			StartWithSlate:          values.GetBool(&ytapi.FlagStartWithSlate),
			EnableClosedCaptions:    values.GetBool(&ytapi.FlagClosedCaptions),
			MonitorStream: &youtube.MonitorStreamInfo{
				EnableMonitorStream:    values.GetBool(&ytapi.FlagMonitorStream),
				BroadcastStreamDelayMs: int64(values.GetUint(&ytapi.FlagBroadcastDelay)),
				ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
					"EnableMonitorStream": &ytapi.FlagMonitorStream,
				}),
			},
			ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
				"EnableDvr":               &ytapi.FlagDvr,
				"EnableLowLatency":        &ytapi.FlagLowLatency,
				"EnableContentEncryption": &ytapi.FlagContentEncryption,
				"EnableEmbed":             &ytapi.FlagEmbed,
				"RecordFromStart":         &ytapi.FlagRecordFromStart,
				"StartWithSlate":          &ytapi.FlagStartWithSlate,
				"EnableClosedCaptions":    &ytapi.FlagClosedCaptions,
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
	_, err := call.Do()
	if err != nil {
		return err
	}

	// TODO: retrieve broadcast again and print out values

	// success
	return nil
}
