/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"
	"github.com/djthorpe/ytapi/ytservice"
	"github.com/djthorpe/ytapi/ytapi"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterSearchFormat(params *ytservice.Params, table *ytapi.Table) error {

	// register parts
	table.RegisterPart("snippet", []ytapi.FieldSpec{
		ytapi.FieldSpec{"id", "Id", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"channel", "Snippet/ChannelId", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"liveBroadcastContent", "Snippet/LiveBroadcastContent", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"channel.title", "Snippet/ChannelTitle", ytservice.FIELD_STRING},
	})

	// set default columns
	table.SetColumns([]string{"id", "title", "description", "publishedAt", "channel", "channel.title", "liveBroadcastContent"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Search.List

func Search(service *ytservice.Service, params *ytservice.Params, table *ytapi.Table) error {

	// create call
	call := service.API.Search.List(strings.Join(table.Parts(), ","))

	// add query term
	if params.IsEmptyQuery() == false {
		call.Q(*params.Query)
	}

	// Perform search, and return results
	return ytapi.DoSearchList(call, table, params.MaxResults)
}
