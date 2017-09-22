package ytapi

/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
)

////////////////////////////////////////////////////////////////////////////////

// Value structure defines a generic value
type Value struct {
	v_string   string
	v_uint     uint64
	v_float    float64
	v_bool     bool
	v_time     time.Time
	v_duration time.Duration
	is_set     bool
	flag       *Flag
}

// Values structure defines a set of values
type Values struct {
	values map[*Flag]*Value
}

////////////////////////////////////////////////////////////////////////////////
// Value implementation

func NewValue(flag *Flag, value reflect.Value) (*Value, error) {
	this := new(Value)
	this.flag = flag

	switch value.Kind() {
	case reflect.String:
		if err := this.SetString(value.String()); err != nil {
			return nil, err
		}
	case reflect.Bool:
		if err := this.SetBool(value.Bool()); err != nil {
			return nil, err
		}
	case reflect.Uint64:
		if err := this.SetUint(value.Uint()); err != nil {
			return nil, err
		}
	case reflect.Int64:
		if err := this.SetInt(value.Int()); err != nil {
			return nil, err
		}
	case reflect.Float64:
		if err := this.SetFloat(value.Float()); err != nil {
			return nil, err
		}
	case reflect.Slice:
		if err := this.SetString(value.String()); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(fmt.Sprint("Invalid value kind in NewValue: ", value.Kind()))
	}

	return this, nil
}

func NewValueWithString(flag *Flag, value string) (*Value, error) {
	this := new(Value)
	this.flag = flag

	if err := this.SetString(value); err != nil {
		return nil, err
	}

	return this, nil
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
		if value == 0 {
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
		if value == 0 {
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

func (this *Value) SetFloat(value float64) error {
	// TODO
	this.v_string = fmt.Sprint(value)
	this.is_set = true
	return nil
}

func (this *Value) SetString(value string) error {
	var err error

	// Don't set if the string value is empty
	if value == "" && this.flag.Type != FLAG_STRING {
		return nil
	}

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
		this.v_string, err = this.flag.asVideo(value)
		break
	case this.flag.Type == FLAG_CHANNEL:
		this.v_string, err = this.flag.asChannel(value)
		break
	case this.flag.Type == FLAG_STREAM:
		this.v_string, err = this.flag.asStream(value)
		break
	case this.flag.Type == FLAG_PLAYLIST:
		this.v_string, err = this.flag.asPlaylist(value)
		break
	case this.flag.Type == FLAG_LANGUAGE:
		this.v_string, err = this.flag.asLanguage(value)
		break
	case this.flag.Type == FLAG_REGION:
		this.v_string, err = this.flag.asRegion(value)
		break
	case this.flag.Type == FLAG_CONTENTOWNER:
		this.v_string, err = this.flag.asContentOwner(value)
		break
	case this.flag.Type == FLAG_TIME:
		this.v_time, err = this.flag.asTime(value)
		break
	case this.flag.Type == FLAG_DURATION:
		this.v_duration, err = this.flag.asDuration(value)
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
	}
	return this.flag.Default
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

func (this *Value) Duration() time.Duration {
	if this.flag.Type != FLAG_DURATION {
		panic("Not a Duration type")
	}
	if this.is_set {
		return this.v_duration
	}
	if this.flag.Default != "" {
		value, err := this.flag.asDuration(this.flag.Default)
		if err != nil {
			panic(err)
		}
		return value
	}
	return 0
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

func (this *Values) Set(value *Value) *Value {
	this.values[value.flag] = value
	return value
}

func (this *Values) Get(flag *Flag) *Value {
	return this.values[flag]
}

func (this *Values) IsSet(flag *Flag) bool {
	v, exists := this.values[flag]
	if exists == false {
		return false
	}
	return v.is_set
}

func (this *Values) IsKindOf(flag *Flag, kind int) bool {
	if this.IsSet(flag) == false {
		return false
	}
	value := this.GetString(flag)
	switch kind {
	case FLAG_VIDEO:
		if _, err := flag.asVideo(value); err != nil {
			return false
		}
	case FLAG_UINT:
		if _, err := flag.asUint(value); err != nil {
			return false
		}
	case FLAG_BOOL:
		if _, err := flag.asBool(value); err != nil {
			return false
		}
	case FLAG_ENUM:
		if _, err := flag.asBool(value); err != nil {
			return false
		}
	case FLAG_CHANNEL:
		if _, err := flag.asChannel(value); err != nil {
			return false
		}
	case FLAG_PLAYLIST:
		if _, err := flag.asPlaylist(value); err != nil {
			return false
		}
	case FLAG_LANGUAGE:
		if _, err := flag.asLanguage(value); err != nil {
			return false
		}
	case FLAG_REGION:
		if _, err := flag.asRegion(value); err != nil {
			return false
		}
	case FLAG_STREAM:
		if _, err := flag.asStream(value); err != nil {
			return false
		}
	case FLAG_CONTENTOWNER:
		if _, err := flag.asContentOwner(value); err != nil {
			return false
		}
	case FLAG_TIME:
		if _, err := flag.asTime(value); err != nil {
			return false
		}
	case FLAG_DURATION:
		if _, err := flag.asDuration(value); err != nil {
			return false
		}
	}
	return true
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
		panic(fmt.Sprint("Missing flag value: ", flag.Name))
	}
	return value.String()
}

func (this *Values) GetStringArray(flag *Flag) []string {
	value, exists := this.values[flag]
	if exists == false {
		panic(fmt.Sprint("Missing flag value: ", flag.Name))
	}
	return strings.FieldsFunc(value.String(), func(c rune) bool {
		return !unicode.IsSpace(c)
	})
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

func (this *Values) GetDuration(flag *Flag) time.Duration {
	value, exists := this.values[flag]
	if exists == false {
		panic("Missing flag")
	}
	return value.Duration()
}

func (this *Values) GetTimeInISOFormat(flag *Flag) string {
	value := this.GetTime(flag).UTC()
	if value.Equal(time.Time{}) {
		return ""
	}
	return value.Format(time.RFC3339Nano)
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
