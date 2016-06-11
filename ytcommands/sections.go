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
			Optional:    []*ytapi.Flag{&ytapi.FlagSectionStyle, &ytapi.FlagLanguage, &ytapi.FlagTitle, &ytapi.FlagSectionPosition},
			Setup:       RegisterChannelSectionFormat,
			Execute:     NewChannelSection,
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

	// set default columns
	table.SetColumns([]string{"position", "title", "type", "style"})

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
	if err = table.Append(response.Items); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Channel Sections List

func NewChannelSection(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Create the body
	body := &youtube.ChannelSection{
		Snippet: &youtube.ChannelSectionSnippet{
			Type:  values.GetString(&ytapi.FlagSectionType),
			Style: values.GetString(&ytapi.FlagSectionStyle),
		},
	}

	// create call and set parameters
	call := service.API.ChannelSections.Insert("snippet", body)
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
