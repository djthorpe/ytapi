/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"errors"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register playlist output

func RegisterPlaylistFormat(params *ytservice.Params,table *ytservice.Table) error {

	// register parts
	table.RegisterPart("id",[]ytservice.FieldSpec{
		ytservice.FieldSpec{ "id","Id",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "kind","Kind",ytservice.FIELD_STRING },
	})
	table.RegisterPart("snippet",[]ytservice.FieldSpec{
		ytservice.FieldSpec{ "title","Snippet/Title",ytservice.FIELD_STRING },
		ytservice.FieldSpec{ "description","Snippet/Description",ytservice.FIELD_STRING },
	})

	// set default columns
	table.SetColumns([]string{ "id","title","description","kind" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Returns set of playlist items for channel

func ListPlaylists(service *ytservice.Service, params *ytservice.Params, output *ytservice.Table) error {
	// create call for channels
	call := service.API.Playlists.List("id,snippet")

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

	// Page through results
	maxresults := params.MaxResults
	nextPageToken := ""
	items := make([]*youtube.Playlist, 0, maxresults)
	for {
		var pagingresults = int64(maxresults) - int64(len(items))
		if pagingresults <= 0 {
			pagingresults = ytservice.YouTubeMaxPagingResults
		} else if pagingresults > ytservice.YouTubeMaxPagingResults {
			pagingresults = ytservice.YouTubeMaxPagingResults
		}
		response, err := call.MaxResults(pagingresults).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		items = append(items, response.Items...)
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	output.AppendColumn("id","id")
	output.AppendColumn("title","Title")
	output.AppendColumn("description","Description")

	for _,item := range(items) {
		row := output.NewRow()

		// id
		row.SetString("id",item.Id)

		// snippet
		row.SetString("title",item.Snippet.Title)
		row.SetString("description",item.Snippet.Description)
	}

	// success
	return nil
}
