/*
Copyright David Thorpe 2015-2016 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"github.com/djthorpe/ytapi/youtubepartner/v1"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////

// Constants
const (
	YouTubeMaxPagingResults = 50
)

////////////////////////////////////////////////////////////////////////////////

func DoSearchList(call *youtube.SearchListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoChannelsList(call *youtube.ChannelsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoVideosList(call *youtube.VideosListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoBroadcastsList(call *youtube.LiveBroadcastsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoStreamsList(call *youtube.LiveStreamsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoPlaylistsList(call *youtube.PlaylistsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoPlaylistItemsList(call *youtube.PlaylistItemsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoClaimsList(call *youtubepartner.ClaimsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// retrieve next page of results
		response, err := call.PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		// TODO: Scope by maxresults
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}


func DoActivityList(call *youtube.ActivitiesListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}


func DoCommentThreadsList(call *youtube.CommentThreadsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}

func DoCommentsList(call *youtube.CommentsListCall, table *Table, maxresults int64) error {
	var numresults int64 = 0
	var nextPageToken string = ""

	// Page through results
	for {
		// test to see if we have all the items we now need
		if maxresults > 0 && numresults >= maxresults {
			break
		}

		// determine how many items we should rerieve in this pass
		var retrieveitems int64 = 0
		if maxresults == 0 {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else if (maxresults - numresults) > YouTubeMaxPagingResults {
			retrieveitems = int64(YouTubeMaxPagingResults)
		} else {
			retrieveitems = (maxresults - numresults)
		}
		response, err := call.MaxResults(retrieveitems).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		if err = table.Append(response.Items); err != nil {
			return err
		}
		numresults += int64(len(response.Items))
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// Success
	return nil
}



