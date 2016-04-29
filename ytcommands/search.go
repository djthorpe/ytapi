/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterSearchFormat(params *ytservice.Params, table *ytapi.Table) error {

	// register parts
	table.RegisterPart("snippet", []ytapi.Flag{
		ytapi.Flag{Name: "id", Path: "Id", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "title", Path: "Snippet/Title", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "description", Path: "Snippet/Description", Type: ytapi.FLAG_STRING},
		ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		ytapi.Flag{Name: "publishedAt", Path: "Snippet/PublishedAt", Type: ytapi.FLAG_TIME},
		ytapi.Flag{Name: "liveBroadcastContent", Path: "Snippet/LiveBroadcastContent", Type: ytapi.FLAG_BOOL},
		ytapi.Flag{Name: "channel.title", Path: "Snippet/ChannelTitle", Type: ytapi.FLAG_STRING},
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
