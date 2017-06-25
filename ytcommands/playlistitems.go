package ytcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
	"fmt"
	"strings"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

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
	call := service.API.PlaylistItems.Insert("snippet,contentDetails", &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlist,
			Position:   values.GetInt(&ytapi.FlagPlaylistPosition),
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
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
	_, err := call.Do(service.CallOptions()...)
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
	call := service.API.PlaylistItems.List("id,snippet").PlaylistId(playlist)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// iterate through, gathering the PlayListItemId
	var playlistItems []*youtube.PlaylistItem = make([]*youtube.PlaylistItem, 0)
	var nextPageToken string
	for {
		response, err := call.PageToken(nextPageToken).Do(service.CallOptions()...)
		if err != nil {
			return err
		}
		for _, playlistItem := range response.Items {
			resource := playlistItem.Snippet.ResourceId
			if resource != nil && resource.Kind == "youtube#video" && resource.VideoId == video {
				playlistItems = append(playlistItems, playlistItem)
			}
		}
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}
	if len(playlistItems) == 0 {
		return errors.New(fmt.Sprint("Video ", video, " not found in playlist ", playlist))
	}

	// delete the items from the playlist
	for _, playlistItem := range playlistItems {
		call := service.API.PlaylistItems.Delete(playlistItem.Id)
		if service.ServiceAccount {
			call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		}
		err := call.Do(service.CallOptions()...)
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
