/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"
	"errors"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register broadcast commands

func RegisterBroadcastCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListBroadcasts",
			Description: "List broadcasts",
			Optional:    []*ytapi.Flag{&ytapi.FlagBroadcastStatus, &ytapi.FlagMaxResults},
			Setup:       RegisterBroadcastFormat,
			Execute:     ListBroadcasts,
		},
		&ytapi.Command{
			Name:        "DeleteBroadcast",
			Description: "Delete broadcast",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Execute:     DeleteBroadcast,
		},
		&ytapi.Command{
			Name:        "TransitionBroadcast",
			Description: "Transition broadcast to another state",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagBroadcastTransition},
			Execute:     TransitionBroadcast,
		},
		&ytapi.Command{
			Name:        "BindBroadcast",
			Description: "Bind or unbind broadcast to stream",
			Optional:    []*ytapi.Flag{&ytapi.FlagContentOwner, &ytapi.FlagStream},
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Execute:     TransitionBroadcast,
		},
		&ytapi.Command{
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

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "broadcast", Path: "Id", Type: ytapi.FLAG_VIDEO},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "scheduledStartTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "scheduledEndTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "actualStartTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "actualEndTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "defaultBroadcast", Path: "Snippet/IsDefaultBroadcast", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "chat", Path: "Snippet/LiveChatId", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "status", Path: "Status/LifeCycleStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "privacyStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "recordingStatus", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "stream", Path: "ContentDetails/BoundStreamId", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "enableMonitorStream", Path: "ContentDetails/MonitorStream/EnableMonitorStream", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "broadcastStreamDelayMs", Path: "ContentDetails/MonitorStream/BroadcastStreamDelayMs", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "embedHtml", Path: "ContentDetails/MonitorStream/EmbedHtml", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "embed", Path: "ContentDetails/EnableEmbed", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "dvr", Path: "ContentDetails/EnableDvr", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "encryption", Path: "ContentDetails/EnableContentEncryption", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "slate", Path: "ContentDetails/StartWithSlate", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "record", Path: "ContentDetails/RecordFromStart", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "captions", Path: "ContentDetails/EnableClosedCaptions", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "captions.type", Path: "ContentDetails/ClosedCaptionsType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "lowlatency", Path: "ContentDetails/EnableLowLatency", Type: ytapi.FLAG_BOOL},
	})

	table.RegisterPart("statistics", []*ytapi.Flag{
		&ytapi.Flag{Name: "chatcount", Path: "Statistics/TotalChatCount", Type: ytapi.FLAG_UINT},
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

	// Set the call parameters
	call := service.API.LiveBroadcasts.List(strings.Join(table.Parts(false), ","))
	call = call.BroadcastType("all")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		if values.IsSet(&ytapi.FlagChannel) {
			call = call.OnBehalfOfContentOwnerChannel(values.GetString(&ytapi.FlagChannel))
		}
	}

	// Status
	if values.IsSet(&ytapi.FlagBroadcastStatus) {
		call = call.BroadcastStatus(values.GetString(&ytapi.FlagBroadcastStatus))
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
