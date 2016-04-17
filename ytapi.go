/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/djthorpe/ytapi/ytcommands"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

var (
	operations = map[string]func(*ytservice.Service, *ytservice.Params, *ytservice.Table) error {
		"Authenticate":   Authenticate,
		"ListVideos":     ytcommands.ListVideos,     // --channel=<id> --maxresults=<n>
		"ListChannels":   ytcommands.ListChannels,   // --channel=<id> --maxresults=<n>
		"ListPlaylists":  ytcommands.ListPlaylists,  // --channel=<id> --maxresults=<n>
		"Search":         ytcommands.Search,         // --channel=<id> --maxresults=<n>
	}
)

var (
	credentialsFolder      = flag.String("credentials", ".credentials", "Folder containing credentials")
	clientsecretFilename   = flag.String("clientsecret", "client_secret.json", "Client secret filename")
	serviceAccountFilename = flag.String("serviceaccount", "service_account.json", "Service account filename")
	defaultsFilename       = flag.String("defaults", "defaults.json", "Defaults filename")
	tokenFilename          = flag.String("authtoken", "oauth_token", "OAuth token filename")

	flagDebug        = flag.Bool("debug", false, "Debug flag")
	flagChannel      = flag.String("channel", "", "Channel ID")
	flagContentOwner = flag.String("contentowner", "", "Content Owner ID")
	flagMaxResults   = flag.Uint64("maxresults", 0, "Maximum results to return (or 0)")
)

const (
	credentialsPathMode = 0700
	crdentialsFileMode  = 0644
)

////////////////////////////////////////////////////////////////////////////////

func userDir() (userDir string) {
	currentUser, _ := user.Current()
	userDir = currentUser.HomeDir
	return
}

func NewParamsFromFile(filename string) (*ytservice.Params, error) {
	var params *ytservice.Params
	var err error

	// if a file exists, then read it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		params = ytservice.NewParams()
	} else {
		params, err = ytservice.NewParamsFromJSON(filename)
	}
	if err != nil {
		return nil, err
	}
	return params,nil
}

func CombineParamsWthFlags(params *ytservice.Params) *ytservice.Params {
	copy := params.Copy()

	// set up parameters
	if len(*flagContentOwner) > 0 {
		copy.ContentOwner = flagContentOwner
	}
	if len(*flagChannel) > 0 {
		copy.Channel = flagChannel
	}
	if *flagMaxResults > 0 {
		copy.MaxResults = *flagMaxResults
	}

	return copy
}

func Authenticate(service *ytservice.Service, params *ytservice.Params,output *ytservice.Table) error {

	// output content owner and channel information
	output.AppendColumn("contentowner","contentowner")
	output.AppendColumn("channel","channel")
	row := output.NewRow()

	if params.IsEmptyContentOwner() == false {
		row.SetString("contentowner",*params.ContentOwner)
	}
	if params.IsEmptyChannel() == false {
		row.SetString("channel",*params.Channel)
	}

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	// enumerate operations
	operation_keys := make([]string, 0, len(operations))
	for key := range operations {
		operation_keys = append(operation_keys, key)
	}

	// Set usage function
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "\n\t%s <flags> <%v>\n\n", filepath.Base(os.Args[0]), strings.Join(operation_keys, "|"))
		fmt.Fprintf(os.Stderr, "Where <flags> are one or more of:\n\n")
		flag.PrintDefaults()
	}

	// Read flags, exit with no operation
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Check operation
	opname := flag.Arg(0)
	_, ok := operations[opname]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Invalid operation: %s\n", opname)
		os.Exit(1)
	}

	// Obtain path for credentials - create if not already made
	credentialsPath := filepath.Join(userDir(), *credentialsFolder)
	if credentialsPathInfo, err := os.Stat(credentialsPath); err != nil || !credentialsPathInfo.IsDir() {
		// if path is missing, try and create the folder
		if err := os.Mkdir(credentialsPath, credentialsPathMode); err != nil {
			fmt.Fprintf(os.Stderr, "Missing credentials folder: %v\n", credentialsPath)
			os.Exit(1)
		}
	}

	var api *ytservice.Service
	var err error

	// Create params object
	defaultsPath := filepath.Join(credentialsPath, *defaultsFilename)
	defaults, err := NewParamsFromFile(defaultsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// if operation is to authenticate, delete existing token
	tokenPath := filepath.Join(credentialsPath, *tokenFilename)
	if opname == "Authenticate" {
		// Delete OAuth token
		if _, err := os.Stat(tokenPath); os.IsNotExist(err) == false {
			err = os.Remove(tokenPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// Authenticate
	serviceAccountPath := filepath.Join(credentialsPath, *serviceAccountFilename)
	clientSecretPath := filepath.Join(credentialsPath, *clientsecretFilename)
	if defaults.IsEmptyContentOwner() == false {
		api, err = ytservice.NewYouTubeServiceFromServiceAccountJSON(serviceAccountPath, defaults, *flagDebug)
	} else {
		api, err = ytservice.NewYouTubeServiceFromClientSecretsJSON(clientSecretPath, tokenPath, defaults, *flagDebug)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// If this isn't a service account and no channel parameter is set, then
	// set the channel parameter
	if api.ServiceAccount == false && defaults.IsValidChannel() == false {
		call := api.API.Channels.List("id").Mine(true)
		response, err := call.MaxResults(1).Do()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if len(response.Items) > 0 {
			defaults.Channel = &response.Items[0].Id
		}
	}

	// Save Defaults
	defaults.Save(defaultsPath,crdentialsFileMode)

	// Combine defaults with command-line flags to make parameters
	params := CombineParamsWthFlags(defaults)

	// Print out parameters
	if *flagDebug {
		fmt.Fprintf(os.Stderr,"parameters=%+v\n",params)
	}

	// Create a table object
	output := ytservice.NewTable([]string{})

	// Run command
	err = operations[opname](api, params, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	err = output.CSV(os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

