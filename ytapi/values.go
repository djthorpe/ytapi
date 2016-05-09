/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"errors"
	"fmt"
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

////////////////////////////////////////////////////////////////////////////////
// Value implementation

func NewValue(flag *Flag,value reflect.Value) (*Value,error) {
	this := new(Value)
	this.flag = flag

	switch(value.Kind()) {
		case reflect.String:
			if err := this.SetString(value.String()); err != nil {
				return nil,err
			}
		case reflect.Bool:
			if err := this.SetBool(value.Bool()); err != nil {
				return nil,err
			}
		case reflect.Uint64:
			if err := this.SetUint(value.Uint()); err != nil {
				return nil,err
			}
		case reflect.Int64:
			if err := this.SetInt(value.Int()); err != nil {
				return nil,err
			}
		case reflect.Slice:
			if err := this.SetString(value.String()); err != nil {
				return nil,err
			}
		default:
			return nil,errors.New(fmt.Sprint("Invalid value kind in NewValue: ",value.Kind()))
	}

	return this,nil
}

func (this *Value) SetBool(value bool) error {
	switch {
	case this.flag.Type == FLAG_UINT:
		if value {
			this.v_uint = 1
		} else {
			this.v_uint = 0
		}
		break
	case this.flag.Type == FLAG_BOOL:
		this.v_bool = value
		break
	default:
		panic("Calling SetBool with invalid type")
	}
	this.v_string = fmt.Sprint(value)
	this.is_set = true
	return nil
}

func (this *Value) SetUint(value uint64) error {
	switch {
	case this.flag.Type == FLAG_UINT:
		this.v_uint = value
		break
	case this.flag.Type == FLAG_BOOL:
		if value==0 {
			this.v_bool = false
		} else {
			this.v_bool = true
		}
		break
	default:
		panic("Calling SetUint with invalid type")
	}
	this.v_string = fmt.Sprint(value)
	this.is_set = true
	return nil
}

func (this *Value) SetInt(value int64) error {
	switch {
	case this.flag.Type == FLAG_UINT:
		if value < 0 {
			return errors.New("SetInt: Setting negative int value in uint")
		}
		this.v_uint = uint64(value)
		break
	case this.flag.Type == FLAG_BOOL:
		if value==0 {
			this.v_bool = false
		} else {
			this.v_bool = true
		}
		break
	default:
		panic("Calling SetInt with invalid type")
	}
	this.v_string = fmt.Sprint(value)
	this.is_set = true
	return nil
}

func (this *Value) SetString(value string) error {
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

// Set conforms Value to the flag interface
func (this *Value) Set(value string) error {
	return this.SetString(value)
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

func (this *Values) Set(value *Value) (*Value) {
	this.values[value.flag] = value
    return value
}

func (this *Values) Get(flag *Flag) (*Value) {
    return this.values[flag]
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
		panic(fmt.Sprint("Missing flag value: ",flag.Name))
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

func (this *Values) GetInt(flag *Flag) int64 {
	value, exists := this.values[flag]
	if exists == false {
		panic("Missing flag")
	}
	return int64(value.Uint64())
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
