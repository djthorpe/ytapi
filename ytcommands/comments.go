package ytcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
	"strings"
)

import (
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////
// Register comments commands

func RegisterCommentsCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListVideoCommentThreads",
			Description: "List comment threads for a video",
			Required:    []*ytapi.Flag{&ytapi.FlagVideo},
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults, &ytapi.FlagCommentFormat, &ytapi.FlagCommentModerationStatus, &ytapi.FlagCommentOrder, &ytapi.FlagSearchQuery},
			Setup:       RegisterCommentThreadFormat,
			Execute:     ListVideoCommentThreads,
		},
		&ytapi.Command{
			Name:        "ListChannelCommentThreads",
			Description: "List comment threads for a channel",
			Required:    []*ytapi.Flag{&ytapi.FlagChannel},
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults, &ytapi.FlagCommentFormat, &ytapi.FlagCommentModerationStatus, &ytapi.FlagCommentOrder, &ytapi.FlagSearchQuery},
			Setup:       RegisterCommentThreadFormat,
			Execute:     ListChannelCommentThreads,
		},
		&ytapi.Command{
			Name:        "ListComments",
			Description: "List comments for a thread",
			Required:    []*ytapi.Flag{&ytapi.FlagCommentThread},
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults, &ytapi.FlagCommentFormat},
			Setup:       RegisterCommentFormat,
			Execute:     ListCommentsForThread,
		},
		&ytapi.Command{
			Name:        "NewComment",
			Description: "Add comment to thread",
			Required:    []*ytapi.Flag{&ytapi.FlagCommentThread, &ytapi.FlagCommentText},
			Setup:       RegisterCommentFormat,
			Execute:     InsertCommentForThread,
		},
		&ytapi.Command{
			Name:        "DeleteComment",
			Description: "Add comment to thread",
			Required:    []*ytapi.Flag{&ytapi.FlagCommentThread},
			Execute:     DeleteComment,
		},
		&ytapi.Command{
			Name:        "MarkCommentAsSpam",
			Description: "Mark comment as spam",
			Required:    []*ytapi.Flag{&ytapi.FlagCommentThread},
			Setup:       RegisterCommentFormat,
			Execute:     MarkCommentAsSpam,
		},
		&ytapi.Command{
			Name:        "SetCommentModerationStatus",
			Description: "Sets the moderation status of a comment",
			Required:    []*ytapi.Flag{&ytapi.FlagCommentThread, &ytapi.FlagCommentModerationStatus2},
			Optional:    []*ytapi.Flag{&ytapi.FlagCommentBanAuthor},
			Setup:       RegisterCommentFormat,
			Execute:     SetCommentModerationStatus,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register output formats

func RegisterCommentThreadFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "thread", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "video", Path: "Snippet/VideoId", Type: ytapi.FLAG_VIDEO},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "can_reply", Path: "Snippet/CanReply", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "reply_count", Path: "Snippet/TotalReplyCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "can_rate", Path: "Snippet/TopLevelComment/Snippet/CanRate", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "viewer_rating", Path: "Snippet/TopLevelComment/Snippet/ViewerRating", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "like_count", Path: "Snippet/TopLevelComment/Snippet/LikeCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "moderation_status", Path: "Snippet/TopLevelComment/Snippet/ModerationStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "published", Path: "Snippet/TopLevelComment/Snippet/PublishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "updated", Path: "Snippet/TopLevelComment/Snippet/UpdatedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "public", Path: "Snippet/IsPublic", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "text", Path: "Snippet/TopLevelComment/Snippet/TextDisplay", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author", Path: "Snippet/TopLevelComment/Snippet/AuthorDisplayName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author_profile_url", Path: "Snippet/TopLevelComment/Snippet/AuthorProfileImageUrl", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author_channel_url", Path: "Snippet/TopLevelComment/Snippet/AuthorChannelUrl", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"thread", "public", "author", "text", "published", "reply_count", "can_reply"})

	// success
	return nil
}

func RegisterCommentFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "thread", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "parent", Path: "Snippet/ParentId", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "video", Path: "Snippet/VideoId", Type: ytapi.FLAG_VIDEO},
		&ytapi.Flag{Name: "channel", Path: "Snippet/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "can_rate", Path: "Snippet/CanRate", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "viewer_rating", Path: "Snippet/ViewerRating", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "like_count", Path: "Snippet/LikeCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "moderation_status", Path: "Snippet/ModerationStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "published", Path: "Snippet/PublishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "updated", Path: "Snippet/UpdatedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "text", Path: "Snippet/TextDisplay", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author", Path: "Snippet/AuthorDisplayName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author_profile_url", Path: "Snippet/AuthorProfileImageUrl", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author_channel_url", Path: "Snippet/AuthorChannelUrl", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"thread", "author", "text", "like_count", "published"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Comment Threads

func ListVideoCommentThreads(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// create call
	call := service.API.CommentThreads.List(strings.Join(table.Parts(false), ","))
	if values.IsSet(&ytapi.FlagVideo) {
		call = call.VideoId(values.GetString(&ytapi.FlagVideo))
	} else {
		return errors.New("Missing -video flag")
	}
	if values.IsSet(&ytapi.FlagCommentFormat) {
		call = call.TextFormat(values.GetString(&ytapi.FlagCommentFormat))
	}
	if values.IsSet(&ytapi.FlagCommentModerationStatus) {
		call = call.ModerationStatus(values.GetString(&ytapi.FlagCommentModerationStatus))
	}
	if values.IsSet(&ytapi.FlagCommentOrder) {
		call = call.Order(values.GetString(&ytapi.FlagCommentOrder))
	}
	if values.IsSet(&ytapi.FlagSearchQuery) {
		call = call.SearchTerms(values.GetString(&ytapi.FlagSearchQuery))
	}

	// Perform search, and return results
	return ytapi.DoCommentThreadsList(call, table, int64(maxresults))
}

func ListChannelCommentThreads(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// create call
	call := service.API.CommentThreads.List(strings.Join(table.Parts(false), ","))
	if values.IsSet(&ytapi.FlagChannel) {
		call = call.ChannelId(values.GetString(&ytapi.FlagChannel))
	} else {
		return errors.New("Missing -channel flag")
	}
	if values.IsSet(&ytapi.FlagCommentFormat) {
		call = call.TextFormat(values.GetString(&ytapi.FlagCommentFormat))
	}
	if values.IsSet(&ytapi.FlagCommentModerationStatus) {
		call = call.ModerationStatus(values.GetString(&ytapi.FlagCommentModerationStatus))
	}
	if values.IsSet(&ytapi.FlagCommentOrder) {
		call = call.Order(values.GetString(&ytapi.FlagCommentOrder))
	}
	if values.IsSet(&ytapi.FlagSearchQuery) {
		call = call.SearchTerms(values.GetString(&ytapi.FlagSearchQuery))
	}

	// Perform search, and return results
	return ytapi.DoCommentThreadsList(call, table, int64(maxresults))
}

////////////////////////////////////////////////////////////////////////////////
// Comments

func ListCommentsForThread(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// create call
	call := service.API.Comments.List(strings.Join(table.Parts(false), ","))
	if values.IsSet(&ytapi.FlagCommentThread) {
		call = call.ParentId(values.GetString(&ytapi.FlagCommentThread))
	} else {
		return errors.New("Missing -thread flag")
	}
	if values.IsSet(&ytapi.FlagCommentFormat) {
		call = call.TextFormat(values.GetString(&ytapi.FlagCommentFormat))
	}

	// Perform search, and return results
	return ytapi.DoCommentsList(call, table, int64(maxresults))
}

func MarkCommentAsSpam(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// create call
	thread := values.GetString(&ytapi.FlagCommentThread)
	call := service.API.Comments.MarkAsSpam(thread)
	err := call.Do(service.CallOptions()...)
	if err != nil {
		return nil
	}

	// return the comment
	call2 := service.API.Comments.List(strings.Join(table.Parts(false), ",")).Id(thread)

	// Perform search, and return results
	return ytapi.DoCommentsList(call2, table, 1)
}

func SetCommentModerationStatus(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// create call
	thread := values.GetString(&ytapi.FlagCommentThread)
	moderation_status := values.GetString(&ytapi.FlagCommentModerationStatus2)
	call := service.API.Comments.SetModerationStatus(thread, moderation_status)

	// Set ban flag
	if values.IsSet(&ytapi.FlagCommentBanAuthor) {
		call.BanAuthor(values.GetBool(&ytapi.FlagCommentBanAuthor))
	}

	// Execute
	if err := call.Do(service.CallOptions()...); err != nil {
		return err
	}

	// return the comment
	call2 := service.API.Comments.List(strings.Join(table.Parts(false), ",")).Id(thread)

	// Perform search, and return results
	return ytapi.DoCommentsList(call2, table, 1)
}

func InsertCommentForThread(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// create call
	call := service.API.Comments.Insert("snippet", &youtube.Comment{
		Snippet: &youtube.CommentSnippet{
			TextOriginal: values.GetString(&ytapi.FlagCommentText),
			ParentId:     values.GetString(&ytapi.FlagCommentThread),
		},
	})

	// Execute
	if response, err := call.Do(); err != nil {
		return err
	} else {
		// List comment
		call := service.API.Comments.List(strings.Join(table.Parts(false), ",")).Id(response.Id)
		// Perform search, and return results
		return ytapi.DoCommentsList(call, table, 1)
	}
}

func DeleteComment(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// create call
	thread := values.GetString(&ytapi.FlagCommentThread)
	call := service.API.Comments.Delete(thread)

	// Execute
	if err := call.Do(service.CallOptions()...); err != nil {
		return err
	}

	return nil
}
