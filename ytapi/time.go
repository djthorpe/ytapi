/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	TimeNow   = regexp.MustCompile("^\\s*(now)\\s*$")
	InHours   = regexp.MustCompile("^\\s*(in)\\s+(\\d+)\\s+(h|hr|hour|hrs|hours)\\s*$")
	InMinutes = regexp.MustCompile("^\\s*(in)\\s+(\\d+)\\s+(m|min|minute|mins|minutes)\\s*$")
	InDays    = regexp.MustCompile("^\\s*(in)\\s+(\\d+)\\s+(d|day|day|days)\\s*$")
	YYYYMMDD  = regexp.MustCompile("^\\s*(\\d{4})-(\\d{2})-(\\d{2})\\s*$")
)

// Parse english-language dates and times. Returns the time,
// a boolean value indicating if the time-part is significant,
// and an error if the time could not be parsed, or nil
func ParseTime(value string) (time.Time, error) {

	// make lowercase
	value = strings.ToLower(value)

	// NOW
	if TimeNow.MatchString(value) {
		return time.Now(), nil
	}
	// IN NNN HOURS
	if submatch := InHours.FindStringSubmatch(value); len(submatch) >= 3 {
		duration, err := strconv.ParseInt(submatch[2], 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Now().Add(time.Duration(duration) * time.Hour), nil
	}
	// IN NNN MINUTES
	if submatch := InMinutes.FindStringSubmatch(value); len(submatch) >= 3 {
		duration, err := strconv.ParseInt(submatch[2], 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Now().Add(time.Duration(duration) * time.Minute), nil
	}
	// IN NNN DAYS
	if submatch := InDays.FindStringSubmatch(value); len(submatch) >= 3 {
		duration, err := strconv.ParseInt(submatch[2], 10, 32)
		if err != nil {
			return time.Time{}, err
		}
		// add days
		return time.Now().AddDate(0, 0, int(duration)), nil
	}

	// error
	return time.Time{}, errors.New("Cannot parse time value")
}
