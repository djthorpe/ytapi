/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"os"
	"errors"
	"strings"
	"strconv"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)


////////////////////////////////////////////////////////////////////////////////

type Thumbnail struct {
	Name    string
	Url     string
	Width   uint64
	Height  uint64
}

////////////////////////////////////////////////////////////////////////////////
// Register video commands

func RegisterVideoCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListVideos",
			Description: "List videos",
			Required:    []*ytapi.Flag{&ytapi.FlagVideoFilter},
			Optional:    []*ytapi.Flag{&ytapi.FlagVideoCategory, &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults},
			Setup:       RegisterVideoFormat,
			Execute:     ListVideos,
		},
		&ytapi.Command{
			Name:        "GetVideo",
			Description: "Get single video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage, &ytapi.FlagRegion},
			Setup:       RegisterVideoFormat,
			Execute:     GetVideo,
		},
		&ytapi.Command{
			Name:        "ListCategories",
			Description: "Get list of video categories",
			Required:    []*ytapi.Flag{ &ytapi.FlagRegion },
			Optional:    []*ytapi.Flag{ &ytapi.FlagLanguage },
			Setup:       RegisterCategoryFormat,
			Execute:     ListCategories,
		},
		&ytapi.Command{
			Name:        "SetVideoThumbnail",
			Description: "Add or update a thumbnail to a video",
			Required:    []*ytapi.Flag{ &ytapi.FlagVideo, &ytapi.FlagFile },
			Setup:       RegisterThumbnailFormat,
			Execute:     SetVideoThumbnail,
		},
		&ytapi.Command{
			Name:        "GetVideoThumbnail",
			Description: "List thumbnails for a video",
			Required:    []*ytapi.Flag{ &ytapi.FlagVideo },
			Setup:       RegisterThumbnailFormat,
			Execute:     GetVideoThumbnail,
		},
	}
}

func RegisterVideoFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "video", Path: "Id", Type: ytapi.FLAG_VIDEO},
	})
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "tags", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "category", Path: "Snippet/CategoryId", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "liveBroadcastContent", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "language", Path: "Snippet/DefaultLanguage", Type: ytapi.FLAG_LANGUAGE},
		&ytapi.Flag{Name: "audioLanguage", Path: "Snippet/DefaultAudioLanguage", Type: ytapi.FLAG_LANGUAGE},
	})
	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "duration", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "dimension", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "definition", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "caption", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "licensedContent", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "regionsAllowed", Path: "ContentDetails/RegionRestriction/Allowed", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "regionsBlocked", Path: "ContentDetails/RegionRestriction/Blocked", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "contentRating", Path: "ContentDetails/ContentRating/YtRating", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "privacyStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "uploadStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "failureReason", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "rejectionReason", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "scheduledPublishAt", Path: "Status/PublishAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "license", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "embeddable", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "publicStatsViewable", Type: ytapi.FLAG_BOOL},
	})
	table.RegisterPart("statistics", []*ytapi.Flag{
		&ytapi.Flag{Name: "viewCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "likeCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "dislikeCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "favoriteCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "commentCount", Type: ytapi.FLAG_UINT},
	})
	table.RegisterPart("liveStreamingDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "actualStartTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "actualEndTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "scheduledStartTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "scheduledEndTime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "activeLiveChatId", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "concurrentViewers", Type: ytapi.FLAG_UINT},
	})

	// set default columns
	table.SetColumns([]string{"video", "title", "privacyStatus"})

	// success
	return nil
}

func RegisterCategoryFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "category", Path: "Id", Type: ytapi.FLAG_UINT},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "assignable", Type: ytapi.FLAG_BOOL},
	})

	// set default columns
	table.SetColumns([]string{"category", "title" })

	// success
	return nil
}

func RegisterThumbnailFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("thumbnails", []*ytapi.Flag{
		&ytapi.Flag{Name: "name", Path: "Name", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "url", Path: "Url", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "width", Path: "Width", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "height", Path: "Height", Type: ytapi.FLAG_UINT},
	})

	// set default columns
	table.SetColumns([]string{"name", "width", "height", "url"})

	// success
	return nil
}


////////////////////////////////////////////////////////////////////////////////
// Videos

func ListVideos(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	parts := strings.Join(table.Parts(false), ",")
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
	if values.IsSet(&ytapi.FlagVideoCategory) {
		category := values.GetUint(&ytapi.FlagVideoCategory)
		call = call.VideoCategoryId(strconv.FormatUint(category,10))
	}
	if filter == "chart" {
		call = call.Chart("mostPopular")
	} else if filter == "like" || filter == "dislike" {
		call = call.MyRating(filter)
	} else {
		return errors.New("Missing or invalid filter flag")
	}

	// Perform search, and return results
	return ytapi.DoVideosList(call, table, int64(maxresults))
}

func GetVideo(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	parts := strings.Join(table.Parts(false), ",")
	language := values.GetString(&ytapi.FlagLanguage)
	region := values.GetString(&ytapi.FlagRegion)
	video := values.GetString(&ytapi.FlagVideo)

	// create call and set parameters
	call := service.API.Videos.List(parts).Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if language != "" {
		call = call.Hl(language)
	}
	if region != "" {
		call = call.RegionCode(region)
	}

	// Perform search, and return results
	return ytapi.DoVideosList(call, table, 0)
}

////////////////////////////////////////////////////////////////////////////////
// Categories

func ListCategories(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// create call and set parameters
	parts := strings.Join(table.Parts(false), ",")
	region := values.GetString(&ytapi.FlagRegion)
	call := service.API.VideoCategories.List(parts)

	// add language parameter
	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.Hl(values.GetString(&ytapi.FlagLanguage))
	}

	// request and response
	response, err := call.RegionCode(region).Do()
	if err != nil {
		return err
	}
	if err = table.Append(response.Items); err != nil {
		return err
	}

	// Success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Thumbnails


func GetVideoThumbnail(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)

	// create call and set parameters
	call := service.API.Videos.List("snippet").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Create Request
	response, err := call.Do()
	if err != nil {
		return err
	}

	// Output Thumbnails
	if len(response.Items) != 1 {
		return errors.New("Not Found")
	}
	thumbnails := map[string]*youtube.Thumbnail{
		"default": response.Items[0].Snippet.Thumbnails.Default,
		"high": response.Items[0].Snippet.Thumbnails.High,
		"maxres": response.Items[0].Snippet.Thumbnails.Maxres,
		"medium": response.Items[0].Snippet.Thumbnails.Medium,
		"standard": response.Items[0].Snippet.Thumbnails.Standard,
	}
	for name,thumbnail := range thumbnails {
		if thumbnail != nil {
			table.Append([]Thumbnail{{name, thumbnail.Url, uint64(thumbnail.Width), uint64(thumbnail.Height) }})
		}
	}

	// Success
	return nil
}


func SetVideoThumbnail(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get Parameters
	video := values.GetString(&ytapi.FlagVideo)
	contentowner := values.GetString(&ytapi.FlagContentOwner)

	// Create Call
	call := service.API.Thumbnails.Set(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Read file
	file, err := os.Open(values.GetString(&ytapi.FlagFile))
	defer file.Close()
	if err != nil {
		return nil
	}

	// Request and Response
	response, err := call.Media(file).Do()
	if err != nil {
		return err
	}

	// Output Thumbnails
	if len(response.Items) != 1 {
		return errors.New("Not Found")
	}
	thumbnails := map[string]*youtube.Thumbnail{
		"default": response.Items[0].Default,
		"high": response.Items[0].High,
		"maxres": response.Items[0].Maxres,
		"medium": response.Items[0].Medium,
		"standard": response.Items[0].Standard,
	}
	for name,thumbnail := range thumbnails {
		if thumbnail != nil {
			table.Append([]Thumbnail{{name, thumbnail.Url, uint64(thumbnail.Width), uint64(thumbnail.Height) }})
		}
	}

	// Success
	return nil
}




