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

	// Frameworks
	"github.com/djthorpe/ytapi/brightcove"
	"github.com/djthorpe/ytapi/brightcoveapi"
	"github.com/djthorpe/ytapi/util"
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

func RegisterCommands() []*util.Command {
	var commands []*util.Command
	commands = append(commands, brightcove.RegisterCMSCommands()...)
	commands = append(commands, brightcove.RegisterYouTubeCommands()...)
	return commands
}

func PrintUsage(dev *os.File, commands []*util.Command) {
	execname := filepath.Base(os.Args[0])
	fmt.Fprintf(dev, "\nUsage of %s:\n\n", execname)
	fmt.Fprintf(dev, "  %s -help\n", execname)
	fmt.Fprintf(dev, "  %s <flags>... <api-call> <arguments>...\n", execname)
	fmt.Fprintln(dev, "")

	fmt.Fprintln(dev, "API Calls:")
	for _, command := range commands {
		fmt.Fprintf(dev, "  %-20s %s\n", command.Name, command.Description)
	}
	fmt.Fprintln(dev, "")
}

func GetCommand(arg string, commands []*util.Command) *util.Command {
	for _, command := range commands {
		if command.Name == arg {
			return command
		}
	}
	return nil
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
	} else if commands := RegisterCommands(); len(commands) == 0 {
		fmt.Fprintln(os.Stderr, "No commands registered")
		os.Exit(-1)
	} else if args := flag.Args(); len(args) == 0 {
		PrintUsage(os.Stderr, commands)
		os.Exit(-1)
	} else if command := GetCommand(args[0], commands); command == nil {
		PrintUsage(os.Stderr, commands)
		os.Exit(-1)
	} else if err := command.ExecBrightcove(client, args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
