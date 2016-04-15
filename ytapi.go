/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
    "path/filepath"
    "os/user"

	"github.com/djthorpe/ytapi/ytservice"
	"github.com/djthorpe/ytapi/ytcommands"
)

////////////////////////////////////////////////////////////////////////////////

var (
	operations = map[string]func(*ytservice.YTService,*ytservice.Defaults)(error){
		"auth":       Authorization,
		"videos":     ytcommands.Channels, // --channel=<id> --maxresults=<n>
		"channels":   ytcommands.Channels, // --channel=<id> --maxresults=<n>
		"broadcasts": ytcommands.Channels, // --channel=<id> --maxresults=<n> --status=<active|all|completed|upcoming>
		"streams":    ytcommands.Channels, // --channel=<id> --maxresults=<n>
		"bind":       ytcommands.Channels, // --video=<id> --stream=<key>
		"unbind":     ytcommands.Channels, // --video=<id>
	}
)

var (
    credentialsFolder      = flag.String("credentials", ".credentials", "Folder containing credentials")
	debug                  = flag.Bool("debug", false, "Debug flag")
	clientsecretFilename   = flag.String("clientsecret", "client_secrets.json", "Client secret filename")
	serviceAccountFilename = flag.String("serviceaccount", "service_account.json", "Service account filename")
	defaultsFilename       = flag.String("defaults", "defaults.json", "Defaults filename")
	tokenFilename          = flag.String("authtoken", "oauth_token", "OAuth token filename")

	paramChannel          = flag.String("channel","","Channel ID")
	paramContentOwner     = flag.String("contentowner", "", "Content Owner ID")
	paramMaxResults       = flag.Uint64("max-results",0,"Maximum results to return (or 0)")
)

const (
    credentialsPathMode = 0700
	crdentialsFileMode = 0644
)

////////////////////////////////////////////////////////////////////////////////

func userDir() (userDir string) {
    currentUser, _ := user.Current()
    userDir = currentUser.HomeDir
    return
}

func NewDefaults(filename string) (*ytservice.Defaults,error) {
	var defaults *ytservice.Defaults
	var err error

	// if a file exists, then read it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		defaults = ytservice.NewDefaults()
	} else {
		defaults,err = ytservice.NewDefaultsFromJSON(filename)
	}
	if err != nil {
		return nil,err
	}

	// set up parameters
	if *debug {
		defaults.Debug = *debug
	}
	if len(*paramContentOwner) > 0 {
		defaults.ContentOwner = paramContentOwner
	}
	if len(*paramChannel) > 0 {
		defaults.Channel = paramChannel
	}
	if *paramMaxResults > 0 {
		defaults.MaxResults = *paramMaxResults
	}

	// return defaults
	return defaults,nil
}

func Authorization(service *ytservice.YTService,defaults *ytservice.Defaults) (error) {
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
            fmt.Fprintf(os.Stderr,"Missing credentials folder: %v\n", credentialsPath)
            os.Exit(1)
        }
    }

	var api *ytservice.YTService
	var err error

	// Create defaults object
	defaultsPath := filepath.Join(credentialsPath, *defaultsFilename)
	defaults,err := NewDefaults(defaultsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// iI operation is to authenticate, delete existing token and save credentials
	tokenPath := filepath.Join(credentialsPath, *tokenFilename)
	if opname == "auth" {
		// Save defaults
		err = defaults.Save(defaultsPath,crdentialsFileMode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		// Delete OAuth token
		if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
			// Do nothing
		} else {
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
	if len(*defaults.ContentOwner) > 0 {
		api, err = ytservice.NewYouTubeServiceFromServiceAccountJSON(serviceAccountPath,defaults)
	} else {
		api, err = ytservice.NewYouTubeServiceFromClientSecretsJSON(clientSecretPath,tokenPath,defaults)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Run command
	err = operations[opname](api,defaults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
