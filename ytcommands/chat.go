package ytcommands

/*
  Copyright David Thorpe 2017 All Rights Reserved
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
// Register chat commands

func RegisterLiveChatCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListChatMessages",
			Description: "List chat messages",
			Required:    []*ytapi.Flag{&ytapi.FlagChat},
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults},
			Setup:       RegisterChatMessageFormat,
			Execute:     ListChatMessages,
		},
		&ytapi.Command{
			Name:        "NewChatMessage",
			Description: "Insert chat message",
			Required:    []*ytapi.Flag{&ytapi.FlagChat, &ytapi.FlagChatText},
			Execute:     InsertChatMessage,
		},
		&ytapi.Command{
			Name:        "DeleteChatMessage",
			Description: "Delete chat message",
			Required:    []*ytapi.Flag{&ytapi.FlagChatMessage},
			Execute:     DeleteChatMessage,
		},
		&ytapi.Command{
			Name:        "ListChatModerators",
			Description: "List chat moderators",
			Required:    []*ytapi.Flag{&ytapi.FlagChat},
			Optional:    []*ytapi.Flag{&ytapi.FlagMaxResults},
			Setup:       RegisterChatModeratorFormat,
			Execute:     ListChatModerators,
		},
		&ytapi.Command{
			Name:        "NewChatModerator",
			Description: "Add chat moderator",
			Required:    []*ytapi.Flag{&ytapi.FlagChat, &ytapi.FlagChannel},
			Execute:     NewChatModerator,
		},
		&ytapi.Command{
			Name:        "DeleteChatModerator",
			Description: "Remove chat moderator",
			Execute:     DeleteChatModerator,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register output formats

func RegisterChatMessageFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "message", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "chat", Path: "Snippet/LiveChatId", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "type", Path: "Snippet/Type", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "published", Path: "Snippet/PublishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "text", Path: "Snippet/DisplayMessage", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "has_display_content", Path: "Snippet/HasDisplayContent", Type: ytapi.FLAG_BOOL},
	})

	table.RegisterPart("authorDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "author", Path: "AuthorDetails/DisplayName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author_channel", Path: "AuthorDetails/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "author_channel_url", Path: "AuthorDetails/ChannelUrl", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "author_verified", Path: "AuthorDetails/IsVerified", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "author_owner", Path: "AuthorDetails/IsChatOwner", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "author_sponsor", Path: "AuthorDetails/IsChatSponsor", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "author_moderator", Path: "AuthorDetails/IsChatModerator", Type: ytapi.FLAG_BOOL},
	})

	// set default columns
	table.SetColumns([]string{"type", "author", "text", "published"})

	// success
	return nil
}

func RegisterChatModeratorFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "moderator", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "channel", Path: "Snippet/ModeratorDetails/ChannelId", Type: ytapi.FLAG_CHANNEL},
		&ytapi.Flag{Name: "url", Path: "Snippet/ModeratorDetails/ChannelUrl", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "name", Path: "Snippet/ModeratorDetails/DisplayName", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "image", Path: "Snippet/ModeratorDetails/ProfileImageUrl", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "chat", Path: "Snippet/LiveChatId", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"moderator", "name"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Return chat-id - lookup if video-id is provided instead

func GetChatId(service *ytservice.Service, values *ytapi.Values) (string, error) {

	// Return error if no chat parameter
	if values.IsSet(&ytapi.FlagChat) == false {
		return "", errors.New("Missing --chat parameter")
	}

	// Get value
	value := values.GetString(&ytapi.FlagChat)

	// Return ID if not of kind video
	if values.IsKindOf(&ytapi.FlagChat, ytapi.FLAG_VIDEO) == false {
		return value, nil
	}

	// Set the call parameters
	call := service.API.LiveBroadcasts.List("snippet")
	call = call.Id(value)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		if values.IsSet(&ytapi.FlagChannel) {
			call = call.OnBehalfOfContentOwnerChannel(values.GetString(&ytapi.FlagChannel))
		}
	}

	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return "", err
	}

	if broadcasts := response.Items; len(broadcasts) != 1 {
		return "", errors.New("Broadcast not found")
	} else {
		return broadcasts[0].Snippet.LiveChatId, nil
	}
}

////////////////////////////////////////////////////////////////////////////////
// ChatMessage methods

func ListChatMessages(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	chat, err := GetChatId(service, values)
	if err != nil {
		return err
	}
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// create call
	call := service.API.LiveChatMessages.List(chat, strings.Join(table.Parts(false), ","))

	// Perform search, and return results
	return ytapi.DoChatMessagesList(call, table, int64(maxresults), service.CallOptions()...)
}

func InsertChatMessage(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	chat, err := GetChatId(service, values)
	if err != nil {
		return err
	}
	text := values.GetString(&ytapi.FlagChatText)

	// create call
	call := service.API.LiveChatMessages.Insert("snippet", &youtube.LiveChatMessage{
		Snippet: &youtube.LiveChatMessageSnippet{
			LiveChatId: chat,
			Type:       "textMessageEvent",
			TextMessageDetails: &youtube.LiveChatTextMessageDetails{
				MessageText: text,
			},
		},
	})

	// Insert chat message, return the message
	_, err = call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Success
	return nil
}

func DeleteChatMessage(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	message := values.GetString(&ytapi.FlagChatMessage)

	// create call
	call := service.API.LiveChatMessages.Delete(message)

	// Delete chat message
	if err := call.Do(service.CallOptions()...); err != nil {
		return err
	}

	// Success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// ChatModerator methods

func ListChatModerators(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	chat, err := GetChatId(service, values)
	if err != nil {
		return err
	}
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// create call
	call := service.API.LiveChatModerators.List(chat, strings.Join(table.Parts(false), ","))

	// Perform search, and return results
	return ytapi.DoChatModeratorsList(call, table, int64(maxresults), service.CallOptions()...)
}

func NewChatModerator(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	channel := values.GetString(&ytapi.FlagChannel)
	chat, err := GetChatId(service, values)
	if err != nil {
		return err
	}

	// create call
	call := service.API.LiveChatModerators.Insert("id", &youtube.LiveChatModerator{
		Snippet: &youtube.LiveChatModeratorSnippet{
			ModeratorDetails: &youtube.ChannelProfileDetails{
				ChannelId: channel,
			},
			LiveChatId: chat,
		},
	})

	// execute
	_, err = call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Success
	return nil
}

func DeleteChatModerator(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Success
	return nil
}
