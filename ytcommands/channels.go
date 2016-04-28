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

func RegisterChannelCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "ListChannels",
			Description: "List channels",
			Optional:    []*ytapi.Flag{ &ytapi.FlagLanguage, &ytapi.FlagMaxResults },
			Setup:       RegisterChannelFormat,
			Execute:     ListChannels,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register channel output format

func RegisterChannelFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []ytapi.FieldSpec{
		ytapi.FieldSpec{"channel", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytapi.FieldSpec{
		ytapi.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytapi.FieldSpec{"countrycode", "Snippet/Country", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"defaultLanguage", "Snippet/DefaultLanguage", ytservice.FIELD_STRING},
	})

	table.RegisterPart("contentDetails", []ytapi.FieldSpec{
		ytapi.FieldSpec{"playlist.likes", "ContentDetails/RelatedPlaylists/Likes", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"playlist.favorites", "ContentDetails/RelatedPlaylists/Favourites", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"playlist.uploads", "ContentDetails/RelatedPlaylists/Uploads", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"playlist.watchHistory", "ContentDetails/RelatedPlaylists/WatchHistory", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"playlist.watchLater", "ContentDetails/RelatedPlaylists/WatchLater", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"googlePlusUserId", "ContentDetails/GooglePlusUserId", ytservice.FIELD_STRING},
	})

	table.RegisterPart("statistics", []ytapi.FieldSpec{
		ytapi.FieldSpec{"viewCount", "Statistics/ViewCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"commentCount", "Statistics/CommentCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"subscriberCount", "Statistics/SubscriberCount", ytservice.FIELD_NUMBER},
		ytapi.FieldSpec{"hiddenSubscriberCount", "Statistics/HiddenSubscriberCount", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"videoCount", "Statistics/VideoCount", ytservice.FIELD_NUMBER},
	})

	table.RegisterPart("status", []ytapi.FieldSpec{
		ytapi.FieldSpec{"privacyStatus", "Status/PrivacyStatus", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"isLinked", "Status/IsLinked", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"longUploadsStatus", "Status/LongUploadsStatus", ytservice.FIELD_STRING},
	})

	table.RegisterPart("auditDetails", []ytapi.FieldSpec{
		ytapi.FieldSpec{"overallGoodStanding", "AuditDetails/OverallGoodStanding", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"communityGuidelinesGoodStanding", "AuditDetails/CommunityGuidelinesGoodStanding", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"copyrightStrikesGoodStanding", "AuditDetails/CopyrightStrikesGoodStanding", ytservice.FIELD_BOOLEAN},
		ytapi.FieldSpec{"contentIdClaimsGoodStanding", "AuditDetails/ContentIdClaimsGoodStanding", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("contentOwnerDetails", []ytapi.FieldSpec{
		ytapi.FieldSpec{"contentowner", "ContentOwnerDetails/ContentOwner", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"timeLinked", "ContentOwnerDetails/TimeLinked", ytservice.FIELD_DATETIME},
	})

	// set default columns
	table.SetColumns([]string{"channel", "title", "description", "publishedAt", "countrycode", "defaultLanguage"})

	// success
	return nil
}

func RegisterLocalizedChannelMetadataFormat(params *ytservice.Params, table *ytapi.Table) error {
	table.RegisterPart("localizations", []ytapi.FieldSpec{
		ytapi.FieldSpec{"language", "Language", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"title", "Title", ytservice.FIELD_STRING},
		ytapi.FieldSpec{"description", "Description", ytservice.FIELD_STRING},
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
	parts := "id,snippet,status" //strings.Join(table.Parts(), ",")

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
	call3 := service.API.Channels.List(strings.Join(table.Parts(), ",")).Id(*params.Channel)
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
