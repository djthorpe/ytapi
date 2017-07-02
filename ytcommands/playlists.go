package ytcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
	"fmt"
	"strings"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register playlist commands

func RegisterPlaylistCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListPlaylists",
			Description: "List playlists for channel",
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage, &ytapi.FlagMaxResults},
			Setup:       RegisterPlaylistFormat,
			Execute:     ListPlaylists,
		},
		&ytapi.Command{
			Name:        "NewPlaylist",
			Description: "Create a new playlist",
			Required:    []*ytapi.Flag{&ytapi.FlagTitle},
			Optional:    []*ytapi.Flag{&ytapi.FlagDescription, &ytapi.FlagPrivacyStatus, &ytapi.FlagLanguage},
			Setup:       RegisterPlaylistFormat,
			Execute:     InsertPlaylist,
		},
		&ytapi.Command{
			Name:        "DeletePlaylist",
			Description: "Delete an existing playlist",
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist},
			Execute:     DeletePlaylist,
		},
		&ytapi.Command{
			Name:        "UpdatePlaylist",
			Description: "Update playlist metadata",
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist},
			Optional:    []*ytapi.Flag{&ytapi.FlagTitle, &ytapi.FlagDescription, &ytapi.FlagPrivacyStatus, &ytapi.FlagLanguage},
			Setup:       RegisterPlaylistFormat,
			Execute:     UpdatePlaylist,
		},
		&ytapi.Command{
			Name:        "ListPlaylistItems",
			Description: "List playlist items for a playlist",
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults},
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist},
			Setup:       RegisterPlaylistItemFormat,
			Execute:     ListPlaylistItems,
		},
		&ytapi.Command{
			Name:        "InsertVideoIntoPlaylist",
			Description: "Inserts a video into a playlist",
			Optional:    []*ytapi.Flag{&ytapi.FlagPlaylistPosition, &ytapi.FlagPlaylistNote},
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist, &ytapi.FlagVideo},
			Setup:       RegisterPlaylistItemFormat,
			Execute:     InsertVideoIntoPlaylist,
		},
		&ytapi.Command{
			Name:        "DeleteVideoFromPlaylist",
			Description: "Deletes a video from a playlist",
			Required:    []*ytapi.Flag{&ytapi.FlagPlaylist, &ytapi.FlagVideo},
			Setup:       RegisterPlaylistItemFormat,
			Execute:     DeletePlaylistVideo,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register playlist output

func RegisterPlaylistFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "playlist", Path: "Id", Type: ytapi.FLAG_PLAYLIST},
	})
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "language", Path: "Snippet/DefaultLanguage", Type: ytapi.FLAG_LANGUAGE},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "channelTitle", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "tags", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "itemCount", Type: ytapi.FLAG_UINT},
	})
	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "privacyStatus", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("player", []*ytapi.Flag{
		&ytapi.Flag{Name: "embedHtml", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"playlist", "title", "description", "itemCount", "privacyStatus"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Register playlist items output

func RegisterPlaylistItemFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "playlistitem", Path: "Id", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channelTitle", Type: ytapi.FLAG_STRING},
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
	table.SetColumns([]string{"position", "title", "description", "video", "privacyStatus"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Returns set of playlists for channel

func ListPlaylists(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	call := service.API.Playlists.List(strings.Join(table.Parts(false), ","))
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	if language := values.GetString(&ytapi.FlagLanguage); language != "" {
		call = call.Hl(language)
	}
	if channel := values.GetString(&ytapi.FlagChannel); channel != "" {
		call = call.ChannelId(channel)
	} else {
		call = call.Mine(true)
	}

	// Perform channels.list and return results
	return ytapi.DoPlaylistsList(call, table, int64(values.GetUint(&ytapi.FlagMaxResults)), service.CallOptions()...)
}

////////////////////////////////////////////////////////////////////////////////
// Creates a new playlist

// TODO: Allow insert of tags

func InsertPlaylist(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Create call, set parameters
	call := service.API.Playlists.Insert("snippet,status", &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:           values.GetString(&ytapi.FlagTitle),
			Description:     values.GetString(&ytapi.FlagDescription),
			DefaultLanguage: values.GetString(&ytapi.FlagLanguage),
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: values.GetString(&ytapi.FlagPrivacyStatus),
		},
	})
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		call = call.OnBehalfOfContentOwnerChannel(values.GetString(&ytapi.FlagChannel))
	}

	// Insert broadcast and get a youtube.Playlist back
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Retrieve playlist
	call2 := service.API.Playlists.List(strings.Join(table.Parts(false), ",")).Id(response.Id)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// Perform channels.list and return results
	return ytapi.DoPlaylistsList(call2, table, 1, service.CallOptions()...)
}

////////////////////////////////////////////////////////////////////////////////
// Update playlist metadata

// TODO: Allow update of tags

func UpdatePlaylist(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	playlist := values.GetString(&ytapi.FlagPlaylist)
	call := service.API.Playlists.List("snippet,status").Id(playlist)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		call = call.OnBehalfOfContentOwnerChannel(values.GetString(&ytapi.FlagChannel))
	}
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}
	if len(response.Items) != 1 {
		return errors.New(fmt.Sprint("Playlist not found: ", playlist))
	}
	metadata := response.Items[0]
	if values.IsSet(&ytapi.FlagTitle) {
		metadata.Snippet.Title = values.GetString(&ytapi.FlagTitle)
	}
	if values.IsSet(&ytapi.FlagDescription) {
		metadata.Snippet.Title = values.GetString(&ytapi.FlagDescription)
	}
	if values.IsSet(&ytapi.FlagLanguage) {
		metadata.Snippet.DefaultLanguage = values.GetString(&ytapi.FlagLanguage)
	}
	if values.IsSet(&ytapi.FlagPrivacyStatus) {
		metadata.Status.PrivacyStatus = values.GetString(&ytapi.FlagPrivacyStatus)
	}
	// do update
	call2 := service.API.Playlists.Update("snippet,status", &youtube.Playlist{
		Id: metadata.Id,
		Snippet: &youtube.PlaylistSnippet{
			Title:           metadata.Snippet.Title,
			Description:     metadata.Snippet.Description,
			DefaultLanguage: metadata.Snippet.DefaultLanguage,
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: metadata.Status.PrivacyStatus,
		},
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	_, err = call2.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Retrieve playlist
	call3 := service.API.Playlists.List(strings.Join(table.Parts(false), ",")).Id(metadata.Id)
	if service.ServiceAccount {
		call3 = call3.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// Perform channels.list and return results
	return ytapi.DoPlaylistsList(call3, table, 1, service.CallOptions()...)
}

////////////////////////////////////////////////////////////////////////////////
// Delete an existing playlist

func DeletePlaylist(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	playlist := values.GetString(&ytapi.FlagPlaylist)

	// Create call, set parameters
	call := service.API.Playlists.Delete(playlist)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// Insert broadcast and get a youtube.Playlist back
	err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Success
	table.Info(fmt.Sprint("Deleted: ", playlist))
	return nil
}
