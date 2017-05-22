/*
  Copyright David Thorpe 2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytreporting

import (
	"errors"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"	
)

////////////////////////////////////////////////////////////////////////////////
// Register chat commands

func RegisterAnalyticsCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ChannelAnalytics",
			Description: "Execute Channel Analytics Query",
			Required:    []*ytapi.Flag{&ytapi.FlagAnalyticsMetrics,&ytapi.FlagAnalyticsPeriod},
			Optional:    []*ytapi.Flag{&ytapi.FlagAnalyticsDimensions,&ytapi.FlagAnalyticsFilter,&ytapi.FlagAnalyticsSort},
			Execute:     RunChannelAnalyticsQuery,
		},
		&ytapi.Command{
			Name:        "ContentOwnerAnalytics",
			Description: "Execute Content Owner Analytics Query",
			Required:    []*ytapi.Flag{&ytapi.FlagContentOwner,&ytapi.FlagAnalyticsMetrics,&ytapi.FlagAnalyticsPeriod},
			Optional:    []*ytapi.Flag{&ytapi.FlagAnalyticsDimensions,&ytapi.FlagAnalyticsFilter,&ytapi.FlagAnalyticsSort},
			Execute:     RunContentOwnerAnalyticsQuery,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Time Periods

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
<year>
<month>-<year>
<year>-<month>
*/

func getTimePeriod(value string) (string, string, error) {
	return  "2017-01-01", "2017-03-31", nil
}

////////////////////////////////////////////////////////////////////////////////
// Run Analytics Query

func RunChannelAnalyticsQuery(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	metrics := values.GetString(&ytapi.FlagAnalyticsMetrics)
	channel := "channel==MINE"
	if values.IsSet(&ytapi.FlagChannel) {
		channel = "channel==" + values.GetString(&ytapi.FlagChannel)
	}
	start_date, end_date, err := getTimePeriod(values.GetString(&ytapi.FlagAnalyticsPeriod))
	if err != nil {
		return err
	}

	// Create the call
	call := service.AAPI.Reports.Query(channel,start_date,end_date,metrics)

	/// Set parameters
	if values.IsSet(&ytapi.FlagAnalyticsDimensions) {
		call = call.Dimensions(values.GetString(&ytapi.FlagAnalyticsDimensions))
	}
	if values.IsSet(&ytapi.FlagAnalyticsFilter) {
		call = call.Filters(values.GetString(&ytapi.FlagAnalyticsFilter))
	}
	if values.IsSet(&ytapi.FlagAnalyticsCurrency) {
		call = call.Currency(values.GetString(&ytapi.FlagAnalyticsCurrency))
	}
	if values.IsSet(&ytapi.FlagAnalyticsSort) {
		call = call.Sort(values.GetString(&ytapi.FlagAnalyticsSort))
	}

	// Execute the call
	response, err := call.Do()
	if err != nil {
		return err
	}

    // Make parts for columns
    parts := make(map[string][]*ytapi.Flag)
    keys := make([]string,0)
    columns := make([]string,0)
    for _, column := range response.ColumnHeaders {
        var flags []*ytapi.Flag

        // retrieve the flags for this column type
        flags,exists := parts[column.ColumnType]
        if exists == false {
            flags = make([]*ytapi.Flag,0)
            parts[column.ColumnType] = flags
            keys = append(keys,column.ColumnType)
        }

        // append the flags and column names
        parts[column.ColumnType] = append(parts[column.ColumnType],&ytapi.Flag{ Name: column.Name, Type: ytapi.FLAG_STRING })
        columns = append(columns,column.Name)
    }

    // Register table columns
    for _, key := range keys {
        table.RegisterPart(key,parts[key])
    }
    table.SetColumns(columns)

    // Append columns
    if err = table.Append(response.Rows); err != nil {
        return err
    }

	// success
	return nil
}

func RunContentOwnerAnalyticsQuery(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {


	// Get parameters
	if values.IsSet(&ytapi.FlagContentOwner) == false {
		return errors.New("Missing -contentowner flag")
	}
	content_owner := "contentOwner==" + values.GetString(&ytapi.FlagContentOwner)
	start_date, end_date, err := getTimePeriod(values.GetString(&ytapi.FlagAnalyticsPeriod))
	if err != nil {
		return err
	}

	// Create the call
	call := service.AAPI.Reports.Query(content_owner,start_date,end_date,values.GetString(&ytapi.FlagAnalyticsMetrics))

	/// Set parameters
	if values.IsSet(&ytapi.FlagAnalyticsDimensions) {
		call = call.Dimensions(values.GetString(&ytapi.FlagAnalyticsDimensions))
	}
	if values.IsSet(&ytapi.FlagAnalyticsFilter) {
		call = call.Filters(values.GetString(&ytapi.FlagAnalyticsFilter))
	}
	if values.IsSet(&ytapi.FlagAnalyticsCurrency) {
		call = call.Currency(values.GetString(&ytapi.FlagAnalyticsCurrency))
	}
	if values.IsSet(&ytapi.FlagAnalyticsSort) {
		call = call.Sort(values.GetString(&ytapi.FlagAnalyticsSort))
	}

	// Execute the call
	response, err := call.Do()
	if err != nil {
		return err
	}

    // Make parts for columns
    parts := make(map[string][]*ytapi.Flag)
    keys := make([]string,0)
    columns := make([]string,0)
    for _, column := range response.ColumnHeaders {
        var flags []*ytapi.Flag

        // retrieve the flags for this column type
        flags,exists := parts[column.ColumnType]
        if exists == false {
            flags = make([]*ytapi.Flag,0)
            parts[column.ColumnType] = flags
            keys = append(keys,column.ColumnType)
        }

        // append the flags and column names
        parts[column.ColumnType] = append(parts[column.ColumnType],&ytapi.Flag{ Name: column.Name, Type: ytapi.FLAG_STRING })
        columns = append(columns,column.Name)
    }

    // Register table columns
    for _, key := range keys {
        table.RegisterPart(key,parts[key])
    }
    table.SetColumns(columns)

    // Append columns
    if err = table.Append(response.Rows); err != nil {
        return err
    }

	// success
	return nil
}

