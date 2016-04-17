/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Returns set of channel items for YouTube service. Can return several, in the
// case of service accounts, or a single one, based on simple OAuth authentication

func ListChannels(service *ytservice.Service, params *ytservice.Params, output *ytservice.Table) error {
	// create call for channels
	call := service.API.Channels.List("id,snippet")

	// set filter parameters
	if params.IsValidChannel() {
		call = call.Id(*params.Channel)
	} else if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).ManagedByMe(true)
	} else {
		call = call.Mine(true)
	}

	// Page through results
	maxresults := params.MaxResults
	nextPageToken := ""
	items := make([]*youtube.Channel, 0, maxresults)
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
