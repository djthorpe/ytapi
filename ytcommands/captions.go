/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"os"
	"io"
	"strings"
	"path/filepath"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register caption commands and format

func RegisterCaptionCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListCaptionTracks",
			Description: "List caption tracks for a video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Setup:       RegisterCaptionFormat,
			Execute:     ListCaptionTracks,
		},
		&ytapi.Command{
			Name:        "AddCaptionTrack",
			Description: "Add Caption Track to video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo,&ytapi.FlagFile,&ytapi.FlagLanguage},
			Optional:    []*ytapi.Flag{&ytapi.FlagCaptionName,&ytapi.FlagCaptionSync,&ytapi.FlagCaptionDraft},
			Setup:       RegisterCaptionFormat,
			Execute:     AddCaptionTrack,
		},
		&ytapi.Command{
			Name:        "DeleteCaptionTrack",
			Description: "Delete Caption Track from video",
			Required:    []*ytapi.Flag{&ytapi.FlagCaption},
			Execute:     DeleteCaptionTrack,
		},
		&ytapi.Command{
			Name:        "DownloadCaptionTrack",
			Description: "Download Caption Track from video",
			Required:    []*ytapi.Flag{ &ytapi.FlagCaption },
			Optional:    []*ytapi.Flag{ &ytapi.FlagCaptionFormat, &ytapi.FlagLanguage, &ytapi.FlagFile },
			Execute:     DownloadCaptionTrack,
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
		&ytapi.Flag{Name: "draft", Path: "Snippet/IsDraft", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "sync", Path: "Snippet/IsAutoSynced", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "audiotype", Path: "Snippet/AudioTrackType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "captions", Path: "Snippet/IsCC", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "largetext", Path: "Snippet/IsLarge", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "easyreader", Path: "Snippet/IsEasyReader", Type: ytapi.FLAG_BOOL},
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

func ListCaptionTracks(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	video := values.GetString(&ytapi.FlagVideo)
	parts := strings.Join(table.Parts(false), ",")

	// create call and set parameters
	call := service.API.Captions.List(parts, video)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
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

func AddCaptionTrack(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	video := values.GetString(&ytapi.FlagVideo)
	language := values.GetString(&ytapi.FlagLanguage)
	parts := strings.Join(table.Parts(false), ",")
	name := values.GetString(&ytapi.FlagCaptionName)
	filename := values.GetString(&ytapi.FlagFile)
	draft := values.GetBool(&ytapi.FlagCaptionDraft)

	// Interpret name from filename if not set
	if values.IsSet(&ytapi.FlagCaptionName) == false {
		name = filepath.Base(filename)
	}

	// Create the call
	call := service.API.Captions.Insert("snippet",&youtube.Caption{
		Snippet: &youtube.CaptionSnippet{
			VideoId: video,
			Language: language,
			Name: name,
			IsDraft: draft,
			ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
				"IsDraft": &ytapi.FlagCaptionDraft,
			}),
		},
	})

	// Set the call parameters
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}
	if values.IsSet(&ytapi.FlagCaptionSync) {
		call = call.Sync(values.GetBool(&ytapi.FlagCaptionSync))
	}

	// Open the caption file file
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	// request and response
	_, err = call.Media(file).Do()
	if err != nil {
		return err
	}

	// Success, list all caption tracks
	call2 := service.API.Captions.List(parts, video)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	}

	// request and response
	response2, err := call2.Do()
	if err != nil {
		return err
	}
	if err = table.Append(response2.Items); err != nil {
		return err
	}

	// Success
	return nil
}

func DeleteCaptionTrack(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	caption := values.GetString(&ytapi.FlagCaption)

	// Create call, set parameters
	call := service.API.Captions.Delete(caption)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Delete the caption track
	err := call.Do()
	if err != nil {
		return err
	}

	// success
	return nil
}


func DownloadCaptionTrack(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	caption := values.GetString(&ytapi.FlagCaption)
	filename := values.GetString(&ytapi.FlagFile)

	// Create call, set parameters
	call := service.API.Captions.Download(caption)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagLanguage) {
		call = call.Tlang(values.GetString(&ytapi.FlagLanguage))
	}

	// Set caption download format, either from -format flag or from
	// the extension of the output filename
	if values.IsSet(&ytapi.FlagCaptionFormat) {
		call = call.Tfmt(values.GetString(&ytapi.FlagCaptionFormat))
	} else if values.IsSet(&ytapi.FlagFile) {
		// Try to determine format from file extension as a fallback
		// filepath.Ext also includes the '.' as the first character
		ext := filepath.Ext(filename)
		formats := strings.Split(ytapi.FlagCaptionFormat.Extra,"|")
		for _,format := range(formats) {
			if ext[1:] == format {
				call = call.Tfmt(format)
			}
		}
	}

	// Download the caption track
	response, err := call.Download()
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Output body to file
	fh := os.Stdout
	if values.IsSet(&ytapi.FlagFile) {
		fh, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer fh.Close()
	}
	_,err = io.Copy(fh,response.Body)
	if err != nil {
		return err
	}

	// success
	return nil
}


