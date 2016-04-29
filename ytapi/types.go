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
}

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
	case this.Type == FLAG_STREAM:
		return "stream"
	case this.Type == FLAG_CHANNEL:
		return "channel"
	case this.Type == FLAG_PLAYLIST:
		return "playlist"
	case this.Type == FLAG_LANGUAGE:
		return "language"
	case this.Type == FLAG_REGION:
		return "region"
	case this.Type == FLAG_CONTENTOWNER:
		return "contentowner"
	case this.Type == FLAG_TIME:
		return "datetime"
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
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\-\\_]{11})$", value)
	if matched == false {
		return "", errors.New(fmt.Sprintf("Malformed video value: %s",value))
	}
	return Video(value), nil
}

func (this *Flag) asStream(value string) (Stream, error) {
	matched, _ := regexp.MatchString("^([a-zA-Z0-9]{4})-([a-zA-Z0-9]{4})-([a-zA-Z0-9]{4})-([a-zA-Z0-9]{4})$", value)
	if matched == false {
		return "", errors.New("Malformed stream value")
	}
	return Stream(value), nil
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
	matched, _ := regexp.MatchString("^(PL|UU|LL|WL|HL)([a-zA-Z0-9]{22})$", value)
	if matched {
		return Playlist(value), nil
	}
	return "", errors.New(fmt.Sprintf("Malformed playlist value: %s",value))
}

func (this *Flag) asLanguage(value string) (Language, error) {
	matched, _ := regexp.MatchString("^([a-z]{2}[a-z]?)$", value)
	if matched {
		return Language(value), nil
	}
	matched, _ = regexp.MatchString("^([a-z]{2})\\-([a-zA-Z0-9]+)$", value)
	if matched {
		return Language(value), nil
	}
	return "", errors.New(fmt.Sprintf("Malformed language value: %s",value))
}

func (this *Flag) asRegion(value string) (Region, error) {
	matched, _ := regexp.MatchString("^([A-Z]{2})$", value)
	if matched {
		return Region(value), nil
	}
	return "", errors.New("Malformed region value")
}

func (this *Flag) asTime(value string) (time.Time, error) {
	datetime, err := util.ParseTime(value)
	if err != nil {
		return time.Time{}, err
	}
	return datetime, nil
}

