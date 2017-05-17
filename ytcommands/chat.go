/*
  Copyright David Thorpe 2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
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
			Required:    []*ytapi.Flag{&ytapi.FlagChat,&ytapi.FlagChatText},
			Execute:     InsertChatMessage,
		},
		&ytapi.Command{
			Name:        "DeleteChatMessage",
			Description: "Delete chat message",
			Required:    []*ytapi.Flag{&ytapi.FlagChatMessage},
			Execute:     DeleteChatMessage,
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
	table.SetColumns([]string{ "type", "author", "text", "published" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// ChatMessage methods

func ListChatMessages(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	chat := values.GetString(&ytapi.FlagChat)
	maxresults := values.GetUint(&ytapi.FlagMaxResults)

	// create call
	call := service.API.LiveChatMessages.List(chat,strings.Join(table.Parts(false), ","))

	// Perform search, and return results
	return ytapi.DoChatMessagesList(call, table, int64(maxresults))
}

func InsertChatMessage(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// Get parameters
	chat := values.GetString(&ytapi.FlagChat)
	text := values.GetString(&ytapi.FlagChatText)

	// create call
	call := service.API.LiveChatMessages.Insert("snippet",&youtube.LiveChatMessage{
		Snippet: &youtube.LiveChatMessageSnippet{
			LiveChatId: chat,
			Type: "textMessageEvent",
			TextMessageDetails: &youtube.LiveChatTextMessageDetails{
				MessageText: text,
			},
		},
	})

	// Insert chat message, return the message
	_, err := call.Do()
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
	if err := call.Do(); err != nil {
		return err
	}

	// Success
	return nil
}




