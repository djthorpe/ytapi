/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"errors"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytcommands"
	"github.com/djthorpe/ytapi/cidcommands"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	// Parse command-line flags
	command, values, err := ytapi.ParseFlags([]*ytapi.RegisterFunction{
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterAuthenticateCommands, Title: "Authentication opetations" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterChannelCommands, Title: "Operations on Channels" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterChannelSectionCommands, Title: "Channel Section operations" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterVideoCommands, Title: "Operations on videos" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterBroadcastCommands, Title: "Operations on Broadcasts" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterStreamCommands, Title: "Operations on Streams" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterCaptionCommands, Title: "Operations on video captions" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterPlaylistCommands, Title: "Operations on Playlists" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterPlaylistItemCommands, Title: "Operations on PlaylistItems" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterLanguageRegionCommands, Title: "Language and Region operations" },
		&ytapi.RegisterFunction{ Callback: ytcommands.RegisterSearchCommands, Title: "Search operations" },
		&ytapi.RegisterFunction{ Callback: cidcommands.RegisterContentOwnerCommands, Title: "Content owner operations" },
		&ytapi.RegisterFunction{ Callback: cidcommands.RegisterPolicyCommands, Title: "Policy operations" },
		&ytapi.RegisterFunction{ Callback: cidcommands.RegisterClaimCommands, Title: "Claim operations" },
	})
	if err != nil {
		// Error occured in command setup
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if command == nil {
		// No command is to be executed, so exit
		os.Exit(0)
	}

	// Read content owner and channel from file
	defaultsPath := ytcommands.GetDefaultsPath(values)
	if err := values.ReadDefaultsFromFile(defaultsPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// call command setup function
	output := ytapi.NewTable()
	if command.Setup != nil {
		if err := command.Setup(values, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	// add and remove fields from the table
	if values.IsSet(&ytapi.FlagFields) {
		fields := strings.Split(values.GetString(&ytapi.FlagFields),",")
		for _,field := range(fields) {
			var err error
			if strings.HasPrefix(field,"+") {
				err = output.AddColumn(field[1:])
			} else if strings.HasPrefix(field,"-") {
				err = output.RemoveColumn(field[1:])
			} else {
				err = errors.New(fmt.Sprint("Unknown field name or snippet: ",field))
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// create the service object
	serviceAccountPath := ytcommands.GetServiceAccountPath(values)
	clientSecretPath := ytcommands.GetClientSecretPath(values)
	tokenPath := ytcommands.GetOAuthTokenPath(values)
	debugFlag := values.GetBool(&ytapi.FlagDebug)

	// if the content owner is set, then create an API object from service
	// account, or else create the API object from client secrets and oauth
	// token
	var service *ytservice.Service
	if values.IsSet(&ytapi.FlagContentOwner) {
		service, err = ytservice.NewYouTubeServiceFromServiceAccountJSON(serviceAccountPath, debugFlag)
	} else {
		service, err = ytservice.NewYouTubeServiceFromClientSecretsJSON(clientSecretPath, tokenPath, debugFlag)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// call command execute function
	if command.Execute != nil {
		if err := command.Execute(service, values, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	// output
	if output.NumberOfColumns() > 0 {
		err = output.ASCII(os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if output.NumberOfRows() > 1 {
			output.Info(fmt.Sprintf("%v items returned",output.NumberOfRows()))
		}

		// TODO: display usage fields
		//ytapi.UsageFields(output)
	}

}

