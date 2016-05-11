/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"fmt"
	"os"
//	"strings"
//	"errors"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytcommands"
	"github.com/djthorpe/ytapi/cidcommands"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
    flags := ytapi.NewFlagSet()

    // Register the commands allowed
    err := flags.RegisterCommands([]*ytapi.RegisterFunction{
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
        &ytapi.RegisterFunction{ Callback: cidcommands.RegisterAssetCommands, Title: "Asset operations" },
        &ytapi.RegisterFunction{ Callback: cidcommands.RegisterReferenceCommands, Title: "Reference operations" },
    })
    if err != nil {
        // Error occured in command setup
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    // Parse command-line flags
    command, err := flags.Parse()
    if err == ytapi.ErrorUsage {
        flags.Usage()
        flags.UsageGlobalFlags()
        if command != nil {
            flags.UsageCommand(command)
            flags.UsageFields()
        } else {
            flags.UsageCommandList()
        }
        os.Exit(0)
    } else if err == ytapi.ErrorClientSecrets {
        // Error occured during parsing of the flags
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    } else if err == ytapi.ErrorServiceAccount {
        // Error occured during parsing of the flags
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    } else if err != nil {
        // Error occured during parsing of the flags
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

	// Read content owner and channel from file if command is not "Authenticate"
	if command.Name != "Authenticate" {
		if err := flags.ReadDefaults(); err != nil {
			// ignore if defaults file doesn't yet exist
			if os.IsNotExist(err) == false {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}
	}

    // Create a service. If the content owner is set, then create an API object
    // from service account, or else create the API object from client secrets
    // and oauth token
    var service *ytservice.Service
    debugFlag := flags.Values.GetBool(&ytapi.FlagDebug)
    if flags.Values.IsSet(&ytapi.FlagContentOwner) {
        service, err = ytservice.NewYouTubeServiceFromServiceAccountJSON(flags.ServiceAccount, debugFlag)
    } else {
        service, err = ytservice.NewYouTubeServiceFromClientSecretsJSON(flags.ClientSecrets,flags.AuthToken, debugFlag)
    }
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

	// Execute command
	if err := flags.ExecuteCommand(command,service); err != nil {
		// Write defaults to file
		if err == ytapi.ErrorWriteDefaults {
			err = flags.WriteDefaults()
		}
		// Check for error
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

    // Display output
    err = flags.DisplayOutput()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

/*
func main2() {
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
        &ytapi.RegisterFunction{ Callback: cidcommands.RegisterAssetCommands, Title: "Asset operations" },
        &ytapi.RegisterFunction{ Callback: cidcommands.RegisterReferenceCommands, Title: "Reference operations" },
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
        // check for service account
        if (command.ServiceAccount==true) && (service.ServiceAccount==false) {
            fmt.Fprintf(os.Stderr, "Error: Requires a service account\n")
            os.Exit(1)
        }
		if err := command.Execute(service, values, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}
*/

