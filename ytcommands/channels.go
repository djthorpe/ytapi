/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"strings"

	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register channel output format

func RegisterChannelFormat(params *ytservice.Params, table *ytservice.Table) error {

	// register parts
	table.RegisterPart("id", []ytservice.FieldSpec{
		ytservice.FieldSpec{"channel", "Id", ytservice.FIELD_STRING},
	})

	table.RegisterPart("snippet", []ytservice.FieldSpec{
		ytservice.FieldSpec{"title", "Snippet/Title", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"description", "Snippet/Description", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"publishedAt", "Snippet/PublishedAt", ytservice.FIELD_DATETIME},
		ytservice.FieldSpec{"countrycode", "Snippet/Country", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"defaultLanguage", "Snippet/DefaultLanguage", ytservice.FIELD_STRING},
	})

	table.RegisterPart("contentDetails", []ytservice.FieldSpec{
		ytservice.FieldSpec{"playlist.likes", "ContentDetails/RelatedPlaylists/Likes", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"playlist.favorites", "ContentDetails/RelatedPlaylists/Favourites", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"playlist.uploads", "ContentDetails/RelatedPlaylists/Uploads", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"playlist.watchHistory", "ContentDetails/RelatedPlaylists/WatchHistory", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"playlist.watchLater", "ContentDetails/RelatedPlaylists/WatchLater", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"googlePlusUserId", "ContentDetails/GooglePlusUserId", ytservice.FIELD_STRING},
	})

	table.RegisterPart("statistics", []ytservice.FieldSpec{
		ytservice.FieldSpec{"viewCount", "Statistics/ViewCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"commentCount", "Statistics/CommentCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"subscriberCount", "Statistics/SubscriberCount", ytservice.FIELD_NUMBER},
		ytservice.FieldSpec{"hiddenSubscriberCount", "Statistics/HiddenSubscriberCount", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"videoCount", "Statistics/VideoCount", ytservice.FIELD_NUMBER},
	})

	table.RegisterPart("status", []ytservice.FieldSpec{
		ytservice.FieldSpec{"privacyStatus", "Status/PrivacyStatus", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"isLinked", "Status/IsLinked", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"longUploadsStatus", "Status/LongUploadsStatus", ytservice.FIELD_STRING},
	})

	table.RegisterPart("auditDetails", []ytservice.FieldSpec{
		ytservice.FieldSpec{"overallGoodStanding", "AuditDetails/OverallGoodStanding", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"communityGuidelinesGoodStanding", "AuditDetails/CommunityGuidelinesGoodStanding", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"copyrightStrikesGoodStanding", "AuditDetails/CopyrightStrikesGoodStanding", ytservice.FIELD_BOOLEAN},
		ytservice.FieldSpec{"contentIdClaimsGoodStanding", "AuditDetails/ContentIdClaimsGoodStanding", ytservice.FIELD_BOOLEAN},
	})

	table.RegisterPart("contentOwnerDetails", []ytservice.FieldSpec{
		ytservice.FieldSpec{"contentowner", "ContentOwnerDetails/ContentOwner", ytservice.FIELD_STRING},
		ytservice.FieldSpec{"timeLinked", "ContentOwnerDetails/TimeLinked", ytservice.FIELD_DATETIME},
	})

	// set default columns
	table.SetColumns([]string{"channel", "title", "description", "publishedAt", "countrycode", "defaultLanguage" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Returns set of channel items for YouTube service. Can return several, in the
// case of service accounts, or a single one, based on simple OAuth authentication

func ListChannels(service *ytservice.Service, params *ytservice.Params, table *ytservice.Table) error {

	// create call
	call := service.API.Channels.List(strings.Join(table.Parts(), ","))

	// set filter parameters
	if service.ServiceAccount {
		call = call.OnBehalfOfContentOwner(*params.ContentOwner).ManagedByMe(true)
	} else if params.IsValidChannel() {
		call = call.Id(*params.Channel)
	} else {
		call = call.Mine(true)
	}

	// Perform channels.list and return results
	return service.DoChannelsList(call, table, params.MaxResults)
}
