/*
Copyright David Thorpe 2015 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"fmt"
	"strings"
	"errors"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterAuthenticateCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "Authenticate",
			Description: "Authenticate against service account or channel",
			Setup:       AuthenticateSetup,
			Execute:     AuthenticateExecute,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Retrieve my channel details

func retrieveChannelDetails(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	call := service.API.Channels.List(strings.Join(table.Parts(false), ","))
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	if values.IsSet(&ytapi.FlagChannel) {
		call = call.Id(values.GetString(&ytapi.FlagChannel))
	} else if(service.ServiceAccount) {
		call = call.ManagedByMe(true)
	} else {
		call = call.Mine(true)
	}

	// Perform search, and return results
	return ytapi.DoChannelsList(call,table,0)
}

////////////////////////////////////////////////////////////////////////////////
// Perform authentication

func AuthenticateSetup(values *ytapi.Values, table *ytapi.Table) error {

	// Disallow -channel parameter without -contentowner parameter
	if values.IsSet(&ytapi.FlagChannel) && values.IsSet(&ytapi.FlagContentOwner)==false {
		return errors.New("Cannot set -channel flag without -contentowner flag")
	}

	// set up output format
	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "channel", Path: "Id", Type: ytapi.FLAG_CHANNEL},
	})
	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "title", Path: "Snippet/Title", Type: ytapi.FLAG_STRING},
	})
	table.RegisterPart("contentOwnerDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "contentowner", Path: "ContentOwnerDetails/ContentOwner", Type: ytapi.FLAG_CONTENTOWNER},
	})
	table.SetColumns([]string{"channel", "title" })

	// success - but ask to remove the authentication token
	return ytapi.ErrorRemoveAuthToken
}

func AuthenticateExecute(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Display authentication settings
	if service.ServiceAccount {
		table.Info(fmt.Sprintf("Service Account: %s",service.ServiceAccountEmail))
		if values.IsSet(&ytapi.FlagContentOwner) {
			table.Info(fmt.Sprint("  Content Owner: ",values.GetString(&ytapi.FlagContentOwner)))
		}
		if values.IsSet(&ytapi.FlagChannel) {
			table.Info(fmt.Sprint("        Channel: ",values.GetString(&ytapi.FlagChannel)))
		}
	}

	// Get channel details
	if err := retrieveChannelDetails(service,values,table); err != nil {
		return err
	}

	// Write defaults
	return ytapi.ErrorWriteDefaults
}

