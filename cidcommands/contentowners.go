package cidcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterContentOwnerCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:           "ListContentOwners",
			Description:    "List content owners",
			ServiceAccount: true,
			Setup:          RegisterContentOwnerFormat,
			Execute:        ListContentOwners,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register content owner format

func RegisterContentOwnerFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "contentowner", Path: "Id", Type: ytapi.FLAG_CONTENTOWNER},
		&ytapi.Flag{Name: "displayName", Path: "DisplayName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "conflictNotificationEmail", Path: "ConflictNotificationEmail", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"contentowner", "displayName", "conflictNotificationEmail"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Content Owners

func ListContentOwners(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Create call and set parameters
	call := service.PAPI.ContentOwners.List()
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Get response
	response, err := call.FetchMine(true).Do(service.CallOptions()...)
	if err != nil {
		return err
	}
	if err = table.Append(response.Items); err != nil {
		return err
	}
	return nil
}
