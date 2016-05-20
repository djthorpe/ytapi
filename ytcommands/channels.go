/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"os"
	"strings"
	"errors"

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
		&ytapi.Command{
			Name:        "SetChannelBanner",
			Description: "Set Channel Banner Image",
			Required:    []*ytapi.Flag{ &ytapi.FlagFile },
			Execute:     SetChannelBanner,
		},
		&ytapi.Command{
			Name:        "GetLocalizedChannelMetadata",
			Description: "Get localized channel metadata",
			Setup:       RegisterLocalizedChannelMetadataFormat,
			Execute:     GetLocalizedChannelMetadata,
		},
		&ytapi.Command{
			Name:        "UpdateLocalizedChannelMetadata",
			Description: "Update localized channel metadata",
			Required:    []*ytapi.Flag{ &ytapi.FlagLanguage, &ytapi.FlagTitle },
			Optional:    []*ytapi.Flag{ &ytapi.FlagDescription },
			Setup:       RegisterLocalizedChannelMetadataFormat,
			Execute:     UpdateLocalizedChannelMetadata,
		},
		&ytapi.Command{
			Name:        "DeleteLocalizedChannelMetadata",
			Description: "Remove localized channel metadata",
			Required:    []*ytapi.Flag{ &ytapi.FlagLanguage },
			Setup:       RegisterLocalizedChannelMetadataFormat,
			Execute:     DeleteLocalizedChannelMetadata,
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
	table.SetColumns([]string{"channel", "title", "description", "defaultLanguage" })

	// success
	return nil
}

func RegisterLocalizedChannelMetadataFormat(values *ytapi.Values, table *ytapi.Table) error {
	table.RegisterPart("localizations", []*ytapi.Flag{
		&ytapi.Flag{Name: "language", Path: "Language", Type: ytapi.FLAG_LANGUAGE},
		&ytapi.Flag{Name: "title", Path: "Title", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "description", Path: "Description", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "default", Path: "Default", Type: ytapi.FLAG_BOOL},
	})

	// set default columns
	table.SetColumns([]string{"language", "title", "description", "default"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Channels

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

func SetChannelBanner(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get Parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
//	channel := values.GetString(&ytapi.FlagChannel)

	// Create Call
	call := service.API.ChannelBanners.Insert(&youtube.ChannelBannerResource{})
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}

	// Read file
	file, err := os.Open(values.GetString(&ytapi.FlagFile))
	defer file.Close()
	if err != nil {
		return err
	}

	// Upload channel banner and retrieve URL for the banner
	response, err := call.Media(file).Do()
	if err != nil {
		return err
	}
	url := response.Url

	// TODO: Retrieve channel
	table.Info(url)

	// success
	return nil
}

func GetLocalizedChannelMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)

	// create call
	call := service.API.Channels.List("localizations")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagChannel) == false {
		call = call.Mine(true)
	} else {
		call = call.Id(channel)
	}

	// Execute call
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("Channel not found")
	}

	// Get localizations
	localizations := response.Items[0].Localizations
	for language, metadata := range localizations {
		table.Append([]Localization{{language, metadata.Title, metadata.Description}})
	}

	// success
	return nil
}

func UpdateLocalizedChannelMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	language := values.GetString(&ytapi.FlagLanguage)

	// create call
	call := service.API.Channels.List("localizations")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagChannel) == false {
		call = call.Mine(true)
	} else {
		call = call.Id(channel)
	}

	// Execute call
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("Channel not found")
	}

	// Update channel localization settings
	metadata, ok := response.Items[0].Localizations[language]
	if ok == false {
		metadata = youtube.ChannelLocalization{ }
	}
	if values.IsSet(&ytapi.FlagTitle) {
		metadata.Title = values.GetString(&ytapi.FlagTitle)
	}
	if values.IsSet(&ytapi.FlagDescription) {
		metadata.Description = values.GetString(&ytapi.FlagDescription)
	}
	response.Items[0].Localizations[language] = metadata

	// update localization
	call2 := service.API.Channels.Update("localizations", &youtube.Channel{
		Id:            channel,
		Localizations: response.Items[0].Localizations,
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}
	_, err = call2.Do()
	if err != nil {
		return err
	}

	// success
	return GetLocalizedChannelMetadata(service,values,table)
}

func DeleteLocalizedChannelMetadata(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// Get parameters
	contentowner := values.GetString(&ytapi.FlagContentOwner)
	channel := values.GetString(&ytapi.FlagChannel)
	language := values.GetString(&ytapi.FlagLanguage)

	// create call
	call := service.API.Channels.List("localizations")
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(contentowner)
	}
	if values.IsSet(&ytapi.FlagChannel) == false {
		call = call.Mine(true)
	} else {
		call = call.Id(channel)
	}

	// Execute
	response, err := call.Do()
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("Channel not found")
	}

	// Update channel localization settings
	_, ok := response.Items[0].Localizations[language]
	if ok == false {
		return errors.New("Localized metadata for language does not exist")
	}
	delete(response.Items[0].Localizations,language)

	// update localization
	call2 := service.API.Channels.Update("localizations", &youtube.Channel{
		Id:            channel,
		Localizations: response.Items[0].Localizations,
	})
	if service.ServiceAccount {
		call2 = call2.OnBehalfOfContentOwner(contentowner)
	}
	_, err = call2.Do()
	if err != nil {
		return err
	}

	// success
	return GetLocalizedChannelMetadata(service,values,table)
}

/*

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

*/



