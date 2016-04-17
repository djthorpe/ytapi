/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"errors"
	"os/user"
	"path/filepath"

	"github.com/djthorpe/ytapi/ytcommands"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

type Operation struct {
	setup func(*ytservice.Params, *ytservice.Table) error
	do    func(*ytservice.Service, *ytservice.Params, *ytservice.Table) error
}

////////////////////////////////////////////////////////////////////////////////

var (
	operations = map[string]Operation {
		"Authenticate":    Operation{ NoOp, Authenticate},
		"ListVideos":      Operation{ NoOp, ytcommands.ListVideos },    // --channel=<id> --maxresults=<n>
		"ListChannels":    Operation{ ytcommands.RegisterChannelFormat, ytcommands.ListChannels },  // --channel=<id> --maxresults=<n>
		"ListPlaylists":   Operation{ ytcommands.RegisterPlaylistFormat, ytcommands.ListPlaylists }, // --channel=<id> --maxresults=<n>
		"ListBroadcasts":  Operation{ ytcommands.RegisterBroadcastFormat, ytcommands.ListBroadcasts }, // --channel=<id> --maxresults=<n>
		"DeleteBroadcast": Operation{ NoOp, ytcommands.DeleteBroadcast }, // --video=<id>
		"ListStreams":     Operation{ ytcommands.RegisterStreamFormat, ytcommands.ListStreams }, // --channel=<id> --maxresults=<n>
		"Search":          Operation{ ytcommands.RegisterSearchFormat, ytcommands.Search },        // --q=<string> --maxresults=<n>
	}
)

var (
	credentialsFolder      = flag.String("credentials", ".credentials", "Folder containing credentials")
	clientsecretFilename   = flag.String("clientsecret", "client_secret.json", "Client secret filename")
	serviceAccountFilename = flag.String("serviceaccount", "service_account.json", "Service account filename")
	defaultsFilename       = flag.String("defaults", "defaults.json", "Defaults filename")
	tokenFilename          = flag.String("authtoken", "oauth_token", "OAuth token filename")

	flagDebug        = flag.Bool("debug", false, "Show API requests and responses on stderr")
	flagChannel      = flag.String("channel", "", "Channel ID")
	flagContentOwner = flag.String("contentowner", "", "Content Owner ID")
	flagMaxResults   = flag.Int64("maxresults", 0, "Maximum results to return (or 0)")
	flagPart         = flag.String("part","","Comma-separated list of parts for response")
	flagOutput       = flag.String("output","ascii","Output type (csv or ascii)")
	flagQuery        = flag.String("q","","Search Query")
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

func CombineParamsWthFlags(params *ytservice.Params) (*ytservice.Params,error) {
	copy := params.Copy()

	// copy --contentowner
	if len(*flagContentOwner) > 0 {
		copy.ContentOwner = flagContentOwner
		if copy.IsValidContentOwner() == false {
			return nil,errors.New("Invalid --contentowner flag")
		}
	}

	// copy --channel
	if len(*flagChannel) > 0 {
		copy.Channel = flagChannel
		if copy.IsValidChannel() == false {
			return nil,errors.New("Invalid --channel flag")
		}
	}

	// copy --maxresults
	if *flagMaxResults < 0 {
		return nil,errors.New("Invalid --maxresults flag")
	}
	if *flagMaxResults > 0 {
		copy.MaxResults = *flagMaxResults
	}

	// copy --q
	if len(*flagQuery) > 0 {
		copy.Query = flagQuery
	}

	return copy,nil
}

func Authenticate(service *ytservice.Service, params *ytservice.Params,output *ytservice.Table) error {

	// output content owner and channel information
	output.AddColumn("contentowner")
	output.AddColumn("channel")
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

func NoOp(params *ytservice.Params,table *ytservice.Table) error {
	// Do Nothing
	return nil
}

func AlterColumnsForParts(parts []string,table *ytservice.Table) error {
	for _,part := range(parts) {
		if strings.HasPrefix(part,"-") {
			if err := table.RemoveColumnsForPart(strings.TrimPrefix(part,"-")); err != nil {
				return err
			}
		} else if strings.HasPrefix(part,"+") {
			if err := table.AddColumnsForPart(strings.TrimPrefix(part,"+")); err != nil {
				return err
			}
		} else {
			return errors.New("part elements must be prefixed with + or -")
		}
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
	params, err := CombineParamsWthFlags(defaults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Print out parameters
	if *flagDebug {
		fmt.Fprintf(os.Stderr,"parameters=%+v\n",params)
	}

	// Create a table object
	output := ytservice.NewTable()

	// Setup
	err = operations[opname].setup(params, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Add and remove parts
	if len(*flagPart) > 0 {
		err = AlterColumnsForParts(strings.Split(*flagPart,","),output)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Execute
	err = operations[opname].do(api, params, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *flagOutput == "ascii" {
		err = output.ASCII(os.Stdout)
	} else if *flagOutput == "csv" {
		err = output.CSV(os.Stdout)
	} else {
		err = errors.New("Invalid --output flag")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

