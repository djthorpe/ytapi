/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package cidcommands

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
//    "github.com/djthorpe/ytapi/youtubepartner/v1"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterReferenceCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListReferences",
			Description: "Get a list of references for an asset",
            ServiceAccount: true,
            Required:    []*ytapi.Flag{ &ytapi.FlagAsset },
			Setup:       RegisterReferenceFormat,
			Execute:     ListReferences,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register format

func RegisterReferenceFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "reference", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{ "reference" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List References

func ListReferences(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
    // TODO

    // Success
	return nil
}

