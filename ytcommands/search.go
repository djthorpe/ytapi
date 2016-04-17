/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterSearchFormat(params *ytservice.Params,table *ytservice.Table) error {

	// register parts
	table.RegisterPart("snippet",[]ytservice.FieldSpec{
		ytservice.FieldSpec{ "id","Id",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "title","Snippet/Title",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "description","Snippet/Description",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "channel","Snippet/ChannelId",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "publishedAt","Snippet/PublishedAt",ytservice.FIELD_DATETIME },
		ytservice.FieldSpec{ "liveBroadcastContent","Snippet/LiveBroadcastContent",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "channel.title","Snippet/ChannelTitle",ytservice.FIELD_STRING },
	})

	// set default columns
	table.SetColumns([]string{ "id","title","description","publishedAt","channel","channel.title","liveBroadcastContent" })

	// success
	return nil
}


////////////////////////////////////////////////////////////////////////////////
// Search.List

func Search(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call
	call := service.API.Search.List(strings.Join(table.Parts(),","))

	// add query term
	if params.IsEmptyQuery() == false {
		call.Q(*params.Query)
	}

	// Perform search, and return results
	return service.DoSearchList(call,table,params.MaxResults)
}

