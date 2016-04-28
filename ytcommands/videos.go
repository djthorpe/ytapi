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

func RegisterVideoFormat(values *ytapi.Values, table *ytservice.Table) error {

	// register parts
	table.RegisterPart("id", []ytservice.FieldSpec{
		ytservice.FieldSpec{"video", "Id", ytservice.FIELD_STRING},
	})

	// snippet
	table.RegisterPart("snippet", []ytservice.FieldSpec{
		ytservice.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"tags", "Snippet/Tags", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"category", "Snippet/CategoryId", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"liveBroadcastContent", "Snippet/LiveBroadcastContent", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"language", "Snippet/DefaultLanguage", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"audioLanguage", "Snippet/DefaultAudioLanguage", ytservice.FIELD_STRING},
	})

	// contentDetails
	table.RegisterPart("contentDetails", []ytservice.FieldSpec{
		ytservice.FieldSpec{"duration", "ContentDetails/Duration", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"dimension", "ContentDetails/Dimension", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"definition", "ContentDetails/Definition", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"caption", "ContentDetails/Caption", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"licensedContent", "ContentDetails/LicensedContent", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"regionsAllowed", "ContentDetails/RegionRestriction/Allowed", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"regionsBlocked", "ContentDetails/RegionRestriction/Blocked", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"contentRating", "ContentDetails/ContentRating/YtRating", ytservice.FIELD_STRING},
	})

	// status
	table.RegisterPart("status", []ytservice.FieldSpec{
		ytservice.FieldSpec{"privacyStatus", "Status/PrivacyStatus", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"uploadStatus", "Status/UploadStatus", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"failureReason", "Status/FailureReason", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"rejectionReason", "Status/RejectionReason", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"scheduledPublishAt", "Status/PublishAt", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"license", "Status/License", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"embeddable", "Status/Embeddable", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"publicStatsViewable", "Status/PublicStatsViewable", ytservice.FIELD_BOOLEAN},
	})

	// statistics
	table.RegisterPart("statistics", []ytservice.FieldSpec{
		ytservice.FieldSpec{"viewcount", "Statistics/ViewCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"likecount", "Statistics/LikeCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"dislikecount", "Statistics/DislikeCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"favoritecount", "Statistics/FavoriteCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"commentcount", "Statistics/CommentCount", ytservice.FIELD_NUMBER},
	})

	// set default columns
	table.SetColumns([]string{"video","title","privacyStatus"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Videos.list

func ListVideos(service *ytservice.Service, values *ytapi.Values, table *ytservice.Table) error {

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
	return service.DoVideosList(call, table, int64(maxresults))
}



