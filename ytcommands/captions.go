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
// Register caption commands

func RegisterCaptionCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListCaptions",
			Description: "List captions",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Setup:       RegisterCaptionFormat,
			Execute:     ListCaptions,
		},
	}
}

func RegisterCaptionFormat(values *ytapi.Values, table *ytservice.Table) error {

	// register parts
	table.RegisterPart("id", []ytservice.FieldSpec{
		ytservice.FieldSpec{"caption", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytservice.FieldSpec{
		ytservice.FieldSpec{"video", "Snippet/VideoId", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"lastUpdated", "Snippet/LastUpdated", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"captiontype", "Snippet/TrackKind", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"language", "Snippet/Language", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"name", "Snippet/Name", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"audiotype", "Snippet/AudioTrackType", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"closedcaptions", "Snippet/IsCC", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"largetext", "Snippet/IsLarge", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"easyreader", "Snippet/IsEasyReader", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"draft", "Snippet/IsDraft", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"autosynced", "Snippet/IsAutoSynced", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"status", "Snippet/Status", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"error", "Snippet/FailureReason", ytservice.FIELD_STRING},
	})

	// set default columns
	table.SetColumns([]string{"caption", "name", "video", "language", "draft"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Captions

func ListCaptions(service *ytservice.Service, values *ytapi.Values, table *ytservice.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)
	parts := "id,snippet"

	// create call and set parameters
	call := service.API.Captions.List(parts,video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
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
