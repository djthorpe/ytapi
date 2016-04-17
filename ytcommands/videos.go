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
// Returns set of video items for YouTube service. Requires the channel
// parameter in order to work

func ListVideos(service *ytservice.Service,params *ytservice.Params, output *ytservice.Table) error {
	// create call for videos
	call := service.API.Videos.List("id,snippet")

	// set filter parameters - requires valid channel name
	if params.IsValidChannel() == false {
		return errors.New("Invalid channel parameter")
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
	items := make([]*youtube.Video, 0, maxresults)
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
