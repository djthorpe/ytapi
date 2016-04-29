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
// Register video commands

func RegisterVideoCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListVideos",
			Description: "List videos",
			Required:    []*ytapi.Flag{&ytapi.FlagVideoFilter},
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults},
			Setup:       RegisterVideoFormat,
			Execute:     ListVideos,
		},
	}
}

func RegisterVideoFormat(values *ytapi.Values, table *ytapi.Table) error {

	// register parts
	table.RegisterPart("id", []ytapi.Flag{
		ytapi.Flag{Name: "video", Path: "Id", Type: ytapi.FLAG_VIDEO},
	})

	// snippet
	table.RegisterPart("snippet", []ytapi.Flag{
		ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "description",Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		ytapi.Flag{Name: "tags", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "category", Path: "Snippet/CategoryId", Type: ytapi.FLAG_UINT},
		ytapi.Flag{Name: "liveBroadcastContent", Type: ytapi.FLAG_BOOL},
		ytapi.Flag{Name: "language", Path: "Snippet/DefaultLanguage", Type: ytapi.FLAG_LANGUAGE},
		ytapi.Flag{Name: "audioLanguage", Path: "Snippet/DefaultAudioLanguage", Type: ytapi.FLAG_LANGUAGE},
	})

	// contentDetails
	table.RegisterPart("contentDetails", []ytapi.Flag{
		ytapi.Flag{Name: "duration", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "dimension", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "definition", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "caption", Type: ytapi.FLAG_BOOL},
		ytapi.Flag{Name: "licensedContent", Type: ytapi.FLAG_BOOL},
		ytapi.Flag{Name: "regionsAllowed", Path: "ContentDetails/RegionRestriction/Allowed", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "regionsBlocked", Path: "ContentDetails/RegionRestriction/Blocked", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "contentRating", Path: "ContentDetails/ContentRating/YtRating", Type: ytapi.FLAG_STRING},
	})

	// status
	table.RegisterPart("status", []ytapi.Flag{
		ytapi.Flag{Name: "privacyStatus", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "uploadStatus", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "failureReason", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "rejectionReason", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "scheduledPublishAt", Path: "Status/PublishAt", Type: ytapi.FLAG_TIME},
		ytapi.Flag{Name: "license", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "embeddable", Type: ytapi.FLAG_BOOL},
		ytapi.Flag{Name: "publicStatsViewable", Type: ytapi.FLAG_BOOL},
	})

	// statistics
	table.RegisterPart("statistics", []ytapi.Flag{
		ytapi.Flag{Name: "viewCount", Type: ytapi.FLAG_UINT},
		ytapi.Flag{Name: "likeCount",  Type: ytapi.FLAG_UINT},
		ytapi.Flag{Name: "dislikeCount", Type: ytapi.FLAG_UINT},
		ytapi.Flag{Name: "favoriteCount", Type: ytapi.FLAG_UINT},
		ytapi.Flag{Name: "commentCount", Type: ytapi.FLAG_UINT},
	})

	// set default columns
	table.SetColumns([]string{"video", "title", "privacyStatus"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Videos.list

func ListVideos(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	parts := "id,snippet,status" //strings.Join(table.Parts(), ",")
	language := values.GetString(&ytapi.FlagLanguage)
	region := values.GetString(&ytapi.FlagRegion)
	filter := values.GetString(&ytapi.FlagVideoFilter)

	// create call and set parameters
	call := service.API.Videos.List(parts)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if language != "" {
		call = call.Hl(language)
	}
	if region != "" {
		call = call.RegionCode(region)
	}
	if filter == "chart" {
		call = call.Chart("mostPopular")
	} else if filter == "like" || filter == "dislike" {
		call = call.MyRating(filter)
	} else {
		return errors.New("Missing or invalid filter flag")
	}
	// TODO: videoCategoryId

	// Perform search, and return results
	return ytapi.DoVideosList(call, table, int64(maxresults))
}
