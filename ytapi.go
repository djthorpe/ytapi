/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/djthorpe/ytapi/cidcommands"
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytcommands"
	"github.com/djthorpe/ytapi/ytservice"
	"github.com/djthorpe/ytapi/ytreporting"
)

////////////////////////////////////////////////////////////////////////////////

func Usage() {
	execname := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "\nUsage of %s:\n\n", execname)
	fmt.Fprintf(os.Stderr, "\t%s -help\n", execname)
	fmt.Fprintf(os.Stderr, "\t%s -help <api-call>\n", execname)
	fmt.Fprintf(os.Stderr, "\t%s <flags> <api-call>\n", execname)
	fmt.Fprintln(os.Stderr, "")
}

func UsageVersion() {
	execname := filepath.Base(os.Args[0])

	fmt.Fprintf(os.Stderr, "%s: Command Line Tool for YouTube API calls\n", execname)
	fmt.Fprintf(os.Stderr, "%s\n", VERSION_URL)
	fmt.Fprintf(os.Stderr, "%s\n\n", VERSION_COPYRIGHT)
	fmt.Fprintf(os.Stderr, "    Author: %s\n", VERSION_AUTHOR)
	if VERSION_TAG != "" {
		fmt.Fprintf(os.Stderr, "       Tag: %s\n", VERSION_TAG)
	}
	if VERSION_BRANCH != "" {
		fmt.Fprintf(os.Stderr, "    Branch: %s\n", VERSION_BRANCH)
	}
	if VERSION_HASH != "" {
		fmt.Fprintf(os.Stderr, "      Hash: %s\n", VERSION_HASH)
	}
	fmt.Fprintf(os.Stderr, "      Date: %s\n", VERSION_DATE)
	fmt.Fprintf(os.Stderr, "Go Version: %s\n", VERSION_GOVERSION)
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	flags := ytapi.NewFlagSet()

	// Register the commands allowed
	err := flags.RegisterCommands([]*ytapi.RegisterFunction{
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterAuthenticateCommands, Title: "Authentication and Installation"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterChannelCommands, Title: "Channels"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterChannelSectionCommands, Title: "Channel Sections"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterVideoCommands, Title: "Videos"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterBroadcastCommands, Title: "Broadcasts"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterStreamCommands, Title: "Streams"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterCuepointsCommands, Title: "Cuepoints"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterCaptionCommands, Title: "Video Caption Tracks"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterPlaylistCommands, Title: "Playlists"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterLanguageRegionCommands, Title: "Language and Regions"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterCommentsCommands, Title: "Comments"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterLiveChatCommands, Title: "Live Chat"},
		&ytapi.RegisterFunction{Callback: ytcommands.RegisterSearchCommands, Title: "Search"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterContentOwnerCommands, Title: "Content Owners"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterPolicyCommands, Title: "Content Owner Policies"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterClaimCommands, Title: "Claims"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterAssetCommands, Title: "Assets"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterReferenceCommands, Title: "Asset References"},
		&ytapi.RegisterFunction{Callback: cidcommands.RegisterValidatorCommands, Title: "Validate Metadata"},
		&ytapi.RegisterFunction{Callback: ytreporting.RegisterAnalyticsCommands, Title: "Reporting"},

	})
	if err != nil {
		// Error occured in command setup
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Parse command-line flags, set up output and determine paths for
	// the configuration
	command, err := flags.Parse()
	if err == ytapi.ErrorEmptyArgs {
		UsageVersion()
		Usage()
		os.Exit(0)
	} else if err == ytapi.ErrorUsage {
		Usage()
		if command != nil {
			flags.UsageGlobalFlags()
			flags.UsageCommand(command)
			flags.UsageFields()
		} else {
			flags.UsageCommandList()
		}
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Add and remove fields from output
	if flags.Values.IsSet(&ytapi.FlagFields) {
		err := flags.SetFields(strings.Split(flags.Values.GetString(&ytapi.FlagFields), ","))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	// install credentials
	if command.Name == "Install" {
		if VERSION_CLIENT_SECRET != "" {
			err = flags.WriteClientSecret(VERSION_CLIENT_SECRET)
		}
		if err == nil && VERSION_SERVICE_ACCOUNT != "" {
			err = flags.WriteServiceAccount(VERSION_SERVICE_ACCOUNT)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
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

	// Open output file
	if err := flags.OpenOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer flags.CloseOutput()

	// Create a service. If the content owner is set, then create an API object
	// from service account, or else create the API object from client secrets
	// and oauth token
	var service *ytservice.Service
	debugFlag := flags.Values.GetBool(&ytapi.FlagDebug)
	if command.Name == "Install" {
		// Empty service object
	} else if flags.Values.IsSet(&ytapi.FlagContentOwner) {
		service, err = ytservice.NewYouTubeServiceFromServiceAccountJSON(flags.ServiceAccount, debugFlag)
	} else {
		service, err = ytservice.NewYouTubeServiceFromClientSecretsJSON(flags.ClientSecrets, flags.AuthToken, debugFlag)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Execute command
	if err := flags.ExecuteCommand(command, service); err != nil {
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
	if err := flags.DisplayOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
