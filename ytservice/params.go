/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"

	"encoding/json"
	"io/ioutil"
)

////////////////////////////////////////////////////////////////////////////////

const (
	FLAG_REQUIRED = 0x0000
	FLAG_OPTIONAL = 0x0001
	FLAG_STRING   = 0x0010
	FLAG_UINT     = 0x0020
	FLAG_ENUM     = 0x0030
	FLAG_VIDEO    = 0x0040
	FLAG_CHANNEL  = 0x0050
	FLAG_PLAYLIST = 0x0060
	FLAG_LANGUAGE = 0x0070
)

////////////////////////////////////////////////////////////////////////////////

// Flag represents a parameter name and type
type Flag struct {
	Name  string
	Type  uint32
	Extra string
}

// Params object stores all the parameters used for making API requests
type Params struct {
	ContentOwner    *string `json:"contentowner,omitempty"`
	Channel         *string `json:"channel,omitempty"`
	Video           *string `json:"-"`
	Stream          *string `json:"-"`
	MaxResults      int64   `json:"-"`
	Query           *string `json:"-"`
	BroadcastStatus *string `json:"-"`
	Language        *string `json:"-"`
	Title           *string `json:"-"`
	Description     *string `json:"-"`

	commands  []string
	SetupHook map[string]func(*Params, *Table) error           // command -> SetupFunc
	ExecHook  map[string]func(*Service, *Params, *Table) error // command -> ExecFunc
	FlagSpec  map[string]map[string]Flag                       // command:flag -> Flag
}

////////////////////////////////////////////////////////////////////////////////

// NewParams returns a new Params object
func NewParams() *Params {
	this := new(Params)
	return this
}

// NewParamsFromJSON returns a params object from a JSON file
func NewParamsFromJSON(filename string) (*Params, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, NewError(ErrorInvalidDefaults, err)
	}
	this := NewParams()
	err = json.Unmarshal(bytes, this)
	if err != nil {
		return nil, NewError(ErrorInvalidDefaults, err)
	}

	// Create maps used in params
	this.commands = make([]string, 0)
	this.FlagSpec = make(map[string]map[string]Flag)
	this.SetupHook = make(map[string]func(*Params, *Table) error)
	this.ExecHook = make(map[string]func(*Service, *Params, *Table) error)

	return this, nil
}

////////////////////////////////////////////////////////////////////////////////

// Register flags and setup/do hooks
func (this *Params) Register(command string, setup func(*Params, *Table) error, exec func(*Service, *Params, *Table) error, flags []Flag) error {

	if _, exists := this.FlagSpec[command]; exists {
		return errors.New(fmt.Sprint("Duplicate command registration: ", command))
	}

	this.commands = append(this.commands, command)
	this.FlagSpec[command] = make(map[string]Flag)
	for _, flag := range flags {
		if _, exists := this.FlagSpec[command][flag.Name]; exists {
			return errors.New(fmt.Sprint("Duplicate flag registration for command ", command, ": ", flag.Name))
		}
		this.FlagSpec[command][flag.Name] = flag
	}

	// Save ExecHook and SetupHook
	this.SetupHook[command] = setup
	this.ExecHook[command] = exec

	return nil
}

func (this *Params) isIgnoreFlag(name string) bool {
	if name == "authtoken" {
		return true
	}
	if name == "clientsecret" {
		return true
	}
	if name == "credentials" {
		return true
	}
	if name == "debug" {
		return true
	}
	if name == "defaults" {
		return true
	}
	if name == "httptest.serve" {
		return true
	}
	if name == "output" {
		return true
	}
	return false
}

// CheckFlags checks all parameters
func (this *Params) CheckFlags(command string) error {
	var err error

	fmt.Printf("FLAGS %+v\n", this.FlagSpec)

	specs, ok := this.FlagSpec[command]
	if !ok {
		return errors.New(fmt.Sprint("Invalid command: ", command))
	}
	flag.VisitAll(func(value *flag.Flag) {
		// if error is caught, return
		if err != nil {
			return
		}
		// ignore some flags
		if this.isIgnoreFlag(value.Name) {
			return
		}
		spec, ok := specs[value.Name]
		if !ok {
			err = errors.New(fmt.Sprint("Invalid flag: ", value.Name))
			return
		}
		// if flag is optional and value is empty, then return
		fmt.Println("FLAG: ", spec, value)
	})

	if err != nil {
		return err
	}
	// success
	return nil
}

// Commands returns the list of command names
func (this *Params) Commands() []string {
	return this.commands
}

func (this *Params) Flags(command string) map[string]Flag {
	return this.FlagSpec[command]
}

////////////////////////////////////////////////////////////////////////////////

// Copy nakes a copy of the object
func (this *Params) Copy() *Params {
	copy := NewParams()
	copy.MaxResults = this.MaxResults
	copy.ContentOwner = this.ContentOwner
	copy.Channel = this.Channel
	copy.Video = this.Video
	copy.Stream = this.Stream
	copy.Query = this.Query
	copy.BroadcastStatus = this.BroadcastStatus
	copy.Language = this.Language
	copy.Title = this.Title
	copy.Description = this.Description
	return copy
}

// Save params object
func (this *Params) Save(filename string, perm os.FileMode) error {
	json, err := json.MarshalIndent(this, "", "  ")
	if err != nil {
		return NewError(ErrorInvalidDefaults, err)
	}
	err = ioutil.WriteFile(filename, json, perm)
	if err != nil {
		return NewError(ErrorInvalidDefaults, err)
	}
	return nil
}

// Return boolean value which indicates if a content owner parameter is missing
func (this *Params) IsEmptyContentOwner() bool {
	if this.ContentOwner == nil {
		return true
	}
	if len(*this.ContentOwner) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid content owner setting
func (this *Params) IsValidContentOwner() bool {
	if this.IsEmptyContentOwner() {
		return false
	}
	matched, _ := regexp.MatchString("^([a-zA-Z0-9]{22})$", *this.ContentOwner)
	return matched
}

// Return boolean value which indicates an empty channel
func (this *Params) IsEmptyChannel() bool {
	if this.Channel == nil {
		return true
	}
	if len(*this.Channel) == 0 {
		return true
	}
	return false
}

// Return boolean value whichindicates a valid content owner setting
func (this *Params) IsValidChannel() bool {
	if this.IsEmptyChannel() {
		return false
	}
	matched, _ := regexp.MatchString("^UC([a-zA-Z0-9\\-]{22})$", *this.Channel)
	return matched
}

// Return boolean value which indicates an empty query
func (this *Params) IsEmptyQuery() bool {
	if this.Query == nil {
		return true
	}
	if len(*this.Query) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates an empty video parameter
func (this *Params) IsEmptyVideo() bool {
	if this.Video == nil {
		return true
	}
	if len(*this.Video) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid video parameter
func (this *Params) IsValidVideo() bool {
	if this.IsEmptyVideo() {
		return false
	}
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\-]{11})$", *this.Video)
	return matched
}

// Return boolean value which indicates an empty stream parameter
func (this *Params) IsEmptyStream() bool {
	if this.Stream == nil {
		return true
	}
	if len(*this.Stream) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid stream parameter
func (this *Params) IsValidStream() bool {
	if this.IsEmptyStream() {
		return false
	}
	matched, _ := regexp.MatchString("^([a-zA-Z0-9]{4})\\-([a-zA-Z0-9]{4})\\-([a-zA-Z0-9]{4})\\-([a-zA-Z0-9]{4})$", *this.Stream)
	return matched
}

// Return boolean value which indicates an empty video parameter
func (this *Params) IsEmptyBroadcastStatus() bool {
	if this.BroadcastStatus == nil {
		return true
	}
	if len(*this.BroadcastStatus) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid video parameter
func (this *Params) IsValidBroadcastStatus() bool {
	if this.IsEmptyBroadcastStatus() {
		return false
	}
	matched, _ := regexp.MatchString("^(all|upcoming|completed|active)$", *this.BroadcastStatus)
	return matched
}

// Return boolean value which indicates an empty title parameter
func (this *Params) IsEmptyTitle() bool {
	if this.Title == nil {
		return true
	}
	if len(*this.Title) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates an empty title parameter
func (this *Params) IsEmptyDescription() bool {
	if this.Description == nil {
		return true
	}
	if len(*this.Description) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates an empty language parameter
func (this *Params) IsEmptyLanguage() bool {
	if this.Language == nil {
		return true
	}
	if len(*this.Language) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid language parameter
func (this *Params) IsValidLanguage() bool {
	if this.IsEmptyLanguage() {
		return false
	}
	matched, _ := regexp.MatchString("^([a-z]{2})$", *this.Language)
	if matched {
		return true
	}
	matched, _ = regexp.MatchString("^([a-z]{2})\\-([a-zA-Z0-9]+)$", *this.Language)
	if matched {
		return true
	}
	return false
}
