/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package cidcommands

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterPolicyCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListPolicies",
			Description: "List policies",
            ServiceAccount: true,
			Optional:    []*ytapi.Flag{&ytapi.FlagPolicyOrder},
			Setup:       RegisterPolicyFormat,
			Execute:     ListPolicies,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register policy format

func RegisterPolicyFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "policy", Path: "Id", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "name", Path: "Name", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Path: "Description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "timeUpdated", Path: "TimeUpdated", Type: ytapi.FLAG_TIME},
	})

	// set default columns
	table.SetColumns([]string{"policy", "name", "description","timeUpdated" })

	// success
	return nil
}


////////////////////////////////////////////////////////////////////////////////
// List Policies

func ListPolicies(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
    // create call and set parameters
	call := service.PAPI.Policies.List()
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	if values.IsSet(&ytapi.FlagPolicyOrder) {
		call = call.Sort(values.GetString(&ytapi.FlagPolicyOrder))
	}

	// perform query
	response,err := call.Do()
	if err != nil {
		return err
	}
	if err = table.Append(response.Items); err != nil {
		return err
	}
	return nil
}
