/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"reflect"
)

////////////////////////////////////////////////////////////////////////////////

// Types
type Video string
type Channel string
type Playlist string
type Language string
type Region string
type Stream string
type ContentOwner string

// Value
type Value struct {
	v_string       string
	v_uint         uint64
	v_bool         bool
	v_video        Video
	v_stream       Stream
	v_channel      Channel
	v_playlist     Playlist
	v_language     Language
	v_region       Region
	v_contentowner ContentOwner
	v_time         time.Time
	is_set         bool
	flag           *Flag
}

// Values
type Values struct {
	values map[*Flag]*Value
}

// Defaults
type Defaults struct {
	ContentOwner string `json:"contentowner,omitempty"`
	Channel      string `json:"channel,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// Value implementation

func NewValue(flag *Flag,value reflect.Value) (*Value,error) {
	this := new(Value)
	this.flag = flag

	switch(value.Kind()) {
		case reflect.String:
			if err := this.Set(value.String()); err != nil {
				return nil,err
			}
		case reflect.Bool:
			return nil,errors.New("Bool value in NewValue")
		default:
			return nil,errors.New("Invalid value in NewValue")
	}

	return this,nil
}

func (this *Value) Set(value string) error {
	var err error

	// Set string value
	this.v_string = value

	// Set other values
	switch {
	case this.flag.Type == FLAG_UINT:
		this.v_uint, err = this.flag.asUint(value)
		break
	case this.flag.Type == FLAG_BOOL:
		this.v_bool, err = this.flag.asBool(value)
		break
	case this.flag.Type == FLAG_ENUM:
		this.v_string, err = this.flag.asEnum(value)
		break
	case this.flag.Type == FLAG_VIDEO:
		this.v_video, err = this.flag.asVideo(value)
		break
	case this.flag.Type == FLAG_CHANNEL:
		this.v_channel, err = this.flag.asChannel(value)
		break
	case this.flag.Type == FLAG_STREAM:
		this.v_stream, err = this.flag.asStream(value)
		break
	case this.flag.Type == FLAG_PLAYLIST:
		this.v_playlist, err = this.flag.asPlaylist(value)
		break
	case this.flag.Type == FLAG_LANGUAGE:
		this.v_language, err = this.flag.asLanguage(value)
		break
	case this.flag.Type == FLAG_REGION:
		this.v_region, err = this.flag.asRegion(value)
		break
	case this.flag.Type == FLAG_CONTENTOWNER:
		this.v_contentowner, err = this.flag.asContentOwner(value)
		break
	case this.flag.Type == FLAG_TIME:
		this.v_time, err = this.flag.asTime(value)
		break
	}

	if err == nil {
		this.is_set = true
	}

	return err
}

func (this *Value) String() string {
	if this.is_set {
		return this.v_string
	} else {
		return this.flag.Default
	}
}

func (this *Value) Bool() bool {
	if this.flag.Type != FLAG_BOOL {
		panic("Not a bool type")
	}
	if this.is_set {
		return this.v_bool
	}
	if this.flag.Default != "" {
		value, err := this.flag.asBool(this.flag.Default)
		if err != nil {
			panic(err)
		}
		return value
	}
	return false
}

func (this *Value) Uint64() uint64 {
	if this.flag.Type != FLAG_UINT {
		panic("Not a Uint type")
	}
	if this.is_set {
		return this.v_uint
	}
	if this.flag.Default != "" {
		value, err := this.flag.asUint(this.flag.Default)
		if err != nil {
			panic(err)
		}
		return value
	}
	return 0
}

func (this *Value) Time() time.Time {
	if this.flag.Type != FLAG_TIME {
		panic("Not a Time type")
	}
	if this.is_set {
		return this.v_time
	}
	if this.flag.Default != "" {
		value, err := this.flag.asTime(this.flag.Default)
		if err != nil {
			panic(err)
		}
		return value
	}
	return time.Time{}
}

func (this *Value) IsBoolFlag() bool {
	return (this.flag.Type == FLAG_BOOL)
}

func (this *Value) IsSet() bool {
	return this.is_set
}

////////////////////////////////////////////////////////////////////////////////
// Values implementation

func NewValues() *Values {
	this := new(Values)
	this.values = make(map[*Flag]*Value, 0)
	return this
}

func (this *Values) ReadDefaultsFromFile(filename string) error {
	var err error
	// if a file exists, then read it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// file doesn't exist', so just return
		return nil
	}
	// read in the file
	bytes, err := ioutil.ReadFile(filename)
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
		err = this.SetDefault(&FlagContentOwner, string(defaults.ContentOwner))
	}
	// Channel
	if err == nil && defaults.Channel != "" {
		err = this.SetDefault(&FlagChannel, string(defaults.Channel))
	}
	if err != nil {
		return err
	}
	// success
	return nil
}

func (this *Values) WriteDefaultsToFile(filename string, perm os.FileMode) error {
	defaults := &Defaults{}
	if this.IsSet(&FlagContentOwner) {
		defaults.ContentOwner = this.GetString(&FlagContentOwner)
	}
	if this.IsSet(&FlagContentOwner) {
		defaults.Channel = this.GetString(&FlagChannel)
	}
	json, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, json, perm)
	if err != nil {
		return err
	}
	return nil
}

func (this *Values) Set(value *Value) {
	this.values[value.flag] = value
}

func (this *Values) IsSet(flag *Flag) bool {
	v, exists := this.values[flag]
	if exists == false {
		return false
	}
	return v.is_set
}

func (this *Values) SetDefault(flag *Flag, value string) error {
	v, exists := this.values[flag]
	if exists == false {
		return errors.New(fmt.Sprint("SetDefault: Invalid flag: ", flag.Name))
	}
	if v.is_set {
		return nil
	}
	return v.Set(value)
}

func (this *Values) GetString(flag *Flag) string {
	value, exists := this.values[flag]
	if exists == false {
		panic("Missing flag")
	}
	return value.String()
}

func (this *Values) GetBool(flag *Flag) bool {
	value, exists := this.values[flag]
	if exists == false {
		panic("Missing flag")
	}
	return value.Bool()
}

func (this *Values) GetUint(flag *Flag) uint64 {
	value, exists := this.values[flag]
	if exists == false {
		panic("Missing flag")
	}
	return value.Uint64()
}

func (this *Values) GetTime(flag *Flag) time.Time {
	value, exists := this.values[flag]
	if exists == false {
		panic("Missing flag")
	}
	return value.Time()
}

func (this *Values) GetTimeInISOFormat(flag *Flag) string {
	value := this.GetTime(flag)
	if value.Equal(time.Time{}) {
		return ""
	} else {
		return value.Format(time.RFC3339)
	}
}

func (this *Values) SetFields(fieldmap map[string]*Flag) []string {
	fields := make([]string, len(fieldmap))
	for k, flag := range fieldmap {
		if this.IsSet(flag) {
			fields = append(fields, k)
		}
	}
	return fields
}
