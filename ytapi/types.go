/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/djthorpe/ytapi/util"
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
	FLAG_REGION
	FLAG_STREAM
	FLAG_CONTENTOWNER
	FLAG_TIME
)

////////////////////////////////////////////////////////////////////////////////

// Defines a flag
type Flag struct {
	Name        string
	Description string
	Type        uint32
	Extra       string
	Default     string
	Path        string
	Array       bool
	added       bool
}

////////////////////////////////////////////////////////////////////////////////
// Flag implementation

func (this *Flag) TypeString() string {
	var suffix string
	if this.Array == true {
		suffix = ",.."
	}
	switch {
	case this.Type == FLAG_STRING:
		return "string" + suffix
	case this.Type == FLAG_UINT:
		return "uint" + suffix
	case this.Type == FLAG_BOOL:
		return "bool" + suffix
	case this.Type == FLAG_ENUM:
		return this.Extra + suffix
	case this.Type == FLAG_VIDEO:
		return "video" + suffix
	case this.Type == FLAG_STREAM:
		return "stream" + suffix
	case this.Type == FLAG_CHANNEL:
		return "channel" + suffix
	case this.Type == FLAG_PLAYLIST:
		return "playlist" + suffix
	case this.Type == FLAG_LANGUAGE:
		return "language" + suffix
	case this.Type == FLAG_REGION:
		return "region" + suffix
	case this.Type == FLAG_CONTENTOWNER:
		return "contentowner" + suffix
	case this.Type == FLAG_TIME:
		return "datetime" + suffix
	default:
		return "other" + suffix
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

func (this *Flag) asVideo(value string) (string, error) {
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\-\\_]{11})$", value)
	if matched == false {
		return "", errors.New(fmt.Sprintf("Malformed video value: %s", value))
	}
	return value, nil
}

func (this *Flag) asStream(value string) (string, error) {
	matched, _ := regexp.MatchString("^([a-zA-Z0-9]{4})-([a-zA-Z0-9]{4})-([a-zA-Z0-9]{4})-([a-zA-Z0-9]{4})$", value)
	if matched {
		return value, nil
	}
	matched2, _ := regexp.MatchString("^([a-zA-Z0-9]{38})$", value)
	if matched2 {
		return value, nil
	}

	return "", errors.New("Malformed stream value")
}

func (this *Flag) asContentOwner(value string) (string, error) {
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\_\\-]{22})$", value)
	if matched == false {
		return "", errors.New("Malformed content owner value")
	}
	return value, nil
}

func (this *Flag) asChannel(value string) (string, error) {
	matched, _ := regexp.MatchString("^UC([a-zA-Z0-9\\-\\_]{22})$", value)
	if matched == false {
		return "", errors.New("Malformed channel value")
	}
	return value, nil
}

func (this *Flag) asPlaylist(value string) (string, error) {
	matched, _ := regexp.MatchString("^(UU|LL|WL|HL|FL)([a-zA-Z0-9\\-]{22})$", value)
	if matched {
		return value, nil
	}
	matched2, _ := regexp.MatchString("^PL([a-zA-Z0-9\\_\\-]{32})$", value)
	if matched2 {
		return value, nil
	}
	matched3, _ := regexp.MatchString("^PL([A-Z0-9]{16})$", value)
	if matched3 {
		return value, nil
	}
	return "", errors.New(fmt.Sprintf("Malformed playlist value: %s", value))
}

func (this *Flag) asLanguage(value string) (string, error) {
	matched, _ := regexp.MatchString("^([a-z]{2}[a-z]?)$", value)
	if matched {
		return value, nil
	}
	matched, _ = regexp.MatchString("^([a-z]{2})\\-([a-zA-Z0-9]+)$", value)
	if matched {
		return value, nil
	}
	return "", errors.New(fmt.Sprintf("Malformed language value: %s", value))
}

func (this *Flag) asRegion(value string) (string, error) {
	matched, _ := regexp.MatchString("^([A-Z]{2})$", value)
	if matched {
		return value, nil
	}
	return "", errors.New("Malformed region value")
}

func (this *Flag) asTime(value string) (time.Time, error) {
	datetime, err := util.ParseTime(value, false)
	if err != nil {
		return time.Time{}, err
	}
	return datetime, nil
}
