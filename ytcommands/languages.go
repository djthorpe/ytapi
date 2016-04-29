/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register language commands

func RegisterLanguageRegionCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListLanguages",
			Description: "List languages",
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage},
			Setup:       RegisterLanguageFormat,
			Execute:     ListLanguages,
		},
		ytapi.Command{
			Name:        "ListRegions",
			Description: "List regions",
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage},
			Setup:       RegisterRegionFormat,
			Execute:     ListRegions,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register output formats

func RegisterLanguageFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []ytapi.Flag{
		ytapi.Flag{Name: "language", Path: "Id", Type: ytapi.FLAG_LANGUAGE},
	})

	table.RegisterPart("snippet", []ytapi.Flag{
		ytapi.Flag{Name: "name", Path: "Snippet/Name", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"language", "name"})

	// success
	return nil
}

func RegisterRegionFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []ytapi.Flag{
		ytapi.Flag{Name: "country", Path: "Id", Type: ytapi.FLAG_REGION},
	})

	table.RegisterPart("snippet", []ytapi.Flag{
		ytapi.Flag{Name: "name", Path: "Snippet/Name", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"country", "name"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Languages.list

func ListLanguages(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	language := values.GetString(&ytapi.FlagLanguage)
	parts := "id,snippet"

	// create call and set parameters
	call := service.API.I18nLanguages.List(parts)
	if language != "" {
		call = call.Hl(language)
	}

	// request and response
	response, err := call.Do()
	if err != nil {
		return err
	}
	if err = table.Append(response.Items); err != nil {
		return err
	}

	// Success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Regions.list

func ListRegions(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	language := values.GetString(&ytapi.FlagLanguage)
	parts := "id,snippet"

	// create call and set parameters
	call := service.API.I18nRegions.List(parts)
	if language != "" {
		call = call.Hl(language)
	}

	// request and response
	response, err := call.Do()
	if err != nil {
		return err
	}
	if err = table.Append(response.Items); err != nil {
		return err
	}

	// Success
	return nil
}
