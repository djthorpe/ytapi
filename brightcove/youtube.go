package brightcove

import (
	"reflect"
	"strings"
	// Frameworks
	"errors"

	"github.com/djthorpe/ytapi/brightcoveapi"
	"github.com/djthorpe/ytapi/util"
)

type YouTubeWebUpload struct {
	Filename          string `json:"filename`
	Channel           string `json:"channel`
	CustomId          string `json:"custom_id`
	Title             string `json:"title"`
	Description       string `json:"description"`
	Keywords          string `json:"keywords"`
	SpokenLanguage    string `json:"spoken_language"`
	CaptionFile       string `json:"caption_file"`
	CaptionLanguage   string `json:"caption_language"`
	Category          string `json:"category"`
	Privacy           string `json:"privacy"`
	NotifySubscribers string `json:"notify_subscribers"`
	StartTime         string `json:"start_time" ytapi:"datetime"`
	EndTime           string `json:"end_time" ytapi:"datetime"`
	CustomThumbnail   string `json:"custom_thumbnail"`
	PlaylistId        string `json:"playlist_id"`
}

////////////////////////////////////////////////////////////////////////////////
// Register YouTube commands

func RegisterYouTubeCommands() []*util.Command {
	return []*util.Command{
		&util.Command{
			Name:        "GetYouTubeWebUploadMetadata",
			Description: "Get video metadata from Brightcove in YouTube format",
			Format:      FormatYouTubeMetdata,
			Brightcove:  GetYouTubeMetdata,
		},
	}
}

/////////////////////////////////////////////////////////////////////

func FormatYouTubeMetdata(output *util.Table) error {
	if err := output.AddColumnsFrom(reflect.ValueOf(&YouTubeWebUpload{})); err != nil {
		return err
	}

	// Success
	return nil
}

func GetYouTubeMetdata(client *brightcoveapi.Client, output *util.Table, args []string) error {
	if len(args) != 0 {
		return errors.New("Too many arguments")
	}
	if videos, err := client.CMS.GetVideos(); err != nil {
		return err
	} else {
		for _, video := range videos {
			youtube := MapYouTubeMetdata(video)
			if err := output.Append([]*YouTubeWebUpload{youtube}); err != nil {
				return err
			}

		}
	}
	// Success
	return nil
}

func MapYouTubeMetdata(video *brightcoveapi.Video) *YouTubeWebUpload {
	yt := &YouTubeWebUpload{}
	if video == nil {
		return nil
	}
	// Convert filename, name
	yt.Filename = video.OriginalFilename
	yt.Title = video.Name
	yt.Description = video.Description
	yt.CustomId = video.Id
	yt.Category = video.Category
	yt.Keywords = strings.Join(video.Tags, "|")
	switch video.State {
	case "ACTIVE":
		yt.Privacy = "unlisted"
	default:
		yt.Privacy = "private"
	}
	return yt
}
