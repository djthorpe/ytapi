/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////

type Thumbnail struct {
	Name   string
	Url    string
	Width  uint64
	Height uint64
}

////////////////////////////////////////////////////////////////////////////////
// Register video commands

func RegisterVideoCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListVideos",
			Description: "List videos",
			Optional:    []*ytapi.Flag{&ytapi.FlagVideoFilter, &ytapi.FlagVideoCategory, &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults},
			Setup:       RegisterVideoFormat,
			Execute:     ListVideos,
		},
		&ytapi.Command{
			Name:        "ListVideosForPlaylist",
			Description: "List videos for a playlist",
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist},
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults},
			Setup:       RegisterVideoFormat,
			Execute:     ListVideosForPlaylist,
		},
		&ytapi.Command{
			Name:        "UploadVideo",
			Description: "Upload a video",
			Required:    []*ytapi.Flag{&ytapi.FlagFile, &ytapi.FlagPrivacyStatus},
			Optional:    []*ytapi.Flag{&ytapi.FlagTitle, &ytapi.FlagDescription, &ytapi.FlagLanguage, &ytapi.FlagVideoCategory},
			Setup:       RegisterVideoFormat,
			Execute:     UploadVideo,
		},
		&ytapi.Command{
			Name:        "DeleteVideo",
			Description: "Remove video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Execute:     DeleteVideo,
		},
		&ytapi.Command{
			Name:        "SetVideoRating",
			Description: "Set like, dislike or remove rating for video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagVideoRating},
			Setup:       RegisterVideoRatingFormat,
			Execute:     SetVideoRating,
		},
		&ytapi.Command{
			Name:        "GetVideoRating",
			Description: "Get rating for video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Setup:       RegisterVideoRatingFormat,
			Execute:     GetVideoRating,
		},
		&ytapi.Command{
			Name:        "GetVideoMetadata",
			Description: "Get metadata for a single video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage, &ytapi.FlagRegion},
			Setup:       RegisterVideoFormat,
			Execute:     GetVideoMetadata,
		},
		&ytapi.Command{
			Name:        "UpdateVideoMetadata",
			Description: "Update metadata for a video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Optional:    []*ytapi.Flag{&ytapi.FlagTitle, &ytapi.FlagDescription, &ytapi.FlagLanguage, &ytapi.FlagVideoCategory},
			Setup:       RegisterVideoFormat,
			Execute:     UpdateVideoMetadata,
		},
		&ytapi.Command{
			Name:        "SetVideoStatus",
			Description: "Set video status (private, public, unlisted, license, embedding or public statistics viewable)",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Optional:    []*ytapi.Flag{&ytapi.FlagPrivacyStatus, &ytapi.FlagEmbeds, &ytapi.FlagLicense, &ytapi.FlagStatsViewable},
			Setup:       RegisterVideoFormat,
			Execute:     SetVideoStatus,
		},
		&ytapi.Command{
			Name:        "SetVideoPublishDate",
			Description: "Set video private and future publish date",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagDate},
			Setup:       RegisterVideoFormat,
			Execute:     SetVideoPublishDate,
		},
		&ytapi.Command{
			Name:        "SetVideoThumbnail",
			Description: "Add or update a thumbnail to a video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagFile},
			Setup:       RegisterThumbnailFormat,
			Execute:     SetVideoThumbnail,
		},
		&ytapi.Command{
			Name:        "GetVideoThumbnail",
			Description: "List thumbnails for a video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Setup:       RegisterThumbnailFormat,
			Execute:     GetVideoThumbnail,
		},
		&ytapi.Command{
			Name:        "GetLocalizedVideoMetadata",
			Description: "Get localized video metadata",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Setup:       RegisterLocalizedVideoMetadataFormat,
			Execute:     GetLocalizedVideoMetadata,
		},
		&ytapi.Command{
			Name:        "UpdateLocalizedVideoMetadata",
			Description: "Update or add localized video metadata",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagLanguage},
			Optional:    []*ytapi.Flag{&ytapi.FlagTitle, &ytapi.FlagDescription},
			Setup:       RegisterLocalizedVideoMetadataFormat,
			Execute:     UpdateLocalizedVideoMetadata,
		},
		&ytapi.Command{
			Name:        "DeleteLocalizedVideoMetadata",
			Description: "Remove localized video metadata",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagLanguage},
			Setup:       RegisterLocalizedVideoMetadataFormat,
			Execute:     DeleteLocalizedVideoMetadata,
		},
		&ytapi.Command{
			Name:        "ListCategories",
			Description: "Get list of video categories",
			Required:    []*ytapi.Flag{&ytapi.FlagRegion},
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage},
			Setup:       RegisterCategoryFormat,
			Execute:     ListCategories,
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
		&ytapi.Flag{Name: "projection", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "privacyStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "uploadStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "failureReason", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "rejectionReason", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "publishAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "license", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "embeds", Path: "Status/Embeddable", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "publicStatsViewable", Type: ytapi.FLAG_BOOL},
	})
	table.RegisterPart("statistics", []*ytapi.Flag{
		&ytapi.Flag{Name: "viewCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "likeCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "dislikeCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "favoriteCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "commentCount", Type: ytapi.FLAG_UINT},
	})
	table.RegisterPart("player", []*ytapi.Flag{
		&ytapi.Flag{Name: "embedHtml", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("recordingDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "location", Path: "RecordingDetails/LocationDescription", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "latitude", Path: "RecordingDetails/Location/Latitude", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "longitude", Path: "RecordingDetails/Location/Latitude", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "altitude", Path: "RecordingDetails/Location/Altitude", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "recordingDate", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("fileDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "fileName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "fileSize", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "fileType", Type: ytapi.FLAG_STRING},
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
	table.SetColumns([]string{"video", "title", "description", "category", "language", "privacyStatus", "publishedAt"})

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
	table.SetColumns([]string{"category", "title"})

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

func RegisterLocalizedVideoMetadataFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("localizations", []*ytapi.Flag{
		&ytapi.Flag{Name: "language", Path: "Language", Type: ytapi.FLAG_LANGUAGE},
		&ytapi.Flag{Name: "title", Path: "Title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Path: "Description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "default", Path: "Default", Type: ytapi.FLAG_BOOL},
	})

	// set default columns
	table.SetColumns([]string{"language", "title", "description", "default"})

	// success
	return nil
}

func RegisterVideoRatingFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "video", Path: "VideoId", Type: ytapi.FLAG_VIDEO},
		&ytapi.Flag{Name: "rating", Path: "Rating", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"video", "rating"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Videos

func UploadVideo(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	filename := values.GetString(&ytapi.FlagFile)
	title := values.GetString(&ytapi.FlagTitle)
	status := values.GetString(&ytapi.FlagPrivacyStatus)
	parts := strings.Join(table.Parts(false), ",")

	// Interpret name from filename if not set
	if values.IsSet(&ytapi.FlagTitle) == false {
		title = filepath.Base(filename)
	}

	// Create the call
	call := service.API.Videos.Insert("snippet,status", &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:           title,
			Description:     values.GetString(&ytapi.FlagDescription),
			CategoryId:      values.GetString(&ytapi.FlagVideoCategory),
			DefaultLanguage: values.GetString(&ytapi.FlagLanguage),
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: status,
		},
	})

	// Set the call parameters
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// Open the caption file file
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	// request and response
	response, err := call.Media(file).Do()
	if err != nil {
		return err
	}

	// now get the video metadata
	call2 := service.API.Videos.List(parts).Id(response.Id)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// Execute
	response2, err := call2.Do()
	if err != nil {
		return err
	}
	if len(response2.Items) != 1 {
		return errors.New("Not Found")
	}

	return table.Append(response2.Items)
}

func ListVideosForPlaylist(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	playlist := values.GetString(&ytapi.FlagPlaylist)
	maxresults := values.GetInt(&ytapi.FlagMaxResults)
	parts := strings.Join(table.Parts(false), ",")

	// create the call
	call := service.API.PlaylistItems.List("snippet").PlaylistId(playlist)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// iterate through to obtain the videos
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(ytapi.YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > ytapi.YouTubeMaxPagingResults {
			retrieveitems = int64(ytapi.YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}

		// Append videos to table
		videos := make([]string, 0)
		for _, item := range response.Items {
			videos = append(videos, item.Snippet.ResourceId.VideoId)
		}

		// Make another call to retrieve videos
		call2 := service.API.Videos.List(parts).Id(strings.Join(videos, ","))
		if service.ServiceAccount {
			call2 = call2.OnBehalfOfContentOwner(contentowner)
		}
		response2, err := call2.Do()
		if err != nil {
			return err
		}
		table.Append(response2.Items)

		// Now iterate to next page
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func ListVideos(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	parts := strings.Join(table.Parts(false), ",")
	language := values.GetString(&ytapi.FlagLanguage)
	region := values.GetString(&ytapi.FlagRegion)
	filter := values.GetString(&ytapi.FlagVideoFilter)

	// if filter is uploads, then switch to Playlist mode
	if filter == "uploads" || filter == "likes" || filter == "favorites" || filter == "watchhistory" || filter == "watchlater" {
		if values.IsSet(&ytapi.FlagVideoCategory) {
			return errors.New("Category cannot be set when listing uploaded videos")
		}
		table.Info("TODO: Get playlist")
		// success
		return nil
	}

	// create call and set parameters
	call := service.API.Videos.List(parts)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.Hl(language)
	}
	if values.IsSet(&ytapi.FlagRegion) {
		call = call.RegionCode(region)
	}
	if values.IsSet(&ytapi.FlagVideoCategory) {
		category := values.GetUint(&ytapi.FlagVideoCategory)
		call = call.VideoCategoryId(strconv.FormatUint(category, 10))
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

func GetVideoMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)
	parts := strings.Join(table.Parts(false), ",")

	// create call and set parameters
	call := service.API.Videos.List(parts).Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.Hl(values.GetString(&ytapi.FlagLanguage))
	}
	if values.IsSet(&ytapi.FlagRegion) {
		call = call.RegionCode(values.GetString(&ytapi.FlagRegion))
	}

	// Execute
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) != 1 {
		return errors.New("Not Found")
	}

	return table.Append(response.Items)
}

func DeleteVideo(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)

	// create call and set parameters
	call := service.API.Videos.Delete(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// execute and return error, if any
	return call.Do()
}

func UpdateVideoMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)

	// Create call and set parameters
	call := service.API.Videos.List("id,snippet").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Execute
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) != 1 {
		return errors.New("Not Found")
	}

	// set metadata
	if values.IsSet(&ytapi.FlagTitle) {
		response.Items[0].Snippet.Title = values.GetString(&ytapi.FlagTitle)
	}
	if values.IsSet(&ytapi.FlagDescription) {
		response.Items[0].Snippet.Description = values.GetString(&ytapi.FlagDescription)
	}
	if values.IsSet(&ytapi.FlagLanguage) {
		response.Items[0].Snippet.DefaultLanguage = values.GetString(&ytapi.FlagLanguage)
	}
	if values.IsSet(&ytapi.FlagVideoCategory) {
		response.Items[0].Snippet.CategoryId = values.GetString(&ytapi.FlagVideoCategory)
	}

	// update video
	call2 := service.API.Videos.Update("snippet", response.Items[0])
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}

	_, err = call2.Do()
	if err != nil {
		return err
	}

	// Success
	return GetVideoMetadata(service, values, table)
}

func SetVideoStatus(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)

	// Create call and set parameters
	call := service.API.Videos.List("id,status").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Execute
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) != 1 {
		return errors.New("Not Found")
	}

	// set metadata
	if values.IsSet(&ytapi.FlagPrivacyStatus) {
		response.Items[0].Status.PrivacyStatus = values.GetString(&ytapi.FlagPrivacyStatus)
	}
	if values.IsSet(&ytapi.FlagEmbeds) {
		response.Items[0].Status.Embeddable = values.GetBool(&ytapi.FlagEmbeds)
	}
	if values.IsSet(&ytapi.FlagLicense) {
		response.Items[0].Status.License = values.GetString(&ytapi.FlagLicense)
	}
	if values.IsSet(&ytapi.FlagStatsViewable) {
		response.Items[0].Status.PublicStatsViewable = values.GetBool(&ytapi.FlagStatsViewable)
	}

	// update video
	call2 := service.API.Videos.Update("status", &youtube.Video{
		Id:     response.Items[0].Id,
		Status: response.Items[0].Status,
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}

	_, err = call2.Do()
	if err != nil {
		return err
	}

	// Success
	return GetVideoMetadata(service, values, table)
}

func SetVideoPublishDate(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)

	// Create call and set parameters
	call := service.API.Videos.List("id,status").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Execute
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) != 1 {
		return errors.New("Not Found")
	}

	// update video
	call2 := service.API.Videos.Update("status", &youtube.Video{
		Id: response.Items[0].Id,
		Status: &youtube.VideoStatus{
			PrivacyStatus: "private",
			PublishAt:     values.GetTimeInISOFormat(&ytapi.FlagDate),
		},
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}

	_, err = call2.Do()
	if err != nil {
		return err
	}

	// Success
	return GetVideoMetadata(service, values, table)
}

func SetVideoRating(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	video := values.GetString(&ytapi.FlagVideo)
	rating := values.GetString(&ytapi.FlagVideoRating)

	// Create call and set parameters
	call := service.API.Videos.Rate(video, rating)

	// Execute
	err := call.Do()
	if err != nil {
		return err
	}

	// Success
	return GetVideoRating(service, values, table)
}

func GetVideoRating(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	video := values.GetString(&ytapi.FlagVideo)

	// Create call and set parameters
	call := service.API.Videos.GetRating(video)

	// Execute
	response, err := call.Do()
	if err != nil {
		return err
	}

	// Success
	return table.Append(response.Items)
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
		"default":  response.Items[0].Snippet.Thumbnails.Default,
		"high":     response.Items[0].Snippet.Thumbnails.High,
		"maxres":   response.Items[0].Snippet.Thumbnails.Maxres,
		"medium":   response.Items[0].Snippet.Thumbnails.Medium,
		"standard": response.Items[0].Snippet.Thumbnails.Standard,
	}
	for name, thumbnail := range thumbnails {
		if thumbnail != nil {
			table.Append([]Thumbnail{{name, thumbnail.Url, uint64(thumbnail.Width), uint64(thumbnail.Height)}})
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
		return err
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
		"default":  response.Items[0].Default,
		"high":     response.Items[0].High,
		"maxres":   response.Items[0].Maxres,
		"medium":   response.Items[0].Medium,
		"standard": response.Items[0].Standard,
	}
	for name, thumbnail := range thumbnails {
		if thumbnail != nil {
			table.Append([]Thumbnail{{name, thumbnail.Url, uint64(thumbnail.Width), uint64(thumbnail.Height)}})
		}
	}

	// Success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Localized Metadata

func GetLocalizedVideoMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)

	// create call and set parameters
	call := service.API.Videos.List("snippet,localizations").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Call
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("Video not found")
	}

	// What is the default language
	defaultLanguage := response.Items[0].Snippet.DefaultLanguage

	// Get localizations
	localizations := response.Items[0].Localizations
	for language, metadata := range localizations {
		table.Append([]Localization{{language, metadata.Title, metadata.Description, defaultLanguage == language}})
	}

	// success
	return nil
}

func UpdateLocalizedVideoMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)
	language := values.GetString(&ytapi.FlagLanguage)

	// create call and set parameters
	call := service.API.Videos.List("localizations").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Call
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("Video not found")
	}

	// Update channel localization settings
	metadata, ok := response.Items[0].Localizations[language]
	if ok == false {
		metadata = youtube.VideoLocalization{}
	}
	if values.IsSet(&ytapi.FlagTitle) {
		metadata.Title = values.GetString(&ytapi.FlagTitle)
	}
	if values.IsSet(&ytapi.FlagDescription) {
		metadata.Description = values.GetString(&ytapi.FlagDescription)
	}
	response.Items[0].Localizations[language] = metadata

	// update localization
	call2 := service.API.Videos.Update("localizations", &youtube.Video{
		Id:            response.Items[0].Id,
		Localizations: response.Items[0].Localizations,
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}
	_, err = call2.Do()
	if err != nil {
		return err
	}

	// success
	return GetLocalizedVideoMetadata(service, values, table)
}

func DeleteLocalizedVideoMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)
	language := values.GetString(&ytapi.FlagLanguage)

	// create call and set parameters
	call := service.API.Videos.List("snippet,localizations").Id(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Call
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("Video not found")
	}

	// Update video localization settings
	_, ok := response.Items[0].Localizations[language]
	if ok == false {
		return errors.New("Localized metadata for language does not exist")
	}
	delete(response.Items[0].Localizations, language)

	// Sanity check for deleting default language
	defaultLanguage := response.Items[0].Snippet.DefaultLanguage
	if defaultLanguage == language {
		return errors.New("You cannot delete the default language metadata")
	}

	// update localization
	call2 := service.API.Videos.Update("localizations", &youtube.Video{
		Id:            response.Items[0].Id,
		Localizations: response.Items[0].Localizations,
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}
	_, err = call2.Do()
	if err != nil {
		return err
	}

	// success
	return GetLocalizedVideoMetadata(service, values, table)

}
