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
)

////////////////////////////////////////////////////////////////////////////////

var (
	operations = map[string]int{
		"videos":     0, // --channel=<id> --maxresults=<n>
		"channels":   1, // --channel=<id> --maxresults=<n>
		"broadcasts": 2, // --channel=<id> --maxresults=<n> --status=<active|all|completed|upcoming>
		"streams":    3, // --channel=<id> --maxresults=<n>
		"bind":       4, // --video=<id> --stream=<key>
		"unbind":     5, // --video=<id>
	}
)

var (
    credentialsFolder = flag.String("credentials", ".credentials", "Folder containing credentials")
)

const (
    credentialsPathMode = 0700
)

////////////////////////////////////////////////////////////////////////////////

func userDir() (userDir string) {
    currentUser, _ := user.Current()
    userDir = currentUser.HomeDir
    return
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

    // Obtain path for credentials
    credentialsPath := filepath.Join(userDir(), *credentialsFolder)
    if credentialsPathInfo, err := os.Stat(credentialsPath); err != nil || !credentialsPathInfo.IsDir() {
        // if path is missing, try and create the folder
        if err := os.Mkdir(credentialsPath, credentialsPathMode); err != nil {
            fmt.Fprintf(os.Stderr,"Missing credentials folder: %v\n", credentialsPath)
            os.Exit(1)
        }
    }
}
