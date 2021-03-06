package ytapi

/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/djthorpe/ytapi/util"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

const (
	credentialsPathMode = 0700
	credentialsFileMode = 0644
)

const (
	filenameClientSecret   = "client_secret.json"
	filenameServiceAccount = "service_account.json"
	filenameAuthToken      = "auth_token"
	filenameDefaults       = "defaults.json"
)

////////////////////////////////////////////////////////////////////////////////

// Command structure defines a command
type Command struct {
	Name           string
	Description    string
	ServiceAccount bool
	Optional       []*Flag
	Required       []*Flag
	Setup          func(*Values, *Table) error
	Execute        func(*ytservice.Service, *Values, *Table) error
}

// Section structure defines a group of commands
type Section struct {
	Title    string
	Commands []*Command
}

// RegisterFunction defines a registration function
type RegisterFunction struct {
	Title    string
	Callback func() []*Command
}

// FlagSet structure defines the main object for execution of a set of instructions
type FlagSet struct {
	flagset        *flag.FlagSet
	sections       []*Section
	Values         *Values
	Input          *Input
	Output         *Table
	ClientSecrets  string
	ServiceAccount string
	AuthToken      string
	Defaults       string
	filehandle     *os.File
}

// Defaults structure defines values read from store
type Defaults struct {
	ContentOwner string `json:"contentowner,omitempty"`
	Channel      string `json:"channel,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////

// Command-line flags
var (
	FlagCredentials                = Flag{Name: "credentials", Description: "Folder containing credentials", Type: FLAG_STRING, Default: ".ytapi"}
	FlagDebug                      = Flag{Name: "debug", Description: "Show API requests and responses on stderr", Type: FLAG_BOOL, Default: "false"}
	FlagQuotaUser                  = Flag{Name: "quotauser", Description: "Set Quota User for API calls", Type: FLAG_STRING}
	FlagTraceToken                 = Flag{Name: "tracetoken", Description: "Set Trace Token for API calls", Type: FLAG_STRING}
	FlagServiceAccount             = Flag{Name: "serviceaccount", Description: "Obtain token using service account information", Type: FLAG_BOOL, Default: "false"}
	FlagScope                      = Flag{Name: "scope", Description: "Permissions granted during authentication", Type: FLAG_ENUM, Default: "data", Extra: "data|dataread|partner|audit|analytics|revenue|all"}
	FlagFields                     = Flag{Name: "fields", Description: "Comma-separated list of output fields", Type: FLAG_STRING}
	FlagInput                      = Flag{Name: "in", Description: "Input filename and/or format (csv)", Type: FLAG_STRING, Default: "csv"}
	FlagOutput                     = Flag{Name: "out", Description: "Output filename and/or format (txt, csv)", Type: FLAG_STRING, Default: "txt"}
	FlagFile                       = Flag{Name: "file", Description: "Filename", Type: FLAG_STRING}
	FlagContentOwner               = Flag{Name: "contentowner", Description: "Content Owner ID", Type: FLAG_CONTENTOWNER}
	FlagChannel                    = Flag{Name: "channel", Description: "Channel ID", Type: FLAG_CHANNEL}
	FlagMaxResults                 = Flag{Name: "maxresults", Description: "Maximum number of results to return", Type: FLAG_UINT, Default: "0"}
	FlagVideo                      = Flag{Name: "video", Description: "Video ID", Type: FLAG_VIDEO}
	FlagPlaylist                   = Flag{Name: "playlist", Description: "Playlist ID", Type: FLAG_PLAYLIST}
	FlagStream                     = Flag{Name: "stream", Description: "Stream ID or Key", Type: FLAG_STREAM}
	FlagLanguage                   = Flag{Name: "language", Description: "Localized language", Type: FLAG_LANGUAGE}
	FlagRegion                     = Flag{Name: "region", Description: "Country region code", Type: FLAG_REGION}
	FlagTitle                      = Flag{Name: "title", Description: "Metadata title", Type: FLAG_STRING}
	FlagDescription                = Flag{Name: "description", Description: "Metadata description", Type: FLAG_STRING}
	FlagEmbeds                     = Flag{Name: "embeds", Description: "Embed player flag", Type: FLAG_BOOL}
	FlagLicense                    = Flag{Name: "license", Description: "Video License", Type: FLAG_ENUM, Extra: "youtube|creativeCommon"}
	FlagStatsViewable              = Flag{Name: "statsviewable", Description: "Extended video statistics are publicly viewable", Type: FLAG_BOOL}
	FlagPrivacyStatus              = Flag{Name: "status", Description: "Privacy status", Type: FLAG_ENUM, Extra: "private|public|unlisted"}
	FlagDate                       = Flag{Name: "date", Description: "Publish date and time", Type: FLAG_TIME}
	FlagBroadcastStatus            = Flag{Name: "status", Description: "Broadcast status", Type: FLAG_ENUM, Extra: "all|upcoming|active|completed"}
	FlagBroadcastTransition        = Flag{Name: "status", Description: "Broadcast transition", Type: FLAG_ENUM, Extra: "complete|live|testing"}
	FlagStartTime                  = Flag{Name: "start", Description: "Scheduled start time", Type: FLAG_TIME}
	FlagEndTime                    = Flag{Name: "end", Description: "Scheduled end time", Type: FLAG_TIME}
	FlagDvr                        = Flag{Name: "dvr", Description: "Enable DVR", Type: FLAG_BOOL}
	FlagContentEncryption          = Flag{Name: "encryption", Description: "Enable content encryption", Type: FLAG_BOOL}
	FlagEmbed                      = Flag{Name: "embed", Description: "Enable embedding", Type: FLAG_BOOL}
	FlagRecordFromStart            = Flag{Name: "record", Description: "Enable recording", Type: FLAG_BOOL}
	FlagStartWithSlate             = Flag{Name: "slate", Description: "Start with slate", Type: FLAG_BOOL}
	FlagClosedCaptions             = Flag{Name: "captions", Description: "Enable closed captions", Type: FLAG_BOOL}
	FlagMonitorStream              = Flag{Name: "monitor", Description: "Enable stream monitoring", Type: FLAG_BOOL}
	FlagBroadcastDelay             = Flag{Name: "delay", Description: "Broadcast delay (ms)", Type: FLAG_UINT}
	FlagLowLatency                 = Flag{Name: "lowlatency", Description: "Enable low latency", Type: FLAG_BOOL}
	FlagProjection                 = Flag{Name: "projection", Description: "Projection format", Type: FLAG_ENUM, Extra: "360|rectangular"}
	FlagCuepointOffset             = Flag{Name: "offset", Description: "Offset time (ms)", Type: FLAG_UINT}
	FlagCuepointDuration           = Flag{Name: "duration", Description: "Duration", Type: FLAG_DURATION}
	FlagCuepointTime               = Flag{Name: "time", Description: "Walltime", Type: FLAG_TIME}
	FlagCuepointPodDuration        = Flag{Name: "podduration", Description: "Ad Pod Duration", Type: FLAG_DURATION}
	FlagVideoFilter                = Flag{Name: "filter", Description: "Video filter", Type: FLAG_ENUM, Extra: "chart|like|dislike|likes|favorites|uploads|watchhistory|watchlater", Default: "uploads"}
	FlagVideoCategory              = Flag{Name: "category", Description: "Video Category", Type: FLAG_UINT}
	FlagVideoRating                = Flag{Name: "rating", Description: "Video Rating", Type: FLAG_ENUM, Extra: "like|dislike|none"}
	FlagCommentThread              = Flag{Name: "thread", Description: "Comment Thread", Type: FLAG_STRING}
	FlagCommentText                = Flag{Name: "text", Description: "Comment Text", Type: FLAG_STRING}
	FlagCommentFormat              = Flag{Name: "format", Description: "Comment Format", Type: FLAG_ENUM, Extra: "plainText|html", Default: "plainText"}
	FlagCommentOrder               = Flag{Name: "order", Description: "Comment order", Type: FLAG_ENUM, Extra: "time|relevance"}
	FlagCommentModerationStatus    = Flag{Name: "status", Description: "Comment moderation status", Type: FLAG_ENUM, Extra: "heldForReview|likelySpam|published"}
	FlagCommentModerationStatus2   = Flag{Name: "status", Description: "Comment moderation status", Type: FLAG_ENUM, Extra: "heldForReview|published|rejected"}
	FlagCommentBanAuthor           = Flag{Name: "ban", Description: "Ban comment author", Type: FLAG_BOOL}
	FlagChat                       = Flag{Name: "chat", Description: "Live Chat or Broadcast ID", Type: FLAG_STRING}
	FlagChatMessage                = Flag{Name: "message", Description: "Live Chat Message", Type: FLAG_STRING}
	FlagChatText                   = Flag{Name: "text", Description: "Chat Text", Type: FLAG_STRING}
	FlagActivityHome               = Flag{Name: "home", Description: "Display Homepage Activity Feed", Type: FLAG_BOOL}
	FlagPlaylistPosition           = Flag{Name: "position", Description: "Playlist position", Type: FLAG_UINT}
	FlagPlaylistNote               = Flag{Name: "note", Description: "Playlist note", Type: FLAG_STRING}
	FlagSectionType                = Flag{Name: "type", Description: "Channel Section type", Type: FLAG_ENUM, Extra: "allPlaylists|completedEvents|likedPlaylists|likes|liveEvents|multipleChannels|multiplePlaylists|popularUploads|postedPlaylists|postedVideos|recentActivity|recentPosts|recentUploads|singlePlaylist|subscriptions|upcomingEvents"}
	FlagSectionStyle               = Flag{Name: "style", Description: "Channel Section style", Type: FLAG_ENUM, Extra: "horizontalRow|verticalList", Default: "horizontalRow"}
	FlagSectionPosition            = Flag{Name: "position", Description: "Channel Section position", Type: FLAG_UINT}
	FlagSearchQuery                = Flag{Name: "q", Description: "Search query", Type: FLAG_STRING}
	FlagSearchOrder                = Flag{Name: "order", Description: "Search order", Type: FLAG_ENUM, Extra: "date|rating|relevance|title|viewcount"}
	FlagSearchChannelOrder         = Flag{Name: "order", Description: "Search order", Type: FLAG_ENUM, Extra: "date|rating|relevance|title|viewcount|videocount"}
	FlagSearchVideo                = Flag{Name: "video", Description: "Related video", Type: FLAG_VIDEO}
	FlagSearchSafe                 = Flag{Name: "safesearch", Description: "Restricted content filter", Type: FLAG_ENUM, Extra: "none|moderate|strict"}
	FlagSearchBroadcastStatus      = Flag{Name: "status", Description: "Broadcast status", Type: FLAG_ENUM, Extra: "completed|live|upcoming"}
	FlagCaption                    = Flag{Name: "caption", Description: "Caption ID", Type: FLAG_STRING}
	FlagCaptionSync                = Flag{Name: "sync", Description: "Automatically synchronize the caption file with the audio track of the video", Type: FLAG_BOOL}
	FlagCaptionDraft               = Flag{Name: "draft", Description: "Status of the caption track", Type: FLAG_BOOL}
	FlagCaptionFormat              = Flag{Name: "format", Description: "Format of the caption track", Type: FLAG_ENUM, Extra: "sbv|scc|srt|ttml|vtt"}
	FlagCaptionName                = Flag{Name: "name", Description: "Name of the caption track", Type: FLAG_STRING}
	FlagPolicy                     = Flag{Name: "policy", Description: "Policy ID", Type: FLAG_STRING}
	FlagPolicyOrder                = Flag{Name: "order", Description: "Policy list order", Type: FLAG_ENUM, Extra: "timeUpdatedAsc|timeUpdatedDesc"}
	FlagClaim                      = Flag{Name: "claim", Description: "Claim ID", Type: FLAG_VIDEO}
	FlagClaimType                  = Flag{Name: "type", Description: "Claim Type: Defaults to audiovisual", Type: FLAG_ENUM, Extra: "audio|visual|audiovisual", Default: "audiovisual"}
	FlagClaimStatus                = Flag{Name: "status", Description: "Claim Status", Type: FLAG_ENUM, Extra: "active|inactive"}
	FlagClaimBlockOutsideOwnership = Flag{Name: "blockoutsideownership", Description: "Block viewing outside ownership regions", Type: FLAG_BOOL}
	FlagAsset                      = Flag{Name: "asset", Description: "Asset ID", Type: FLAG_STRING}
	FlagAssetFilter                = Flag{Name: "filter", Description: "Retrieve computed asset information or my asset information", Type: FLAG_ENUM, Extra: "none|effective|mine", Default: "none"}
	FlagUploader                   = Flag{Name: "uploader", Description: "Content Owner Uploader Name", Type: FLAG_STRING}
	FlagStreamResolution           = Flag{Name: "resolution", Description: "Stream Resolution", Type: FLAG_ENUM, Extra: "2160p_hfr|2160p|1440p_hfr|1440p|1080p_hfr|1080p|720p_hfr|720p|480p|360p|240p", Default: "1080p"}
	FlagStreamType                 = Flag{Name: "type", Description: "Stream Ingestion Type", Type: FLAG_ENUM, Extra: "rtmp|dash", Default: "rtmp"}
	FlagStreamReusable             = Flag{Name: "reusable", Description: "Stream Reusable", Type: FLAG_BOOL}
	FlagAnalyticsPeriod            = Flag{Name: "period", Description: "Time period", Type: FLAG_STRING, Default: "last28Days"}
	FlagAnalyticsMetrics           = Flag{Name: "metrics", Description: "Analytics Metrics", Type: FLAG_STRING}
	FlagAnalyticsDimensions        = Flag{Name: "dimensions", Description: "Analytics Dimensions", Type: FLAG_STRING}
	FlagAnalyticsFilter            = Flag{Name: "filter", Description: "Analytics Filters", Type: FLAG_STRING}
	FlagAnalyticsSort              = Flag{Name: "sort", Description: "Data sort order", Type: FLAG_STRING}
	FlagAnalyticsCurrency          = Flag{Name: "currency", Description: "Reporting Currency", Type: FLAG_STRING}
	FlagAnalyticsIncludeSystem     = Flag{Name: "system", Description: "Include system managed reports", Type: FLAG_BOOL}
)

////////////////////////////////////////////////////////////////////////////////
// Global variables
var (

	// Global flags which are defined for every invocation of the tool
	GlobalFlags = []*Flag{
		&FlagDebug, &FlagQuotaUser, &FlagTraceToken, &FlagCredentials,
		&FlagContentOwner, &FlagChannel,
		&FlagFields, &FlagOutput, &FlagInput,
	}

	// Variety of error conditions
	ErrorEmptyArgs        = errors.New("No Arguments")
	ErrorUsage            = errors.New("Display usage information")
	ErrorWriteDefaults    = errors.New("Write Defaults to file")
	ErrorWriteCredentials = errors.New("Write Credentials")
	ErrorRemoveAuthToken  = errors.New("Remove OAuth token")
	ErrorServiceAccount   = errors.New("Invalid service account or missing content owner")
)

////////////////////////////////////////////////////////////////////////////////
// FlagSet implementation

// NewFlagSet returns a new set of flags
func NewFlagSet() *FlagSet {
	this := new(FlagSet)
	this.flagset = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	this.flagset.SetOutput(ioutil.Discard)
	this.sections = make([]*Section, 0)
	this.Values = NewValues()
	this.Output = NewTable()
	this.Input = NewInput()
	return this
}

// AddFlag adds a flag to the flagset
func (this *FlagSet) AddFlag(flag *Flag) error {

	// Skip if flag has already been added
	if flag.added {
		return nil
	}

	// check for flag name clash
	if this.flagset.Lookup(flag.Name) != nil {
		return errors.New(fmt.Sprint("Duplicate flag: ", flag.Name))
	}

	// set flag
	value := this.Values.Set(&Value{flag: flag})
	this.flagset.Var(value, flag.Name, flag.Description)

	// mark as added
	flag.added = true

	// success
	return nil
}

func (this *FlagSet) AddFlags(flags []*Flag) error {
	for _, flag := range flags {
		if err := this.AddFlag(flag); err != nil {
			return err
		}
	}
	// success
	return nil
}

func (this *FlagSet) RegisterCommands(funcs []*RegisterFunction) error {

	// call functions to retrieve sets of commands
	commands := make(map[string]bool, 0)
	for _, f := range funcs {
		section := &Section{
			Title:    f.Title,
			Commands: make([]*Command, 0),
		}
		for _, c := range f.Callback() {
			// check for existing command
			if _, exists := commands[c.Name]; exists {
				return errors.New(fmt.Sprint("Duplicate command: ", c.Name))
			}
			// or else add to list of sections
			section.Commands = append(section.Commands, c)
			commands[c.Name] = true
		}
		this.sections = append(this.sections, section)
	}

	// Success
	return nil
}

func (this *FlagSet) GetCommandFromName(name string) (*Command, error) {
	for _, section := range this.sections {
		for _, command := range section.Commands {
			if command.Name == name {
				return command, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprint("Invalid command: ", name))
}

func (this *FlagSet) setPaths() error {
	// get credentials path, make it if it doesn't exist
	credentialsPath, exists := util.ResolvePath(this.Values.GetString(&FlagCredentials), util.UserDir())
	if exists == false {
		if err := os.Mkdir(credentialsPath, credentialsPathMode); err != nil {
			return err
		}
	}

	// client secrets
	clientSecretsPath, _ := util.ResolvePath(filenameClientSecret, credentialsPath)
	this.ClientSecrets = clientSecretsPath

	// service account
	serviceAccountPath, _ := util.ResolvePath(filenameServiceAccount, credentialsPath)
	this.ServiceAccount = serviceAccountPath

	// oauth token
	authTokenPath, _ := util.ResolvePath(filenameAuthToken, credentialsPath)
	this.AuthToken = authTokenPath

	// defaults file
	defaultsPath, _ := util.ResolvePath(filenameDefaults, credentialsPath)
	this.Defaults = defaultsPath

	// success
	return nil
}

func (this *FlagSet) Parse() (*Command, error) {
	var command *Command
	var err error

	// Add global flags
	if err := this.AddFlags(GlobalFlags); err != nil {
		return nil, err
	}

	// Determine the command which will be used to add additional flags
	if len(os.Args) < 2 {
		return nil, ErrorEmptyArgs
	}
	lastarg := os.Args[len(os.Args)-1]
	if strings.HasPrefix(lastarg, "-") == false {
		command, err = this.GetCommandFromName(lastarg)
		if err != nil {
			return nil, err
		}
	}

	// Add additional optional and required flags
	if command != nil {
		if err := this.AddFlags(command.Required); err != nil {
			return command, err
		}
		// Skip an optional flag if it was already marked as required
		if err := this.AddFlags(command.Optional); err != nil {
			return command, err
		}
	}

	// Set empty usage function
	this.flagset.Usage = func() {}

	// Set flag values
	err = this.flagset.Parse(os.Args[1:])

	// Set paths for various files
	err2 := this.setPaths()
	if err == nil && err2 != nil {
		err = err2
	}

	// Check for -help on command line
	if this.flagset.NArg() == 0 && err == flag.ErrHelp {
		return nil, ErrorUsage
	}
	// Check for none or too many arguments
	if this.flagset.NArg() == 0 {
		return nil, errors.New("Missing command")
	}
	if this.flagset.NArg() > 1 && err == nil {
		return nil, errors.New(fmt.Sprint("Too many arguments on command line: ", this.flagset.Args()))
	}
	if command == nil {
		return nil, ErrorUsage
	}

	// setup output
	if err := this.SetupCommand(command); err != nil {
		return command, err
	}

	// check for -help <Command>
	if err == flag.ErrHelp {
		return command, ErrorUsage
	}
	// general caught errors
	if err != nil {
		return command, err
	}

	// Look for missing required flags
	for _, flag := range command.Required {
		if this.Values.IsSet(flag) == false {
			return nil, errors.New(fmt.Sprint("Missing required flag: ", flag.Name))
		}
	}

	// success
	return command, nil
}

////////////////////////////////////////////////////////////////////////////////
// Usage

func (this *FlagSet) UsageGlobalFlags() {
	// Output globals
	fmt.Fprintf(os.Stderr, "\nGlobal flags:\n\n")
	for _, f := range GlobalFlags {
		// skip content owner flag if no service account
		if f == &FlagContentOwner && this.ServiceAccount == "" {
			continue
		}
		// output flag description
		fmt.Fprintf(os.Stderr, "\t-%s <%s>\n\t\t%s", f.Name, f.TypeString(), f.Description)
		if f.Default != "" {
			fmt.Fprintf(os.Stderr, " (default: \"%s\")", f.Default)
		}
		fmt.Fprint(os.Stderr, "\n")
	}
}

// Output a list of flags for a command.
func (this *FlagSet) UsageCommand(command *Command) {
	fmt.Fprintf(os.Stderr, "\nFlags for %s:\n\n", command.Name)
	// Output flags
	for _, f := range command.Required {
		fmt.Fprintf(os.Stderr, "\t-%s <%s>\n\t\t%s, required\n", f.Name, f.TypeString(), f.Description)
	}
	for _, f := range command.Optional {
		fmt.Fprintf(os.Stderr, "\t-%s <%s>\n\t\t%s, optional\n", f.Name, f.TypeString(), f.Description)
	}
}

// Output a list of possible commands, grouped by section. Will not display
// commands which are only to be used for service accounts, if no service
// account is installed
func (this *FlagSet) UsageCommandList() {
	for _, section := range this.sections {
		var commands []*Command = make([]*Command, 0)
		for _, command := range section.Commands {
			if command.ServiceAccount && this.ServiceAccount == "" {
				continue
			}
			commands = append(commands, command)
		}
		if len(commands) == 0 {
			continue
		}
		fmt.Fprintf(os.Stderr, "\n%s:\n\n", section.Title)
		for _, command := range commands {
			fmt.Fprintf(os.Stderr, "\t%s\n\t\t%s\n", command.Name, command.Description)
		}
	}
}

func (this *FlagSet) UsageFields() {
	for _, part := range this.Output.Parts(true) {
		fields := this.Output.FieldsForPart(part)
		if len(fields) == 0 {
			continue
		}
		fmt.Fprintf(os.Stderr, "\nFields for '%s':\n\n", part)
		for _, field := range fields {
			fmt.Fprintf(os.Stderr, "\t%s (%s)\n", field.Name, field.TypeString())
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// Read and write Defaults

func (this *FlagSet) ReadDefaults() error {
	var err error
	// if a file exists, then read it
	if _, err := os.Stat(this.Defaults); os.IsNotExist(err) {
		// file doesn't exist, so just return
		return err
	}
	// read in the file
	bytes, err := ioutil.ReadFile(this.Defaults)
	if err != nil {
		return err
	}
	defaults := &Defaults{}
	err = json.Unmarshal(bytes, defaults)
	if err != nil {
		return err
	}
	// ContentOwner
	if err == nil && defaults.ContentOwner != "" {
		err = this.Values.SetDefault(&FlagContentOwner, string(defaults.ContentOwner))
	}
	// Channel
	if err == nil && defaults.Channel != "" {
		err = this.Values.SetDefault(&FlagChannel, string(defaults.Channel))
	}
	if err != nil {
		return err
	}
	// success
	return nil
}

func (this *FlagSet) WriteDefaults() error {
	defaults := &Defaults{}
	if this.Values.IsSet(&FlagContentOwner) {
		defaults.ContentOwner = this.Values.GetString(&FlagContentOwner)
	}
	if this.Values.IsSet(&FlagChannel) && this.Values.IsSet(&FlagContentOwner) {
		defaults.Channel = this.Values.GetString(&FlagChannel)
	}
	json, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(this.Defaults, json, credentialsFileMode)
	if err != nil {
		return err
	}
	return nil
}

func (this *FlagSet) RemoveAuthToken() error {
	if this.AuthToken == "" {
		return nil
	}
	if err := util.DeleteFileIfExists(this.AuthToken); err != nil {
		return err
	}
	return nil
}

func (this *FlagSet) WriteClientSecret(data64 string) error {
	if this.ClientSecrets == "" {
		return errors.New("Invalid client secret filename")
	}
	data, err := base64.StdEncoding.DecodeString(data64)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(this.ClientSecrets, data, credentialsFileMode)
	if err != nil {
		return err
	}
	return nil
}

func (this *FlagSet) WriteServiceAccount(data64 string) error {
	if this.ServiceAccount == "" {
		return errors.New("Invalid service account filename")
	}
	data, err := base64.StdEncoding.DecodeString(data64)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(this.ServiceAccount, data, credentialsFileMode)
	if err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Add and remove fields

func (this *FlagSet) SetFields(fields []string) error {
	// Check fields - all of which should start with +- or with no +-, but not mixed
	is_plusminus := false
	is_setfields := false
	for _, field := range fields {
		field := strings.TrimSpace(field)
		if field == "" {
			continue
		}
		if strings.HasPrefix(field, "+") || strings.HasPrefix(field, "-") {
			if is_setfields {
				return errors.New(fmt.Sprintf("Invalid field name or snippet \"%s\", cannot have prefix of '+' or '-'", field))
			}
			is_plusminus = true
		} else {
			if is_plusminus {
				return errors.New(fmt.Sprintf("Invalid field name or snippet \"%s\", must have prefix of '+' or '-'", field))
			}
			is_setfields = true
		}
	}
	// Sanity check
	if is_plusminus == false && is_setfields == false {
		return errors.New("Missing fields")
	}
	// Remove existing field columns if we're setting the fields from scratch
	if is_setfields {
		this.Output.SetColumns([]string{})
	}
	// Add and remove columns
	for _, field := range fields {
		var err error
		if is_setfields {
			err = this.Output.AddFieldOrPart(field)
		} else if strings.HasPrefix(field, "+") {
			err = this.Output.AddFieldOrPart(field[1:])
		} else if strings.HasPrefix(field, "-") {
			err = this.Output.RemoveFieldOrPart(field[1:])
		}
		if err != nil {
			return err
		}
	}
	// Success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Open & Close output

func (this *FlagSet) OpenOutput() error {
	output := this.Values.GetString(&FlagOutput)
	format := OUTPUT_ASCII
	ext := filepath.Ext(output)
	if ext == "" {
		ext = output
	}

	switch ext {
	case "txt":
		format = OUTPUT_ASCII
	case ".txt":
		format = OUTPUT_ASCII
	case "csv":
		format = OUTPUT_CSV
	case ".csv":
		format = OUTPUT_CSV
	default:
		return errors.New("Invalid output format, should be csv,txt")
	}
	// Open output file or use stdout
	if output != ext {
		fh, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, credentialsFileMode)
		if err != nil {
			return err
		}
		this.Output.SetDataFormat(fh, format)
		this.filehandle = fh
	} else {
		this.Output.SetDataFormat(os.Stdout, format)
	}
	// Success
	return nil
}

func (this *FlagSet) CloseOutput() error {
	if this.filehandle != nil {
		return this.filehandle.Close()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Read input

func (this *FlagSet) ReadInput() error {
	// Determine input filename and format
	input := this.Values.GetString(&FlagInput)
	ext := filepath.Ext(input)
	format := INPUT_CSV
	switch {
	case input == "csv" || input == "CSV":
		input = "-"
	case ext == ".csv" || ext == ".CSV":
		format = INPUT_CSV
	case input == "-":
		format = INPUT_CSV
	default:
		return errors.New("Invalid input format, only files with .csv extension supported")
	}

	// Open input file
	if input == "-" || input == "" {
		this.Input.SetDataSource(os.Stdin, format)
	} else {
		if fh, err := os.Open(input); err != nil {
			return err
		} else {
			this.Input.SetDataSource(fh, format)
		}
	}
	defer this.Input.Close()

	// Read the data
	err := this.Input.ReadAll()
	if err != nil {
		return err
	}

	// Success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Setup & Execute command, display output

func (this *FlagSet) SetupCommand(command *Command) error {
	// if command is nil, then return without execution
	if command == nil || command.Setup == nil {
		return nil
	}
	err := command.Setup(this.Values, this.Output)
	if err == ErrorRemoveAuthToken {
		err = this.RemoveAuthToken()
	}
	return err
}

func (this *FlagSet) ExecuteCommand(command *Command, service *ytservice.Service) error {
	// if command is nil, then return without execution
	if command == nil || command.Execute == nil {
		return nil
	}

	// check for service account and return error if command can't be executed
	if command.ServiceAccount && (service.ServiceAccount == false) {
		return ErrorServiceAccount
	}

	// execute the command
	if err := command.Execute(service, this.Values, this.Output); err != nil {
		return err
	}

	// return success
	return nil
}

func (this *FlagSet) DisplayOutput() error {
	if this.Output.NumberOfColumns() > 0 {
		if err := this.Output.DataOutput(); err != nil {
			return err
		}
	}
	if this.Output.NumberOfRows() > 1 {
		this.Output.Info(fmt.Sprintf("%v items returned", this.Output.NumberOfRows()))
	}
	// success
	return nil
}
