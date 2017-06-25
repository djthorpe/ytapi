package cidcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
)

import (
	"github.com/djthorpe/ytapi/youtubepartner/v1"
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterCuepointsCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:           "InsertCuepoint",
			Description:    "Inserts a cuepoint into a live broadcast",
			ServiceAccount: true,
			Required:       []*ytapi.Flag{&ytapi.FlagVideo},
			Optional:       []*ytapi.Flag{&ytapi.FlagCuepointOffset, &ytapi.FlagCuepointDuration, &ytapi.FlagCuepointTime},
			Setup:          RegisterCuepointFormat,
			Execute:        InsertCuepoint,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register content owner format

func RegisterCuepointFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "id", Path: "Id", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "kind", Path: "Kind", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "video", Path: "BroadcastId", Type: ytapi.FLAG_VIDEO},
	})

	table.RegisterPart("settings", []*ytapi.Flag{
		&ytapi.Flag{Name: "offset_ms", Path: "Settings/OffsetTimeMs", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "walltime", Path: "Settings/Walltime", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "type", Path: "Settings/CueType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "duration_secs", Path: "Settings/DurationSecs", Type: ytapi.FLAG_UINT},
	})

	// set default columns
	table.SetColumns([]string{"id", "video", "type", "duration_secs"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Content Owners

func InsertCuepoint(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	if values.IsSet(&ytapi.FlagChannel) == false {
		return errors.New("Missing channel value")
	}
	call := service.PAPI.LiveCuepoints.Insert(values.GetString(&ytapi.FlagChannel), &youtubepartner.LiveCuepoint{
		BroadcastId: values.GetString(&ytapi.FlagVideo),
		Settings: &youtubepartner.CuepointSettings{
			CueType:      "ad",
			OffsetTimeMs: values.GetInt(&ytapi.FlagCuepointOffset),
			Walltime:     values.GetTimeInISOFormat(&ytapi.FlagCuepointTime),
			DurationSecs: values.GetInt(&ytapi.FlagCuepointDuration),
		},
	})
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}
	if err = table.Append([]*youtubepartner.LiveCuepoint{response}); err != nil {
		return err
	}
	return nil
}
