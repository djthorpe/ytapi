/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"errors"
	"strings"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register channel section commands

func RegisterChannelSectionCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListChannelSections",
			Description: "List channel sections",
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage},
			Setup:       RegisterChannelSectionFormat,
			Execute:     ListChannelSections,
		},
		&ytapi.Command{
			Name:        "NewChannelSection",
			Description: "Create a new channel section",
			Required:    []*ytapi.Flag{&ytapi.FlagSectionType},
			Optional: []*ytapi.Flag{
				&ytapi.FlagSectionStyle, &ytapi.FlagLanguage, &ytapi.FlagTitle,
				&ytapi.FlagSectionPosition, &ytapi.FlagPlaylist,
			},
			Setup:   RegisterChannelSectionFormat,
			Execute: NewChannelSection,
		},
		&ytapi.Command{
			Name:        "UpdateChannelSection",
			Description: "Update channel section",
			Required: []*ytapi.Flag{
				&ytapi.FlagSectionPosition,
			},
			Optional: []*ytapi.Flag{
				&ytapi.FlagSectionStyle, &ytapi.FlagLanguage, &ytapi.FlagTitle,
			},
			Setup:   RegisterChannelSectionFormat,
			Execute: UpdateChannelSection,
		},
		&ytapi.Command{
			Name:        "DeleteChannelSection",
			Description: "Delete a channel section",
			Required:    []*ytapi.Flag{&ytapi.FlagSectionPosition},
			Setup:       RegisterChannelSectionFormat,
			Execute:     DeleteChannelSection,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register channel section output format

func RegisterChannelSectionFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "section", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "type", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "style", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "position", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "language", Path: "Snippet/DefaultLanguage", Type: ytapi.FLAG_LANGUAGE},
	})

	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "playlists", Type: ytapi.FLAG_STRING, Array: true},
		&ytapi.Flag{Name: "channels", Type: ytapi.FLAG_STRING, Array: true},
	})

	// set default columns
	table.SetColumns([]string{"position", "title", "style", "language", "playlists", "channels"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func sectionFromPosition(service *ytservice.Service, values *ytapi.Values) (string, error) {

	// obtain position parameter
	position := values.GetUint(&ytapi.FlagSectionPosition)

	// create call and set parameters
	call := service.API.ChannelSections.List("snippet")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	if values.IsSet(&ytapi.FlagChannel) {
		call = call.ChannelId(values.GetString(&ytapi.FlagChannel))
	} else {
		call = call.Mine(true)
	}

	// obtain the sections
	response, err := call.Do()
	if err != nil {
		return "", err
	}

	for _, resource := range response.Items {
		if int64(position) == *resource.Snippet.Position {
			return resource.Id, nil
		}
	}
	return "", errors.New("Channel Section not found")
}

////////////////////////////////////////////////////////////////////////////////
// Channel Sections List

func ListChannelSections(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	parts := strings.Join(table.Parts(false), ",")

	// create call and set parameters
	call := service.API.ChannelSections.List(parts)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagChannel) {
		call = call.ChannelId(values.GetString(&ytapi.FlagChannel))
	} else {
		call = call.Mine(true)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	// Fudge the title column where the type isn't multipleChannels or multiplePlaylists
	for _, resource := range response.Items {
		switch resource.Snippet.Type {
		case "allPlaylists":
			resource.Snippet.Title = "Playlists by [Channel Name]"
		case "singlePlaylist":
			resource.Snippet.Title = "[Single Playlist]"
		case "likes":
			resource.Snippet.Title = "Liked videos"
		case "liveEvents":
			resource.Snippet.Title = "Live Events"
		case "completedEvents":
			resource.Snippet.Title = "Completed Events"
		case "upcomingEvents":
			resource.Snippet.Title = "Upcoming Events"
		case "likedPlaylists":
			resource.Snippet.Title = "Saved playlists"
		case "multipleChannels":
			resource.Snippet.Title = resource.Snippet.Title
		case "multiplePlaylists":
			resource.Snippet.Title = resource.Snippet.Title
		case "popularUploads":
			resource.Snippet.Title = "Popular uploads"
		case "postedPlaylists":
			resource.Snippet.Title = "Posted playlists"
		case "postedVideos":
			resource.Snippet.Title = "Posted videos"
		case "recentActivity":
			resource.Snippet.Title = "Recent activities"
		case "recentPosts":
			resource.Snippet.Title = "Recent posts"
		case "recentUploads":
			resource.Snippet.Title = "Uploads"
		case "subscriptions":
			resource.Snippet.Title = "Subscriptions"
		default:
			resource.Snippet.Title = resource.Snippet.Type
		}
	}

	if err = table.Append(response.Items); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Channel Sections List

func NewChannelSection(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Set parameters
	sectionType := values.GetString(&ytapi.FlagSectionType)
	part := "snippet"

	// Create the body
	body := &youtube.ChannelSection{
		Snippet: &youtube.ChannelSectionSnippet{
			Type:  sectionType,
			Style: values.GetString(&ytapi.FlagSectionStyle),
		},
	}

	// Title, Position and Language
	if values.IsSet(&ytapi.FlagLanguage) {
		body.Snippet.DefaultLanguage = values.GetString(&ytapi.FlagLanguage)
	}
	if values.IsSet(&ytapi.FlagTitle) {
		if sectionType != "multiplePlaylists" && sectionType != "multipleChannels" {
			return errors.New("Title can only be set when type is multiplePlaylists or multipleChannels")
		}
		body.Snippet.Title = values.GetString(&ytapi.FlagTitle)
	}
	if values.IsSet(&ytapi.FlagSectionPosition) {
		position := values.GetInt(&ytapi.FlagSectionPosition)
		body.Snippet.Position = &position
	}

	// Single Playlist
	if sectionType == "singlePlaylist" {
		if values.IsSet(&ytapi.FlagPlaylist) == false {
			return errors.New("Required flag: playlist")
		}
		part = "snippet,contentDetails"
		body.ContentDetails = &youtube.ChannelSectionContentDetails{
			Playlists: []string{
				values.GetString(&ytapi.FlagPlaylist),
			},
		}
	}

	// create call and set parameters
	call := service.API.ChannelSections.Insert(part, body)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		if values.IsSet(&ytapi.FlagChannel) {
			call = call.OnBehalfOfContentOwnerChannel(values.GetString(&ytapi.FlagChannel))
		}
	}

	// create the call
	_, err := call.Do()
	if err != nil {
		return err
	}

	// success
	return ListChannelSections(service, values, table)
}

////////////////////////////////////////////////////////////////////////////////
// Delete Channel Section

func DeleteChannelSection(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// obtain Channel Section ID from position
	section, err := sectionFromPosition(service, values)
	if err != nil {
		return err
	}

	// Perform Delete
	call := service.API.ChannelSections.Delete(section)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	err = call.Do()
	if err != nil {
		return err
	}

	// success
	return ListChannelSections(service, values, table)
}

////////////////////////////////////////////////////////////////////////////////
// Update Channel Section

func UpdateChannelSection(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// obtain Channel Section ID from position
	_, err := sectionFromPosition(service, values)
	if err != nil {
		return err
	}

	// TODO

	// success
	return ListChannelSections(service, values, table)
}
