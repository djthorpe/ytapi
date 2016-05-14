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

func RegisterCaptionCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListCaptions",
			Description: "List captions",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Setup:       RegisterCaptionFormat,
			Execute:     ListCaptions,
		},
	}
}

func RegisterCaptionFormat(values *ytapi.Values, table *ytapi.Table) error {

	// register parts
	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "caption", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "video", Path: "Snippet/VideoId", Type: ytapi.FLAG_VIDEO},
		&ytapi.Flag{Name: "lastUpdated", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "captiontype", Path: "Snippet/TrackKind", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "language", Type: ytapi.FLAG_LANGUAGE},
		&ytapi.Flag{Name: "name", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "audiotype", Path: "Snippet/AudioTrackType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "captions", Path: "Snippet/IsCC", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "largetext", Path: "Snippet/IsLarge", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "easyreader", Path: "Snippet/IsEasyReader", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "draft", Path: "Snippet/IsDraft", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "autosynced", Path: "Snippet/IsAutoSynced", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "status", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "error", Path: "Snippet/FailureReason", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"caption", "name", "video", "language", "draft"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Captions

func ListCaptions(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	video := values.GetString(&ytapi.FlagVideo)
	parts := "id,snippet"

	// create call and set parameters
	call := service.API.Captions.List(parts, video)
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
