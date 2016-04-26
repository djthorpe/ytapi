/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"fmt"
	"os"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytcommands"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	// Parse command-line flags
	command, values, err := ytapi.ParseFlags([]ytapi.RegisterFunction{
		ytcommands.RegisterAuthenticateCommands,
		ytcommands.RegisterBroadcastCommands,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if command == nil {
		os.Exit(2)
	}

	// Read content owner and channel from file
	defaultsPath := ytcommands.GetDefaultsPath(values)
	if err := values.ReadDefaultsFromFile(defaultsPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// call command setup function
	output := ytservice.NewTable()
	if command.Setup != nil {
		if err := command.Setup(values, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
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
	err = output.ASCII(os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

/*

var (
	operations = map[string]Operation{
		"Authenticate":    Operation{NoOp, Authenticate},
		"ListVideos":      Operation{NoOp, ytcommands.ListVideos},                                   // --channel=<id> --maxresults=<n>
		"ListBroadcasts":  Operation{ytcommands.RegisterBroadcastFormat, ytcommands.ListBroadcasts}, // --channel=<id> --maxresults=<n>
		"DeleteBroadcast": Operation{NoOp, ytcommands.DeleteBroadcast},                              // --video=<id>
		"ListStreams":     Operation{ytcommands.RegisterStreamFormat, ytcommands.ListStreams},       // --channel=<id> --maxresults=<n>
		"DeleteStream":    Operation{NoOp, ytcommands.DeleteStream},                                 // --stream=<id>
		"Search":          Operation{ytcommands.RegisterSearchFormat, ytcommands.Search},            // --q=<string> --maxresults=<n>

		// Channels
		"ListChannels":                   Operation{ytcommands.RegisterChannelFormat, ytcommands.ListChannels}, // --channel=<id> --maxresults=<n>
		"ListLocalizedChannelMetadata":   Operation{ytcommands.RegisterLocalizedChannelMetadataFormat, ytcommands.ListLocalizedChannelMetadata},
		"UpdateChannelMetadata":          Operation{ytcommands.RegisterChannelFormat, ytcommands.UpdateChannelMetadata},                           // -hl=<string> -title=<string> -description=<string>
		"UpdateLocalizedChannelMetadata": Operation{ytcommands.RegisterLocalizedChannelMetadataFormat, ytcommands.UpdateLocalizedChannelMetadata}, // -hl=<string> -title=<string> -description=<string>

		// Playlists
		"ListPlaylists":                   Operation{ytcommands.RegisterPlaylistFormat, ytcommands.ListPlaylists}, // -maxresults=<n>
		"ListLocalizedPlaylistMetadata":   Operation{ytcommands.RegisterPlaylistFormat, ytcommands.ListPlaylists}, // -playlist=<id>
		"CreatePlaylist":                  Operation{ytcommands.RegisterPlaylistFormat, ytcommands.ListPlaylists}, // -hl=<string> -title=<string> -description=<string> -privacystatus=(public|private|unlisted) -tags=
		"DeletePlaylist":                  Operation{ytcommands.RegisterPlaylistFormat, ytcommands.ListPlaylists}, // -playlist=<id>
		"UpdatePlaylistMetadata":          Operation{ytcommands.RegisterPlaylistFormat, ytcommands.UpdatePlaylistMetadata}, // -playlist=<id> -hl=<string> --title=<string> --description=<string> --status=(public|private|unlisted)
		"UpdateLocalizedPlaylistMetadata": Operation{ytcommands.RegisterPlaylistFormat, ytcommands.ListPlaylists}, // --playlist=<id>  -hl=<string> --title=<string> --description=<string>

	}
)

*/
