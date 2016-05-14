/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
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
// Channel Sections List

func ListChannelSections(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	parts := strings.Join(table.Parts(false), ",")

	// create call and set parameters
	call := service.API.ChannelSections.List(parts)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if channel != "" {
		call = call.ChannelId(channel)
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
