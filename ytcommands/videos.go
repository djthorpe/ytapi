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
			Required:    []*ytapi.Flag{ &ytapi.FlagVideoFilter },
			Optional:    []*ytapi.Flag{ &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults },
			Setup:       RegisterVideoFormat,
			Execute:     ListVideos,
		},
	}
}

func RegisterVideoFormat(values *ytapi.Values, table *ytapi.Table) error {

	// register parts
	table.RegisterPart("id", []ytapi.FieldSpec{
		ytapi.FieldSpec{"video", "Id", ytservice.FIELD_STRING},
	})

	// snippet
	table.RegisterPart("snippet", []ytapi.FieldSpec{
		ytapi.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"tags", "Snippet/Tags", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"category", "Snippet/CategoryId", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"liveBroadcastContent", "Snippet/LiveBroadcastContent", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"language", "Snippet/DefaultLanguage", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"audioLanguage", "Snippet/DefaultAudioLanguage", ytservice.FIELD_STRING},
	})

	// contentDetails
	table.RegisterPart("contentDetails", []ytapi.FieldSpec{
		ytapi.FieldSpec{"duration", "ContentDetails/Duration", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"dimension", "ContentDetails/Dimension", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"definition", "ContentDetails/Definition", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"caption", "ContentDetails/Caption", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"licensedContent", "ContentDetails/LicensedContent", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"regionsAllowed", "ContentDetails/RegionRestriction/Allowed", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"regionsBlocked", "ContentDetails/RegionRestriction/Blocked", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"contentRating", "ContentDetails/ContentRating/YtRating", ytservice.FIELD_STRING},
	})

	// status
	table.RegisterPart("status", []ytapi.FieldSpec{
		ytapi.FieldSpec{"privacyStatus", "Status/PrivacyStatus", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"uploadStatus", "Status/UploadStatus", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"failureReason", "Status/FailureReason", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"rejectionReason", "Status/RejectionReason", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"scheduledPublishAt", "Status/PublishAt", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"license", "Status/License", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"embeddable", "Status/Embeddable", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"publicStatsViewable", "Status/PublicStatsViewable", ytservice.FIELD_BOOLEAN},
	})

	// statistics
	table.RegisterPart("statistics", []ytapi.FieldSpec{
		ytapi.FieldSpec{"viewcount", "Statistics/ViewCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"likecount", "Statistics/LikeCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"dislikecount", "Statistics/DislikeCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"favoritecount", "Statistics/FavoriteCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"commentcount", "Statistics/CommentCount", ytservice.FIELD_NUMBER},
	})

	// set default columns
	table.SetColumns([]string{"video","title","privacyStatus"})

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



