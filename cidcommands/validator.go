package cidcommands

/*
  Copyright David Thorpe 2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"io/ioutil"
)

import (
	"github.com/djthorpe/ytapi/youtubepartner/v1"
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterValidatorCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:           "ValidateMetadata",
			Description:    "Validate Metadata File",
			ServiceAccount: true,
			Required:       []*ytapi.Flag{&ytapi.FlagUploader, &ytapi.FlagFile},
			Setup:          RegisterValidateFormat,
			Execute:        ValidateMetadata,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register content owner format

func RegisterValidateFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("error", []*ytapi.Flag{
		&ytapi.Flag{Name: "severity", Path: "Severity", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "message", Path: "Message", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "messageCode", Path: "MessageCode", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "lineNumber", Path: "LineNumber", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "columnNumber", Path: "ColumnNumber", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "columnName", Path: "ColumnName", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"severity", "message", "lineNumber"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Validate Metadata

func ValidateMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get filename
	filename := values.GetString(&ytapi.FlagFile)

	// Open the metadata file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Create call and set parameters
	// TODO: assume UTF-8 for the moment, need to fix this not to assume that
	call := service.PAPI.Validator.Validate(&youtubepartner.ValidateRequest{
		UploaderName: values.GetString(&ytapi.FlagUploader),
		Content:      string(bytes),
	})
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Get response
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// The response will be errors
	if err = table.Append(response.Errors); err != nil {
		return err
	}

	return nil

	return nil
}
