/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package cidcommands

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
    "github.com/djthorpe/ytapi/youtubepartner/v1"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterAssetCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "GetAsset",
			Description: "Get a single asset",
            ServiceAccount: true,
            Required:    []*ytapi.Flag{ &ytapi.FlagAsset },
            Optional:    []*ytapi.Flag{ &ytapi.FlagAssetFilter },
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

    // TODO: part should be "effective" if the filter is affective
    metadataPart := "Metadata"
    if values.GetString(&ytapi.FlagAssetFilter) == "effective" {
        metadataPart = "MetadataEffective"
    }
	table.RegisterPart(metadataPart, []*ytapi.Flag{
		&ytapi.Flag{Name: "customId", Type: ytapi.FLAG_STRING },
        &ytapi.Flag{Name: "title", Type: ytapi.FLAG_STRING },
        &ytapi.Flag{Name: "description", Type: ytapi.FLAG_STRING },
	})

	// set default columns
	table.SetColumns([]string{ "asset","type","timeCreated","title","customId" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Assets

func GetAsset(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	call := service.PAPI.Assets.Get(values.GetString(&ytapi.FlagAsset))
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

    // set filters
    if values.IsSet(&ytapi.FlagAssetFilter) {
        call = call.FetchMatchPolicy(values.GetString(&ytapi.FlagAssetFilter))
        call = call.FetchMetadata(values.GetString(&ytapi.FlagAssetFilter))
        call = call.FetchOwnership(values.GetString(&ytapi.FlagAssetFilter))
    }

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

