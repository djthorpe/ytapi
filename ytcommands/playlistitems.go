/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"
	"errors"
	"fmt"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)


////////////////////////////////////////////////////////////////////////////////
// Register playlist item commands

func RegisterPlaylistItemCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListPlaylistItems",
			Description: "List playlist items for a playlist",
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults},
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist},
			Setup:       RegisterPlaylistItemFormat,
			Execute:     ListPlaylistItems,
		},
		ytapi.Command{
			Name:        "InsertVideoIntoPlaylist",
			Description: "Inserts a video into a playlist",
			Optional:    []*ytapi.Flag{&ytapi.FlagPlaylistPosition,&ytapi.FlagPlaylistNote},
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist,&ytapi.FlagVideo},
			Setup:       RegisterPlaylistItemFormat,
			Execute:     InsertVideoIntoPlaylist,
		},
		ytapi.Command{
			Name:        "DeleteVideoFromPlaylist",
			Description: "Deletes a video from a playlist",
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist,&ytapi.FlagVideo},
			Setup:       RegisterPlaylistItemFormat,
			Execute:     DeletePlaylistVideo,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register playlist output

func RegisterPlaylistItemFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "playlistitem", Path: "Id", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channelTitle",  Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "playlist", Path: "Snippet/PlaylistId", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "position", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "kind", Path: "Snippet/ResourceId/Kind", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "video", Path: "Snippet/ResourceId/VideoId", Type: ytapi.FLAG_VIDEO},
	})
	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "startAt", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "endAt", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "note", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "privacyStatus", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"position", "title", "description","video","privacyStatus"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Returns set of playlist items for channel

func ListPlaylistItems(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	playlist := values.GetString(&ytapi.FlagPlaylist)
	call := service.API.PlaylistItems.List(strings.Join(table.Parts(false), ",")).PlaylistId(playlist)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// Perform channels.list and return results
	return ytapi.DoPlaylistItemsList(call, table, int64(values.GetUint(&ytapi.FlagMaxResults)))
}

////////////////////////////////////////////////////////////////////////////////
// Inserts a video into a playlist

func InsertVideoIntoPlaylist(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	playlist := values.GetString(&ytapi.FlagPlaylist)
	video := values.GetString(&ytapi.FlagVideo)
	call := service.API.PlaylistItems.Insert("snippet,contentDetails",&youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlist,
			Position: values.GetInt(&ytapi.FlagPlaylistPosition),
			ResourceId: &youtube.ResourceId{
				Kind: "youtube#video",
				VideoId: video,
			},
			ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
				"Position": &ytapi.FlagPlaylistPosition,
			}),
		},
		ContentDetails: &youtube.PlaylistItemContentDetails{
			Note: values.GetString(&ytapi.FlagPlaylistNote),
		},
	})
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	_, err := call.Do()
	if err != nil {
		return err
	}
	// List playlist items
	call2 := service.API.PlaylistItems.List(strings.Join(table.Parts(false), ",")).PlaylistId(playlist)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	// Perform playlistitems.list and return results
	return ytapi.DoPlaylistItemsList(call2, table, 0)
}

////////////////////////////////////////////////////////////////////////////////
// Deletes a video from a playlist

func DeletePlaylistVideo(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	playlist := values.GetString(&ytapi.FlagPlaylist)
	video := values.GetString(&ytapi.FlagVideo)
	call := service.API.PlaylistItems.List("id").PlaylistId(playlist).VideoId(video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// iterate through, gathering the PlayListItemId
	var playlistItems []*youtube.PlaylistItem = make([]*youtube.PlaylistItem,0)
	var nextPageToken string
	for {
		response, err := call.PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		for _,playlistItem := range(response.Items) {
			playlistItems = append(playlistItems,playlistItem)
		}
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}
	if len(playlistItems) == 0 {
		return errors.New(fmt.Sprint("Video ",video," not found in playlist ",playlist))
	}

	// delete the items from the playlist
	for _,playlistItem := range(playlistItems) {
		call := service.API.PlaylistItems.Delete(playlistItem.Id)
		err := call.Do()
		if err != nil {
			return err
		}
	}

	// List playlist items
	call2 := service.API.PlaylistItems.List(strings.Join(table.Parts(false), ",")).PlaylistId(playlist)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	// Perform playlistitems.list and return results
	return ytapi.DoPlaylistItemsList(call2, table, 0)
}



