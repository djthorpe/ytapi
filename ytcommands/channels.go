/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////

type Localization struct {
	Language    string
	Title       string
	Description string
}

////////////////////////////////////////////////////////////////////////////////
// Register channel commands

func RegisterChannelCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "ListChannels",
			Description: "List channels",
			Optional:    []*ytapi.Flag{&ytapi.FlagLanguage, &ytapi.FlagMaxResults},
			Setup:       RegisterChannelFormat,
			Execute:     ListChannels,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register channel output format

func RegisterChannelFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "channel", Path: "Id", Type: ytapi.FLAG_CHANNEL},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "title", Path: "Snippet/Title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Path: "Snippet/Description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "publishedAt", Path: "Snippet/PublishedAt", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "countrycode", Path: "Snippet/Country", Type: ytapi.FLAG_REGION},
		&ytapi.Flag{Name: "defaultLanguage", Path: "Snippet/DefaultLanguage", Type: ytapi.FLAG_LANGUAGE},
	})

	table.RegisterPart("contentDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "playlistLikes", Path: "ContentDetails/RelatedPlaylists/Likes", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "playlistFavorites", Path: "ContentDetails/RelatedPlaylists/Favourites", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "playlistUploads", Path: "ContentDetails/RelatedPlaylists/Uploads", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "playlistWatchHistory", Path: "ContentDetails/RelatedPlaylists/WatchHistory", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "playlistWatchLater", Path: "ContentDetails/RelatedPlaylists/WatchLater", Type: ytapi.FLAG_PLAYLIST},
		&ytapi.Flag{Name: "googlePlusUserId", Path: "ContentDetails/GooglePlusUserId", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("statistics", []*ytapi.Flag{
		&ytapi.Flag{Name: "viewCount", Path: "Statistics/ViewCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "commentCount", Path: "Statistics/CommentCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "subscriberCount", Path: "Statistics/SubscriberCount", Type: ytapi.FLAG_UINT},
		&ytapi.Flag{Name: "hiddenSubscriberCount", Path: "Statistics/HiddenSubscriberCount", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "videoCount", Path: "Statistics/VideoCount", Type: ytapi.FLAG_UINT},
	})

	table.RegisterPart("status", []*ytapi.Flag{
		&ytapi.Flag{Name: "privacyStatus", Path: "Status/PrivacyStatus", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "isLinked", Path: "Status/IsLinked", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "longUploadsStatus", Path: "Status/LongUploadsStatus", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("auditDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "overallGoodStanding", Path: "AuditDetails/OverallGoodStanding", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "communityGuidelinesGoodStanding", Path: "AuditDetails/CommunityGuidelinesGoodStanding", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "copyrightStrikesGoodStanding", Path: "AuditDetails/CopyrightStrikesGoodStanding", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "contentIdClaimsGoodStanding", Path: "AuditDetails/ContentIdClaimsGoodStanding", Type: ytapi.FLAG_BOOL},
	})

	table.RegisterPart("contentOwnerDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "contentowner", Path: "ContentOwnerDetails/ContentOwner", Type: ytapi.FLAG_CONTENTOWNER},
		&ytapi.Flag{Name: "timeLinked", Path: "ContentOwnerDetails/TimeLinked", Type: ytapi.FLAG_TIME},
	})

	// set default columns
	table.SetColumns([]string{"channel", "title", "description", "publishedAt"})

	// success
	return nil
}

func RegisterLocalizedChannelMetadataFormat(params *ytservice.Params, table *ytapi.Table) error {
	table.RegisterPart("localizations", []*ytapi.Flag{
		&ytapi.Flag{Name: "language", Path: "Language", Type: ytapi.FLAG_LANGUAGE},
		&ytapi.Flag{Name: "title", Path: "Title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Path: "Description", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"language", "title", "description"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Channels.list

func ListChannels(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	maxresults := values.GetUint(&ytapi.FlagMaxResults)
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	parts := strings.Join(table.Parts(false), ",")

	// create call and set parameters
	call := service.API.Channels.List(parts)
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
		if channel == "" {
			call = call.ManagedByMe(true)
		} else {
			call = call.Id(channel)
		}
	} else if channel != "" {
		call = call.Id(channel)
	} else {
		call = call.Mine(true)
	}

	// Perform search, and return results
	return ytapi.DoChannelsList(call, table, int64(maxresults))
}

func ListLocalizedChannelMetadata(service *ytservice.Service, params *ytservice.Params, table *ytapi.Table) error {

	// Check channel parameter
	if params.IsValidChannel() == false {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}

	// create call
	call := service.API.Channels.List("localizations")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner)
	}
	response, err := call.Id(*params.Channel).Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}
	if len(response.Items) == 0 {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}

	// Get localizations
	localizations := response.Items[0].Localizations
	for language, metadata := range localizations {
		table.Append([]Localization{{language, metadata.Title, metadata.Description}})
	}

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Set channel metadata

func UpdateChannelMetadata(service *ytservice.Service, params *ytservice.Params, table *ytapi.Table) error {

	// Check channel parameter
	if params.IsValidChannel() == false {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}

	// Retrieve banding settings
	call := service.API.Channels.List("brandingSettings")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner)
	}
	response, err := call.Id(*params.Channel).Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}
	if len(response.Items) == 0 {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}

	// Set language, title and description in youtube.Channel
	channel := response.Items[0]
	if params.IsValidLanguage() {
		channel.BrandingSettings.Channel.DefaultLanguage = *params.Language
	}
	if params.IsEmptyTitle() == false {
		channel.BrandingSettings.Channel.Title = *params.Title
	}
	if params.IsEmptyDescription() == false {
		channel.BrandingSettings.Channel.Description = *params.Description
	}

	// Update branding settings
	call2 := service.API.Channels.Update("brandingSettings", channel)
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(*params.ContentOwner)
	}
	_, err = call2.Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}

	// Retrieve channel again
	call3 := service.API.Channels.List(strings.Join(table.Parts(false), ",")).Id(*params.Channel)
	if service.ServiceAccount {
		call3 = call3.OnBehalfOfContentOwner(*params.ContentOwner).ManagedByMe(true)
	}

	// Perform channels.list and return results
	return ytapi.DoChannelsList(call3, table, 1)
}

func UpdateLocalizedChannelMetadata(service *ytservice.Service, params *ytservice.Params, table *ytapi.Table) error {

	// Check channel parameter
	if params.IsValidChannel() == false {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}
	// Check language parameter
	if params.IsValidLanguage() == false {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}

	// retrieve localizations information from the channel
	call := service.API.Channels.List("id,localizations")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner)
	}
	response, err := call.Id(*params.Channel).Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}
	if len(response.Items) == 0 {
		return ytservice.NewError(ytservice.ErrorBadParameter, nil)
	}

	// edit localizations
	localizations := response.Items[0].Localizations
	metadata := youtube.ChannelLocalization{}
	if _, ok := localizations[*params.Language]; ok {
		metadata = localizations[*params.Language]
	}
	if params.IsEmptyTitle() == false {
		metadata.Title = *params.Title
	}
	if params.IsEmptyDescription() == false {
		metadata.Description = *params.Description
	}
	localizations[*params.Language] = metadata

	// update localization
	call2 := service.API.Channels.Update("localizations", &youtube.Channel{
		Id:            *params.Channel,
		Localizations: localizations,
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(*params.ContentOwner)
	}
	_, err = call2.Do()
	if err != nil {
		return ytservice.NewError(ytservice.ErrorResponse, err)
	}

	// Call list
	return ListLocalizedChannelMetadata(service, params, table)
}
