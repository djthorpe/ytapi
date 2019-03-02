/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/djthorpe/ytapi/brightcoveapi"
)

////////////////////////////////////////////////////////////////////////////////

const (
	DEFAULTS_FILENAME = "brightcove.json"
)

var (
	FlagCredentials = flag.String("credentials", ".ytapi", "Folder containing credentials")
	FlagDebug       = flag.Bool("debug", false, "Show API requests and responses on stderr")
)

////////////////////////////////////////////////////////////////////////////////

func PathCredentials(folder string) (string, error) {
	if filepath.IsAbs(folder) == false {
		if user, err := user.Current(); err != nil {
			return "", err
		} else {
			folder = filepath.Join(user.HomeDir, folder)
		}
	}
	if stat, err := os.Stat(folder); os.IsNotExist(err) {
		return "", err
	} else if stat.IsDir() == false {
		return "", errors.New("Invalid credentials")
	} else {
		return filepath.Join(folder, DEFAULTS_FILENAME), nil
	}
}

func ClientOptions(debug bool) ([]brightcoveapi.ClientOption, error) {
	options := make([]brightcoveapi.ClientOption, 0, 5)
	options = append(options, brightcoveapi.WithDebug(debug))
	return options, nil
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()

	if credentials, err := PathCredentials(*FlagCredentials); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if options, err := ClientOptions(*FlagDebug); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if client, err := brightcoveapi.NewClient(credentials, options...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else {
		fmt.Println(client)
	}
}
