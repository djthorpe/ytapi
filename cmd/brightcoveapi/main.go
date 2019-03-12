/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package main

import (
	"errors"
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
	FLAG_CREDENTIALS  = "credentials"
	FLAG_DEBUG        = "debug"
	FLAG_TIMEOUT      = "timeout"
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

func ClientOptions(flagset *util.FlagSet) ([]brightcoveapi.ClientOption, error) {
	options := make([]brightcoveapi.ClientOption, 0, 5)
	options = append(options, brightcoveapi.WithDebug(flagset.GetBool(FLAG_DEBUG)))
	options = append(options, brightcoveapi.WithTimeout(flagset.GetDuration(FLAG_TIMEOUT)))
	return options, nil
}

func RegisterCommands() []*util.Command {
	var commands []*util.Command
	commands = append(commands, brightcove.RegisterCMSCommands()...)
	commands = append(commands, brightcove.RegisterYouTubeCommands()...)
	return commands
}

func PrintUsage(dev *os.File, commands []*util.Command, flags *util.FlagSet) {
	args := flags.Args()
	if len(args) == 1 {
		if command := GetCommand(args[0], commands); command != nil {
			PrintUsageCommand(dev, command, flags)
			return
		}
	}

	fmt.Fprintf(dev, "\nUsage of %s:\n\n", flags.Name())
	fmt.Fprintf(dev, "  %s -help\n", flags.Name())
	fmt.Fprintf(dev, "  %s -help <api-call>\n", flags.Name())
	fmt.Fprintf(dev, "  %s (<flag> <flag> ...) <api-call> (<arg> <arg>,...)\n", flags.Name())
	fmt.Fprintln(dev, "")

	fmt.Fprintln(dev, "API Calls:")
	for _, command := range commands {
		fmt.Fprintf(dev, "  %-30s %s\n", command.Name, command.Description)
	}
	fmt.Fprintln(dev, "")

	PrintAllFlags(dev, flags)
}

func PrintAllFlags(dev *os.File, flagset *util.FlagSet) {
	for _, scope := range []util.ScopeType{util.SCOPE_REQUIRED, util.SCOPE_OPTIONAL, util.SCOPE_GLOBAL} {
		if flags := flagset.FlagsForScope(scope); len(flags) > 0 {
			fmt.Fprintf(dev, "%v flags:\n", scope)
			for _, flag := range flags {
				var value string
				if flag.String() != "" {
					value = fmt.Sprintf(" (default \"%v\")", flag.String())
				}
				fmt.Fprintf(dev, "  %-30s %s%s\n", fmt.Sprintf("-%s=%s", flag.Name(), flag.Type()), flag.Description(), value)
			}
			fmt.Fprintln(dev, "")
		}
	}
}

func PrintUsageCommand(dev *os.File, command *util.Command, flags *util.FlagSet) {
	// Create a table object
	table := util.NewTable()

	// Format
	if command.Format != nil {
		if err := command.Format(command, flags, table); err != nil {
			fmt.Fprintln(dev, err)
			return
		}
	}

	fmt.Fprintf(dev, "%s: %s\n", command.Name, command.Description)
	fmt.Fprintln(dev, "")

	if command.Usage != "" {
		fmt.Fprintf(dev, "\nUsage of %s:\n", command.Name)
		fmt.Fprintf(dev, "  %s (<flag> <flag> ...) %s\n", flags.Name(), command.Usage)
		fmt.Fprintf(dev, "  %s -help %s\n", flags.Name(), command.Usage)
		fmt.Fprintln(dev, "")
	}

	PrintAllFlags(dev, flags)
}

func GetCommand(arg string, commands []*util.Command) *util.Command {
	for _, command := range commands {
		if command.Name == arg {
			return command
		}
	}
	return nil
}

// RegisterFlags registers all the global flags
func RegisterFlags(flags *util.FlagSet) error {
	if err := flags.String(FLAG_CREDENTIALS, ".ytapi", "Folder containing credentials", util.SCOPE_GLOBAL); err != nil {
		return err
	}
	if err := flags.Bool(FLAG_DEBUG, false, "Show API requests and responses on stderr", util.SCOPE_GLOBAL); err != nil {
		return err
	}
	if err := flags.Duration(FLAG_TIMEOUT, 0, "Request timeout", util.SCOPE_GLOBAL); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	flags := util.NewFlagSet(filepath.Base(os.Args[0]))
	if err := RegisterFlags(flags); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if commands := RegisterCommands(); len(commands) == 0 {
		fmt.Fprintln(os.Stderr, "No commands registered")
		os.Exit(-1)
	} else if err := flags.Parse(); err == util.ErrHelpRequested {
		PrintUsage(os.Stderr, commands, flags)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if flags.NArg() == 0 {
		PrintUsage(os.Stderr, commands, flags)
	} else if credentials, err := PathCredentials(flags.GetString(FLAG_CREDENTIALS)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if options, err := ClientOptions(flags); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if client, err := brightcoveapi.NewClient(credentials, options...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if command := GetCommand(flags.Args()[0], commands); command == nil {
		PrintUsage(os.Stderr, commands, flags)
	} else if err := command.ExecBrightcove(client, flags); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
