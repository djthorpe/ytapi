/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	Empty     = regexp.MustCompile("^\\s*$")
	TimeNow   = regexp.MustCompile("^\\s*(NOW)\\s*$")
	InHours   = regexp.MustCompile("^\\s*(IN)\\s+(\\d+)\\s+(H|HR|HOUR|HRS|HOURS)\\s*$")
	InMinutes = regexp.MustCompile("^\\s*(IN)\\s+(\\d+)\\s+(M|MIN|MINUTE|MINS|MINUTES)\\s*$")
	InDays    = regexp.MustCompile("^\\s*(IN)\\s+(\\d+)\\s+(D|DAY|DAYS)\\s*$")
	YYYYMMDD  = regexp.MustCompile("^\\s*(\\d{4})-(\\d{2})-(\\d{2})\\s*$")
	ISO       = time.RFC3339Nano
)

// Parse english-language dates and times. Returns the time,
// a boolean value indicating if the time-part is significant,
// and an error if the time could not be parsed, or nil
// The 'notnull' boolean flag can be used to allow empty string,
// in which case time.Time{} is returned
func ParseTime(value string, notnull bool) (time.Time, error) {

	// EMPTY
	if (notnull == false) && Empty.MatchString(value) {
		return time.Time{}, nil
	}

	// make lowercase
	value = strings.ToUpper(value)

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
	// ISO format
	t, err := time.Parse(ISO, value)
	if err == nil {
		return t, nil
	}

	// error
	return time.Time{}, errors.New(fmt.Sprint("Cannot parse time value: \"", value, "\""))
}

func ParseDuration(value string, notnull bool) (time.Duration, error) {
	value = strings.TrimSpace(value)
	if notnull == false && value == "" {
		return 0, nil
	}
	return time.ParseDuration(value)
}
