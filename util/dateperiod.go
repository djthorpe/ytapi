/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"regexp"
	"strings"
	"time"
	"errors"
	"strconv"
)

var (
	Year = regexp.MustCompile("^\\s*(\\d{4})\\s*$") // YYYY
	YearMonth = regexp.MustCompile("^\\s*(\\d{4})-(\\d{1,2})\\s*$") // YYYY-MM
	MonthYear = regexp.MustCompile("^\\s*(\\d{1,2})-(\\d{4})\\s*$") // MM-YYYY
	NullTime = time.Time{}
)


/* TO SUPPORT:
lifetime
thisWeek
lastWeek
last7Days
first7Days
thisMonth
lastMonth
last28Days
last30Days
first28Days
thisQuarter
lastQuarter
last90Days
first90Days
thisYear
lastYear
last365Days
first365Days
<month>
*/


func parseYearMonth(year_value string,month_value string) (time.Time, time.Time, error) {
	var year, month int64
	var err error

	// Compute start date
	if year,err = strconv.ParseInt(year_value,10,0); err != nil {
		return NullTime,NullTime,err
	}
	if month,err = strconv.ParseInt(month_value,10,0); err != nil {
		return NullTime,NullTime,err
	}
	// Sanity check year and month
	if year < 2000 || month < 1 || month > 12 {
		return NullTime,NullTime,errors.New("Invalid year or month value")
	}
	// Set start and end times
	start_time := time.Date(int(year),time.Month(month),1,0,0,0,0,time.Local)
	end_time := start_time.AddDate(0,1,-1)
	// Return success
	return start_time,end_time,nil
}

func parseYear(year_value string) (time.Time, time.Time, error) {
	var year int64
	var err error

	// Compute start date
	if year,err = strconv.ParseInt(year_value,10,0); err != nil {
		return NullTime,NullTime,err
	}
	// Sanity check year and month
	if year < 2000 {
		return NullTime,NullTime,errors.New("Invalid year value")
	}
	// Set start and end times
	start_time := time.Date(int(year),time.Month(1),1,0,0,0,0,time.Local)
	end_time := start_time.AddDate(1,0,-1)
	// Return success
	return start_time,end_time,nil
}

////////////////////////////////////////////////////////////////////////////////

func ParseDatePeriod(value string) (time.Time, time.Time, error) {
	// YYYY-MM
	if s := YearMonth.FindStringSubmatch(strings.ToUpper(value)); len(s) == 3 {
		// Return
		return parseYearMonth(s[1],s[2])
	}

	// MM-YYYY
	if s := MonthYear.FindStringSubmatch(strings.ToUpper(value)); len(s) == 3 {
		// Return
		return parseYearMonth(s[2],s[1])
	}

	// YYYY
	if s := Year.FindStringSubmatch(strings.ToUpper(value)); len(s) == 2 {
		// Return
		return parseYear(s[1])
	}

	return NullTime,NullTime,errors.New("Unable to interpret date period")
}
