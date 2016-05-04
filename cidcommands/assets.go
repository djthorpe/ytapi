/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package cidcommands

import (
	"errors"
	
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
    "github.com/djthorpe/ytapi/youtubepartner/v1"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterAssetCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "GetAsset",
			Description: "Get a single asset",
            Required:    []*ytapi.Flag{ &ytapi.FlagAsset },
			Setup:       RegisterAssetFormat,
			Execute:     GetAsset,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register format

func RegisterAssetFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "asset", Path: "Id", Type: ytapi.FLAG_STRING},
        &ytapi.Flag{Name: "type", Path: "Type", Type: ytapi.FLAG_STRING},
        &ytapi.Flag{Name: "timeCreated", Path: "TimeCreated", Type: ytapi.FLAG_TIME},
	})

	// set default columns
	table.SetColumns([]string{ "asset","type","timeCreated" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Assets

func GetAsset(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	if service.ServiceAccount == false {
		return errors.New("No service account authenticated")
	}

	// create call and set parameters
	call := service.PAPI.Assets.Get(values.GetString(&ytapi.FlagAsset))
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

    // get assets
    response, err := call.Do()
    if err != nil {
        return err
    }
    if err = table.Append([]*youtubepartner.Asset{ response }); err != nil {
        return err
    }

	// Success
	return nil
}

