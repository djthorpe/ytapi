/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

const (
	FLAG_STRING = iota
	FLAG_UINT
	FLAG_BOOL
	FLAG_ENUM
	FLAG_VIDEO
	FLAG_CHANNEL
	FLAG_PLAYLIST
	FLAG_LANGUAGE
	FLAG_CONTENTOWNER
)

////////////////////////////////////////////////////////////////////////////////

// Defines a flag
type Flag struct {
	Name        string
	Description string
	Type        uint32
	Extra       string
	Default     string
}

// Defines a command
type Command struct {
	Name        string
	Description string
	Optional    []*Flag
	Required    []*Flag
	Setup       func(*Values, *ytservice.Table) error
	Execute     func(*ytservice.Service, *Values, *ytservice.Table) error
}

// Registration function
type RegisterFunction func() []Command

////////////////////////////////////////////////////////////////////////////////

// Command-line flags
var (
    FlagDebug           = Flag{Name: "debug", Description: "Show API requests and responses on stderr", Type: FLAG_BOOL, Default: "false"}
	FlagCredentials     = Flag{Name: "credentials", Description: "Folder containing credentials", Type: FLAG_STRING, Default: ".credentials"}
	FlagDefaults        = Flag{Name: "defaults", Description: "Defaults filename", Type: FLAG_STRING, Default: "defaults.json"}
	FlagClientSecret    = Flag{Name: "clientsecret", Description: "Client Secret filename", Type: FLAG_STRING, Default: "client_secret.json"}
	FlagServiceAccount  = Flag{Name: "serviceaccount", Description: "Service Account filename", Type: FLAG_STRING, Default: "service_account.json"}
	FlagAuthToken       = Flag{Name: "authtoken", Description: "OAuth token filename", Type: FLAG_STRING, Default: "oauth_token"}
	FlagContentOwner    = Flag{Name: "contentowner", Description: "Content Owner ID", Type: FLAG_CONTENTOWNER}
	FlagChannel         = Flag{Name: "channel", Description: "Channel ID", Type: FLAG_CHANNEL}
	FlagVideo           = Flag{Name: "video", Description: "Video ID", Type: FLAG_VIDEO}
	FlagBroadcastStatus = Flag{Name: "status", Description: "Broadcast Status", Type: FLAG_ENUM, Extra: "all|upcoming|live|completed"}
    FlagMaxResults      = Flag{Name: "maxresults", Description: "Maximum number of results to return", Type: FLAG_UINT, Default: "0"}
	FlagTitle           = Flag{Name: "title", Description: "Metadata Title", Type: FLAG_STRING}
	FlagDescription     = Flag{Name: "descriptioon", Description: "Metadata Description", Type: FLAG_STRING}
)

// Global variables
var (
	globalflags = []*Flag{&FlagDebug, &FlagCredentials, &FlagDefaults, &FlagClientSecret, &FlagServiceAccount, &FlagAuthToken}
	flagvalues  = make(map[string]*Value, 0)
)

////////////////////////////////////////////////////////////////////////////////
// Flag implementation

func (this *Flag) TypeString() string {
	switch {
	case this.Type == FLAG_STRING:
		return "string"
	case this.Type == FLAG_UINT:
		return "uint"
	case this.Type == FLAG_BOOL:
		return "bool"
	case this.Type == FLAG_ENUM:
		return this.Extra
	case this.Type == FLAG_VIDEO:
		return "video"
	case this.Type == FLAG_CHANNEL:
		return "channel"
	case this.Type == FLAG_PLAYLIST:
		return "playlist"
	case this.Type == FLAG_LANGUAGE:
		return "language"
    case this.Type == FLAG_CONTENTOWNER:
        return "contentowner"
	default:
		return "other"
	}
}

func (this *Flag) asUint(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

func (this *Flag) asBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func (this *Flag) asEnum(value string) (string, error) {
	tags := strings.Split(this.Extra, "|")
	if len(tags) < 2 {
		return "", errors.New("Missing or invalid 'Extra' field")
	}
	for _, tag := range tags {
		if tag == value {
			return value, nil
		}
	}
	return "", errors.New(fmt.Sprint("Value should be one of: ", strings.Join(tags, ",")))
}

func (this *Flag) asVideo(value string) (Video, error) {
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\-]{11})$", value)
	if matched == false {
		return "", errors.New("Malformed video value")
	}
	return Video(value), nil
}

func (this *Flag) asContentOwner(value string) (ContentOwner, error) {
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\_]{22})$", value)
	if matched == false {
		return "", errors.New("Malformed content owner value")
	}
	return ContentOwner(value), nil
}

func (this *Flag) asChannel(value string) (Channel, error) {
	matched, _ := regexp.MatchString("^UC([a-zA-Z0-9\\-]{22})$", value)
	if matched == false {
		return "", errors.New("Malformed channel value")
	}
	return Channel(value), nil
}

func (this *Flag) asPlaylist(value string) (Playlist, error) {
	// TODO: Fix this
	matched, _ := regexp.MatchString("^PL([a-zA-Z0-9\\-]{22})$", value)
	if matched == false {
		return "", errors.New("Malformed playlist value")
	}
	return Playlist(value), nil
}

func (this *Flag) asLanguage(value string) (Language, error) {
	matched, _ := regexp.MatchString("^([a-z]{2})$", value)
	if matched {
		return Language(value), nil
	}
	matched, _ = regexp.MatchString("^([a-z]{2})\\-([a-zA-Z0-9]+)$", value)
	if matched {
		return Language(value), nil
	}
	return "", errors.New("Malformed language value")
}

////////////////////////////////////////////////////////////////////////////////
// Usage

func usage() {
	execname := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", execname)
	fmt.Fprintf(os.Stderr, "\t%s -help\n", execname)
	fmt.Fprintf(os.Stderr, "\t%s -help <command>\n", execname)
	fmt.Fprintf(os.Stderr, "\t%s <flags> <command>\n", execname)

	// Output globals
	fmt.Fprintf(os.Stderr, "\nGlobal flags:\n\n")
	for _, f := range globalflags {
		fmt.Fprintf(os.Stderr, "\t-%s <%s>\n\t\t%s\n", f.Name, f.TypeString(), f.Description)
	}
}

func usageListCommands(commands map[string]Command) {
	fmt.Fprintf(os.Stderr, "Commands:\n\n")
	// Output commands
	for name, command := range commands {
		fmt.Fprintf(os.Stderr, "\t%s\n\t\t%s\n", name, command.Description)
	}
}

func usageListCommandFlags(command Command) {
	fmt.Fprintf(os.Stderr, "\nFlags for %s:\n\n", command.Name)
	// Output flags
	for _, f := range command.Required {
		fmt.Fprintf(os.Stderr, "\t-%s <%s>\n\t\t%s, required\n", f.Name, f.TypeString(), f.Description)
	}
	for _, f := range command.Optional {
		fmt.Fprintf(os.Stderr, "\t-%s <%s>\n\t\t%s, optional\n", f.Name, f.TypeString(), f.Description)
	}
}

////////////////////////////////////////////////////////////////////////////////
// Parse Flags

func addFlags(flagset *flag.FlagSet, flags []*Flag) error {
	for _, f := range flags {
		if _, exists := flagvalues[f.Name]; exists {
			return errors.New(fmt.Sprint("Duplicate flag: ", f.Name))
		}
		flagvalues[f.Name] = &Value{flag: f}
		flagset.Var(flagvalues[f.Name], f.Name, f.Description)
	}
	// success
	return nil
}

// Parse arguments on the command line
func ParseFlags(funcs []RegisterFunction) (*Command, *Values, error) {

	commands := make(map[string]Command, 0)
	flagset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	// register global flags
	if err := addFlags(flagset, globalflags); err != nil {
		return nil, nil, err
	}

	// call functions to retrieve sets of commands
	for _, f := range funcs {
		for _, c := range f() {
			if _, exists := commands[c.Name]; exists {
				return nil, nil, errors.New(fmt.Sprint("Duplicate command: ", c.Name))
			}
			commands[c.Name] = c
		}
	}

	// Retrieve last element of arguments, which is the API command
	if len(os.Args) < 2 {
		usage()
		return nil, nil, nil
	}

	var command Command

	// Get command from command line
	lastarg := os.Args[len(os.Args)-1]
	if strings.HasPrefix(lastarg, "-") == false {
		var exists bool
		command, exists = commands[lastarg]
		if exists == false {
			return nil, nil, errors.New(fmt.Sprint("Invalid command: ", lastarg))
		}

		// Register flag values for command
		if err := addFlags(flagset, command.Optional); err != nil {
			return nil, nil, err
		}
		if err := addFlags(flagset, command.Required); err != nil {
			return nil, nil, err
		}
	}

	// Set empty usage function
	flagset.Usage = func() {}

	// Set flags
	err := flagset.Parse(os.Args[1:])

	// Check for -help on command line
	if flagset.NArg() == 0 && err == flag.ErrHelp {
		usage()
		usageListCommands(commands)
		return nil, nil, nil
	}

	// Check for none or too many arguments
	if flagset.NArg() == 0 {
		return nil, nil, errors.New("Missing command")
	}
	if flagset.NArg() > 1 && err == nil {
		return nil, nil, errors.New(fmt.Sprint("Too many arguments on command line: ", flagset.Args()))
	}

	// Check for -help on command line, or any other error from the Parse() method
	if err == flag.ErrHelp {
		usage()
		usageListCommandFlags(command)
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}

	// Check for empty command
	if command.Name == "" {
		usage()
		return nil, nil, nil
	}

	// Look for missing required flags
	for _, flag := range command.Required {
		v, exists := flagvalues[flag.Name]
		if exists == false || v.IsSet() == false {
			return nil, nil, errors.New(fmt.Sprint("Missing required flag: ", flag.Name))
		}
	}

	// Copy parameters into the parameters
	values := NewValues()
	for _, value := range flagvalues {
		values.Set(value)
	}

	// Return success
	return &command, values, nil
}
