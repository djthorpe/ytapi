/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
    "time"
    "regexp"
    "errors"
    "strconv"
)

var (
    TimeNow = regexp.MustCompile("^\\s*(now)\\s*$")
    InHours = regexp.MustCompile("^\\s*(in)\\s+(\\d+)\\s+(h|hr|hour|hrs|hours)\\s*$")
    InMinutes = regexp.MustCompile("^\\s*(in)\\s+(\\d+)\\s+(m|min|minute|mins|minutes)\\s*$")
    InDays = regexp.MustCompile("^\\s*(in)\\s+(\\d+)\\s+(d|day|day|days)\\s*$")
    YYYYMMDD = regexp.MustCompile("^\\s*(\\d{4})-(\\d{2})-(\\d{2})\\s*$")
)

// Parse english-language dates and times. Returns the time,
// a boolean value indicating if the time-part is significant,
// and an error if the time could not be parsed, or nil
func ParseTime(value string) (time.Time,bool,error) {
    // NOW
    if TimeNow.MatchString(value) {
        return time.Now(),true,nil
    }
    // IN NNN HOURS
    if submatch := InHours.FindStringSubmatch(value);  len(submatch) >= 3 {
        duration,err := strconv.ParseInt(submatch[2],10,64)
        if err != nil {
            return time.Time{},false,err
        }
        return time.Now().Add(time.Duration(duration) * time.Hour),true,nil
    }
    // IN NNN MINUTES
    if submatch := InMinutes.FindStringSubmatch(value);  len(submatch) >= 3 {
        duration,err := strconv.ParseInt(submatch[2],10,64)
        if err != nil {
            return time.Time{},false,err
        }
        return time.Now().Add(time.Duration(duration) * time.Minute),true,nil
    }
    // IN NNN DAYS
    if submatch := InDays.FindStringSubmatch(value);  len(submatch) >= 3 {
        duration,err := strconv.ParseInt(submatch[2],10,64)
        if err != nil {
            return time.Time{},false,err
        }
        // add days
        return time.Now().AddDate(0,0,duration),false,nil
    }

    // error
    return time.Time{},false,errors.New("Cannot parse time value")
}

// Parse start time. When the time isn't significant, the time is set to
// 00:00 local time
func ParseStartTime(value string) (time.Time,error) {
    datetime,time_set,err := ParseTime(value)
    if time_set == false {
        
    }
}
