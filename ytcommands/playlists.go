/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"errors"
	"strings"

	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register playlist output

func RegisterPlaylistFormat(params *ytservice.Params, table *ytservice.Table) error {

	// register parts
	table.RegisterPart("id", []ytservice.FieldSpec{
		ytservice.FieldSpec{"id", "Id", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"kind", "Kind", ytservice.FIELD_STRING},
	})
	table.RegisterPart("snippet", []ytservice.FieldSpec{
		ytservice.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"defaultLanguage", "Snippet/DefaultLanguage", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"tags", "Snippet/Tags", ytservice.FIELD_STRING},

	})

	// set default columns
	table.SetColumns([]string{ "id", "title", "description" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Returns set of playlist items for channel

func ListPlaylists(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call for fetching playlists
	call := service.API.Playlists.List(strings.Join(table.Parts(), ","))

	// set filter parameters
	if params.IsValidChannel() == false {
		return errors.New("Invalid channel parameter")
	} else {
		call = call.ChannelId(*params.Channel)
	}
	if service.ServiceAccount {
		if params.IsValidContentOwner() == false {
			return errors.New("Invalid content owner parameter")
		}
		call = call.OnBehalfOfContentOwner(*params.ContentOwner)
	}
	if params.IsValidLanguage() {
		call = call.Hl(*params.Language)
	}

	// Perform channels.list and return results
	return service.DoPlaylistsList(call, table, params.MaxResults)
}

func UpdatePlaylistMetadata(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call for fetching the playlist
	call := service.API.Playlists.List("snippet")
	if service.ServiceAccount {
		if params.IsValidContentOwner() == false {
			return errors.New("Invalid content owner parameter")
		}
		call = call.OnBehalfOfContentOwner(*params.ContentOwner)
	}

	// TODO

	return nil
}
