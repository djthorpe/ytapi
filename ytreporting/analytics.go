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
			Name:        "ChannelAnalyticsQuery",
			Description: "Execute Channel Analytics Query",
			Required:    []*ytapi.Flag{&ytapi.FlagAnalyticsMetrics},
			Optional:    []*ytapi.Flag{&ytapi.FlagAnalyticsDimensions},
			Execute:     RunChannelAnalyticsQuery,
		},
		&ytapi.Command{
			Name:        "ContentOwnerAnalyticsQuery",
			Description: "Execute Content Owner Analytics Query",
			Required:    []*ytapi.Flag{&ytapi.FlagContentOwner,&ytapi.FlagAnalyticsMetrics},
			Optional:    []*ytapi.Flag{&ytapi.FlagAnalyticsDimensions},
			Execute:     RunContentOwnerAnalyticsQuery,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Run Analytics Query

func RunChannelAnalyticsQuery(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	start_date := "2017-01-01"
	end_date := "2017-03-31"

	// Get parameters
	metrics := values.GetString(&ytapi.FlagAnalyticsMetrics)
	channel := "channel==MINE"
	if values.IsSet(&ytapi.FlagChannel) {
		channel = "channel==" + values.GetString(&ytapi.FlagChannel)
	}

	// Create the call
	call := service.AAPI.Reports.Query(channel,start_date,end_date,metrics)
	if values.IsSet(&ytapi.FlagAnalyticsDimensions) {
		call = call.Dimensions(values.GetString(&ytapi.FlagAnalyticsDimensions))
	}

	// Execute the call
	response, err := call.Do()
	if err != nil {
		return err
	}

    // Make parts for columns
    parts := make(map[string][]*ytapi.Flag)
    keys := make([]string,0)
    for _, column := range response.ColumnHeaders {
        var flags []*ytapi.Flag

        // retrieve the flags for this column type
        flags,exists := parts[column.ColumnType]
        if exists == false {
            flags = make([]*ytapi.Flag,0)
            parts[column.ColumnType] = flags
            keys = append(keys,column.ColumnType)
        }

        // append the flags
        parts[column.ColumnType] = append(parts[column.ColumnType],&ytapi.Flag{ Name: column.Name, Type: ytapi.FLAG_STRING })
    }

    // Register with table
    for _, key := range keys {
        table.RegisterPart(key,parts[key])
    }

    table.Append(response.Rows)

	// success
	return nil
}

func RunContentOwnerAnalyticsQuery(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	if values.IsSet(&ytapi.FlagContentOwner) == false {
		return errors.New("Missing -contentowner flag")
	}

	start_date := "2017-01-01"
	end_date := "2017-03-31"

	// Get parameters
	metrics := values.GetString(&ytapi.FlagAnalyticsMetrics)
	content_owner := "contentOwner==" + values.GetString(&ytapi.FlagContentOwner)

	// Create the call
	call := service.AAPI.Reports.Query(content_owner,start_date,end_date,metrics)
	if values.IsSet(&ytapi.FlagAnalyticsDimensions) {
		call = call.Dimensions(values.GetString(&ytapi.FlagAnalyticsDimensions))
	}

	// Execute the call
	_, err := call.Do()
	if err != nil {
		return err
	}
	// success
	return nil
}

