package ytcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"strings"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register search commands

func RegisterSearchCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "SearchVideos",
			Description: "Search for videos based on text query or related video",
			Optional:    []*ytapi.Flag{&ytapi.FlagSearchQuery, &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults, &ytapi.FlagSearchOrder, &ytapi.FlagSearchVideo, &ytapi.FlagSearchSafe},
			Setup:       RegisterVideoSearchFormat,
			Execute:     VideoSearch,
		},
		&ytapi.Command{
			Name:        "SearchBroadcasts",
			Description: "Search for live broadcasts",
			Optional:    []*ytapi.Flag{&ytapi.FlagSearchQuery, &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults, &ytapi.FlagSearchOrder, &ytapi.FlagSearchSafe, &ytapi.FlagSearchBroadcastStatus},
			Setup:       RegisterBroadcastSearchFormat,
			Execute:     BroadcastSearch,
		},
		&ytapi.Command{
			Name:        "SearchPlaylists",
			Description: "Search for playlists based on text query",
			Optional:    []*ytapi.Flag{&ytapi.FlagSearchQuery, &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults, &ytapi.FlagSearchOrder, &ytapi.FlagSearchSafe},
			Setup:       RegisterPlaylistSearchFormat,
			Execute:     PlaylistSearch,
		},
		&ytapi.Command{
			Name:        "SearchChannels",
			Description: "Search for channels",
			Optional:    []*ytapi.Flag{&ytapi.FlagSearchQuery, &ytapi.FlagLanguage, &ytapi.FlagRegion, &ytapi.FlagMaxResults, &ytapi.FlagSearchChannelOrder, &ytapi.FlagSearchSafe},
			Setup:       RegisterChannelSearchFormat,
			Execute:     ChannelSearch,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterVideoSearchFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "video", Path: "Id/VideoId", Type: ytapi.FLAG_VIDEO},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "liveBroadcastContent", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channelTitle", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"video", "title", "description", "channelTitle"})

	// success
	return nil
}

func RegisterPlaylistSearchFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "playlist", Path: "Id/PlaylistId", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "channelTitle", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"playlist", "title", "description", "channelTitle"})

	// success
	return nil
}

func RegisterChannelSearchFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "channel", Path: "Id/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "publishedAt", Type: ytapi.FLAG_TIME},
	})

	// set default columns
	table.SetColumns([]string{"channel", "title", "description"})

	// success
	return nil
}

func RegisterBroadcastSearchFormat(values *ytapi.Values, table *ytapi.Table) error {
	err := RegisterVideoSearchFormat(values, table)
	if err != nil {
		return nil
	}

	// set default columns
	table.SetColumns([]string{"video", "title", "description", "channelTitle", "liveBroadcastContent"})

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Video Search

func VideoSearch(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	call := service.API.Search.List(strings.Join(table.Parts(false), ","))
	call = call.Q(values.GetString(&ytapi.FlagSearchQuery))
	call = call.Type("video")

	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.RelevanceLanguage(values.GetString(&ytapi.FlagLanguage))
	}
	if values.IsSet(&ytapi.FlagRegion) {
		call = call.RegionCode(values.GetString(&ytapi.FlagRegion))
	}
	if values.IsSet(&ytapi.FlagSearchOrder) {
		call = call.Order(values.GetString(&ytapi.FlagSearchOrder))
	}
	if values.IsSet(&ytapi.FlagSearchSafe) {
		call = call.SafeSearch(values.GetString(&ytapi.FlagSearchSafe))
	}
	if values.IsSet(&ytapi.FlagSearchVideo) {
		call = call.RelatedToVideoId(values.GetString(&ytapi.FlagSearchVideo))
	}

	// Perform search, and return results
	return ytapi.DoSearchList(call, table, values.GetInt(&ytapi.FlagMaxResults))
}

////////////////////////////////////////////////////////////////////////////////
// Broadcast Search

func BroadcastSearch(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	call := service.API.Search.List(strings.Join(table.Parts(false), ","))
	call = call.Q(values.GetString(&ytapi.FlagSearchQuery))
	call = call.Type("video")

	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.RelevanceLanguage(values.GetString(&ytapi.FlagLanguage))
	}
	if values.IsSet(&ytapi.FlagRegion) {
		call = call.RegionCode(values.GetString(&ytapi.FlagRegion))
	}
	if values.IsSet(&ytapi.FlagSearchOrder) {
		call = call.Order(values.GetString(&ytapi.FlagSearchOrder))
	}
	if values.IsSet(&ytapi.FlagSearchSafe) {
		call = call.SafeSearch(values.GetString(&ytapi.FlagSearchSafe))
	}
	if values.IsSet(&ytapi.FlagSearchBroadcastStatus) {
		call = call.EventType(values.GetString(&ytapi.FlagSearchBroadcastStatus))
	} else {
		call = call.EventType("live")
	}

	// Perform search, and return results
	return ytapi.DoSearchList(call, table, values.GetInt(&ytapi.FlagMaxResults))
}

////////////////////////////////////////////////////////////////////////////////
// Playlist Search

func PlaylistSearch(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	call := service.API.Search.List(strings.Join(table.Parts(false), ","))
	call = call.Q(values.GetString(&ytapi.FlagSearchQuery))
	call = call.Type("playlist")

	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.RelevanceLanguage(values.GetString(&ytapi.FlagLanguage))
	}
	if values.IsSet(&ytapi.FlagRegion) {
		call = call.RegionCode(values.GetString(&ytapi.FlagRegion))
	}
	if values.IsSet(&ytapi.FlagSearchOrder) {
		call = call.Order(values.GetString(&ytapi.FlagSearchOrder))
	}
	if values.IsSet(&ytapi.FlagSearchSafe) {
		call = call.SafeSearch(values.GetString(&ytapi.FlagSearchSafe))
	}

	// Perform search, and return results
	return ytapi.DoSearchList(call, table, values.GetInt(&ytapi.FlagMaxResults))
}

////////////////////////////////////////////////////////////////////////////////
// Channel Search

func ChannelSearch(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	call := service.API.Search.List(strings.Join(table.Parts(false), ","))
	call = call.Q(values.GetString(&ytapi.FlagSearchQuery))
	call = call.Type("channel")

	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.RelevanceLanguage(values.GetString(&ytapi.FlagLanguage))
	}
	if values.IsSet(&ytapi.FlagRegion) {
		call = call.RegionCode(values.GetString(&ytapi.FlagRegion))
	}
	if values.IsSet(&ytapi.FlagSearchChannelOrder) {
		call = call.Order(values.GetString(&ytapi.FlagSearchChannelOrder))
	}
	if values.IsSet(&ytapi.FlagSearchSafe) {
		call = call.SafeSearch(values.GetString(&ytapi.FlagSearchSafe))
	}

	// Perform search, and return results
	return ytapi.DoSearchList(call, table, values.GetInt(&ytapi.FlagMaxResults))
}
