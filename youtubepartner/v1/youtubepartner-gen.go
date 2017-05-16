// Package youtubepartner provides access to the Youtube Content ID API.
//
// See https://developers.google.com/youtube/partner/
//
// Usage example:
//
//   import "google.golang.org/api/youtubepartner/v1"
//   ...
//   youtubepartnerService, err := youtubepartner.New(oauthHttpClient)
package youtubepartner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "youtubePartner:v1"
const apiName = "youtubePartner"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/youtube/partner/v1/"

// OAuth2 scopes used by this API.
const (
	// View and manage your assets and associated content on YouTube
	YoutubepartnerScope = "https://www.googleapis.com/auth/youtubepartner"

	// View content owner account details from YouTube.
	YoutubepartnerContentOwnerReadonlyScope = "https://www.googleapis.com/auth/youtubepartner-content-owner-readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.AssetLabels = NewAssetLabelsService(s)
	s.AssetMatchPolicy = NewAssetMatchPolicyService(s)
	s.AssetRelationships = NewAssetRelationshipsService(s)
	s.AssetSearch = NewAssetSearchService(s)
	s.AssetShares = NewAssetSharesService(s)
	s.Assets = NewAssetsService(s)
	s.Campaigns = NewCampaignsService(s)
	s.ClaimHistory = NewClaimHistoryService(s)
	s.ClaimSearch = NewClaimSearchService(s)
	s.Claims = NewClaimsService(s)
	s.ContentOwnerAdvertisingOptions = NewContentOwnerAdvertisingOptionsService(s)
	s.ContentOwners = NewContentOwnersService(s)
	s.LiveCuepoints = NewLiveCuepointsService(s)
	s.MetadataHistory = NewMetadataHistoryService(s)
	s.Orders = NewOrdersService(s)
	s.Ownership = NewOwnershipService(s)
	s.OwnershipHistory = NewOwnershipHistoryService(s)
	s.Package = NewPackageService(s)
	s.Policies = NewPoliciesService(s)
	s.Publishers = NewPublishersService(s)
	s.ReferenceConflicts = NewReferenceConflictsService(s)
	s.References = NewReferencesService(s)
	s.Validator = NewValidatorService(s)
	s.VideoAdvertisingOptions = NewVideoAdvertisingOptionsService(s)
	s.Whitelists = NewWhitelistsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	AssetLabels *AssetLabelsService

	AssetMatchPolicy *AssetMatchPolicyService

	AssetRelationships *AssetRelationshipsService

	AssetSearch *AssetSearchService

	AssetShares *AssetSharesService

	Assets *AssetsService

	Campaigns *CampaignsService

	ClaimHistory *ClaimHistoryService

	ClaimSearch *ClaimSearchService

	Claims *ClaimsService

	ContentOwnerAdvertisingOptions *ContentOwnerAdvertisingOptionsService

	ContentOwners *ContentOwnersService

	LiveCuepoints *LiveCuepointsService

	MetadataHistory *MetadataHistoryService

	Orders *OrdersService

	Ownership *OwnershipService

	OwnershipHistory *OwnershipHistoryService

	Package *PackageService

	Policies *PoliciesService

	Publishers *PublishersService

	ReferenceConflicts *ReferenceConflictsService

	References *ReferencesService

	Validator *ValidatorService

	VideoAdvertisingOptions *VideoAdvertisingOptionsService

	Whitelists *WhitelistsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAssetLabelsService(s *Service) *AssetLabelsService {
	rs := &AssetLabelsService{s: s}
	return rs
}

type AssetLabelsService struct {
	s *Service
}

func NewAssetMatchPolicyService(s *Service) *AssetMatchPolicyService {
	rs := &AssetMatchPolicyService{s: s}
	return rs
}

type AssetMatchPolicyService struct {
	s *Service
}

func NewAssetRelationshipsService(s *Service) *AssetRelationshipsService {
	rs := &AssetRelationshipsService{s: s}
	return rs
}

type AssetRelationshipsService struct {
	s *Service
}

func NewAssetSearchService(s *Service) *AssetSearchService {
	rs := &AssetSearchService{s: s}
	return rs
}

type AssetSearchService struct {
	s *Service
}

func NewAssetSharesService(s *Service) *AssetSharesService {
	rs := &AssetSharesService{s: s}
	return rs
}

type AssetSharesService struct {
	s *Service
}

func NewAssetsService(s *Service) *AssetsService {
	rs := &AssetsService{s: s}
	return rs
}

type AssetsService struct {
	s *Service
}

func NewCampaignsService(s *Service) *CampaignsService {
	rs := &CampaignsService{s: s}
	return rs
}

type CampaignsService struct {
	s *Service
}

func NewClaimHistoryService(s *Service) *ClaimHistoryService {
	rs := &ClaimHistoryService{s: s}
	return rs
}

type ClaimHistoryService struct {
	s *Service
}

func NewClaimSearchService(s *Service) *ClaimSearchService {
	rs := &ClaimSearchService{s: s}
	return rs
}

type ClaimSearchService struct {
	s *Service
}

func NewClaimsService(s *Service) *ClaimsService {
	rs := &ClaimsService{s: s}
	return rs
}

type ClaimsService struct {
	s *Service
}

func NewContentOwnerAdvertisingOptionsService(s *Service) *ContentOwnerAdvertisingOptionsService {
	rs := &ContentOwnerAdvertisingOptionsService{s: s}
	return rs
}

type ContentOwnerAdvertisingOptionsService struct {
	s *Service
}

func NewContentOwnersService(s *Service) *ContentOwnersService {
	rs := &ContentOwnersService{s: s}
	return rs
}

type ContentOwnersService struct {
	s *Service
}

func NewLiveCuepointsService(s *Service) *LiveCuepointsService {
	rs := &LiveCuepointsService{s: s}
	return rs
}

type LiveCuepointsService struct {
	s *Service
}

func NewMetadataHistoryService(s *Service) *MetadataHistoryService {
	rs := &MetadataHistoryService{s: s}
	return rs
}

type MetadataHistoryService struct {
	s *Service
}

func NewOrdersService(s *Service) *OrdersService {
	rs := &OrdersService{s: s}
	return rs
}

type OrdersService struct {
	s *Service
}

func NewOwnershipService(s *Service) *OwnershipService {
	rs := &OwnershipService{s: s}
	return rs
}

type OwnershipService struct {
	s *Service
}

func NewOwnershipHistoryService(s *Service) *OwnershipHistoryService {
	rs := &OwnershipHistoryService{s: s}
	return rs
}

type OwnershipHistoryService struct {
	s *Service
}

func NewPackageService(s *Service) *PackageService {
	rs := &PackageService{s: s}
	return rs
}

type PackageService struct {
	s *Service
}

func NewPoliciesService(s *Service) *PoliciesService {
	rs := &PoliciesService{s: s}
	return rs
}

type PoliciesService struct {
	s *Service
}

func NewPublishersService(s *Service) *PublishersService {
	rs := &PublishersService{s: s}
	return rs
}

type PublishersService struct {
	s *Service
}

func NewReferenceConflictsService(s *Service) *ReferenceConflictsService {
	rs := &ReferenceConflictsService{s: s}
	return rs
}

type ReferenceConflictsService struct {
	s *Service
}

func NewReferencesService(s *Service) *ReferencesService {
	rs := &ReferencesService{s: s}
	return rs
}

type ReferencesService struct {
	s *Service
}

func NewValidatorService(s *Service) *ValidatorService {
	rs := &ValidatorService{s: s}
	return rs
}

type ValidatorService struct {
	s *Service
}

func NewVideoAdvertisingOptionsService(s *Service) *VideoAdvertisingOptionsService {
	rs := &VideoAdvertisingOptionsService{s: s}
	return rs
}

type VideoAdvertisingOptionsService struct {
	s *Service
}

func NewWhitelistsService(s *Service) *WhitelistsService {
	rs := &WhitelistsService{s: s}
	return rs
}

type WhitelistsService struct {
	s *Service
}

type AdBreak struct {
	// MidrollSeconds: The time of the ad break specified as the number of
	// seconds after the start of the video when the break occurs.
	MidrollSeconds int64 `json:"midrollSeconds,omitempty"`

	// Position: The point at which the break occurs during the video
	// playback.
	Position string `json:"position,omitempty"`

	// Slot: A list of ad slots that occur in an ad break. Ad slots let you
	// specify the number of ads that should run in each break.
	Slot []*AdSlot `json:"slot,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MidrollSeconds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MidrollSeconds") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AdBreak) MarshalJSON() ([]byte, error) {
	type noMethod AdBreak
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AdSlot struct {
	// Id: A value that identifies the ad slot to the ad server.
	Id string `json:"id,omitempty"`

	// Type: The type of ad that runs in the slot. The value may affect
	// YouTube's fallback behavior if the third-party platform does not
	// return ads.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AdSlot) MarshalJSON() ([]byte, error) {
	type noMethod AdSlot
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AllowedAdvertisingOptions struct {
	// AdsOnEmbeds: This setting indicates whether the partner can display
	// ads when videos run in an embedded player.
	AdsOnEmbeds bool `json:"adsOnEmbeds,omitempty"`

	// Kind: This property identifies the resource type. Its value is
	// youtubePartner#allowedAdvertisingOptions.
	Kind string `json:"kind,omitempty"`

	// LicAdFormats: A list of ad formats that the partner is allowed to use
	// for its uploaded videos.
	LicAdFormats []string `json:"licAdFormats,omitempty"`

	// UgcAdFormats: A list of ad formats that the partner is allowed to use
	// for claimed, user-uploaded content.
	UgcAdFormats []string `json:"ugcAdFormats,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdsOnEmbeds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdsOnEmbeds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AllowedAdvertisingOptions) MarshalJSON() ([]byte, error) {
	type noMethod AllowedAdvertisingOptions
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Asset struct {
	// AliasId: A list of asset IDs that can be used to refer to the asset.
	// The list contains values if the asset represents multiple constituent
	// assets that have been merged. In that case, any of the asset IDs
	// originally assigned to the constituent assets could be used to update
	// the master, or synthesized, asset.
	AliasId []string `json:"aliasId,omitempty"`

	// Id: An ID that YouTube assigns and uses to uniquely identify the
	// asset.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For asset resources, the value is
	// youtubePartner#asset.
	Kind string `json:"kind,omitempty"`

	// Label: A list of asset labels on the asset.
	Label []string `json:"label,omitempty"`

	// MatchPolicy: The matchPolicy object contains information about the
	// asset's match policy, which YouTube applies to user-uploaded videos
	// that match the asset.
	MatchPolicy *AssetMatchPolicy `json:"matchPolicy,omitempty"`

	MatchPolicyEffective *AssetMatchPolicy `json:"matchPolicyEffective,omitempty"`

	MatchPolicyMine *AssetMatchPolicy `json:"matchPolicyMine,omitempty"`

	// Metadata: The metadata object contains information that identifies
	// and describes the asset. This information could be used to search for
	// the asset or to eliminate duplication within YouTube's database.
	Metadata *Metadata `json:"metadata,omitempty"`

	MetadataEffective *Metadata `json:"metadataEffective,omitempty"`

	MetadataMine *Metadata `json:"metadataMine,omitempty"`

	// Ownership: The ownership object identifies an asset's owners and
	// provides additional details about their ownership, such as the
	// territories where they own the asset.
	Ownership *RightsOwnership `json:"ownership,omitempty"`

	// OwnershipConflicts: The ownershipConflicts object contains
	// information about the asset's ownership conflicts.
	OwnershipConflicts *OwnershipConflicts `json:"ownershipConflicts,omitempty"`

	OwnershipEffective *RightsOwnership `json:"ownershipEffective,omitempty"`

	OwnershipMine *RightsOwnership `json:"ownershipMine,omitempty"`

	// Status: The asset's status.
	Status string `json:"status,omitempty"`

	// TimeCreated: The date and time the asset was created. The value is
	// specified in RFC 3339 (YYYY-MM-DDThh:mm:ss.000Z) format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// Type: The asset's type. This value determines the metadata fields
	// that you can set for the asset. In addition, certain API functions
	// may only be supported for specific types of assets. For example,
	// composition assets may have more complex ownership data than other
	// types of assets.
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AliasId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AliasId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Asset) MarshalJSON() ([]byte, error) {
	type noMethod Asset
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetLabel struct {
	// Kind: The type of the API resource. For assetLabel resources, this
	// value is youtubePartner#assetLabel.
	Kind string `json:"kind,omitempty"`

	// LabelName: Name of the asset label.
	LabelName string `json:"labelName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetLabel) MarshalJSON() ([]byte, error) {
	type noMethod AssetLabel
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetLabelListResponse struct {
	// Items: A list of assetLabel resources that match the request
	// criteria.
	Items []*AssetLabel `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#assetLabelList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetLabelListResponse) MarshalJSON() ([]byte, error) {
	type noMethod AssetLabelListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetListResponse struct {
	// Items: A list of asset resources that match the request criteria.
	Items []*Asset `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#assetList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetListResponse) MarshalJSON() ([]byte, error) {
	type noMethod AssetListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetMatchPolicy struct {
	// Kind: The type of the API resource. Value:
	// youtubePartner#assetMatchPolicy.
	Kind string `json:"kind,omitempty"`

	// PolicyId: A value that uniquely identifies the Policy resource that
	// YouTube applies to user-uploaded videos that match the asset.
	PolicyId string `json:"policyId,omitempty"`

	// Rules: A list of rules that collectively define the policy that the
	// content owner wants to apply to user-uploaded videos that match the
	// asset. Each rule specifies the action that YouTube should take and
	// may optionally specify the conditions under which that action is
	// enforced.
	Rules []*PolicyRule `json:"rules,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetMatchPolicy) MarshalJSON() ([]byte, error) {
	type noMethod AssetMatchPolicy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetRelationship struct {
	// ChildAssetId: The ID of the child (contained) asset.
	ChildAssetId string `json:"childAssetId,omitempty"`

	// Id: A value that YouTube assigns and uses to uniquely identify the
	// asset relationship.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For this resource, the value is
	// youtubePartner#assetRelationship.
	Kind string `json:"kind,omitempty"`

	// ParentAssetId: The ID of the parent (containing) asset.
	ParentAssetId string `json:"parentAssetId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ChildAssetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ChildAssetId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetRelationship) MarshalJSON() ([]byte, error) {
	type noMethod AssetRelationship
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetRelationshipListResponse struct {
	// Items: A list of assetRelationship resources that match the request
	// criteria.
	Items []*AssetRelationship `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#assetRelationshipList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetRelationshipListResponse) MarshalJSON() ([]byte, error) {
	type noMethod AssetRelationshipListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetSearchResponse struct {
	// Items: A list of asset resources that match the request criteria.
	Items []*AssetSnippet `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#assetSnippetList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetSearchResponse) MarshalJSON() ([]byte, error) {
	type noMethod AssetSearchResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetShare struct {
	// Kind: The type of the API resource. For this resource, the value is
	// youtubePartner#assetShare.
	Kind string `json:"kind,omitempty"`

	// ShareId: A value that YouTube assigns and uses to uniquely identify
	// the asset share.
	ShareId string `json:"shareId,omitempty"`

	// ViewId: A value that YouTube assigns and uses to uniquely identify
	// the asset view.
	ViewId string `json:"viewId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetShare) MarshalJSON() ([]byte, error) {
	type noMethod AssetShare
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetShareListResponse struct {
	// Items: An assetShare resource that matches the request criteria.
	Items []*AssetShare `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#assetShareList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetShareListResponse) MarshalJSON() ([]byte, error) {
	type noMethod AssetShareListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AssetSnippet struct {
	// CustomId: Custom ID assigned by the content owner to this asset.
	CustomId string `json:"customId,omitempty"`

	// Id: An ID that YouTube assigns and uses to uniquely identify the
	// asset.
	Id string `json:"id,omitempty"`

	// Isrc: The ISRC (International Standard Recording Code) for this
	// asset.
	Isrc string `json:"isrc,omitempty"`

	// Iswc: The ISWC (International Standard Musical Work Code) for this
	// asset.
	Iswc string `json:"iswc,omitempty"`

	// Kind: The type of the API resource. For this operation, the value is
	// youtubePartner#assetSnippet.
	Kind string `json:"kind,omitempty"`

	// TimeCreated: The date and time the asset was created. The value is
	// specified in RFC 3339 (YYYY-MM-DDThh:mm:ss.000Z) format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// Title: Title of this asset.
	Title string `json:"title,omitempty"`

	// Type: The asset's type. This value determines which metadata fields
	// might be included in the metadata object.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CustomId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetSnippet) MarshalJSON() ([]byte, error) {
	type noMethod AssetSnippet
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Campaign struct {
	// CampaignData: The campaignData object contains details like the
	// campaign's start and end dates, target and source.
	CampaignData *CampaignData `json:"campaignData,omitempty"`

	// Id: The unique ID that YouTube uses to identify the campaign.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For campaign resources, this
	// value is youtubePartner#campaign.
	Kind string `json:"kind,omitempty"`

	// Status: The status of the campaign.
	Status string `json:"status,omitempty"`

	// TimeCreated: The time the campaign was created.
	TimeCreated string `json:"timeCreated,omitempty"`

	// TimeLastModified: The time the campaign was last modified.
	TimeLastModified string `json:"timeLastModified,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CampaignData") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CampaignData") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Campaign) MarshalJSON() ([]byte, error) {
	type noMethod Campaign
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CampaignData struct {
	// CampaignSource: The campaignSource object contains information about
	// the assets for which the campaign will generate links.
	CampaignSource *CampaignSource `json:"campaignSource,omitempty"`

	// ExpireTime: The time at which the campaign should expire. Do not
	// specify a value if the campaign has no expiration time.
	ExpireTime string `json:"expireTime,omitempty"`

	// Name: The user-given name of the campaign.
	Name string `json:"name,omitempty"`

	// PromotedContent: A list of videos or channels that will be linked to
	// from claimed videos that are included in the campaign.
	PromotedContent []*PromotedContent `json:"promotedContent,omitempty"`

	// StartTime: The time at which the campaign should start. Do not
	// specify a value if the campaign should start immediately.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CampaignSource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CampaignSource") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CampaignData) MarshalJSON() ([]byte, error) {
	type noMethod CampaignData
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CampaignList struct {
	// Items: A list of campaigns.
	Items []*Campaign `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#campaignList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CampaignList) MarshalJSON() ([]byte, error) {
	type noMethod CampaignList
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CampaignSource struct {
	// SourceType: The type of the campaign source.
	SourceType string `json:"sourceType,omitempty"`

	// SourceValue: A list of values of the campaign source.
	SourceValue []string `json:"sourceValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SourceType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SourceType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CampaignSource) MarshalJSON() ([]byte, error) {
	type noMethod CampaignSource
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CampaignTargetLink struct {
	// TargetId: The channel ID or video ID of the link target.
	TargetId string `json:"targetId,omitempty"`

	// TargetType: Indicates whether the link target is a channel or video.
	TargetType string `json:"targetType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TargetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TargetId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CampaignTargetLink) MarshalJSON() ([]byte, error) {
	type noMethod CampaignTargetLink
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Claim struct {
	// AppliedPolicy: The applied policy for the viewing owner on the claim.
	// This might not be the same as the final claim policy on the video as
	// it does not consider other partners' policy of the same claim.
	AppliedPolicy *Policy `json:"appliedPolicy,omitempty"`

	// AssetId: The unique YouTube asset ID that identifies the asset
	// associated with the claim.
	AssetId string `json:"assetId,omitempty"`

	// BlockOutsideOwnership: Indicates whether or not the claimed video
	// should be blocked anywhere it is not explicitly owned.
	BlockOutsideOwnership bool `json:"blockOutsideOwnership,omitempty"`

	// ContentType: This value indicates whether the claim covers the audio,
	// video, or audiovisual portion of the claimed content.
	ContentType string `json:"contentType,omitempty"`

	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// claim.
	Id string `json:"id,omitempty"`

	// IsPartnerUploaded: Indicates whether or not the claim is a partner
	// uploaded claim.
	IsPartnerUploaded bool `json:"isPartnerUploaded,omitempty"`

	// Kind: The type of the API resource. For claim resources, this value
	// is youtubePartner#claim.
	Kind string `json:"kind,omitempty"`

	// MatchInfo: If this claim was auto-generated based on a provided
	// reference, this section will provide details of the match that
	// generated the claim.
	MatchInfo *ClaimMatchInfo `json:"matchInfo,omitempty"`

	Origin *ClaimOrigin `json:"origin,omitempty"`

	// Policy: The policy provided by the viewing owner on the claim.
	Policy *Policy `json:"policy,omitempty"`

	// Status: The claim's status. When updating a claim, you can update its
	// status from active to inactive to effectively release the claim, but
	// the API does not support other updates to a claim's status.
	Status string `json:"status,omitempty"`

	// TimeCreated: The time the claim was created.
	TimeCreated string `json:"timeCreated,omitempty"`

	// VideoId: The unique YouTube video ID that identifies the video
	// associated with the claim.
	VideoId string `json:"videoId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppliedPolicy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppliedPolicy") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Claim) MarshalJSON() ([]byte, error) {
	type noMethod Claim
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimMatchInfo: If this claim was auto-generated based on a provided
// reference, this section will provide details of the match that
// generated the claim.
type ClaimMatchInfo struct {
	// LongestMatch: Details of the longest match between the reference and
	// the user video.
	LongestMatch *ClaimMatchInfoLongestMatch `json:"longestMatch,omitempty"`

	// MatchSegments: Details about each match segment. Each item in the
	// list contains information about one match segment associated with the
	// claim. It is possible to have multiple match segments. For example,
	// if the audio and video content of an uploaded video match that of a
	// reference video, there would be two match segments. One segment would
	// describe the audio match and the other would describe the video
	// match.
	MatchSegments []*MatchSegment `json:"matchSegments,omitempty"`

	// ReferenceId: The reference ID that generated this match.
	ReferenceId string `json:"referenceId,omitempty"`

	// TotalMatch: Details of the total amount of reference and user video
	// content which matched each other. Note these two values may differ if
	// either the reference or the user video contains a loop.
	TotalMatch *ClaimMatchInfoTotalMatch `json:"totalMatch,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LongestMatch") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LongestMatch") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimMatchInfo) MarshalJSON() ([]byte, error) {
	type noMethod ClaimMatchInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimMatchInfoLongestMatch: Details of the longest match between the
// reference and the user video.
type ClaimMatchInfoLongestMatch struct {
	// DurationSecs: The duration of the longest match between the reference
	// and the user video.
	DurationSecs uint64 `json:"durationSecs,omitempty,string"`

	// ReferenceOffset: The offset in seconds into the reference at which
	// the longest match began.
	ReferenceOffset uint64 `json:"referenceOffset,omitempty,string"`

	// UserVideoOffset: The offset in seconds into the user video at which
	// the longest match began.
	UserVideoOffset uint64 `json:"userVideoOffset,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "DurationSecs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DurationSecs") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimMatchInfoLongestMatch) MarshalJSON() ([]byte, error) {
	type noMethod ClaimMatchInfoLongestMatch
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimMatchInfoTotalMatch: Details of the total amount of reference
// and user video content which matched each other. Note these two
// values may differ if either the reference or the user video contains
// a loop.
type ClaimMatchInfoTotalMatch struct {
	// ReferenceDurationSecs: The total amount of content in the reference
	// which matched the user video in seconds.
	ReferenceDurationSecs uint64 `json:"referenceDurationSecs,omitempty,string"`

	// UserVideoDurationSecs: The total amount of content in the user video
	// which matched the reference in seconds.
	UserVideoDurationSecs uint64 `json:"userVideoDurationSecs,omitempty,string"`

	// ForceSendFields is a list of field names (e.g.
	// "ReferenceDurationSecs") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ReferenceDurationSecs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClaimMatchInfoTotalMatch) MarshalJSON() ([]byte, error) {
	type noMethod ClaimMatchInfoTotalMatch
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimOrigin struct {
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Source") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Source") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimOrigin) MarshalJSON() ([]byte, error) {
	type noMethod ClaimOrigin
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimEvent struct {
	// Kind: The type of the API resource. For claimEvent resources, this
	// value is youtubePartner#claimEvent.
	Kind string `json:"kind,omitempty"`

	// Reason: Reason of the event.
	Reason string `json:"reason,omitempty"`

	// Source: Data related to source of the event.
	Source *ClaimEventSource `json:"source,omitempty"`

	// Time: The time when the event occurred.
	Time string `json:"time,omitempty"`

	// Type: Type of the event.
	Type string `json:"type,omitempty"`

	// TypeDetails: Details of event's type.
	TypeDetails *ClaimEventTypeDetails `json:"typeDetails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimEvent) MarshalJSON() ([]byte, error) {
	type noMethod ClaimEvent
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimEventSource: Data related to source of the event.
type ClaimEventSource struct {
	// ContentOwnerId: Id of content owner that initiated the event.
	ContentOwnerId string `json:"contentOwnerId,omitempty"`

	// Type: Type of the event source.
	Type string `json:"type,omitempty"`

	// UserEmail: Email of user who initiated the event.
	UserEmail string `json:"userEmail,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContentOwnerId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentOwnerId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClaimEventSource) MarshalJSON() ([]byte, error) {
	type noMethod ClaimEventSource
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClaimEventTypeDetails: Details of event's type.
type ClaimEventTypeDetails struct {
	// AppealExplanation: Appeal explanations for dispute_appeal event.
	AppealExplanation string `json:"appealExplanation,omitempty"`

	// DisputeNotes: Dispute notes for dispute_create events.
	DisputeNotes string `json:"disputeNotes,omitempty"`

	// DisputeReason: Dispute reason for dispute_create and dispute_appeal
	// events.
	DisputeReason string `json:"disputeReason,omitempty"`

	// UpdateStatus: Status that was a result of update for claim_update
	// event.
	UpdateStatus string `json:"updateStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppealExplanation")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppealExplanation") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClaimEventTypeDetails) MarshalJSON() ([]byte, error) {
	type noMethod ClaimEventTypeDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimHistory struct {
	// Event: A list of claim history events.
	Event []*ClaimEvent `json:"event,omitempty"`

	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// claim.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For claimHistory resources, this
	// value is youtubePartner#claimHistory.
	Kind string `json:"kind,omitempty"`

	// UploaderChannelId: The external channel id of claimed video's
	// uploader.
	UploaderChannelId string `json:"uploaderChannelId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Event") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Event") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimHistory) MarshalJSON() ([]byte, error) {
	type noMethod ClaimHistory
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimListResponse struct {
	// Items: A list of claims that match the request criteria.
	Items []*Claim `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#claimList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PreviousPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PreviousPageToken string `json:"previousPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ClaimListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimSearchResponse struct {
	// Items: A list of claims that match the request criteria.
	Items []*ClaimSnippet `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#claimSnippetList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PreviousPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PreviousPageToken string `json:"previousPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimSearchResponse) MarshalJSON() ([]byte, error) {
	type noMethod ClaimSearchResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimSnippet struct {
	// AssetId: The unique YouTube asset ID that identifies the asset
	// associated with the claim.
	AssetId string `json:"assetId,omitempty"`

	// ContentType: This value indicates whether the claim covers the audio,
	// video, or audiovisual portion of the claimed content.
	ContentType string `json:"contentType,omitempty"`

	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// claim.
	Id string `json:"id,omitempty"`

	// IsPartnerUploaded: Indicates whether or not the claim is a partner
	// uploaded claim.
	IsPartnerUploaded bool `json:"isPartnerUploaded,omitempty"`

	// Kind: The type of the API resource. For this operation, the value is
	// youtubePartner#claimSnippet.
	Kind string `json:"kind,omitempty"`

	Origin *ClaimSnippetOrigin `json:"origin,omitempty"`

	// Status: The claim's status.
	Status string `json:"status,omitempty"`

	// ThirdPartyClaim: Indicates that this is a third party claim.
	ThirdPartyClaim bool `json:"thirdPartyClaim,omitempty"`

	// TimeCreated: The time the claim was created.
	TimeCreated string `json:"timeCreated,omitempty"`

	// TimeStatusLastModified: The time the claim status and/or status
	// detail was last modified.
	TimeStatusLastModified string `json:"timeStatusLastModified,omitempty"`

	// VideoId: The unique YouTube video ID that identifies the video
	// associated with the claim.
	VideoId string `json:"videoId,omitempty"`

	// VideoTitle: The title of the claimed video.
	VideoTitle string `json:"videoTitle,omitempty"`

	// VideoViews: Number of views for the claimed video.
	VideoViews uint64 `json:"videoViews,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "AssetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AssetId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimSnippet) MarshalJSON() ([]byte, error) {
	type noMethod ClaimSnippet
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimSnippetOrigin struct {
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Source") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Source") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClaimSnippetOrigin) MarshalJSON() ([]byte, error) {
	type noMethod ClaimSnippetOrigin
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClaimedVideoDefaults struct {
	// AutoGeneratedBreaks: Set this property to true to enable
	// automatically generated breaks for a newly claimed video longer than
	// 10 minutes. The first partner that claims the video sets its default
	// advertising options to that video.
	// claimedVideoOptions.auto_generated_breaks_default
	AutoGeneratedBreaks bool `json:"autoGeneratedBreaks,omitempty"`

	// ChannelOverride: Set this property to true to indicate that the
	// channel's claimedVideoOptions can override the content owner's
	// claimedVideoOptions.
	ChannelOverride bool `json:"channelOverride,omitempty"`

	// Kind: Identifies this resource as default options for newly claimed
	// video. Value: "youtubePartner#claimedVideoDefaults".
	Kind string `json:"kind,omitempty"`

	// NewVideoDefaults: A list of ad formats that could be used as the
	// default settings for a newly claimed video. The first partner that
	// claims the video sets its default advertising options to that video.
	NewVideoDefaults []string `json:"newVideoDefaults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AutoGeneratedBreaks")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutoGeneratedBreaks") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClaimedVideoDefaults) MarshalJSON() ([]byte, error) {
	type noMethod ClaimedVideoDefaults
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Conditions struct {
	// ContentMatchType: This match condition specifies whether the user- or
	// partner-uploaded content needs to match the audio, video or
	// audiovisual content of a reference file for the rule to apply.
	ContentMatchType []string `json:"contentMatchType,omitempty"`

	// MatchDuration: This match condition specifies an amount of time that
	// the user- or partner- uploaded content needs to match a reference
	// file for the rule to apply.
	MatchDuration []*IntervalCondition `json:"matchDuration,omitempty"`

	// MatchPercent: This match condition specifies a percentage of the
	// user- or partner-uploaded content that needs to match a reference
	// file for the rule to apply.
	MatchPercent []*IntervalCondition `json:"matchPercent,omitempty"`

	// ReferenceDuration: This match condition indicates that the reference
	// must be a certain duration for the rule to apply.
	ReferenceDuration []*IntervalCondition `json:"referenceDuration,omitempty"`

	// ReferencePercent: This match condition indicates that the specified
	// percentage of a reference file must match the user- or
	// partner-uploaded content for the rule to apply.
	ReferencePercent []*IntervalCondition `json:"referencePercent,omitempty"`

	// RequiredTerritories: This watch condition specifies where users are
	// (or or not) allowed to watch (or listen to) an asset. YouTube
	// determines whether the condition is satisfied based on the user's
	// location.
	RequiredTerritories *TerritoryCondition `json:"requiredTerritories,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContentMatchType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentMatchType") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Conditions) MarshalJSON() ([]byte, error) {
	type noMethod Conditions
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ConflictingOwnership struct {
	// Owner: The ID of the conflicting asset's owner.
	Owner string `json:"owner,omitempty"`

	// Ratio: The percentage of the asset that the owner controls or
	// administers.
	Ratio float64 `json:"ratio,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Owner") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Owner") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ConflictingOwnership) MarshalJSON() ([]byte, error) {
	type noMethod ConflictingOwnership
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *ConflictingOwnership) UnmarshalJSON(data []byte) error {
	type noMethod ConflictingOwnership
	var s1 struct {
		Ratio gensupport.JSONFloat64 `json:"ratio"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Ratio = float64(s1.Ratio)
	return nil
}

type ContentOwner struct {
	// ConflictNotificationEmail: The email address visible to other
	// partners for use in managing asset ownership conflicts.
	ConflictNotificationEmail string `json:"conflictNotificationEmail,omitempty"`

	// DisplayName: The content owner's display name.
	DisplayName string `json:"displayName,omitempty"`

	// DisputeNotificationEmails: The email address(es) to which YouTube
	// sends claim dispute notifications and possible claim notifications.
	DisputeNotificationEmails []string `json:"disputeNotificationEmails,omitempty"`

	// FingerprintReportNotificationEmails: The email address(es) to which
	// YouTube sends fingerprint reports.
	FingerprintReportNotificationEmails []string `json:"fingerprintReportNotificationEmails,omitempty"`

	// Id: A unique ID that YouTube uses to identify the content owner.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For content owner resources, the
	// value is youtubePartner#contentOwner.
	Kind string `json:"kind,omitempty"`

	// PrimaryNotificationEmails: The email address(es) to which YouTube
	// sends CMS account details and report notifications.
	PrimaryNotificationEmails []string `json:"primaryNotificationEmails,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ConflictNotificationEmail") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "ConflictNotificationEmail") to include in API requests with the JSON
	// null value. By default, fields with empty values are omitted from API
	// requests. However, any field with an empty value appearing in
	// NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContentOwner) MarshalJSON() ([]byte, error) {
	type noMethod ContentOwner
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ContentOwnerAdvertisingOption struct {
	// AllowedOptions: This object identifies the ad formats that the
	// content owner is allowed to use.
	AllowedOptions *AllowedAdvertisingOptions `json:"allowedOptions,omitempty"`

	// ClaimedVideoOptions: This object identifies the advertising options
	// used by default for the content owner's newly claimed videos.
	ClaimedVideoOptions *ClaimedVideoDefaults `json:"claimedVideoOptions,omitempty"`

	// Id: The value that YouTube uses to uniquely identify the content
	// owner.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For this resource, the value is
	// youtubePartner#contentOwnerAdvertisingOption.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AllowedOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllowedOptions") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ContentOwnerAdvertisingOption) MarshalJSON() ([]byte, error) {
	type noMethod ContentOwnerAdvertisingOption
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ContentOwnerListResponse struct {
	// Items: A list of content owners that match the request criteria.
	Items []*ContentOwner `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#contentOwnerList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContentOwnerListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ContentOwnerListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CountriesRestriction struct {
	// AdFormats: A list of ad formats that can be used in the specified
	// countries.
	AdFormats []string `json:"adFormats,omitempty"`

	// Territories: A list of ISO 3166-1 alpha-2 country codes that identify
	// the countries where ads are enabled.
	Territories []string `json:"territories,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdFormats") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdFormats") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CountriesRestriction) MarshalJSON() ([]byte, error) {
	type noMethod CountriesRestriction
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CuepointSettings struct {
	// CueType: The cuepoint's type. See the Getting started guide for an
	// explanation of the different types of cuepoints. Also see the Life of
	// a broadcast document for best practices about inserting cuepoints
	// during your broadcast.
	CueType string `json:"cueType,omitempty"`

	// DurationSecs: The cuepoint's duration, in seconds. This value must be
	// specified if the cueType is ad and is ignored otherwise.
	DurationSecs int64 `json:"durationSecs,omitempty"`

	// OffsetTimeMs: This value specifies a point in time in the video when
	// viewers should see an ad or in-stream slate. The property value
	// identifies a time offset, in milliseconds, from the beginning of the
	// monitor stream. Though measured in milliseconds, the value is
	// actually an approximation, and YouTube will insert the cuepoint as
	// closely as possible to that time. You should not specify a value for
	// this parameter if your broadcast does not have a monitor
	// stream.
	//
	// This property's default value is 0, which indicates that the cuepoint
	// should be inserted as soon as possible. If your broadcast stream is
	// not delayed, then 0 is also the only valid value. However, if your
	// broadcast stream is delayed, then the property value can specify the
	// time when the cuepoint should be inserted. See the Getting started
	// guide for more details.
	//
	// Note: If your broadcast had a testing phase, the offset is measured
	// from the time that the testing phase began.
	OffsetTimeMs int64 `json:"offsetTimeMs,omitempty,string"`

	// Walltime: This value specifies the wall clock time at which the
	// cuepoint should be inserted. The value is specified in ISO 8601
	// (YYYY-MM-DDThh:mm:ss.sssZ) format.
	Walltime string `json:"walltime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CueType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CueType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CuepointSettings) MarshalJSON() ([]byte, error) {
	type noMethod CuepointSettings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Date struct {
	// Day: The date's day. The value should be an integer between 1 and 31.
	// Note that some day-month combinations are not valid.
	Day int64 `json:"day,omitempty"`

	// Month: The date's month. The value should be an integer between 1 and
	// 12.
	Month int64 `json:"month,omitempty"`

	// Year: The date's year in the Gregorian Calendar. Assumed to be A.D.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type noMethod Date
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DateRange struct {
	// End: The end date (inclusive) for the date range. This value is
	// required for video-on-demand (VOD) orders and optional for electronic
	// sell-through (EST) orders.
	End *Date `json:"end,omitempty"`

	// Kind: Identifies this resource as order date range. Value:
	// "youtubePartner#dateRange".
	Kind string `json:"kind,omitempty"`

	// Start: The start date for the date range. This value is required for
	// all date ranges.
	Start *Date `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g. "End") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "End") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DateRange) MarshalJSON() ([]byte, error) {
	type noMethod DateRange
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ExcludedInterval struct {
	// High: The end (inclusive) time in seconds of the time window. The
	// value can be any value greater than low. If high is greater than the
	// length of the reference, the interval between low and the end of the
	// reference will be excluded. Every interval must specify a value for
	// this field.
	High float64 `json:"high,omitempty"`

	// Low: The start (inclusive) time in seconds of the time window. The
	// value can be any value between 0 and high. Every interval must
	// specify a value for this field.
	Low float64 `json:"low,omitempty"`

	// Origin: The source of the request to exclude the interval from
	// Content ID matching.
	Origin string `json:"origin,omitempty"`

	// TimeCreated: The date and time that the exclusion was created. The
	// value is specified in RFC 3339 (YYYY-MM-DDThh:mm:ss.000Z) format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// ForceSendFields is a list of field names (e.g. "High") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "High") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExcludedInterval) MarshalJSON() ([]byte, error) {
	type noMethod ExcludedInterval
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *ExcludedInterval) UnmarshalJSON(data []byte) error {
	type noMethod ExcludedInterval
	var s1 struct {
		High gensupport.JSONFloat64 `json:"high"`
		Low  gensupport.JSONFloat64 `json:"low"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.High = float64(s1.High)
	s.Low = float64(s1.Low)
	return nil
}

type IntervalCondition struct {
	// High: The maximum (inclusive) allowed value for the condition to be
	// satisfied. The default value is ∞.
	High float64 `json:"high,omitempty"`

	// Low: The minimum (inclusive) allowed value for the condition to be
	// satisfied. The default value is -∞.
	Low float64 `json:"low,omitempty"`

	// ForceSendFields is a list of field names (e.g. "High") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "High") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntervalCondition) MarshalJSON() ([]byte, error) {
	type noMethod IntervalCondition
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *IntervalCondition) UnmarshalJSON(data []byte) error {
	type noMethod IntervalCondition
	var s1 struct {
		High gensupport.JSONFloat64 `json:"high"`
		Low  gensupport.JSONFloat64 `json:"low"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.High = float64(s1.High)
	s.Low = float64(s1.Low)
	return nil
}

type LiveCuepoint struct {
	// BroadcastId: The ID that YouTube assigns to uniquely identify the
	// broadcast into which the cuepoint is being inserted.
	BroadcastId string `json:"broadcastId,omitempty"`

	// Id: A value that YouTube assigns to uniquely identify the cuepoint.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For liveCuepoint resources, the
	// value is youtubePartner#liveCuepoint.
	Kind string `json:"kind,omitempty"`

	// Settings: The settings object defines the cuepoint's settings.
	Settings *CuepointSettings `json:"settings,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BroadcastId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BroadcastId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LiveCuepoint) MarshalJSON() ([]byte, error) {
	type noMethod LiveCuepoint
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type MatchSegment struct {
	// Channel: Identifies the manner in which the claimed video matches the
	// reference video.
	Channel string `json:"channel,omitempty"`

	// ReferenceSegment: The reference_segment object contains information
	// about the matched portion of the reference content.
	ReferenceSegment *Segment `json:"reference_segment,omitempty"`

	// VideoSegment: The video_segment object contains information about the
	// matched portion of the claimed video.
	VideoSegment *Segment `json:"video_segment,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Channel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Channel") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MatchSegment) MarshalJSON() ([]byte, error) {
	type noMethod MatchSegment
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Metadata struct {
	// Actor: A list that identifies actors associated with the asset. You
	// can specify up to 50 actors for an asset.
	Actor []string `json:"actor,omitempty"`

	// Album: The album on which a sound recording asset is included. This
	// field is only valid for sound recording assets and has a maximum
	// length of 255 bytes.
	Album string `json:"album,omitempty"`

	// Artist: The artist associated with a music video or sound recording
	// asset. This field is only valid for music video and sound recording
	// assets. It is required for sound recordings included in the AudioSwap
	// program.
	Artist []string `json:"artist,omitempty"`

	// Broadcaster: Identifies the network or channel that originally
	// broadcast a show or a season of a show. This field should only be
	// included for an asset if the broadcaster associated with the asset is
	// different from the partner uploading the asset to YouTube. Note that
	// a show may have multiple broadcasters; for example, a show may switch
	// networks between seasons.
	Broadcaster []string `json:"broadcaster,omitempty"`

	// Category: Category of this asset.
	Category string `json:"category,omitempty"`

	// ContentType: The type of video content that the asset represents.
	// This field is only valid for movie and episode assets, and is
	// required for the following types of those assets:
	// - Episode assets that are linked to a show
	// - Movie assets that appear in YouTube's Movies category
	ContentType string `json:"contentType,omitempty"`

	// CopyrightDate: The date copyright for this asset was established. *
	CopyrightDate *Date `json:"copyrightDate,omitempty"`

	// CustomId: A unique value that you, the metadata provider, use to
	// identify an asset. The value could be a unique ID that you created
	// for the asset or a standard identifier, such as an ISRC. The value
	// has a maximum length of 64 bytes and may contain alphanumeric
	// characters, hyphens (-), underscores (_), periods (.), "at" symbols
	// (@), or forward slashes (/).
	CustomId string `json:"customId,omitempty"`

	// Description: A description of the asset. The description may be
	// displayed on YouTube or in CMS. This field has a maximum length of
	// 5,000 bytes.
	Description string `json:"description,omitempty"`

	// Director: A list that identifies directors associated with the asset.
	// You can specify up to 50 directors for an asset.
	Director []string `json:"director,omitempty"`

	// Eidr: The Entertainment Identifier Registry (EIDR) assigned to an
	// asset. This value is only used for episode and movie assets and is
	// optional in both cases. The value contains a standard prefix for EIDR
	// registry, followed by a forward slash, a 20-character hexadecimal
	// string, and an alphanumeric (0-9A-Z) check character.
	Eidr string `json:"eidr,omitempty"`

	// EndYear: The last year that a television show aired. This value is
	// only used for show assets, for which it is optional. Do not specify a
	// value if new show episodes are still being created.
	EndYear int64 `json:"endYear,omitempty"`

	// EpisodeNumber: The episode number associated with an episode asset.
	// This field is required for and only used for episode assets that are
	// linked to show assets. It has a maximum length of 5 bytes.
	EpisodeNumber string `json:"episodeNumber,omitempty"`

	// EpisodesAreUntitled: This value indicates that the episodes
	// associated with a particular show asset or a particular season asset
	// are untitled. An untitled show (or season) has episodes which are
	// identified by their episode number or date. If this field is set to
	// true, then YouTube will optimize the title displayed for associated
	// episodes.
	EpisodesAreUntitled bool `json:"episodesAreUntitled,omitempty"`

	// Genre: This field specifies a genre that can be used to categorize an
	// asset. Assets may be categorized in more than one genre, and YouTube
	// uses different sets of genres to categorize different types of
	// assets. For example, Soaps might be a valid genre for a show but not
	// for a movie or sound recording.
	// - Show assets
	// - Movie assets that appear in YouTube's Movies category
	// - Sound recordings included in the AudioSwap program
	Genre []string `json:"genre,omitempty"`

	// Grid: The GRID (Global Release Identifier) of a music video or sound
	// recording. This field's value must contain exactly 18 alphanumeric
	// characters.
	Grid string `json:"grid,omitempty"`

	// Hfa: The six-character Harry Fox Agency (HFA) song code issued to
	// uniquely identify a composition. This value is only valid for
	// composition assets.
	Hfa string `json:"hfa,omitempty"`

	// InfoUrl: An official URL associated with the asset. This field has a
	// maximum length of 1536 bytes. Please do not submit a 1537-byte URL.
	// Your efforts would be futile.
	InfoUrl string `json:"infoUrl,omitempty"`

	// Isan: The ISAN (International Standard Audiovisual Number) for the
	// asset. This value is only used for episode and movie assets and is
	// optional in both cases. The value contains 26 characters, which
	// includes the 24 hexadecimal characters of the ISAN as well as two
	// check characters, in the following format:
	// - The first 16 characters in the tag value contain hexadecimal
	// characters specifying the 'root' and 'episode' components of the
	// ISAN.
	// - The seventeenth character is a check character (a letter from A-Z).
	//
	// - Characters 18 to 25 are the remaining eight characters of the ISAN,
	// which specify the 'version' component of the ISAN.
	// - The twenty-sixth character is another check character (A-Z).
	Isan string `json:"isan,omitempty"`

	// Isrc: The ISRC (International Standard Recording Code) of a music
	// video or sound recording asset. This field's value must contain
	// exactly 12 alphanumeric characters.
	Isrc string `json:"isrc,omitempty"`

	// Iswc: The ISWC (International Standard Musical Work Code) for a
	// composition asset. The field's value must contain exactly 11
	// characters in the format of a letter (T) followed by 10 digits.
	Iswc string `json:"iswc,omitempty"`

	// Keyword: A list of up to 100 keywords associated with a show asset.
	// This field is required for and also only used for show assets.
	Keyword []string `json:"keyword,omitempty"`

	// Label: The record label that released a sound recording asset. This
	// field is only valid for sound recording assets and has a maximum
	// length of 255 bytes.
	Label string `json:"label,omitempty"`

	// Notes: Additional information that does not map directly to one of
	// the other metadata fields. This field has a maximum length of 255
	// bytes.
	Notes string `json:"notes,omitempty"`

	// OriginalReleaseMedium: The method by which people first had the
	// opportunity to see a video asset. This value is only used for episode
	// and movie assets. It is required for the assets listed below and
	// otherwise optional:
	// - Episode assets that are linked to a show
	// - Movie assets that appear in YouTube's Movies category
	OriginalReleaseMedium string `json:"originalReleaseMedium,omitempty"`

	// Producer: A list that identifies producers of the asset. You can
	// specify up to 50 producers for an asset.
	Producer []string `json:"producer,omitempty"`

	// Ratings: A list of ratings that an asset received. The rating must be
	// valid under the specified rating system.
	Ratings []*Rating `json:"ratings,omitempty"`

	// ReleaseDate: The date that an asset was publicly released. For season
	// assets, this value specifies the first date that the season aired.
	// Dates prior to the year 1902 are not supported. This value is valid
	// for episode, season, movie, music video, and sound recording assets.
	// It is required for the assets listed below and otherwise optional:
	//
	// - Episode assets that are linked to a show
	// - Movie assets that appear in YouTube's Movies category
	ReleaseDate *Date `json:"releaseDate,omitempty"`

	// SeasonNumber: The season number that identifies a season asset, or
	// the season number that is associated with an episode asset. This
	// field has a maximum length of 5 bytes.
	SeasonNumber string `json:"seasonNumber,omitempty"`

	// ShowCustomId: The customId of the show asset that a season or episode
	// asset is associated with. It is required for season and episode
	// assets that appear in the Shows category on YouTube, and it is not
	// valid for other types of assets. This field has a maximum length of
	// 64 bytes and may contain alphanumeric characters, hyphens (-),
	// underscores (_), periods (.), "at" symbols (@), or forward slashes
	// (/).
	ShowCustomId string `json:"showCustomId,omitempty"`

	// ShowTitle: The name of the show that an episode asset is associated
	// with. Note: This tag is only used for and valid for episodes that are
	// not associated with show assets and enables those assets to still
	// display a show title in the asset metadata section of CMS. This field
	// has a maximum length of 120 bytes.
	ShowTitle string `json:"showTitle,omitempty"`

	// SpokenLanguage: The video's primary spoken language. The value can be
	// any ISO 639-1 two-letter language code. This value is only used for
	// episode and movie assets and is not valid for other types of assets.
	SpokenLanguage string `json:"spokenLanguage,omitempty"`

	// StartYear: The first year that a television show aired. This value is
	// required for and also only used for show assets.
	StartYear int64 `json:"startYear,omitempty"`

	// SubtitledLanguage: A list of languages for which the video has either
	// a separate caption track or burnt-in captions that are part of the
	// video. Each value in the list should be an ISO 639-1 two-letter
	// language code. This value is only used for episode and movie assets
	// and is not valid for other types of assets.
	SubtitledLanguage []string `json:"subtitledLanguage,omitempty"`

	// Title: The asset's title or name. The value has a maximum length of
	// 255 bytes. This value is required for the assets listed below and
	// optional for all other assets:
	// - Show assets
	// - Episode assets that are linked to a show
	// - Movie assets that appear in YouTube's Movies category
	// - Sound recordings included in the AudioSwap program
	Title string `json:"title,omitempty"`

	// TmsId: TMS (Tribune Media Systems) ID for the asset.
	TmsId string `json:"tmsId,omitempty"`

	// TotalEpisodesExpected: Specifies the total number of full-length
	// episodes in the season. This value is used only for season assets.
	TotalEpisodesExpected int64 `json:"totalEpisodesExpected,omitempty"`

	// Upc: The UPC (Universal Product Code) associated with the asset.
	Upc string `json:"upc,omitempty"`

	// Writer: A list that identifies writers associated with the asset. You
	// can specify up to 50 writers for an asset.
	Writer []string `json:"writer,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Actor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Actor") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Metadata) MarshalJSON() ([]byte, error) {
	type noMethod Metadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type MetadataHistory struct {
	// Kind: The type of the API resource. For metadata history resources,
	// the value is youtubePartner#metadataHistory.
	Kind string `json:"kind,omitempty"`

	// Metadata: The metadata object contains the metadata provided by the
	// specified source (origination) at the specified time (timeProvided).
	Metadata *Metadata `json:"metadata,omitempty"`

	// Origination: The origination object contains information that
	// describes the metadata source.
	Origination *Origination `json:"origination,omitempty"`

	// TimeProvided: The time the metadata was provided.
	TimeProvided string `json:"timeProvided,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetadataHistory) MarshalJSON() ([]byte, error) {
	type noMethod MetadataHistory
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type MetadataHistoryListResponse struct {
	// Items: A list of metadata history (youtubePartner#metadataHistory)
	// resources that match the request criteria.
	Items []*MetadataHistory `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#metadataHistoryList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetadataHistoryListResponse) MarshalJSON() ([]byte, error) {
	type noMethod MetadataHistoryListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Order struct {
	// AvailGroupId: Links an order to the avails associated with it.
	AvailGroupId string `json:"availGroupId,omitempty"`

	// ChannelId: Channel ID - identifies the channel this order and video
	// are associated with
	ChannelId string `json:"channelId,omitempty"`

	// ContentType: Type of content possible values are
	// - MOVIE
	// - SHOW
	ContentType string `json:"contentType,omitempty"`

	// Country: Two letter country code for the order only countries where
	// YouTube does transactional business are allowed.
	Country string `json:"country,omitempty"`

	// CustomId: Secondary id to be used to identify content in other
	// systems like partner database
	CustomId string `json:"customId,omitempty"`

	// DvdReleaseDate: Date when this content was first made available on
	// DVD
	DvdReleaseDate *Date `json:"dvdReleaseDate,omitempty"`

	// EstDates: Range of time content is to be available for rental.
	EstDates *DateRange `json:"estDates,omitempty"`

	// Events: History log of events for this order
	Events []*StateCompleted `json:"events,omitempty"`

	// Id: Order Id unique identifier for an order.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this resource as order. Value:
	// "youtubePartner#order".
	Kind string `json:"kind,omitempty"`

	// Movie: Title if the order is type movie.
	Movie string `json:"movie,omitempty"`

	// OriginalReleaseDate: Date when this content was first made available
	// to the public
	OriginalReleaseDate *Date `json:"originalReleaseDate,omitempty"`

	// Priority: The priority for the order in the QC review queue once the
	// content is ready for QC.
	Priority string `json:"priority,omitempty"`

	// ProductionHouse: Post production house that is to process this order
	ProductionHouse string `json:"productionHouse,omitempty"`

	// PurchaseOrder: Youtube purchase order reference for the post
	// production house.
	PurchaseOrder string `json:"purchaseOrder,omitempty"`

	// Requirements: Minumim set of requirements for this order to be
	// complete such as is a trailer required.
	Requirements *Requirements `json:"requirements,omitempty"`

	// Show: Details of a show, show name, season number, episode etc.
	Show *ShowDetails `json:"show,omitempty"`

	// Status: The order's status.
	Status string `json:"status,omitempty"`

	// VideoId: Video ID the video that this order is associated with if
	// any.
	VideoId string `json:"videoId,omitempty"`

	// VodDates: Range of time content is to be available for purchase.
	VodDates *DateRange `json:"vodDates,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AvailGroupId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvailGroupId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Order) MarshalJSON() ([]byte, error) {
	type noMethod Order
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type OrderListResponse struct {
	// Items: A list of orders that match the request criteria.
	Items []*Order `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#orderList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PreviousPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PreviousPageToken string `json:"previousPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OrderListResponse) MarshalJSON() ([]byte, error) {
	type noMethod OrderListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Origination struct {
	// Owner: The content owner who provided the metadata or ownership
	// information.
	Owner string `json:"owner,omitempty"`

	// Source: The mechanism by which the piece of metadata, ownership or
	// relationship information was provided.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Owner") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Owner") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Origination) MarshalJSON() ([]byte, error) {
	type noMethod Origination
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type OwnershipConflicts struct {
	// General: A list that identifies ownership conflicts of an asset and
	// the territories where conflicting ownership is inserted.
	General []*TerritoryConflicts `json:"general,omitempty"`

	// Kind: The type of the API resource. For ownershipConflicts resources,
	// the value is youtubePartner#ownershipConflicts.
	Kind string `json:"kind,omitempty"`

	// Mechanical: A list that identifies ownership conflicts of the
	// mechanical rights for a composition asset and the territories where
	// conflicting ownership is inserted.
	Mechanical []*TerritoryConflicts `json:"mechanical,omitempty"`

	// Performance: A list that identifies ownership conflicts of the
	// performance rights for a composition asset and the territories where
	// conflicting ownership is inserted.
	Performance []*TerritoryConflicts `json:"performance,omitempty"`

	// Synchronization: A list that identifies ownership conflicts of the
	// synchronization rights for a composition asset and the territories
	// where conflicting ownership is inserted.
	Synchronization []*TerritoryConflicts `json:"synchronization,omitempty"`

	// ForceSendFields is a list of field names (e.g. "General") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "General") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OwnershipConflicts) MarshalJSON() ([]byte, error) {
	type noMethod OwnershipConflicts
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type OwnershipHistoryListResponse struct {
	// Items: A list of ownership history (youtubePartner#ownershipHistory)
	// resources that match the request criteria.
	Items []*RightsOwnershipHistory `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#ownershipHistoryList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OwnershipHistoryListResponse) MarshalJSON() ([]byte, error) {
	type noMethod OwnershipHistoryListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Package struct {
	// Content: The package's metadata file contents.
	Content string `json:"content,omitempty"`

	// CustomId: The list of customer IDs.
	CustomId []string `json:"custom_id,omitempty"`

	// Id: An ID that YouTube assigns and uses to uniquely identify the
	// package.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For package resources, this value
	// is youtubePartner#package.
	Kind string `json:"kind,omitempty"`

	// Locale: The desired locale of the error messages as defined in BCP 47
	// (http://tools.ietf.org/html/bcp47). For example, "en-US" or "de". If
	// not specified we will return the error messages in English ("en").
	Locale string `json:"locale,omitempty"`

	// Name: The package name.
	Name string `json:"name,omitempty"`

	// Status: The package status.
	Status string `json:"status,omitempty"`

	// TimeCreated: The package creation time. The value is specified in RFC
	// 3339 (YYYY-MM-DDThh:mm:ss.000Z) format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// Type: The package type.
	Type string `json:"type,omitempty"`

	// UploaderName: The uploader name.
	UploaderName string `json:"uploaderName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Package) MarshalJSON() ([]byte, error) {
	type noMethod Package
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PackageInsertResponse struct {
	// Errors: The list of errors and/or warnings.
	Errors []*ValidateError `json:"errors,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#packageInsert.
	Kind string `json:"kind,omitempty"`

	// Resource: The package resource.
	Resource *Package `json:"resource,omitempty"`

	// Status: The package insert status. Indicates whether the insert
	// operation completed successfully or identifies the general cause of
	// failure. For most cases where the insert operation failed, the errors
	// are described in the API response's errors object. However, if the
	// operation failed because the package contained non-metadata files,
	// the errors object is not included in the response.
	Status string `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PackageInsertResponse) MarshalJSON() ([]byte, error) {
	type noMethod PackageInsertResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PageInfo struct {
	// ResultsPerPage: The number of results included in the API response.
	ResultsPerPage int64 `json:"resultsPerPage,omitempty"`

	// StartIndex: The index of the first item in the API response.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results in the result set.
	TotalResults int64 `json:"totalResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResultsPerPage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResultsPerPage") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PageInfo) MarshalJSON() ([]byte, error) {
	type noMethod PageInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Policy struct {
	// Description: The policy's description.
	Description string `json:"description,omitempty"`

	// Id: A value that YouTube assigns and uses to uniquely identify the
	// policy.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this as a policy. Value: "youtubePartner#policy"
	Kind string `json:"kind,omitempty"`

	// Name: The policy's name.
	Name string `json:"name,omitempty"`

	// Rules: A list of rules that specify the action that YouTube should
	// take and may optionally specify the conditions under which that
	// action is enforced.
	Rules []*PolicyRule `json:"rules,omitempty"`

	// TimeUpdated: The time the policy was updated.
	TimeUpdated string `json:"timeUpdated,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type noMethod Policy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PolicyList struct {
	// Items: A list of saved policies.
	Items []*Policy `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#policyList.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PolicyList) MarshalJSON() ([]byte, error) {
	type noMethod PolicyList
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PolicyRule struct {
	// Action: The policy that YouTube should enforce if the rule's
	// conditions are all valid for an asset or for an attempt to view that
	// asset on YouTube.
	Action string `json:"action,omitempty"`

	// Conditions: A set of conditions that must be met for the rule's
	// action (and subactions) to be enforced. For a rule to be valid, all
	// of its conditions must be met.
	Conditions *Conditions `json:"conditions,omitempty"`

	// Subaction: A list of additional actions that YouTube should take if
	// the conditions in the rule are met.
	Subaction []string `json:"subaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PolicyRule) MarshalJSON() ([]byte, error) {
	type noMethod PolicyRule
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PromotedContent struct {
	// Link: A list of link targets that will be used to generate the
	// annotation link that appears on videos included in the  campaign. If
	// more than one link is specified, the link that is displayed to
	// viewers will be randomly selected from the list.
	Link []*CampaignTargetLink `json:"link,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Link") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Link") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PromotedContent) MarshalJSON() ([]byte, error) {
	type noMethod PromotedContent
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Publisher struct {
	// CaeNumber: The publisher's unique CAE (Compositeur, Auteur and
	// Editeur) number.
	CaeNumber string `json:"caeNumber,omitempty"`

	// Id: A value that YouTube assigns and uses to uniquely identify the
	// publisher.
	Id string `json:"id,omitempty"`

	// IpiNumber: The publisher's unique IPI (Interested Parties
	// Information) code.
	IpiNumber string `json:"ipiNumber,omitempty"`

	// Kind: The type of the API resource. For this resource, the value is
	// youtubePartner#publisher.
	Kind string `json:"kind,omitempty"`

	// Name: The publisher's name.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CaeNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CaeNumber") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Publisher) MarshalJSON() ([]byte, error) {
	type noMethod Publisher
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PublisherList struct {
	// Items: A list of publishers that match the request criteria.
	Items []*Publisher `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#publisherList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PublisherList) MarshalJSON() ([]byte, error) {
	type noMethod PublisherList
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Rating struct {
	// Rating: The rating that the asset received.
	Rating string `json:"rating,omitempty"`

	// RatingSystem: The rating system associated with the rating.
	RatingSystem string `json:"ratingSystem,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Rating") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Rating") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Rating) MarshalJSON() ([]byte, error) {
	type noMethod Rating
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Reference struct {
	// AssetId: The ID that uniquely identifies the asset that the reference
	// is associated with.
	AssetId string `json:"assetId,omitempty"`

	// AudioswapEnabled: Set this field's value to true to indicate that the
	// reference content should be included in YouTube's AudioSwap program.
	AudioswapEnabled bool `json:"audioswapEnabled,omitempty"`

	// ClaimId: This field is present if the reference was created by
	// associating an asset with an existing YouTube video that was uploaded
	// to a YouTube channel linked to your CMS account. In that case, this
	// field contains the ID of the claim representing the resulting
	// association between the asset and the video.
	ClaimId string `json:"claimId,omitempty"`

	// ContentType: The type of content that the reference represents.
	ContentType string `json:"contentType,omitempty"`

	// DuplicateLeader: The ID that uniquely identifies the reference that
	// this reference duplicates. This field is only present if the
	// reference's status is inactive with reason
	// REASON_DUPLICATE_FOR_OWNERS.
	DuplicateLeader string `json:"duplicateLeader,omitempty"`

	// ExcludedIntervals: The list of time intervals from this reference
	// that will be ignored during the match process.
	ExcludedIntervals []*ExcludedInterval `json:"excludedIntervals,omitempty"`

	// FpDirect: When uploading a reference, set this value to true to
	// indicate that the reference is a pre-generated fingerprint.
	FpDirect bool `json:"fpDirect,omitempty"`

	// HashCode: The MD5 hashcode of the reference content.
	HashCode string `json:"hashCode,omitempty"`

	// Id: A value that YouTube assigns and uses to uniquely identify a
	// reference.
	Id string `json:"id,omitempty"`

	// IgnoreFpMatch: Set this value to true to indicate that the reference
	// should not be used to generate claims. This field is only used on
	// AudioSwap references.
	IgnoreFpMatch bool `json:"ignoreFpMatch,omitempty"`

	// Kind: The type of the API resource. For reference resources, the
	// value is youtubePartner#reference.
	Kind string `json:"kind,omitempty"`

	// Length: The length of the reference in seconds.
	Length float64 `json:"length,omitempty"`

	// Origination: The origination object contains information that
	// describes the reference source.
	Origination *Origination `json:"origination,omitempty"`

	// Status: The reference's status.
	Status string `json:"status,omitempty"`

	// StatusReason: An explanation of how a reference entered its current
	// state. This value is only present if the reference's status is either
	// inactive or deleted.
	StatusReason string `json:"statusReason,omitempty"`

	// Urgent: Set this value to true to indicate that YouTube should
	// prioritize Content ID processing for a video file. YouTube processes
	// urgent video files before other files that are not marked as urgent.
	// This setting is primarily used for videos of live events or other
	// videos that require time-sensitive processing. The sooner YouTube
	// completes Content ID processing for a video, the sooner YouTube can
	// match user-uploaded videos to that video.
	//
	// Note that marking all of your files as urgent could delay processing
	// for those files.
	Urgent bool `json:"urgent,omitempty"`

	// VideoId: This field is present if the reference was created by
	// associating an asset with an existing YouTube video that was uploaded
	// to a YouTube channel linked to your CMS account. In that case, this
	// field contains the ID of the source video.
	VideoId string `json:"videoId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AssetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AssetId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Reference) MarshalJSON() ([]byte, error) {
	type noMethod Reference
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Reference) UnmarshalJSON(data []byte) error {
	type noMethod Reference
	var s1 struct {
		Length gensupport.JSONFloat64 `json:"length"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Length = float64(s1.Length)
	return nil
}

type ReferenceConflict struct {
	// ConflictingReferenceId: An id of a conflicting reference.
	ConflictingReferenceId string `json:"conflictingReferenceId,omitempty"`

	// ExpiryTime: Conflict review expiry time.
	ExpiryTime string `json:"expiryTime,omitempty"`

	// Id: A value that YouTube assigns and uses to uniquely identify a
	// reference conflict.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For referenceConflict resources,
	// the value is youtubePartner#referenceConflict.
	Kind string `json:"kind,omitempty"`

	// Matches: The list of matches between conflicting and original
	// references at the time of conflict creation.
	Matches []*ReferenceConflictMatch `json:"matches,omitempty"`

	// OriginalReferenceId: An id of an original reference.
	OriginalReferenceId string `json:"originalReferenceId,omitempty"`

	// Status: The referenceConflict's status.
	Status string `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ConflictingReferenceId") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConflictingReferenceId")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ReferenceConflict) MarshalJSON() ([]byte, error) {
	type noMethod ReferenceConflict
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ReferenceConflictListResponse struct {
	// Items: A list of reference conflicts that match the request criteria.
	Items []*ReferenceConflict `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#referenceConflictList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ReferenceConflictListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ReferenceConflictListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ReferenceConflictMatch struct {
	// ConflictingReferenceOffsetMs: Conflicting reference offset in
	// milliseconds.
	ConflictingReferenceOffsetMs int64 `json:"conflicting_reference_offset_ms,omitempty,string"`

	// LengthMs: Match length in milliseconds.
	LengthMs int64 `json:"length_ms,omitempty,string"`

	// OriginalReferenceOffsetMs: Original reference offset in milliseconds.
	OriginalReferenceOffsetMs int64 `json:"original_reference_offset_ms,omitempty,string"`

	// Type: The referenceConflictMatch's type.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ConflictingReferenceOffsetMs") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "ConflictingReferenceOffsetMs") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ReferenceConflictMatch) MarshalJSON() ([]byte, error) {
	type noMethod ReferenceConflictMatch
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ReferenceListResponse struct {
	// Items: A list of references that match the request criteria.
	Items []*Reference `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#referenceList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ReferenceListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ReferenceListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Requirements struct {
	// Caption: This value indicates whether the order requires closed
	// captions.
	Caption bool `json:"caption,omitempty"`

	// HdTranscode: This value indicates whether the order requires
	// HD-quality video.
	HdTranscode bool `json:"hdTranscode,omitempty"`

	// PosterArt: This value indicates whether the order requires poster
	// artwork.
	PosterArt bool `json:"posterArt,omitempty"`

	// SpotlightArt: This value indicates whether the order requires
	// spotlight artwork.
	SpotlightArt bool `json:"spotlightArt,omitempty"`

	// SpotlightReview: This value indicates whether the spotlight artwork
	// for the order needs to be reviewed.
	SpotlightReview bool `json:"spotlightReview,omitempty"`

	// Trailer: This value indicates whether the order requires a trailer.
	Trailer bool `json:"trailer,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Caption") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Caption") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Requirements) MarshalJSON() ([]byte, error) {
	type noMethod Requirements
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type RightsOwnership struct {
	// General: A list that identifies the owners of an asset and the
	// territories where each owner has ownership. General asset ownership
	// is used for all types of assets and is the only type of ownership
	// data that can be provided for assets that are not
	// compositions.
	//
	// Note: You cannot specify general ownership rights and also specify
	// either mechanical, performance, or synchronization rights.
	General []*TerritoryOwners `json:"general,omitempty"`

	// Kind: The type of the API resource. For rightsOwnership resources,
	// the value is youtubePartner#rightsOwnership.
	Kind string `json:"kind,omitempty"`

	// Mechanical: A list that identifies owners of the mechanical rights
	// for a composition asset.
	Mechanical []*TerritoryOwners `json:"mechanical,omitempty"`

	// Performance: A list that identifies owners of the performance rights
	// for a composition asset.
	Performance []*TerritoryOwners `json:"performance,omitempty"`

	// Synchronization: A list that identifies owners of the synchronization
	// rights for a composition asset.
	Synchronization []*TerritoryOwners `json:"synchronization,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "General") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "General") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RightsOwnership) MarshalJSON() ([]byte, error) {
	type noMethod RightsOwnership
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type RightsOwnershipHistory struct {
	// Kind: The type of the API resource. For ownership history resources,
	// the value is youtubePartner#rightsOwnershipHistory.
	Kind string `json:"kind,omitempty"`

	// Origination: The origination object contains information that
	// describes the metadata source.
	Origination *Origination `json:"origination,omitempty"`

	// Ownership: The ownership object contains the ownership data provided
	// by the specified source (origination) at the specified time
	// (timeProvided).
	Ownership *RightsOwnership `json:"ownership,omitempty"`

	// TimeProvided: The time that the ownership data was provided.
	TimeProvided string `json:"timeProvided,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RightsOwnershipHistory) MarshalJSON() ([]byte, error) {
	type noMethod RightsOwnershipHistory
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Segment struct {
	// Duration: The duration of the segment in milliseconds.
	Duration uint64 `json:"duration,omitempty,string"`

	// Kind: The type of the API resource. For segment resources, the value
	// is youtubePartner#segment.
	Kind string `json:"kind,omitempty"`

	// Start: The start time of the segment, measured in milliseconds from
	// the beginning.
	Start uint64 `json:"start,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Duration") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Duration") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Segment) MarshalJSON() ([]byte, error) {
	type noMethod Segment
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ShowDetails struct {
	// EpisodeNumber: The episode number associated with the episode.
	EpisodeNumber string `json:"episodeNumber,omitempty"`

	// EpisodeTitle: The episode's title.
	EpisodeTitle string `json:"episodeTitle,omitempty"`

	// SeasonNumber: The season number associated with the episode.
	SeasonNumber string `json:"seasonNumber,omitempty"`

	// Title: The show's title
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EpisodeNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EpisodeNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ShowDetails) MarshalJSON() ([]byte, error) {
	type noMethod ShowDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type StateCompleted struct {
	// State: The state that the order entered.
	State string `json:"state,omitempty"`

	// TimeCompleted: The time that the state transition occurred.
	TimeCompleted int64 `json:"timeCompleted,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "State") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "State") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StateCompleted) MarshalJSON() ([]byte, error) {
	type noMethod StateCompleted
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TerritoryCondition struct {
	// Territories: A list of territories. Each territory is an ISO 3166
	// two-letter country code..
	Territories []string `json:"territories,omitempty"`

	// Type: This field indicates whether the associated policy rule is or
	// is not valid in the specified territories.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Territories") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Territories") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TerritoryCondition) MarshalJSON() ([]byte, error) {
	type noMethod TerritoryCondition
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TerritoryConflicts struct {
	// ConflictingOwnership: A list of conflicting ownerships.
	ConflictingOwnership []*ConflictingOwnership `json:"conflictingOwnership,omitempty"`

	// Territory: A territories where the ownership conflict is present.
	// Territory is an ISO 3166 two-letter country code..
	Territory string `json:"territory,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "ConflictingOwnership") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConflictingOwnership") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TerritoryConflicts) MarshalJSON() ([]byte, error) {
	type noMethod TerritoryConflicts
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TerritoryOwners struct {
	// Owner: The name of the asset's owner or rights administrator.
	Owner string `json:"owner,omitempty"`

	// Publisher: The name of the asset's publisher. This field is only used
	// for composition assets, and it is used when the asset owner is not
	// known to have a formal relationship established with YouTube.
	Publisher string `json:"publisher,omitempty"`

	// Ratio: The percentage of the asset that the owner controls or
	// administers. For composition assets, the value can be any value
	// between 0 and 100 inclusive. For all other assets, the only valid
	// values are 100, which indicates that the owner completely owns the
	// asset in the specified territories, and 0, which indicates that you
	// are removing ownership of the asset in the specified territories.
	Ratio float64 `json:"ratio,omitempty"`

	// Territories: A list of territories where the owner owns (or does not
	// own) the asset. Each territory is an ISO 3166 two-letter country
	// code..
	Territories []string `json:"territories,omitempty"`

	// Type: This field indicates whether the ownership data applies or does
	// not apply in the specified territories.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Owner") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Owner") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TerritoryOwners) MarshalJSON() ([]byte, error) {
	type noMethod TerritoryOwners
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *TerritoryOwners) UnmarshalJSON(data []byte) error {
	type noMethod TerritoryOwners
	var s1 struct {
		Ratio gensupport.JSONFloat64 `json:"ratio"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Ratio = float64(s1.Ratio)
	return nil
}

type ValidateError struct {
	// ColumnName: The column name where the error occurred.
	ColumnName string `json:"columnName,omitempty"`

	// ColumnNumber: The column number where the error occurred (1-based).
	ColumnNumber int64 `json:"columnNumber,omitempty"`

	// LineNumber: The line number where the error occurred (1-based).
	LineNumber int64 `json:"lineNumber,omitempty"`

	// Message: The error message.
	Message string `json:"message,omitempty"`

	// MessageCode: The code for the error message (if one exists).
	MessageCode int64 `json:"messageCode,omitempty"`

	// Severity: The error severity.
	Severity string `json:"severity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ColumnName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ColumnName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ValidateError) MarshalJSON() ([]byte, error) {
	type noMethod ValidateError
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ValidateRequest struct {
	// Content: The metadata file contents.
	Content string `json:"content,omitempty"`

	// Kind: The type of the API resource. For this operation, the value is
	// youtubePartner#validateRequest.
	Kind string `json:"kind,omitempty"`

	// Locale: The desired locale of the error messages as defined in BCP 47
	// (http://tools.ietf.org/html/bcp47). For example, "en-US" or "de". If
	// not specified we will return the error messages in English ("en").
	Locale string `json:"locale,omitempty"`

	// UploaderName: The uploader name.
	UploaderName string `json:"uploaderName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ValidateRequest) MarshalJSON() ([]byte, error) {
	type noMethod ValidateRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ValidateResponse struct {
	// Errors: The list of errors and/or warnings.
	Errors []*ValidateError `json:"errors,omitempty"`

	// Kind: The type of the API resource. For this operation, the value is
	// youtubePartner#validateResponse.
	Kind string `json:"kind,omitempty"`

	// Status: The validation status.
	Status string `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ValidateResponse) MarshalJSON() ([]byte, error) {
	type noMethod ValidateResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type VideoAdvertisingOption struct {
	// AdBreaks: A list of times when YouTube can show an in-stream
	// advertisement during playback of the video.
	AdBreaks []*AdBreak `json:"adBreaks,omitempty"`

	// AdFormats: A list of ad formats that the video is allowed to show.
	AdFormats []string `json:"adFormats,omitempty"`

	// AutoGeneratedBreaks: Enables this video for automatically generated
	// midroll breaks.
	AutoGeneratedBreaks bool `json:"autoGeneratedBreaks,omitempty"`

	// BreakPosition: The point at which the break occurs during the video
	// playback.
	BreakPosition []string `json:"breakPosition,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the video
	// associated with the advertising settings.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For this resource, the value is
	// youtubePartner#videoAdvertisingOption.
	Kind string `json:"kind,omitempty"`

	// TpAdServerVideoId: A value that uniquely identifies the video to the
	// third-party ad server.
	TpAdServerVideoId string `json:"tpAdServerVideoId,omitempty"`

	// TpTargetingUrl: The base URL for a third-party ad server from which
	// YouTube can retrieve in-stream ads for the video.
	TpTargetingUrl string `json:"tpTargetingUrl,omitempty"`

	// TpUrlParameters: A parameter string to append to the end of the
	// request to the third-party ad server.
	TpUrlParameters string `json:"tpUrlParameters,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AdBreaks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdBreaks") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VideoAdvertisingOption) MarshalJSON() ([]byte, error) {
	type noMethod VideoAdvertisingOption
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type VideoAdvertisingOptionGetEnabledAdsResponse struct {
	// AdBreaks: A list of ad breaks that occur in a claimed YouTube video.
	AdBreaks []*AdBreak `json:"adBreaks,omitempty"`

	// AdsOnEmbeds: This field indicates whether YouTube can show ads when
	// the video is played in an embedded player.
	AdsOnEmbeds bool `json:"adsOnEmbeds,omitempty"`

	// CountriesRestriction: A list that identifies the countries where ads
	// can run and the types of ads allowed in those countries.
	CountriesRestriction []*CountriesRestriction `json:"countriesRestriction,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the claimed video.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For this resource, the value is
	// youtubePartner#videoAdvertisingOptionGetEnabledAds.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AdBreaks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdBreaks") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VideoAdvertisingOptionGetEnabledAdsResponse) MarshalJSON() ([]byte, error) {
	type noMethod VideoAdvertisingOptionGetEnabledAdsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Whitelist struct {
	// Id: The YouTube channel ID that uniquely identifies the whitelisted
	// channel.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For whitelist resources, this
	// value is youtubePartner#whitelist.
	Kind string `json:"kind,omitempty"`

	// Title: Title of the whitelisted YouTube channel.
	Title string `json:"title,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Whitelist) MarshalJSON() ([]byte, error) {
	type noMethod Whitelist
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type WhitelistListResponse struct {
	// Items: A list of whitelist resources that match the request criteria.
	Items []*Whitelist `json:"items,omitempty"`

	// Kind: The type of the API response. For this operation, the value is
	// youtubePartner#whitelistList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WhitelistListResponse) MarshalJSON() ([]byte, error) {
	type noMethod WhitelistListResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "youtubePartner.assetLabels.insert":

type AssetLabelsInsertCall struct {
	s          *Service
	assetlabel *AssetLabel
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Insert an asset label for an owner.
func (r *AssetLabelsService) Insert(assetlabel *AssetLabel) *AssetLabelsInsertCall {
	c := &AssetLabelsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetlabel = assetlabel
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetLabelsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetLabelsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetLabelsInsertCall) Fields(s ...googleapi.Field) *AssetLabelsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetLabelsInsertCall) Context(ctx context.Context) *AssetLabelsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetLabelsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetLabelsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.assetlabel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetLabels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetLabels.insert" call.
// Exactly one of *AssetLabel or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AssetLabel.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AssetLabelsInsertCall) Do(opts ...googleapi.CallOption) (*AssetLabel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetLabel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert an asset label for an owner.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.assetLabels.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetLabels",
	//   "request": {
	//     "$ref": "AssetLabel"
	//   },
	//   "response": {
	//     "$ref": "AssetLabel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetLabels.list":

type AssetLabelsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of all asset labels for an owner.
func (r *AssetLabelsService) List() *AssetLabelsListCall {
	c := &AssetLabelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// LabelPrefix sets the optional parameter "labelPrefix": The
// labelPrefix parameter identifies the prefix of asset labels to
// retrieve.
func (c *AssetLabelsListCall) LabelPrefix(labelPrefix string) *AssetLabelsListCall {
	c.urlParams_.Set("labelPrefix", labelPrefix)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetLabelsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetLabelsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Q sets the optional parameter "q": The q parameter specifies the
// query string to use to filter search results. YouTube searches for
// the query string in the labelName field of asset labels.
func (c *AssetLabelsListCall) Q(q string) *AssetLabelsListCall {
	c.urlParams_.Set("q", q)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetLabelsListCall) Fields(s ...googleapi.Field) *AssetLabelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetLabelsListCall) IfNoneMatch(entityTag string) *AssetLabelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetLabelsListCall) Context(ctx context.Context) *AssetLabelsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetLabelsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetLabelsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetLabels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetLabels.list" call.
// Exactly one of *AssetLabelListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetLabelListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetLabelsListCall) Do(opts ...googleapi.CallOption) (*AssetLabelListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetLabelListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of all asset labels for an owner.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assetLabels.list",
	//   "parameters": {
	//     "labelPrefix": {
	//       "description": "The labelPrefix parameter identifies the prefix of asset labels to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "The q parameter specifies the query string to use to filter search results. YouTube searches for the query string in the labelName field of asset labels.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetLabels",
	//   "response": {
	//     "$ref": "AssetLabelListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetMatchPolicy.get":

type AssetMatchPolicyGetCall struct {
	s            *Service
	assetId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the match policy assigned to the specified asset by
// the content owner associated with the authenticated user. This
// information is only accessible to an owner of the asset.
func (r *AssetMatchPolicyService) Get(assetId string) *AssetMatchPolicyGetCall {
	c := &AssetMatchPolicyGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetMatchPolicyGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetMatchPolicyGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetMatchPolicyGetCall) Fields(s ...googleapi.Field) *AssetMatchPolicyGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetMatchPolicyGetCall) IfNoneMatch(entityTag string) *AssetMatchPolicyGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetMatchPolicyGetCall) Context(ctx context.Context) *AssetMatchPolicyGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetMatchPolicyGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetMatchPolicyGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}/matchPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetMatchPolicy.get" call.
// Exactly one of *AssetMatchPolicy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetMatchPolicy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetMatchPolicyGetCall) Do(opts ...googleapi.CallOption) (*AssetMatchPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetMatchPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the match policy assigned to the specified asset by the content owner associated with the authenticated user. This information is only accessible to an owner of the asset.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assetMatchPolicy.get",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving the match policy.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}/matchPolicy",
	//   "response": {
	//     "$ref": "AssetMatchPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetMatchPolicy.patch":

type AssetMatchPolicyPatchCall struct {
	s                *Service
	assetId          string
	assetmatchpolicy *AssetMatchPolicy
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Patch: Updates the asset's match policy. If an asset has multiple
// owners, each owner may set its own match policy for the asset.
// YouTube then computes the match policy that is actually applied for
// the asset based on the territories where each owner owns the asset.
// This method supports patch semantics.
func (r *AssetMatchPolicyService) Patch(assetId string, assetmatchpolicy *AssetMatchPolicy) *AssetMatchPolicyPatchCall {
	c := &AssetMatchPolicyPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	c.assetmatchpolicy = assetmatchpolicy
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetMatchPolicyPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetMatchPolicyPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetMatchPolicyPatchCall) Fields(s ...googleapi.Field) *AssetMatchPolicyPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetMatchPolicyPatchCall) Context(ctx context.Context) *AssetMatchPolicyPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetMatchPolicyPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetMatchPolicyPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.assetmatchpolicy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}/matchPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetMatchPolicy.patch" call.
// Exactly one of *AssetMatchPolicy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetMatchPolicy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetMatchPolicyPatchCall) Do(opts ...googleapi.CallOption) (*AssetMatchPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetMatchPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the asset's match policy. If an asset has multiple owners, each owner may set its own match policy for the asset. YouTube then computes the match policy that is actually applied for the asset based on the territories where each owner owns the asset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.assetMatchPolicy.patch",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving the match policy.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}/matchPolicy",
	//   "request": {
	//     "$ref": "AssetMatchPolicy"
	//   },
	//   "response": {
	//     "$ref": "AssetMatchPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetMatchPolicy.update":

type AssetMatchPolicyUpdateCall struct {
	s                *Service
	assetId          string
	assetmatchpolicy *AssetMatchPolicy
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Update: Updates the asset's match policy. If an asset has multiple
// owners, each owner may set its own match policy for the asset.
// YouTube then computes the match policy that is actually applied for
// the asset based on the territories where each owner owns the asset.
func (r *AssetMatchPolicyService) Update(assetId string, assetmatchpolicy *AssetMatchPolicy) *AssetMatchPolicyUpdateCall {
	c := &AssetMatchPolicyUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	c.assetmatchpolicy = assetmatchpolicy
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetMatchPolicyUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetMatchPolicyUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetMatchPolicyUpdateCall) Fields(s ...googleapi.Field) *AssetMatchPolicyUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetMatchPolicyUpdateCall) Context(ctx context.Context) *AssetMatchPolicyUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetMatchPolicyUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetMatchPolicyUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.assetmatchpolicy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}/matchPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetMatchPolicy.update" call.
// Exactly one of *AssetMatchPolicy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetMatchPolicy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetMatchPolicyUpdateCall) Do(opts ...googleapi.CallOption) (*AssetMatchPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetMatchPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the asset's match policy. If an asset has multiple owners, each owner may set its own match policy for the asset. YouTube then computes the match policy that is actually applied for the asset based on the territories where each owner owns the asset.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.assetMatchPolicy.update",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving the match policy.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}/matchPolicy",
	//   "request": {
	//     "$ref": "AssetMatchPolicy"
	//   },
	//   "response": {
	//     "$ref": "AssetMatchPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetRelationships.delete":

type AssetRelationshipsDeleteCall struct {
	s                   *Service
	assetRelationshipId string
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Delete: Deletes a relationship between two assets.
func (r *AssetRelationshipsService) Delete(assetRelationshipId string) *AssetRelationshipsDeleteCall {
	c := &AssetRelationshipsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetRelationshipId = assetRelationshipId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetRelationshipsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetRelationshipsDeleteCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetRelationshipsDeleteCall) Fields(s ...googleapi.Field) *AssetRelationshipsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetRelationshipsDeleteCall) Context(ctx context.Context) *AssetRelationshipsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetRelationshipsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetRelationshipsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetRelationships/{assetRelationshipId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetRelationshipId": c.assetRelationshipId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetRelationships.delete" call.
func (c *AssetRelationshipsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a relationship between two assets.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubePartner.assetRelationships.delete",
	//   "parameterOrder": [
	//     "assetRelationshipId"
	//   ],
	//   "parameters": {
	//     "assetRelationshipId": {
	//       "description": "The assetRelationshipId parameter specifies a value that uniquely identifies the relationship you are deleting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetRelationships/{assetRelationshipId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetRelationships.insert":

type AssetRelationshipsInsertCall struct {
	s                 *Service
	assetrelationship *AssetRelationship
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Insert: Creates a relationship that links two assets.
func (r *AssetRelationshipsService) Insert(assetrelationship *AssetRelationship) *AssetRelationshipsInsertCall {
	c := &AssetRelationshipsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetrelationship = assetrelationship
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetRelationshipsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetRelationshipsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetRelationshipsInsertCall) Fields(s ...googleapi.Field) *AssetRelationshipsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetRelationshipsInsertCall) Context(ctx context.Context) *AssetRelationshipsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetRelationshipsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetRelationshipsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.assetrelationship)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetRelationships")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetRelationships.insert" call.
// Exactly one of *AssetRelationship or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetRelationship.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetRelationshipsInsertCall) Do(opts ...googleapi.CallOption) (*AssetRelationship, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetRelationship{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a relationship that links two assets.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.assetRelationships.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetRelationships",
	//   "request": {
	//     "$ref": "AssetRelationship"
	//   },
	//   "response": {
	//     "$ref": "AssetRelationship"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assetRelationships.list":

type AssetRelationshipsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of relationships for a given asset. The list
// contains relationships where the specified asset is either the parent
// (embedding) or child (embedded) asset in the relationship.
func (r *AssetRelationshipsService) List(assetId string) *AssetRelationshipsListCall {
	c := &AssetRelationshipsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("assetId", assetId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetRelationshipsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetRelationshipsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of search results.
func (c *AssetRelationshipsListCall) PageToken(pageToken string) *AssetRelationshipsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetRelationshipsListCall) Fields(s ...googleapi.Field) *AssetRelationshipsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetRelationshipsListCall) IfNoneMatch(entityTag string) *AssetRelationshipsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetRelationshipsListCall) Context(ctx context.Context) *AssetRelationshipsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetRelationshipsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetRelationshipsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetRelationships")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetRelationships.list" call.
// Exactly one of *AssetRelationshipListResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *AssetRelationshipListResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetRelationshipsListCall) Do(opts ...googleapi.CallOption) (*AssetRelationshipListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetRelationshipListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of relationships for a given asset. The list contains relationships where the specified asset is either the parent (embedding) or child (embedded) asset in the relationship.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assetRelationships.list",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the asset ID of the asset for which you are retrieving relationships.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetRelationships",
	//   "response": {
	//     "$ref": "AssetRelationshipListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AssetRelationshipsListCall) Pages(ctx context.Context, f func(*AssetRelationshipListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.assetSearch.list":

type AssetSearchListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Searches for assets based on asset metadata. The method can
// retrieve all assets or only assets owned by the content owner. This
// method mimics the functionality of the advanced search feature on the
// Assets page in CMS.
func (r *AssetSearchService) List() *AssetSearchListCall {
	c := &AssetSearchListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": The
// createdAfter parameter restricts the set of returned assets to ones
// originally created on or after the specified datetime. For example:
// 2015-01-29T23:00:00Z
func (c *AssetSearchListCall) CreatedAfter(createdAfter string) *AssetSearchListCall {
	c.urlParams_.Set("createdAfter", createdAfter)
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": The
// createdBefore parameter restricts the set of returned assets to ones
// originally created on or before the specified datetime. For example:
// 2015-01-29T23:00:00Z
func (c *AssetSearchListCall) CreatedBefore(createdBefore string) *AssetSearchListCall {
	c.urlParams_.Set("createdBefore", createdBefore)
	return c
}

// HasConflicts sets the optional parameter "hasConflicts": The
// hasConflicts parameter enables you to only retrieve assets that have
// ownership conflicts. The only valid value is true. Setting the
// parameter value to false does not affect the results.
func (c *AssetSearchListCall) HasConflicts(hasConflicts bool) *AssetSearchListCall {
	c.urlParams_.Set("hasConflicts", fmt.Sprint(hasConflicts))
	return c
}

// IncludeAnyProvidedlabel sets the optional parameter
// "includeAnyProvidedlabel": If includeAnyProvidedlabel parameter is
// set to true, will search for assets that contain any of the provided
// labels; else will search for assets that contain all the provided
// labels.
func (c *AssetSearchListCall) IncludeAnyProvidedlabel(includeAnyProvidedlabel bool) *AssetSearchListCall {
	c.urlParams_.Set("includeAnyProvidedlabel", fmt.Sprint(includeAnyProvidedlabel))
	return c
}

// Isrcs sets the optional parameter "isrcs": A comma-separated list of
// up to 50 ISRCs. If you specify a value for this parameter, the API
// server ignores any values set for the following parameters: q,
// includeAnyProvidedLabel, hasConflicts, labels, metadataSearchFields,
// sort, and type.
func (c *AssetSearchListCall) Isrcs(isrcs string) *AssetSearchListCall {
	c.urlParams_.Set("isrcs", isrcs)
	return c
}

// Labels sets the optional parameter "labels": The labels parameter
// specifies the assets with certain asset labels that you want to
// retrieve. The parameter value is a comma-separated list of asset
// labels.
func (c *AssetSearchListCall) Labels(labels string) *AssetSearchListCall {
	c.urlParams_.Set("labels", labels)
	return c
}

// MetadataSearchFields sets the optional parameter
// "metadataSearchFields": The metadataSearchField parameter specifies
// which metadata fields to search by. It is a comma-separated list of
// metadata field and value pairs connected by colon(:). For example:
// customId:my_custom_id,artist:Dandexx
func (c *AssetSearchListCall) MetadataSearchFields(metadataSearchFields string) *AssetSearchListCall {
	c.urlParams_.Set("metadataSearchFields", metadataSearchFields)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetSearchListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetSearchListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// OwnershipRestriction sets the optional parameter
// "ownershipRestriction": The ownershipRestriction parameter specifies
// the ownership filtering option for the search. By default the search
// is performed in the assets owned by currently authenticated user
// only.
//
// Possible values:
//   "mine" - Find assets owned by the current user that match the
// search query. This is the default behavior.
//   "none" - Find all assets that match the search query, regardless of
// owner.
func (c *AssetSearchListCall) OwnershipRestriction(ownershipRestriction string) *AssetSearchListCall {
	c.urlParams_.Set("ownershipRestriction", ownershipRestriction)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of search results.
func (c *AssetSearchListCall) PageToken(pageToken string) *AssetSearchListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Q sets the optional parameter "q": YouTube searches within the id,
// type, and customId fields for all assets as well as in numerous other
// metadata fields – such as actor, album, director, isrc, and tmsId
// – that vary for different types of assets (movies, music videos,
// etc.).
func (c *AssetSearchListCall) Q(q string) *AssetSearchListCall {
	c.urlParams_.Set("q", q)
	return c
}

// Sort sets the optional parameter "sort": The sort parameter specifies
// how the search results should be sorted. Note that results are always
// sorted in descending order.
//
// Possible values:
//   "claims" - Sort by the number of claims for each asset.
//   "time" - Sort by the modification time for each asset. This is the
// default value.
//   "views" - Sort by the approximate daily views for each asset.
func (c *AssetSearchListCall) Sort(sort string) *AssetSearchListCall {
	c.urlParams_.Set("sort", sort)
	return c
}

// Type sets the optional parameter "type": The type parameter specifies
// the types of assets that you want to retrieve. The parameter value is
// a comma-separated list of asset types.
func (c *AssetSearchListCall) Type(type_ string) *AssetSearchListCall {
	c.urlParams_.Set("type", type_)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetSearchListCall) Fields(s ...googleapi.Field) *AssetSearchListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetSearchListCall) IfNoneMatch(entityTag string) *AssetSearchListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetSearchListCall) Context(ctx context.Context) *AssetSearchListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetSearchListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetSearchListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetSearch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetSearch.list" call.
// Exactly one of *AssetSearchResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetSearchResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetSearchListCall) Do(opts ...googleapi.CallOption) (*AssetSearchResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetSearchResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for assets based on asset metadata. The method can retrieve all assets or only assets owned by the content owner. This method mimics the functionality of the advanced search feature on the Assets page in CMS.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assetSearch.list",
	//   "parameters": {
	//     "createdAfter": {
	//       "description": "The createdAfter parameter restricts the set of returned assets to ones originally created on or after the specified datetime. For example: 2015-01-29T23:00:00Z",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "The createdBefore parameter restricts the set of returned assets to ones originally created on or before the specified datetime. For example: 2015-01-29T23:00:00Z",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hasConflicts": {
	//       "description": "The hasConflicts parameter enables you to only retrieve assets that have ownership conflicts. The only valid value is true. Setting the parameter value to false does not affect the results.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "includeAnyProvidedlabel": {
	//       "description": "If includeAnyProvidedlabel parameter is set to true, will search for assets that contain any of the provided labels; else will search for assets that contain all the provided labels.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "isrcs": {
	//       "description": "A comma-separated list of up to 50 ISRCs. If you specify a value for this parameter, the API server ignores any values set for the following parameters: q, includeAnyProvidedLabel, hasConflicts, labels, metadataSearchFields, sort, and type.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "labels": {
	//       "description": "The labels parameter specifies the assets with certain asset labels that you want to retrieve. The parameter value is a comma-separated list of asset labels.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "metadataSearchFields": {
	//       "description": "The metadataSearchField parameter specifies which metadata fields to search by. It is a comma-separated list of metadata field and value pairs connected by colon(:). For example: customId:my_custom_id,artist:Dandexx",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ownershipRestriction": {
	//       "default": "RESTRICT_MINE",
	//       "description": "The ownershipRestriction parameter specifies the ownership filtering option for the search. By default the search is performed in the assets owned by currently authenticated user only.",
	//       "enum": [
	//         "mine",
	//         "none"
	//       ],
	//       "enumDescriptions": [
	//         "Find assets owned by the current user that match the search query. This is the default behavior.",
	//         "Find all assets that match the search query, regardless of owner."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "YouTube searches within the id, type, and customId fields for all assets as well as in numerous other metadata fields – such as actor, album, director, isrc, and tmsId – that vary for different types of assets (movies, music videos, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The sort parameter specifies how the search results should be sorted. Note that results are always sorted in descending order.",
	//       "enum": [
	//         "claims",
	//         "time",
	//         "views"
	//       ],
	//       "enumDescriptions": [
	//         "Sort by the number of claims for each asset.",
	//         "Sort by the modification time for each asset. This is the default value.",
	//         "Sort by the approximate daily views for each asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "The type parameter specifies the types of assets that you want to retrieve. The parameter value is a comma-separated list of asset types.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetSearch",
	//   "response": {
	//     "$ref": "AssetSearchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AssetSearchListCall) Pages(ctx context.Context, f func(*AssetSearchResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.assetShares.list":

type AssetSharesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: This method either retrieves a list of asset shares the partner
// owns and that map to a specified asset view ID or it retrieves a list
// of asset views associated with a specified asset share ID owned by
// the partner.
func (r *AssetSharesService) List(assetId string) *AssetSharesListCall {
	c := &AssetSharesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("assetId", assetId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetSharesListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetSharesListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of search results.
func (c *AssetSharesListCall) PageToken(pageToken string) *AssetSharesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetSharesListCall) Fields(s ...googleapi.Field) *AssetSharesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetSharesListCall) IfNoneMatch(entityTag string) *AssetSharesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetSharesListCall) Context(ctx context.Context) *AssetSharesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetSharesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetSharesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assetShares")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assetShares.list" call.
// Exactly one of *AssetShareListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetShareListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetSharesListCall) Do(opts ...googleapi.CallOption) (*AssetShareListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetShareListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "This method either retrieves a list of asset shares the partner owns and that map to a specified asset view ID or it retrieves a list of asset views associated with a specified asset share ID owned by the partner.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assetShares.list",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the asset ID for which you are retrieving data. The parameter can be an asset view ID or an asset share ID. \n- If the value is an asset view ID, the API response identifies any asset share ids mapped to the asset view.\n- If the value is an asset share ID, the API response identifies any asset view ids that maps to that asset share.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assetShares",
	//   "response": {
	//     "$ref": "AssetShareListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AssetSharesListCall) Pages(ctx context.Context, f func(*AssetShareListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.assets.get":

type AssetsGetCall struct {
	s            *Service
	assetId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the metadata for the specified asset. Note that if the
// request identifies an asset that has been merged with another asset,
// meaning that YouTube identified the requested asset as a duplicate,
// then the request retrieves the merged, or synthesized, asset.
func (r *AssetsService) Get(assetId string) *AssetsGetCall {
	c := &AssetsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	return c
}

// FetchMatchPolicy sets the optional parameter "fetchMatchPolicy": The
// fetchMatchPolicy parameter specifies the version of the asset's match
// policy that should be returned in the API response.
func (c *AssetsGetCall) FetchMatchPolicy(fetchMatchPolicy string) *AssetsGetCall {
	c.urlParams_.Set("fetchMatchPolicy", fetchMatchPolicy)
	return c
}

// FetchMetadata sets the optional parameter "fetchMetadata": The
// fetchMetadata parameter specifies the version of the asset's metadata
// that should be returned in the API response. In some cases, YouTube
// receives metadata for an asset from multiple sources, such as when
// different partners own the asset in different territories.
func (c *AssetsGetCall) FetchMetadata(fetchMetadata string) *AssetsGetCall {
	c.urlParams_.Set("fetchMetadata", fetchMetadata)
	return c
}

// FetchOwnership sets the optional parameter "fetchOwnership": The
// fetchOwnership parameter specifies the version of the asset's
// ownership data that should be returned in the API response. As with
// asset metadata, YouTube can receive asset ownership data from
// multiple sources.
func (c *AssetsGetCall) FetchOwnership(fetchOwnership string) *AssetsGetCall {
	c.urlParams_.Set("fetchOwnership", fetchOwnership)
	return c
}

// FetchOwnershipConflicts sets the optional parameter
// "fetchOwnershipConflicts": The fetchOwnershipConflicts parameter
// allows you to retrieve information about ownership conflicts.
func (c *AssetsGetCall) FetchOwnershipConflicts(fetchOwnershipConflicts bool) *AssetsGetCall {
	c.urlParams_.Set("fetchOwnershipConflicts", fmt.Sprint(fetchOwnershipConflicts))
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsGetCall) Fields(s ...googleapi.Field) *AssetsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetsGetCall) IfNoneMatch(entityTag string) *AssetsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsGetCall) Context(ctx context.Context) *AssetsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assets.get" call.
// Exactly one of *Asset or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Asset.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AssetsGetCall) Do(opts ...googleapi.CallOption) (*Asset, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Asset{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the metadata for the specified asset. Note that if the request identifies an asset that has been merged with another asset, meaning that YouTube identified the requested asset as a duplicate, then the request retrieves the merged, or synthesized, asset.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assets.get",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fetchMatchPolicy": {
	//       "description": "The fetchMatchPolicy parameter specifies the version of the asset's match policy that should be returned in the API response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchMetadata": {
	//       "description": "The fetchMetadata parameter specifies the version of the asset's metadata that should be returned in the API response. In some cases, YouTube receives metadata for an asset from multiple sources, such as when different partners own the asset in different territories.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchOwnership": {
	//       "description": "The fetchOwnership parameter specifies the version of the asset's ownership data that should be returned in the API response. As with asset metadata, YouTube can receive asset ownership data from multiple sources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchOwnershipConflicts": {
	//       "description": "The fetchOwnershipConflicts parameter allows you to retrieve information about ownership conflicts.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}",
	//   "response": {
	//     "$ref": "Asset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assets.insert":

type AssetsInsertCall struct {
	s          *Service
	asset      *Asset
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Inserts an asset with the specified metadata. After inserting
// an asset, you can set its ownership data and match policy.
func (r *AssetsService) Insert(asset *Asset) *AssetsInsertCall {
	c := &AssetsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.asset = asset
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsInsertCall) Fields(s ...googleapi.Field) *AssetsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsInsertCall) Context(ctx context.Context) *AssetsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.asset)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assets.insert" call.
// Exactly one of *Asset or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Asset.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AssetsInsertCall) Do(opts ...googleapi.CallOption) (*Asset, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Asset{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts an asset with the specified metadata. After inserting an asset, you can set its ownership data and match policy.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.assets.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets",
	//   "request": {
	//     "$ref": "Asset"
	//   },
	//   "response": {
	//     "$ref": "Asset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assets.list":

type AssetsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of assets based on asset metadata. The method
// can retrieve all assets or only assets owned by the content
// owner.
//
// Note that in cases where duplicate assets have been merged, the API
// response only contains the synthesized asset. (It does not contain
// the constituent assets that were merged into the synthesized asset.)
func (r *AssetsService) List(id string) *AssetsListCall {
	c := &AssetsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("id", id)
	return c
}

// FetchMatchPolicy sets the optional parameter "fetchMatchPolicy": The
// fetchMatchPolicy parameter specifies the version of the asset's match
// policy that should be returned in the API response.
func (c *AssetsListCall) FetchMatchPolicy(fetchMatchPolicy string) *AssetsListCall {
	c.urlParams_.Set("fetchMatchPolicy", fetchMatchPolicy)
	return c
}

// FetchMetadata sets the optional parameter "fetchMetadata": The
// fetchMetadata parameter specifies the version of the asset's metadata
// that should be returned in the API response. In some cases, YouTube
// receives metadata for an asset from multiple sources, such as when
// different partners own the asset in different territories.
func (c *AssetsListCall) FetchMetadata(fetchMetadata string) *AssetsListCall {
	c.urlParams_.Set("fetchMetadata", fetchMetadata)
	return c
}

// FetchOwnership sets the optional parameter "fetchOwnership": The
// fetchOwnership parameter specifies the version of the asset's
// ownership data that should be returned in the API response. As with
// asset metadata, YouTube can receive asset ownership data from
// multiple sources.
func (c *AssetsListCall) FetchOwnership(fetchOwnership string) *AssetsListCall {
	c.urlParams_.Set("fetchOwnership", fetchOwnership)
	return c
}

// FetchOwnershipConflicts sets the optional parameter
// "fetchOwnershipConflicts": The fetchOwnershipConflicts parameter
// allows you to retrieve information about ownership conflicts.
func (c *AssetsListCall) FetchOwnershipConflicts(fetchOwnershipConflicts bool) *AssetsListCall {
	c.urlParams_.Set("fetchOwnershipConflicts", fmt.Sprint(fetchOwnershipConflicts))
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsListCall) Fields(s ...googleapi.Field) *AssetsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetsListCall) IfNoneMatch(entityTag string) *AssetsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsListCall) Context(ctx context.Context) *AssetsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assets.list" call.
// Exactly one of *AssetListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AssetListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetsListCall) Do(opts ...googleapi.CallOption) (*AssetListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AssetListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of assets based on asset metadata. The method can retrieve all assets or only assets owned by the content owner.\n\nNote that in cases where duplicate assets have been merged, the API response only contains the synthesized asset. (It does not contain the constituent assets that were merged into the synthesized asset.)",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.assets.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "fetchMatchPolicy": {
	//       "description": "The fetchMatchPolicy parameter specifies the version of the asset's match policy that should be returned in the API response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchMetadata": {
	//       "description": "The fetchMetadata parameter specifies the version of the asset's metadata that should be returned in the API response. In some cases, YouTube receives metadata for an asset from multiple sources, such as when different partners own the asset in different territories.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchOwnership": {
	//       "description": "The fetchOwnership parameter specifies the version of the asset's ownership data that should be returned in the API response. As with asset metadata, YouTube can receive asset ownership data from multiple sources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fetchOwnershipConflicts": {
	//       "description": "The fetchOwnershipConflicts parameter allows you to retrieve information about ownership conflicts.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of YouTube Asset IDs that identify the assets you want to retrieve. As noted in the method description, if you try to retrieve an asset that YouTube identified as a duplicate and merged with another asset, the API response only returns the synthesized asset. In that case, the aliasId property in the asset resource specifies a list of other asset IDs that can be used to identify that asset.\n\nAlso note that the API response does not contain duplicates. As such, if your request identifies three asset IDs, and all of those have been merged into a single asset, then the API response identifies one matching asset.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets",
	//   "response": {
	//     "$ref": "AssetListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assets.patch":

type AssetsPatchCall struct {
	s          *Service
	assetId    string
	asset      *Asset
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the metadata for the specified asset. This method
// supports patch semantics.
func (r *AssetsService) Patch(assetId string, asset *Asset) *AssetsPatchCall {
	c := &AssetsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	c.asset = asset
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetsPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetsPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsPatchCall) Fields(s ...googleapi.Field) *AssetsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsPatchCall) Context(ctx context.Context) *AssetsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.asset)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assets.patch" call.
// Exactly one of *Asset or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Asset.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AssetsPatchCall) Do(opts ...googleapi.CallOption) (*Asset, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Asset{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the metadata for the specified asset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.assets.patch",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}",
	//   "request": {
	//     "$ref": "Asset"
	//   },
	//   "response": {
	//     "$ref": "Asset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.assets.update":

type AssetsUpdateCall struct {
	s          *Service
	assetId    string
	asset      *Asset
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates the metadata for the specified asset.
func (r *AssetsService) Update(assetId string, asset *Asset) *AssetsUpdateCall {
	c := &AssetsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	c.asset = asset
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *AssetsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *AssetsUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsUpdateCall) Fields(s ...googleapi.Field) *AssetsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsUpdateCall) Context(ctx context.Context) *AssetsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.asset)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.assets.update" call.
// Exactly one of *Asset or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Asset.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AssetsUpdateCall) Do(opts ...googleapi.CallOption) (*Asset, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Asset{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the metadata for the specified asset.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.assets.update",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}",
	//   "request": {
	//     "$ref": "Asset"
	//   },
	//   "response": {
	//     "$ref": "Asset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.campaigns.delete":

type CampaignsDeleteCall struct {
	s          *Service
	campaignId string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a specified campaign for an owner.
func (r *CampaignsService) Delete(campaignId string) *CampaignsDeleteCall {
	c := &CampaignsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.campaignId = campaignId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *CampaignsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *CampaignsDeleteCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsDeleteCall) Fields(s ...googleapi.Field) *CampaignsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CampaignsDeleteCall) Context(ctx context.Context) *CampaignsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CampaignsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CampaignsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "campaigns/{campaignId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"campaignId": c.campaignId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.campaigns.delete" call.
func (c *CampaignsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a specified campaign for an owner.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubePartner.campaigns.delete",
	//   "parameterOrder": [
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "The campaignId parameter specifies the YouTube campaign ID of the campaign being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "campaigns/{campaignId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.campaigns.get":

type CampaignsGetCall struct {
	s            *Service
	campaignId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves a particular campaign for an owner.
func (r *CampaignsService) Get(campaignId string) *CampaignsGetCall {
	c := &CampaignsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.campaignId = campaignId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *CampaignsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *CampaignsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsGetCall) Fields(s ...googleapi.Field) *CampaignsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CampaignsGetCall) IfNoneMatch(entityTag string) *CampaignsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CampaignsGetCall) Context(ctx context.Context) *CampaignsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CampaignsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CampaignsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "campaigns/{campaignId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"campaignId": c.campaignId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.campaigns.get" call.
// Exactly one of *Campaign or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Campaign.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CampaignsGetCall) Do(opts ...googleapi.CallOption) (*Campaign, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Campaign{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a particular campaign for an owner.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.campaigns.get",
	//   "parameterOrder": [
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "The campaignId parameter specifies the YouTube campaign ID of the campaign being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "campaigns/{campaignId}",
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.campaigns.insert":

type CampaignsInsertCall struct {
	s          *Service
	campaign   *Campaign
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Insert a new campaign for an owner using the specified
// campaign data.
func (r *CampaignsService) Insert(campaign *Campaign) *CampaignsInsertCall {
	c := &CampaignsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.campaign = campaign
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *CampaignsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *CampaignsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsInsertCall) Fields(s ...googleapi.Field) *CampaignsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CampaignsInsertCall) Context(ctx context.Context) *CampaignsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CampaignsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CampaignsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.campaign)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "campaigns")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.campaigns.insert" call.
// Exactly one of *Campaign or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Campaign.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CampaignsInsertCall) Do(opts ...googleapi.CallOption) (*Campaign, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Campaign{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a new campaign for an owner using the specified campaign data.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.campaigns.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "campaigns",
	//   "request": {
	//     "$ref": "Campaign"
	//   },
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.campaigns.list":

type CampaignsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of campaigns for an owner.
func (r *CampaignsService) List() *CampaignsListCall {
	c := &CampaignsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *CampaignsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *CampaignsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. For example, set this parameter to the value of
// the nextPageToken value from the previous API response to retrieve
// the next page of search results.
func (c *CampaignsListCall) PageToken(pageToken string) *CampaignsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsListCall) Fields(s ...googleapi.Field) *CampaignsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CampaignsListCall) IfNoneMatch(entityTag string) *CampaignsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CampaignsListCall) Context(ctx context.Context) *CampaignsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CampaignsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CampaignsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "campaigns")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.campaigns.list" call.
// Exactly one of *CampaignList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CampaignList.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CampaignsListCall) Do(opts ...googleapi.CallOption) (*CampaignList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CampaignList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of campaigns for an owner.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.campaigns.list",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. For example, set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "campaigns",
	//   "response": {
	//     "$ref": "CampaignList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.campaigns.patch":

type CampaignsPatchCall struct {
	s          *Service
	campaignId string
	campaign   *Campaign
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Update the data for a specific campaign. This method supports
// patch semantics.
func (r *CampaignsService) Patch(campaignId string, campaign *Campaign) *CampaignsPatchCall {
	c := &CampaignsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.campaignId = campaignId
	c.campaign = campaign
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *CampaignsPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *CampaignsPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsPatchCall) Fields(s ...googleapi.Field) *CampaignsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CampaignsPatchCall) Context(ctx context.Context) *CampaignsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CampaignsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CampaignsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.campaign)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "campaigns/{campaignId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"campaignId": c.campaignId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.campaigns.patch" call.
// Exactly one of *Campaign or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Campaign.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CampaignsPatchCall) Do(opts ...googleapi.CallOption) (*Campaign, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Campaign{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the data for a specific campaign. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.campaigns.patch",
	//   "parameterOrder": [
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "The campaignId parameter specifies the YouTube campaign ID of the campaign being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "campaigns/{campaignId}",
	//   "request": {
	//     "$ref": "Campaign"
	//   },
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.campaigns.update":

type CampaignsUpdateCall struct {
	s          *Service
	campaignId string
	campaign   *Campaign
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Update the data for a specific campaign.
func (r *CampaignsService) Update(campaignId string, campaign *Campaign) *CampaignsUpdateCall {
	c := &CampaignsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.campaignId = campaignId
	c.campaign = campaign
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *CampaignsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *CampaignsUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsUpdateCall) Fields(s ...googleapi.Field) *CampaignsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CampaignsUpdateCall) Context(ctx context.Context) *CampaignsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CampaignsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CampaignsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.campaign)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "campaigns/{campaignId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"campaignId": c.campaignId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.campaigns.update" call.
// Exactly one of *Campaign or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Campaign.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CampaignsUpdateCall) Do(opts ...googleapi.CallOption) (*Campaign, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Campaign{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the data for a specific campaign.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.campaigns.update",
	//   "parameterOrder": [
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "The campaignId parameter specifies the YouTube campaign ID of the campaign being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "campaigns/{campaignId}",
	//   "request": {
	//     "$ref": "Campaign"
	//   },
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.claimHistory.get":

type ClaimHistoryGetCall struct {
	s            *Service
	claimId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the claim history for a specified claim.
func (r *ClaimHistoryService) Get(claimId string) *ClaimHistoryGetCall {
	c := &ClaimHistoryGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.claimId = claimId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimHistoryGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimHistoryGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimHistoryGetCall) Fields(s ...googleapi.Field) *ClaimHistoryGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClaimHistoryGetCall) IfNoneMatch(entityTag string) *ClaimHistoryGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimHistoryGetCall) Context(ctx context.Context) *ClaimHistoryGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimHistoryGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimHistoryGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claimHistory/{claimId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"claimId": c.claimId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claimHistory.get" call.
// Exactly one of *ClaimHistory or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ClaimHistory.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ClaimHistoryGetCall) Do(opts ...googleapi.CallOption) (*ClaimHistory, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClaimHistory{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the claim history for a specified claim.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.claimHistory.get",
	//   "parameterOrder": [
	//     "claimId"
	//   ],
	//   "parameters": {
	//     "claimId": {
	//       "description": "The claimId parameter specifies the YouTube claim ID of the claim for which you are retrieving the claim history.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claimHistory/{claimId}",
	//   "response": {
	//     "$ref": "ClaimHistory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.claimSearch.list":

type ClaimSearchListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of claims that match the search criteria. You
// can search for claims that are associated with a specific asset or
// video or that match a specified query string.
func (r *ClaimSearchService) List() *ClaimSearchListCall {
	c := &ClaimSearchListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// AssetId sets the optional parameter "assetId": The assetId parameter
// specifies the YouTube asset ID of the asset for which you are
// retrieving claims.
func (c *ClaimSearchListCall) AssetId(assetId string) *ClaimSearchListCall {
	c.urlParams_.Set("assetId", assetId)
	return c
}

// ContentType sets the optional parameter "contentType": The
// contentType parameter specifies the content type of claims that you
// want to retrieve.
//
// Possible values:
//   "audio" - Restrict results to audio-only claims.
//   "audiovisual" - Restrict results to audiovisual claims.
//   "visual" - Restrict results to video-only claims.
func (c *ClaimSearchListCall) ContentType(contentType string) *ClaimSearchListCall {
	c.urlParams_.Set("contentType", contentType)
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": The
// createdAfter parameter allows you to restrict the set of returned
// claims to ones created on or after the specified date (inclusive).
func (c *ClaimSearchListCall) CreatedAfter(createdAfter string) *ClaimSearchListCall {
	c.urlParams_.Set("createdAfter", createdAfter)
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": The
// createdBefore parameter allows you to restrict the set of returned
// claims to ones created before the specified date (exclusive).
func (c *ClaimSearchListCall) CreatedBefore(createdBefore string) *ClaimSearchListCall {
	c.urlParams_.Set("createdBefore", createdBefore)
	return c
}

// InactiveReasons sets the optional parameter "inactiveReasons": The
// inactiveReasons parameter allows you to specify what kind of inactive
// claims you want to find based on the reasons why the claims became
// inactive.
func (c *ClaimSearchListCall) InactiveReasons(inactiveReasons string) *ClaimSearchListCall {
	c.urlParams_.Set("inactiveReasons", inactiveReasons)
	return c
}

// IncludeThirdPartyClaims sets the optional parameter
// "includeThirdPartyClaims": Used along with the videoId parameter this
// parameter determines whether or not to include third party claims in
// the search results.
func (c *ClaimSearchListCall) IncludeThirdPartyClaims(includeThirdPartyClaims bool) *ClaimSearchListCall {
	c.urlParams_.Set("includeThirdPartyClaims", fmt.Sprint(includeThirdPartyClaims))
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimSearchListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimSearchListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Origin sets the optional parameter "origin": The origin parameter
// specifies the origins you want to find claims for. It is a
// comma-separated list of origin values.
func (c *ClaimSearchListCall) Origin(origin string) *ClaimSearchListCall {
	c.urlParams_.Set("origin", origin)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. For example, set this parameter to the value of
// the nextPageToken value from the previous API response to retrieve
// the next page of search results.
func (c *ClaimSearchListCall) PageToken(pageToken string) *ClaimSearchListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// PartnerUploaded sets the optional parameter "partnerUploaded": The
// partnerUploaded parameter specifies whether you want to filter your
// search results to only partner uploaded or non partner uploaded
// claims.
func (c *ClaimSearchListCall) PartnerUploaded(partnerUploaded bool) *ClaimSearchListCall {
	c.urlParams_.Set("partnerUploaded", fmt.Sprint(partnerUploaded))
	return c
}

// Q sets the optional parameter "q": The q parameter specifies the
// query string to use to filter search results. YouTube searches for
// the query string in the following claim fields: video_title,
// video_keywords, user_name, isrc, iswc, grid, custom_id, and in the
// content owner's email address.
func (c *ClaimSearchListCall) Q(q string) *ClaimSearchListCall {
	c.urlParams_.Set("q", q)
	return c
}

// ReferenceId sets the optional parameter "referenceId": The
// referenceId parameter specifies the YouTube reference ID of the
// reference for which you are retrieving claims.
func (c *ClaimSearchListCall) ReferenceId(referenceId string) *ClaimSearchListCall {
	c.urlParams_.Set("referenceId", referenceId)
	return c
}

// Sort sets the optional parameter "sort": The sort parameter specifies
// the method that will be used to order resources in the API response.
// The default value is date. However, if the status parameter value is
// either appealed, disputed, pending, potential, or routedForReview,
// then results will be sorted by the time that the claim review period
// expires.
//
// Possible values:
//   "date" - Resources are sorted in reverse chronological order (from
// newest to oldest) based on the dates they were created. This value is
// not applicable if the status parameter is set to any of the following
// values: appealed, disputed, pending, potential, or routedForReview.
//   "viewCount" - Resources are sorted from highest to lowest number of
// views for the claimed content. This value is not applicable if the
// status parameter is set to any of the following values: appealed,
// disputed, pending, potential, or routedForReview.
func (c *ClaimSearchListCall) Sort(sort string) *ClaimSearchListCall {
	c.urlParams_.Set("sort", sort)
	return c
}

// Status sets the optional parameter "status": The status parameter
// restricts your results to only claims in the specified status.
//
// Possible values:
//   "active" - Restrict results to claims with active status.
//   "appealed" - Restrict results to claims with appealed status.
//   "disputed" - Restrict results to claims with disputed status.
//   "inactive" - Restrict results to claims with inactive status.
//   "pending" - Restrict results to claims with pending status.
//   "potential" - Restrict results to claims with potetial status.
//   "routedForReview" - Restrict results to claims that require review
// based on a match policy rule.
//   "takedown" - Restrict results to claims with takedown status.
func (c *ClaimSearchListCall) Status(status string) *ClaimSearchListCall {
	c.urlParams_.Set("status", status)
	return c
}

// StatusModifiedAfter sets the optional parameter
// "statusModifiedAfter": The statusModifiedAfter parameter allows you
// to restrict the result set to only include claims that have had their
// status modified on or after the specified date (inclusive). The date
// specified must be on or after June 30, 2016 (2016-06-30). The
// parameter value's format is YYYY-MM-DD.
func (c *ClaimSearchListCall) StatusModifiedAfter(statusModifiedAfter string) *ClaimSearchListCall {
	c.urlParams_.Set("statusModifiedAfter", statusModifiedAfter)
	return c
}

// StatusModifiedBefore sets the optional parameter
// "statusModifiedBefore": The statusModifiedBefore parameter allows you
// to restrict the result set to only include claims that have had their
// status modified before the specified date (exclusive). The date
// specified must be on or after July 1, 2016 (2016-07-01). The
// parameter value's format is YYYY-MM-DD.
func (c *ClaimSearchListCall) StatusModifiedBefore(statusModifiedBefore string) *ClaimSearchListCall {
	c.urlParams_.Set("statusModifiedBefore", statusModifiedBefore)
	return c
}

// VideoId sets the optional parameter "videoId": The videoId parameter
// specifies comma-separated list of YouTube video IDs for which you are
// retrieving claims.
func (c *ClaimSearchListCall) VideoId(videoId string) *ClaimSearchListCall {
	c.urlParams_.Set("videoId", videoId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimSearchListCall) Fields(s ...googleapi.Field) *ClaimSearchListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClaimSearchListCall) IfNoneMatch(entityTag string) *ClaimSearchListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimSearchListCall) Context(ctx context.Context) *ClaimSearchListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimSearchListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimSearchListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claimSearch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claimSearch.list" call.
// Exactly one of *ClaimSearchResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClaimSearchResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ClaimSearchListCall) Do(opts ...googleapi.CallOption) (*ClaimSearchResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClaimSearchResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of claims that match the search criteria. You can search for claims that are associated with a specific asset or video or that match a specified query string.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.claimSearch.list",
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving claims.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "contentType": {
	//       "description": "The contentType parameter specifies the content type of claims that you want to retrieve.",
	//       "enum": [
	//         "audio",
	//         "audiovisual",
	//         "visual"
	//       ],
	//       "enumDescriptions": [
	//         "Restrict results to audio-only claims.",
	//         "Restrict results to audiovisual claims.",
	//         "Restrict results to video-only claims."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "The createdAfter parameter allows you to restrict the set of returned claims to ones created on or after the specified date (inclusive).",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "The createdBefore parameter allows you to restrict the set of returned claims to ones created before the specified date (exclusive).",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "inactiveReasons": {
	//       "description": "The inactiveReasons parameter allows you to specify what kind of inactive claims you want to find based on the reasons why the claims became inactive.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "includeThirdPartyClaims": {
	//       "description": "Used along with the videoId parameter this parameter determines whether or not to include third party claims in the search results.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "origin": {
	//       "description": "The origin parameter specifies the origins you want to find claims for. It is a comma-separated list of origin values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. For example, set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "partnerUploaded": {
	//       "description": "The partnerUploaded parameter specifies whether you want to filter your search results to only partner uploaded or non partner uploaded claims.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "q": {
	//       "description": "The q parameter specifies the query string to use to filter search results. YouTube searches for the query string in the following claim fields: video_title, video_keywords, user_name, isrc, iswc, grid, custom_id, and in the content owner's email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceId": {
	//       "description": "The referenceId parameter specifies the YouTube reference ID of the reference for which you are retrieving claims.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The sort parameter specifies the method that will be used to order resources in the API response. The default value is date. However, if the status parameter value is either appealed, disputed, pending, potential, or routedForReview, then results will be sorted by the time that the claim review period expires.",
	//       "enum": [
	//         "date",
	//         "viewCount"
	//       ],
	//       "enumDescriptions": [
	//         "Resources are sorted in reverse chronological order (from newest to oldest) based on the dates they were created. This value is not applicable if the status parameter is set to any of the following values: appealed, disputed, pending, potential, or routedForReview.",
	//         "Resources are sorted from highest to lowest number of views for the claimed content. This value is not applicable if the status parameter is set to any of the following values: appealed, disputed, pending, potential, or routedForReview."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "The status parameter restricts your results to only claims in the specified status.",
	//       "enum": [
	//         "active",
	//         "appealed",
	//         "disputed",
	//         "inactive",
	//         "pending",
	//         "potential",
	//         "routedForReview",
	//         "takedown"
	//       ],
	//       "enumDescriptions": [
	//         "Restrict results to claims with active status.",
	//         "Restrict results to claims with appealed status.",
	//         "Restrict results to claims with disputed status.",
	//         "Restrict results to claims with inactive status.",
	//         "Restrict results to claims with pending status.",
	//         "Restrict results to claims with potetial status.",
	//         "Restrict results to claims that require review based on a match policy rule.",
	//         "Restrict results to claims with takedown status."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "statusModifiedAfter": {
	//       "description": "The statusModifiedAfter parameter allows you to restrict the result set to only include claims that have had their status modified on or after the specified date (inclusive). The date specified must be on or after June 30, 2016 (2016-06-30). The parameter value's format is YYYY-MM-DD.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "statusModifiedBefore": {
	//       "description": "The statusModifiedBefore parameter allows you to restrict the result set to only include claims that have had their status modified before the specified date (exclusive). The date specified must be on or after July 1, 2016 (2016-07-01). The parameter value's format is YYYY-MM-DD.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "The videoId parameter specifies comma-separated list of YouTube video IDs for which you are retrieving claims.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claimSearch",
	//   "response": {
	//     "$ref": "ClaimSearchResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ClaimSearchListCall) Pages(ctx context.Context, f func(*ClaimSearchResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.claims.get":

type ClaimsGetCall struct {
	s            *Service
	claimId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves a specific claim by ID.
func (r *ClaimsService) Get(claimId string) *ClaimsGetCall {
	c := &ClaimsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.claimId = claimId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimsGetCall) Fields(s ...googleapi.Field) *ClaimsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClaimsGetCall) IfNoneMatch(entityTag string) *ClaimsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimsGetCall) Context(ctx context.Context) *ClaimsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claims/{claimId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"claimId": c.claimId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claims.get" call.
// Exactly one of *Claim or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Claim.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ClaimsGetCall) Do(opts ...googleapi.CallOption) (*Claim, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Claim{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a specific claim by ID.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.claims.get",
	//   "parameterOrder": [
	//     "claimId"
	//   ],
	//   "parameters": {
	//     "claimId": {
	//       "description": "The claimId parameter specifies the claim ID of the claim being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claims/{claimId}",
	//   "response": {
	//     "$ref": "Claim"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.claims.insert":

type ClaimsInsertCall struct {
	s          *Service
	claim      *Claim
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a claim. The video being claimed must have been
// uploaded to a channel associated with the same content owner as the
// API user sending the request. You can set the claim's policy in any
// of the following ways:
// - Use the claim resource's policy property to identify a saved policy
// by its unique ID.
// - Use the claim resource's policy property to specify a custom set of
// rules.
func (r *ClaimsService) Insert(claim *Claim) *ClaimsInsertCall {
	c := &ClaimsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.claim = claim
	return c
}

// IsManualClaim sets the optional parameter "isManualClaim": restricted
func (c *ClaimsInsertCall) IsManualClaim(isManualClaim bool) *ClaimsInsertCall {
	c.urlParams_.Set("isManualClaim", fmt.Sprint(isManualClaim))
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimsInsertCall) Fields(s ...googleapi.Field) *ClaimsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimsInsertCall) Context(ctx context.Context) *ClaimsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.claim)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claims")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claims.insert" call.
// Exactly one of *Claim or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Claim.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ClaimsInsertCall) Do(opts ...googleapi.CallOption) (*Claim, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Claim{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a claim. The video being claimed must have been uploaded to a channel associated with the same content owner as the API user sending the request. You can set the claim's policy in any of the following ways:\n- Use the claim resource's policy property to identify a saved policy by its unique ID.\n- Use the claim resource's policy property to specify a custom set of rules.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.claims.insert",
	//   "parameters": {
	//     "isManualClaim": {
	//       "description": "restricted",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claims",
	//   "request": {
	//     "$ref": "Claim"
	//   },
	//   "response": {
	//     "$ref": "Claim"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.claims.list":

type ClaimsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of claims administered by the content owner
// associated with the currently authenticated user. Results are sorted
// in descending order of creation time.
func (r *ClaimsService) List() *ClaimsListCall {
	c := &ClaimsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// AssetId sets the optional parameter "assetId": Use the
// claimSearch.list method's assetId parameter to search for claim
// snippets by asset ID. You can then retrieve the claim resources for
// those claims by using this method's id parameter to specify a
// comma-separated list of claim IDs.
func (c *ClaimsListCall) AssetId(assetId string) *ClaimsListCall {
	c.urlParams_.Set("assetId", assetId)
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// list of comma-separated YouTube claim IDs to retrieve.
func (c *ClaimsListCall) Id(id string) *ClaimsListCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. For example, set this parameter to the value of
// the nextPageToken value from the previous API response to retrieve
// the next page of search results.
func (c *ClaimsListCall) PageToken(pageToken string) *ClaimsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Q sets the optional parameter "q": Use the claimSearch.list method's
// q parameter to search for claim snippets that match a particular
// query string. You can then retrieve the claim resources for those
// claims by using this method's id parameter to specify a
// comma-separated list of claim IDs.
func (c *ClaimsListCall) Q(q string) *ClaimsListCall {
	c.urlParams_.Set("q", q)
	return c
}

// VideoId sets the optional parameter "videoId": Use the
// claimSearch.list method's videoId parameter to search for claim
// snippets by video ID. You can then retrieve the claim resources for
// those claims by using this method's id parameter to specify a
// comma-separated list of claim IDs.
func (c *ClaimsListCall) VideoId(videoId string) *ClaimsListCall {
	c.urlParams_.Set("videoId", videoId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimsListCall) Fields(s ...googleapi.Field) *ClaimsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClaimsListCall) IfNoneMatch(entityTag string) *ClaimsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimsListCall) Context(ctx context.Context) *ClaimsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claims")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claims.list" call.
// Exactly one of *ClaimListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClaimListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ClaimsListCall) Do(opts ...googleapi.CallOption) (*ClaimListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClaimListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of claims administered by the content owner associated with the currently authenticated user. Results are sorted in descending order of creation time.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.claims.list",
	//   "parameters": {
	//     "assetId": {
	//       "description": "Use the claimSearch.list method's assetId parameter to search for claim snippets by asset ID. You can then retrieve the claim resources for those claims by using this method's id parameter to specify a comma-separated list of claim IDs.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a list of comma-separated YouTube claim IDs to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. For example, set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Use the claimSearch.list method's q parameter to search for claim snippets that match a particular query string. You can then retrieve the claim resources for those claims by using this method's id parameter to specify a comma-separated list of claim IDs.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "Use the claimSearch.list method's videoId parameter to search for claim snippets by video ID. You can then retrieve the claim resources for those claims by using this method's id parameter to specify a comma-separated list of claim IDs.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claims",
	//   "response": {
	//     "$ref": "ClaimListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ClaimsListCall) Pages(ctx context.Context, f func(*ClaimListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.claims.patch":

type ClaimsPatchCall struct {
	s          *Service
	claimId    string
	claim      *Claim
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates an existing claim by either changing its policy or its
// status. You can update a claim's status from active to inactive to
// effectively release the claim. This method supports patch semantics.
func (r *ClaimsService) Patch(claimId string, claim *Claim) *ClaimsPatchCall {
	c := &ClaimsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.claimId = claimId
	c.claim = claim
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimsPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimsPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimsPatchCall) Fields(s ...googleapi.Field) *ClaimsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimsPatchCall) Context(ctx context.Context) *ClaimsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.claim)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claims/{claimId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"claimId": c.claimId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claims.patch" call.
// Exactly one of *Claim or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Claim.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ClaimsPatchCall) Do(opts ...googleapi.CallOption) (*Claim, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Claim{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing claim by either changing its policy or its status. You can update a claim's status from active to inactive to effectively release the claim. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.claims.patch",
	//   "parameterOrder": [
	//     "claimId"
	//   ],
	//   "parameters": {
	//     "claimId": {
	//       "description": "The claimId parameter specifies the claim ID of the claim being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claims/{claimId}",
	//   "request": {
	//     "$ref": "Claim"
	//   },
	//   "response": {
	//     "$ref": "Claim"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.claims.update":

type ClaimsUpdateCall struct {
	s          *Service
	claimId    string
	claim      *Claim
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates an existing claim by either changing its policy or
// its status. You can update a claim's status from active to inactive
// to effectively release the claim.
func (r *ClaimsService) Update(claimId string, claim *Claim) *ClaimsUpdateCall {
	c := &ClaimsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.claimId = claimId
	c.claim = claim
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ClaimsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ClaimsUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClaimsUpdateCall) Fields(s ...googleapi.Field) *ClaimsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClaimsUpdateCall) Context(ctx context.Context) *ClaimsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClaimsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClaimsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.claim)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "claims/{claimId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"claimId": c.claimId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.claims.update" call.
// Exactly one of *Claim or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Claim.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ClaimsUpdateCall) Do(opts ...googleapi.CallOption) (*Claim, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Claim{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing claim by either changing its policy or its status. You can update a claim's status from active to inactive to effectively release the claim.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.claims.update",
	//   "parameterOrder": [
	//     "claimId"
	//   ],
	//   "parameters": {
	//     "claimId": {
	//       "description": "The claimId parameter specifies the claim ID of the claim being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "claims/{claimId}",
	//   "request": {
	//     "$ref": "Claim"
	//   },
	//   "response": {
	//     "$ref": "Claim"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.contentOwnerAdvertisingOptions.get":

type ContentOwnerAdvertisingOptionsGetCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves advertising options for the content owner associated
// with the authenticated user.
func (r *ContentOwnerAdvertisingOptionsService) Get() *ContentOwnerAdvertisingOptionsGetCall {
	c := &ContentOwnerAdvertisingOptionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ContentOwnerAdvertisingOptionsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ContentOwnerAdvertisingOptionsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentOwnerAdvertisingOptionsGetCall) Fields(s ...googleapi.Field) *ContentOwnerAdvertisingOptionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ContentOwnerAdvertisingOptionsGetCall) IfNoneMatch(entityTag string) *ContentOwnerAdvertisingOptionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ContentOwnerAdvertisingOptionsGetCall) Context(ctx context.Context) *ContentOwnerAdvertisingOptionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ContentOwnerAdvertisingOptionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ContentOwnerAdvertisingOptionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "contentOwnerAdvertisingOptions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.contentOwnerAdvertisingOptions.get" call.
// Exactly one of *ContentOwnerAdvertisingOption or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ContentOwnerAdvertisingOption.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ContentOwnerAdvertisingOptionsGetCall) Do(opts ...googleapi.CallOption) (*ContentOwnerAdvertisingOption, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ContentOwnerAdvertisingOption{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves advertising options for the content owner associated with the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.contentOwnerAdvertisingOptions.get",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "contentOwnerAdvertisingOptions",
	//   "response": {
	//     "$ref": "ContentOwnerAdvertisingOption"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.contentOwnerAdvertisingOptions.patch":

type ContentOwnerAdvertisingOptionsPatchCall struct {
	s                             *Service
	contentowneradvertisingoption *ContentOwnerAdvertisingOption
	urlParams_                    gensupport.URLParams
	ctx_                          context.Context
	header_                       http.Header
}

// Patch: Updates advertising options for the content owner associated
// with the authenticated API user. This method supports patch
// semantics.
func (r *ContentOwnerAdvertisingOptionsService) Patch(contentowneradvertisingoption *ContentOwnerAdvertisingOption) *ContentOwnerAdvertisingOptionsPatchCall {
	c := &ContentOwnerAdvertisingOptionsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.contentowneradvertisingoption = contentowneradvertisingoption
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ContentOwnerAdvertisingOptionsPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ContentOwnerAdvertisingOptionsPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentOwnerAdvertisingOptionsPatchCall) Fields(s ...googleapi.Field) *ContentOwnerAdvertisingOptionsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ContentOwnerAdvertisingOptionsPatchCall) Context(ctx context.Context) *ContentOwnerAdvertisingOptionsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ContentOwnerAdvertisingOptionsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ContentOwnerAdvertisingOptionsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.contentowneradvertisingoption)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "contentOwnerAdvertisingOptions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.contentOwnerAdvertisingOptions.patch" call.
// Exactly one of *ContentOwnerAdvertisingOption or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ContentOwnerAdvertisingOption.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ContentOwnerAdvertisingOptionsPatchCall) Do(opts ...googleapi.CallOption) (*ContentOwnerAdvertisingOption, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ContentOwnerAdvertisingOption{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates advertising options for the content owner associated with the authenticated API user. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.contentOwnerAdvertisingOptions.patch",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "contentOwnerAdvertisingOptions",
	//   "request": {
	//     "$ref": "ContentOwnerAdvertisingOption"
	//   },
	//   "response": {
	//     "$ref": "ContentOwnerAdvertisingOption"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.contentOwnerAdvertisingOptions.update":

type ContentOwnerAdvertisingOptionsUpdateCall struct {
	s                             *Service
	contentowneradvertisingoption *ContentOwnerAdvertisingOption
	urlParams_                    gensupport.URLParams
	ctx_                          context.Context
	header_                       http.Header
}

// Update: Updates advertising options for the content owner associated
// with the authenticated API user.
func (r *ContentOwnerAdvertisingOptionsService) Update(contentowneradvertisingoption *ContentOwnerAdvertisingOption) *ContentOwnerAdvertisingOptionsUpdateCall {
	c := &ContentOwnerAdvertisingOptionsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.contentowneradvertisingoption = contentowneradvertisingoption
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ContentOwnerAdvertisingOptionsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ContentOwnerAdvertisingOptionsUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentOwnerAdvertisingOptionsUpdateCall) Fields(s ...googleapi.Field) *ContentOwnerAdvertisingOptionsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ContentOwnerAdvertisingOptionsUpdateCall) Context(ctx context.Context) *ContentOwnerAdvertisingOptionsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ContentOwnerAdvertisingOptionsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ContentOwnerAdvertisingOptionsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.contentowneradvertisingoption)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "contentOwnerAdvertisingOptions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.contentOwnerAdvertisingOptions.update" call.
// Exactly one of *ContentOwnerAdvertisingOption or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ContentOwnerAdvertisingOption.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ContentOwnerAdvertisingOptionsUpdateCall) Do(opts ...googleapi.CallOption) (*ContentOwnerAdvertisingOption, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ContentOwnerAdvertisingOption{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates advertising options for the content owner associated with the authenticated API user.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.contentOwnerAdvertisingOptions.update",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "contentOwnerAdvertisingOptions",
	//   "request": {
	//     "$ref": "ContentOwnerAdvertisingOption"
	//   },
	//   "response": {
	//     "$ref": "ContentOwnerAdvertisingOption"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.contentOwners.get":

type ContentOwnersGetCall struct {
	s              *Service
	contentOwnerId string
	urlParams_     gensupport.URLParams
	ifNoneMatch_   string
	ctx_           context.Context
	header_        http.Header
}

// Get: Retrieves information about the specified content owner.
func (r *ContentOwnersService) Get(contentOwnerId string) *ContentOwnersGetCall {
	c := &ContentOwnersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.contentOwnerId = contentOwnerId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ContentOwnersGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ContentOwnersGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentOwnersGetCall) Fields(s ...googleapi.Field) *ContentOwnersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ContentOwnersGetCall) IfNoneMatch(entityTag string) *ContentOwnersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ContentOwnersGetCall) Context(ctx context.Context) *ContentOwnersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ContentOwnersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ContentOwnersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "contentOwners/{contentOwnerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"contentOwnerId": c.contentOwnerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.contentOwners.get" call.
// Exactly one of *ContentOwner or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ContentOwner.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ContentOwnersGetCall) Do(opts ...googleapi.CallOption) (*ContentOwner, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ContentOwner{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about the specified content owner.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.contentOwners.get",
	//   "parameterOrder": [
	//     "contentOwnerId"
	//   ],
	//   "parameters": {
	//     "contentOwnerId": {
	//       "description": "The contentOwnerId parameter specifies a value that uniquely identifies the content owner.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "contentOwners/{contentOwnerId}",
	//   "response": {
	//     "$ref": "ContentOwner"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.contentOwners.list":

type ContentOwnersListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of content owners that match the request
// criteria.
func (r *ContentOwnersService) List() *ContentOwnersListCall {
	c := &ContentOwnersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// FetchMine sets the optional parameter "fetchMine": The fetchMine
// parameter restricts the result set to content owners associated with
// the currently authenticated API user.
func (c *ContentOwnersListCall) FetchMine(fetchMine bool) *ContentOwnersListCall {
	c.urlParams_.Set("fetchMine", fmt.Sprint(fetchMine))
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of YouTube content owner IDs to retrieve.
func (c *ContentOwnersListCall) Id(id string) *ContentOwnersListCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ContentOwnersListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ContentOwnersListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentOwnersListCall) Fields(s ...googleapi.Field) *ContentOwnersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ContentOwnersListCall) IfNoneMatch(entityTag string) *ContentOwnersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ContentOwnersListCall) Context(ctx context.Context) *ContentOwnersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ContentOwnersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ContentOwnersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "contentOwners")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.contentOwners.list" call.
// Exactly one of *ContentOwnerListResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ContentOwnerListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ContentOwnersListCall) Do(opts ...googleapi.CallOption) (*ContentOwnerListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ContentOwnerListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of content owners that match the request criteria.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.contentOwners.list",
	//   "parameters": {
	//     "fetchMine": {
	//       "description": "The fetchMine parameter restricts the result set to content owners associated with the currently authenticated API user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of YouTube content owner IDs to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "contentOwners",
	//   "response": {
	//     "$ref": "ContentOwnerListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/youtubepartner-content-owner-readonly"
	//   ]
	// }

}

// method id "youtubePartner.liveCuepoints.insert":

type LiveCuepointsInsertCall struct {
	s            *Service
	livecuepoint *LiveCuepoint
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Insert: Inserts a cuepoint into a live broadcast.
func (r *LiveCuepointsService) Insert(channelId string, livecuepoint *LiveCuepoint) *LiveCuepointsInsertCall {
	c := &LiveCuepointsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("channelId", channelId)
	c.livecuepoint = livecuepoint
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners. You can obtain the content owner ID that
// will serve as the parameter value by calling the YouTube Content ID
// API's contentOwners.list method.
func (c *LiveCuepointsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *LiveCuepointsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LiveCuepointsInsertCall) Fields(s ...googleapi.Field) *LiveCuepointsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LiveCuepointsInsertCall) Context(ctx context.Context) *LiveCuepointsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *LiveCuepointsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *LiveCuepointsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.livecuepoint)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "liveCuepoints")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.liveCuepoints.insert" call.
// Exactly one of *LiveCuepoint or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *LiveCuepoint.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *LiveCuepointsInsertCall) Do(opts ...googleapi.CallOption) (*LiveCuepoint, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &LiveCuepoint{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a cuepoint into a live broadcast.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.liveCuepoints.insert",
	//   "parameterOrder": [
	//     "channelId"
	//   ],
	//   "parameters": {
	//     "channelId": {
	//       "description": "The channelId parameter identifies the channel that owns the broadcast into which the cuepoint is being inserted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners. You can obtain the content owner ID that will serve as the parameter value by calling the YouTube Content ID API's contentOwners.list method.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "liveCuepoints",
	//   "request": {
	//     "$ref": "LiveCuepoint"
	//   },
	//   "response": {
	//     "$ref": "LiveCuepoint"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.metadataHistory.list":

type MetadataHistoryListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of all metadata provided for an asset,
// regardless of which content owner provided the data.
func (r *MetadataHistoryService) List(assetId string) *MetadataHistoryListCall {
	c := &MetadataHistoryListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("assetId", assetId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *MetadataHistoryListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *MetadataHistoryListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetadataHistoryListCall) Fields(s ...googleapi.Field) *MetadataHistoryListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MetadataHistoryListCall) IfNoneMatch(entityTag string) *MetadataHistoryListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MetadataHistoryListCall) Context(ctx context.Context) *MetadataHistoryListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *MetadataHistoryListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *MetadataHistoryListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "metadataHistory")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.metadataHistory.list" call.
// Exactly one of *MetadataHistoryListResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *MetadataHistoryListResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MetadataHistoryListCall) Do(opts ...googleapi.CallOption) (*MetadataHistoryListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &MetadataHistoryListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of all metadata provided for an asset, regardless of which content owner provided the data.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.metadataHistory.list",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving a metadata history.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "metadataHistory",
	//   "response": {
	//     "$ref": "MetadataHistoryListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.orders.delete":

type OrdersDeleteCall struct {
	s          *Service
	orderId    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an order, which moves orders to inactive state and
// removes any associated video.
func (r *OrdersService) Delete(orderId string) *OrdersDeleteCall {
	c := &OrdersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.orderId = orderId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": ContentOwnerId that super admin acts in
// behalf of.
func (c *OrdersDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OrdersDeleteCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrdersDeleteCall) Fields(s ...googleapi.Field) *OrdersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrdersDeleteCall) Context(ctx context.Context) *OrdersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrdersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrdersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "orders/{orderId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"orderId": c.orderId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.orders.delete" call.
func (c *OrdersDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Delete an order, which moves orders to inactive state and removes any associated video.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubePartner.orders.delete",
	//   "parameterOrder": [
	//     "orderId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "ContentOwnerId that super admin acts in behalf of.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderId": {
	//       "description": "Id of the order to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "orders/{orderId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.orders.get":

type OrdersGetCall struct {
	s            *Service
	orderId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieve the details of an existing order.
func (r *OrdersService) Get(orderId string) *OrdersGetCall {
	c := &OrdersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.orderId = orderId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": ContentOnwerId that super admin acts in
// behalf of.
func (c *OrdersGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OrdersGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrdersGetCall) Fields(s ...googleapi.Field) *OrdersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrdersGetCall) IfNoneMatch(entityTag string) *OrdersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrdersGetCall) Context(ctx context.Context) *OrdersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrdersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrdersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "orders/{orderId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"orderId": c.orderId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.orders.get" call.
// Exactly one of *Order or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Order.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrdersGetCall) Do(opts ...googleapi.CallOption) (*Order, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Order{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve the details of an existing order.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.orders.get",
	//   "parameterOrder": [
	//     "orderId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "ContentOnwerId that super admin acts in behalf of.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderId": {
	//       "description": "The id of the order.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "orders/{orderId}",
	//   "response": {
	//     "$ref": "Order"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.orders.insert":

type OrdersInsertCall struct {
	s          *Service
	order      *Order
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a new basic order entry in the YouTube premium asset
// order management system. You must supply at least a country and
// channel in the new order.
func (r *OrdersService) Insert(order *Order) *OrdersInsertCall {
	c := &OrdersInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.order = order
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": ContentOnwerId that super admin acts in
// behalf of.
func (c *OrdersInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OrdersInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrdersInsertCall) Fields(s ...googleapi.Field) *OrdersInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrdersInsertCall) Context(ctx context.Context) *OrdersInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrdersInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrdersInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.order)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "orders")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.orders.insert" call.
// Exactly one of *Order or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Order.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrdersInsertCall) Do(opts ...googleapi.CallOption) (*Order, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Order{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new basic order entry in the YouTube premium asset order management system. You must supply at least a country and channel in the new order.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.orders.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "ContentOnwerId that super admin acts in behalf of.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "orders",
	//   "request": {
	//     "$ref": "Order"
	//   },
	//   "response": {
	//     "$ref": "Order"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.orders.list":

type OrdersListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Return a list of orders, filtered by the parameters below, may
// return more than a single page of results.
func (r *OrdersService) List() *OrdersListCall {
	c := &OrdersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// ChannelId sets the optional parameter "channelId": Filter results to
// only a specific channel ( use encrypted ID).
func (c *OrdersListCall) ChannelId(channelId string) *OrdersListCall {
	c.urlParams_.Set("channelId", channelId)
	return c
}

// ContentType sets the optional parameter "contentType": Filter the
// results by type, possible values are SHOW or MOVIE.
func (c *OrdersListCall) ContentType(contentType string) *OrdersListCall {
	c.urlParams_.Set("contentType", contentType)
	return c
}

// Country sets the optional parameter "country": Filter results by
// country, two letter ISO country codes are used.
func (c *OrdersListCall) Country(country string) *OrdersListCall {
	c.urlParams_.Set("country", country)
	return c
}

// CustomId sets the optional parameter "customId": Filter result by
// orders that have this custom ID.
func (c *OrdersListCall) CustomId(customId string) *OrdersListCall {
	c.urlParams_.Set("customId", customId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": ContentOnwerId that super admin acts in
// behalf of.
func (c *OrdersListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OrdersListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token is an optional value that is only used to page through large
// result sets.
//
// - To retrieve the next page of results for a request, set this
// parameter to the value of the nextPageToken value from the previous
// response.
// - To get the previous page of results, set this parameter to the
// value of the previousPageToken value from the previous response.
func (c *OrdersListCall) PageToken(pageToken string) *OrdersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Priority sets the optional parameter "priority": Filter results by
// priority. P0, P1, P2, P3 and P4 are the acceptable options.
func (c *OrdersListCall) Priority(priority string) *OrdersListCall {
	c.urlParams_.Set("priority", priority)
	return c
}

// ProductionHouse sets the optional parameter "productionHouse": Filter
// results by a particular production house. Specified by the name of
// the production house.
func (c *OrdersListCall) ProductionHouse(productionHouse string) *OrdersListCall {
	c.urlParams_.Set("productionHouse", productionHouse)
	return c
}

// Q sets the optional parameter "q": Filter results to only orders that
// contain this string in the title.
func (c *OrdersListCall) Q(q string) *OrdersListCall {
	c.urlParams_.Set("q", q)
	return c
}

// Status sets the optional parameter "status": Filter results to have
// this status, available options are STATUS_AVAILED, STATUS_ORDERED,
// STATUS_RECEIVED, STATUS_READY_FOR_QC, STATUS_MOC_FIX,
// STATUS_PARTNER_FIX, STATUS_YOUTUBE_FIX, STATUS_QC_APPROVED,
// STATUS_INACTIVE, STATUS_INGESTION_COMPLETE, STATUS_REORDERED
func (c *OrdersListCall) Status(status string) *OrdersListCall {
	c.urlParams_.Set("status", status)
	return c
}

// VideoId sets the optional parameter "videoId": Filter results to
// orders that are associated with this YouTube external video id.
func (c *OrdersListCall) VideoId(videoId string) *OrdersListCall {
	c.urlParams_.Set("videoId", videoId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrdersListCall) Fields(s ...googleapi.Field) *OrdersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrdersListCall) IfNoneMatch(entityTag string) *OrdersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrdersListCall) Context(ctx context.Context) *OrdersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrdersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrdersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "orders")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.orders.list" call.
// Exactly one of *OrderListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *OrderListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrdersListCall) Do(opts ...googleapi.CallOption) (*OrderListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrderListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return a list of orders, filtered by the parameters below, may return more than a single page of results.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.orders.list",
	//   "parameters": {
	//     "channelId": {
	//       "description": "Filter results to only a specific channel ( use encrypted ID).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "contentType": {
	//       "description": "Filter the results by type, possible values are SHOW or MOVIE.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "country": {
	//       "description": "Filter results by country, two letter ISO country codes are used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customId": {
	//       "description": "Filter result by orders that have this custom ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "ContentOnwerId that super admin acts in behalf of.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token is an optional value that is only used to page through large result sets.\n\n- To retrieve the next page of results for a request, set this parameter to the value of the nextPageToken value from the previous response.\n- To get the previous page of results, set this parameter to the value of the previousPageToken value from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "priority": {
	//       "description": "Filter results by priority. P0, P1, P2, P3 and P4 are the acceptable options.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productionHouse": {
	//       "description": "Filter results by a particular production house. Specified by the name of the production house.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Filter results to only orders that contain this string in the title.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "Filter results to have this status, available options are STATUS_AVAILED, STATUS_ORDERED, STATUS_RECEIVED, STATUS_READY_FOR_QC, STATUS_MOC_FIX, STATUS_PARTNER_FIX, STATUS_YOUTUBE_FIX, STATUS_QC_APPROVED, STATUS_INACTIVE, STATUS_INGESTION_COMPLETE, STATUS_REORDERED",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "Filter results to orders that are associated with this YouTube external video id.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "orders",
	//   "response": {
	//     "$ref": "OrderListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrdersListCall) Pages(ctx context.Context, f func(*OrderListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.orders.patch":

type OrdersPatchCall struct {
	s          *Service
	orderId    string
	order      *Order
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Update the values in an existing order. This method supports
// patch semantics.
func (r *OrdersService) Patch(orderId string, order *Order) *OrdersPatchCall {
	c := &OrdersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.orderId = orderId
	c.order = order
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": ContentOwnerId that super admin acts in
// behalf of.
func (c *OrdersPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OrdersPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrdersPatchCall) Fields(s ...googleapi.Field) *OrdersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrdersPatchCall) Context(ctx context.Context) *OrdersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrdersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrdersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.order)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "orders/{orderId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"orderId": c.orderId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.orders.patch" call.
// Exactly one of *Order or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Order.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrdersPatchCall) Do(opts ...googleapi.CallOption) (*Order, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Order{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the values in an existing order. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.orders.patch",
	//   "parameterOrder": [
	//     "orderId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "ContentOwnerId that super admin acts in behalf of.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderId": {
	//       "description": "The id of the order.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "orders/{orderId}",
	//   "request": {
	//     "$ref": "Order"
	//   },
	//   "response": {
	//     "$ref": "Order"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.orders.update":

type OrdersUpdateCall struct {
	s          *Service
	orderId    string
	order      *Order
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Update the values in an existing order.
func (r *OrdersService) Update(orderId string, order *Order) *OrdersUpdateCall {
	c := &OrdersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.orderId = orderId
	c.order = order
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": ContentOwnerId that super admin acts in
// behalf of.
func (c *OrdersUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OrdersUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrdersUpdateCall) Fields(s ...googleapi.Field) *OrdersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrdersUpdateCall) Context(ctx context.Context) *OrdersUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrdersUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrdersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.order)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "orders/{orderId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"orderId": c.orderId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.orders.update" call.
// Exactly one of *Order or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Order.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrdersUpdateCall) Do(opts ...googleapi.CallOption) (*Order, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Order{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update the values in an existing order.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.orders.update",
	//   "parameterOrder": [
	//     "orderId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "ContentOwnerId that super admin acts in behalf of.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderId": {
	//       "description": "The id of the order.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "orders/{orderId}",
	//   "request": {
	//     "$ref": "Order"
	//   },
	//   "response": {
	//     "$ref": "Order"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.ownership.get":

type OwnershipGetCall struct {
	s            *Service
	assetId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the ownership data provided for the specified asset by
// the content owner associated with the authenticated user.
func (r *OwnershipService) Get(assetId string) *OwnershipGetCall {
	c := &OwnershipGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *OwnershipGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OwnershipGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OwnershipGetCall) Fields(s ...googleapi.Field) *OwnershipGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OwnershipGetCall) IfNoneMatch(entityTag string) *OwnershipGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OwnershipGetCall) Context(ctx context.Context) *OwnershipGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OwnershipGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OwnershipGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}/ownership")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.ownership.get" call.
// Exactly one of *RightsOwnership or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RightsOwnership.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OwnershipGetCall) Do(opts ...googleapi.CallOption) (*RightsOwnership, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RightsOwnership{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the ownership data provided for the specified asset by the content owner associated with the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.ownership.get",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID for which you are retrieving ownership data.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}/ownership",
	//   "response": {
	//     "$ref": "RightsOwnership"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.ownership.patch":

type OwnershipPatchCall struct {
	s               *Service
	assetId         string
	rightsownership *RightsOwnership
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Patch: Provides new ownership information for the specified asset.
// Note that YouTube may receive ownership information from multiple
// sources. For example, if an asset has multiple owners, each owner
// might send ownership data for the asset. YouTube algorithmically
// combines the ownership data received from all of those sources to
// generate the asset's canonical ownership data, which should provide
// the most comprehensive and accurate representation of the asset's
// ownership. This method supports patch semantics.
func (r *OwnershipService) Patch(assetId string, rightsownership *RightsOwnership) *OwnershipPatchCall {
	c := &OwnershipPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	c.rightsownership = rightsownership
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *OwnershipPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OwnershipPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OwnershipPatchCall) Fields(s ...googleapi.Field) *OwnershipPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OwnershipPatchCall) Context(ctx context.Context) *OwnershipPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OwnershipPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OwnershipPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rightsownership)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}/ownership")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.ownership.patch" call.
// Exactly one of *RightsOwnership or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RightsOwnership.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OwnershipPatchCall) Do(opts ...googleapi.CallOption) (*RightsOwnership, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RightsOwnership{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Provides new ownership information for the specified asset. Note that YouTube may receive ownership information from multiple sources. For example, if an asset has multiple owners, each owner might send ownership data for the asset. YouTube algorithmically combines the ownership data received from all of those sources to generate the asset's canonical ownership data, which should provide the most comprehensive and accurate representation of the asset's ownership. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.ownership.patch",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}/ownership",
	//   "request": {
	//     "$ref": "RightsOwnership"
	//   },
	//   "response": {
	//     "$ref": "RightsOwnership"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.ownership.update":

type OwnershipUpdateCall struct {
	s               *Service
	assetId         string
	rightsownership *RightsOwnership
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Update: Provides new ownership information for the specified asset.
// Note that YouTube may receive ownership information from multiple
// sources. For example, if an asset has multiple owners, each owner
// might send ownership data for the asset. YouTube algorithmically
// combines the ownership data received from all of those sources to
// generate the asset's canonical ownership data, which should provide
// the most comprehensive and accurate representation of the asset's
// ownership.
func (r *OwnershipService) Update(assetId string, rightsownership *RightsOwnership) *OwnershipUpdateCall {
	c := &OwnershipUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.assetId = assetId
	c.rightsownership = rightsownership
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *OwnershipUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OwnershipUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OwnershipUpdateCall) Fields(s ...googleapi.Field) *OwnershipUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OwnershipUpdateCall) Context(ctx context.Context) *OwnershipUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OwnershipUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OwnershipUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rightsownership)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{assetId}/ownership")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"assetId": c.assetId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.ownership.update" call.
// Exactly one of *RightsOwnership or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RightsOwnership.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OwnershipUpdateCall) Do(opts ...googleapi.CallOption) (*RightsOwnership, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RightsOwnership{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Provides new ownership information for the specified asset. Note that YouTube may receive ownership information from multiple sources. For example, if an asset has multiple owners, each owner might send ownership data for the asset. YouTube algorithmically combines the ownership data received from all of those sources to generate the asset's canonical ownership data, which should provide the most comprehensive and accurate representation of the asset's ownership.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.ownership.update",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{assetId}/ownership",
	//   "request": {
	//     "$ref": "RightsOwnership"
	//   },
	//   "response": {
	//     "$ref": "RightsOwnership"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.ownershipHistory.list":

type OwnershipHistoryListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of the ownership data for an asset, regardless
// of which content owner provided the data. The list only includes the
// most recent ownership data for each content owner. However, if the
// content owner has submitted ownership data through multiple data
// sources (API, content feeds, etc.), the list will contain the most
// recent data for each content owner and data source.
func (r *OwnershipHistoryService) List(assetId string) *OwnershipHistoryListCall {
	c := &OwnershipHistoryListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("assetId", assetId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *OwnershipHistoryListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *OwnershipHistoryListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OwnershipHistoryListCall) Fields(s ...googleapi.Field) *OwnershipHistoryListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OwnershipHistoryListCall) IfNoneMatch(entityTag string) *OwnershipHistoryListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OwnershipHistoryListCall) Context(ctx context.Context) *OwnershipHistoryListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OwnershipHistoryListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OwnershipHistoryListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "ownershipHistory")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.ownershipHistory.list" call.
// Exactly one of *OwnershipHistoryListResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *OwnershipHistoryListResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OwnershipHistoryListCall) Do(opts ...googleapi.CallOption) (*OwnershipHistoryListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OwnershipHistoryListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of the ownership data for an asset, regardless of which content owner provided the data. The list only includes the most recent ownership data for each content owner. However, if the content owner has submitted ownership data through multiple data sources (API, content feeds, etc.), the list will contain the most recent data for each content owner and data source.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.ownershipHistory.list",
	//   "parameterOrder": [
	//     "assetId"
	//   ],
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving an ownership data history.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "ownershipHistory",
	//   "response": {
	//     "$ref": "OwnershipHistoryListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.package.get":

type PackageGetCall struct {
	s            *Service
	packageId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves information for the specified package.
func (r *PackageService) Get(packageId string) *PackageGetCall {
	c := &PackageGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.packageId = packageId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PackageGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PackageGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PackageGetCall) Fields(s ...googleapi.Field) *PackageGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PackageGetCall) IfNoneMatch(entityTag string) *PackageGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PackageGetCall) Context(ctx context.Context) *PackageGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PackageGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PackageGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "package/{packageId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"packageId": c.packageId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.package.get" call.
// Exactly one of *Package or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Package.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PackageGetCall) Do(opts ...googleapi.CallOption) (*Package, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Package{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information for the specified package.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.package.get",
	//   "parameterOrder": [
	//     "packageId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "packageId": {
	//       "description": "The packageId parameter specifies the Content Delivery package ID of the package being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "package/{packageId}",
	//   "response": {
	//     "$ref": "Package"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.package.insert":

type PackageInsertCall struct {
	s          *Service
	package_   *Package
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Inserts a metadata-only package.
func (r *PackageService) Insert(package_ *Package) *PackageInsertCall {
	c := &PackageInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.package_ = package_
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PackageInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PackageInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PackageInsertCall) Fields(s ...googleapi.Field) *PackageInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PackageInsertCall) Context(ctx context.Context) *PackageInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PackageInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PackageInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.package_)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "package")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.package.insert" call.
// Exactly one of *PackageInsertResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *PackageInsertResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PackageInsertCall) Do(opts ...googleapi.CallOption) (*PackageInsertResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PackageInsertResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a metadata-only package.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.package.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "package",
	//   "request": {
	//     "$ref": "Package"
	//   },
	//   "response": {
	//     "$ref": "PackageInsertResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.policies.get":

type PoliciesGetCall struct {
	s            *Service
	policyId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified saved policy.
func (r *PoliciesService) Get(policyId string) *PoliciesGetCall {
	c := &PoliciesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.policyId = policyId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PoliciesGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PoliciesGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PoliciesGetCall) Fields(s ...googleapi.Field) *PoliciesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PoliciesGetCall) IfNoneMatch(entityTag string) *PoliciesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PoliciesGetCall) Context(ctx context.Context) *PoliciesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PoliciesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PoliciesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "policies/{policyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"policyId": c.policyId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.policies.get" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PoliciesGetCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified saved policy.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.policies.get",
	//   "parameterOrder": [
	//     "policyId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "policyId": {
	//       "description": "The policyId parameter specifies a value that uniquely identifies the policy being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "policies/{policyId}",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.policies.insert":

type PoliciesInsertCall struct {
	s          *Service
	policy     *Policy
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a saved policy.
func (r *PoliciesService) Insert(policy *Policy) *PoliciesInsertCall {
	c := &PoliciesInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.policy = policy
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PoliciesInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PoliciesInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PoliciesInsertCall) Fields(s ...googleapi.Field) *PoliciesInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PoliciesInsertCall) Context(ctx context.Context) *PoliciesInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PoliciesInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PoliciesInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.policy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "policies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.policies.insert" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PoliciesInsertCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a saved policy.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.policies.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "policies",
	//   "request": {
	//     "$ref": "Policy"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.policies.list":

type PoliciesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of the content owner's saved policies.
func (r *PoliciesService) List() *PoliciesListCall {
	c := &PoliciesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of saved policy IDs to retrieve. Only policies
// belonging to the currently authenticated content owner will be
// available.
func (c *PoliciesListCall) Id(id string) *PoliciesListCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PoliciesListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PoliciesListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Sort sets the optional parameter "sort": The sort parameter specifies
// how the search results should be sorted.
//
// Possible values:
//   "timeUpdatedAsc" - Sort by the update time ascending.
//   "timeUpdatedDesc" - Sort by the update time descending.
func (c *PoliciesListCall) Sort(sort string) *PoliciesListCall {
	c.urlParams_.Set("sort", sort)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PoliciesListCall) Fields(s ...googleapi.Field) *PoliciesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PoliciesListCall) IfNoneMatch(entityTag string) *PoliciesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PoliciesListCall) Context(ctx context.Context) *PoliciesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PoliciesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PoliciesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "policies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.policies.list" call.
// Exactly one of *PolicyList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *PolicyList.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *PoliciesListCall) Do(opts ...googleapi.CallOption) (*PolicyList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PolicyList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of the content owner's saved policies.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.policies.list",
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of saved policy IDs to retrieve. Only policies belonging to the currently authenticated content owner will be available.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The sort parameter specifies how the search results should be sorted.",
	//       "enum": [
	//         "timeUpdatedAsc",
	//         "timeUpdatedDesc"
	//       ],
	//       "enumDescriptions": [
	//         "Sort by the update time ascending.",
	//         "Sort by the update time descending."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "policies",
	//   "response": {
	//     "$ref": "PolicyList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.policies.patch":

type PoliciesPatchCall struct {
	s          *Service
	policyId   string
	policy     *Policy
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the specified saved policy. This method supports patch
// semantics.
func (r *PoliciesService) Patch(policyId string, policy *Policy) *PoliciesPatchCall {
	c := &PoliciesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.policyId = policyId
	c.policy = policy
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PoliciesPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PoliciesPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PoliciesPatchCall) Fields(s ...googleapi.Field) *PoliciesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PoliciesPatchCall) Context(ctx context.Context) *PoliciesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PoliciesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PoliciesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.policy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "policies/{policyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"policyId": c.policyId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.policies.patch" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PoliciesPatchCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified saved policy. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.policies.patch",
	//   "parameterOrder": [
	//     "policyId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "policyId": {
	//       "description": "The policyId parameter specifies a value that uniquely identifies the policy being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "policies/{policyId}",
	//   "request": {
	//     "$ref": "Policy"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.policies.update":

type PoliciesUpdateCall struct {
	s          *Service
	policyId   string
	policy     *Policy
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates the specified saved policy.
func (r *PoliciesService) Update(policyId string, policy *Policy) *PoliciesUpdateCall {
	c := &PoliciesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.policyId = policyId
	c.policy = policy
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PoliciesUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PoliciesUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PoliciesUpdateCall) Fields(s ...googleapi.Field) *PoliciesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PoliciesUpdateCall) Context(ctx context.Context) *PoliciesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PoliciesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PoliciesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.policy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "policies/{policyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"policyId": c.policyId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.policies.update" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *PoliciesUpdateCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified saved policy.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.policies.update",
	//   "parameterOrder": [
	//     "policyId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "policyId": {
	//       "description": "The policyId parameter specifies a value that uniquely identifies the policy being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "policies/{policyId}",
	//   "request": {
	//     "$ref": "Policy"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.publishers.get":

type PublishersGetCall struct {
	s            *Service
	publisherId  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves information about the specified publisher.
func (r *PublishersService) Get(publisherId string) *PublishersGetCall {
	c := &PublishersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.publisherId = publisherId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PublishersGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PublishersGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PublishersGetCall) Fields(s ...googleapi.Field) *PublishersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PublishersGetCall) IfNoneMatch(entityTag string) *PublishersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PublishersGetCall) Context(ctx context.Context) *PublishersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PublishersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PublishersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "publishers/{publisherId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"publisherId": c.publisherId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.publishers.get" call.
// Exactly one of *Publisher or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Publisher.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *PublishersGetCall) Do(opts ...googleapi.CallOption) (*Publisher, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Publisher{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about the specified publisher.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.publishers.get",
	//   "parameterOrder": [
	//     "publisherId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "publisherId": {
	//       "description": "The publisherId parameter specifies a publisher ID that uniquely identifies the publisher being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "publishers/{publisherId}",
	//   "response": {
	//     "$ref": "Publisher"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.publishers.list":

type PublishersListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of publishers that match the request criteria.
// This method is analogous to a publisher search function.
func (r *PublishersService) List() *PublishersListCall {
	c := &PublishersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// CaeNumber sets the optional parameter "caeNumber": The caeNumber
// parameter specifies the CAE number of the publisher that you want to
// retrieve.
func (c *PublishersListCall) CaeNumber(caeNumber string) *PublishersListCall {
	c.urlParams_.Set("caeNumber", caeNumber)
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of YouTube publisher IDs to retrieve.
func (c *PublishersListCall) Id(id string) *PublishersListCall {
	c.urlParams_.Set("id", id)
	return c
}

// IpiNumber sets the optional parameter "ipiNumber": The ipiNumber
// parameter specifies the IPI number of the publisher that you want to
// retrieve.
func (c *PublishersListCall) IpiNumber(ipiNumber string) *PublishersListCall {
	c.urlParams_.Set("ipiNumber", ipiNumber)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maxResults
// parameter specifies the maximum number of results to return per page.
func (c *PublishersListCall) MaxResults(maxResults int64) *PublishersListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// NamePrefix sets the optional parameter "namePrefix": The namePrefix
// parameter indicates that the API should only return publishers whose
// name starts with this prefix.
func (c *PublishersListCall) NamePrefix(namePrefix string) *PublishersListCall {
	c.urlParams_.Set("namePrefix", namePrefix)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *PublishersListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *PublishersListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of search results.
func (c *PublishersListCall) PageToken(pageToken string) *PublishersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PublishersListCall) Fields(s ...googleapi.Field) *PublishersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PublishersListCall) IfNoneMatch(entityTag string) *PublishersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PublishersListCall) Context(ctx context.Context) *PublishersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PublishersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PublishersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "publishers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.publishers.list" call.
// Exactly one of *PublisherList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *PublisherList.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PublishersListCall) Do(opts ...googleapi.CallOption) (*PublisherList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PublisherList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of publishers that match the request criteria. This method is analogous to a publisher search function.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.publishers.list",
	//   "parameters": {
	//     "caeNumber": {
	//       "description": "The caeNumber parameter specifies the CAE number of the publisher that you want to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of YouTube publisher IDs to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ipiNumber": {
	//       "description": "The ipiNumber parameter specifies the IPI number of the publisher that you want to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "25",
	//       "description": "The maxResults parameter specifies the maximum number of results to return per page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "namePrefix": {
	//       "description": "The namePrefix parameter indicates that the API should only return publishers whose name starts with this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "publishers",
	//   "response": {
	//     "$ref": "PublisherList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *PublishersListCall) Pages(ctx context.Context, f func(*PublisherList) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.referenceConflicts.get":

type ReferenceConflictsGetCall struct {
	s                   *Service
	referenceConflictId string
	urlParams_          gensupport.URLParams
	ifNoneMatch_        string
	ctx_                context.Context
	header_             http.Header
}

// Get: Retrieves information about the specified reference conflict.
func (r *ReferenceConflictsService) Get(referenceConflictId string) *ReferenceConflictsGetCall {
	c := &ReferenceConflictsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.referenceConflictId = referenceConflictId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferenceConflictsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferenceConflictsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferenceConflictsGetCall) Fields(s ...googleapi.Field) *ReferenceConflictsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReferenceConflictsGetCall) IfNoneMatch(entityTag string) *ReferenceConflictsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReferenceConflictsGetCall) Context(ctx context.Context) *ReferenceConflictsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferenceConflictsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferenceConflictsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "referenceConflicts/{referenceConflictId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"referenceConflictId": c.referenceConflictId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.referenceConflicts.get" call.
// Exactly one of *ReferenceConflict or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ReferenceConflict.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ReferenceConflictsGetCall) Do(opts ...googleapi.CallOption) (*ReferenceConflict, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ReferenceConflict{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about the specified reference conflict.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.referenceConflicts.get",
	//   "parameterOrder": [
	//     "referenceConflictId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceConflictId": {
	//       "description": "The referenceConflictId parameter specifies the YouTube reference conflict ID of the reference conflict being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "referenceConflicts/{referenceConflictId}",
	//   "response": {
	//     "$ref": "ReferenceConflict"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.referenceConflicts.list":

type ReferenceConflictsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of unresolved reference conflicts.
func (r *ReferenceConflictsService) List() *ReferenceConflictsListCall {
	c := &ReferenceConflictsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferenceConflictsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferenceConflictsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of search results.
func (c *ReferenceConflictsListCall) PageToken(pageToken string) *ReferenceConflictsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferenceConflictsListCall) Fields(s ...googleapi.Field) *ReferenceConflictsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReferenceConflictsListCall) IfNoneMatch(entityTag string) *ReferenceConflictsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReferenceConflictsListCall) Context(ctx context.Context) *ReferenceConflictsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferenceConflictsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferenceConflictsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "referenceConflicts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.referenceConflicts.list" call.
// Exactly one of *ReferenceConflictListResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ReferenceConflictListResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ReferenceConflictsListCall) Do(opts ...googleapi.CallOption) (*ReferenceConflictListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ReferenceConflictListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of unresolved reference conflicts.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.referenceConflicts.list",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "referenceConflicts",
	//   "response": {
	//     "$ref": "ReferenceConflictListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ReferenceConflictsListCall) Pages(ctx context.Context, f func(*ReferenceConflictListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.references.get":

type ReferencesGetCall struct {
	s            *Service
	referenceId  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves information about the specified reference.
func (r *ReferencesService) Get(referenceId string) *ReferencesGetCall {
	c := &ReferencesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.referenceId = referenceId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferencesGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferencesGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesGetCall) Fields(s ...googleapi.Field) *ReferencesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReferencesGetCall) IfNoneMatch(entityTag string) *ReferencesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReferencesGetCall) Context(ctx context.Context) *ReferencesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferencesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferencesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "references/{referenceId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"referenceId": c.referenceId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.references.get" call.
// Exactly one of *Reference or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Reference.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ReferencesGetCall) Do(opts ...googleapi.CallOption) (*Reference, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Reference{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about the specified reference.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.references.get",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceId": {
	//       "description": "The referenceId parameter specifies the YouTube reference ID of the reference being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "references/{referenceId}",
	//   "response": {
	//     "$ref": "Reference"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.references.insert":

type ReferencesInsertCall struct {
	s                *Service
	reference        *Reference
	urlParams_       gensupport.URLParams
	media_           io.Reader
	mediaBuffer_     *gensupport.MediaBuffer
	mediaType_       string
	mediaSize_       int64 // mediaSize, if known.  Used only for calls to progressUpdater_.
	progressUpdater_ googleapi.ProgressUpdater
	ctx_             context.Context
	header_          http.Header
}

// Insert: Creates a reference in one of the following ways:
// - If your request is uploading a reference file, YouTube creates the
// reference from the provided content. You can provide either a
// video/audio file or a pre-generated fingerprint. If you are providing
// a pre-generated fingerprint, set the reference resource's fpDirect
// property to true in the request body. In this flow, you can use
// either the multipart or resumable upload flows to provide the
// reference content.
// - If you want to create a reference using a claimed video as the
// reference content, use the claimId parameter to identify the claim.
func (r *ReferencesService) Insert(reference *Reference) *ReferencesInsertCall {
	c := &ReferencesInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.reference = reference
	return c
}

// ClaimId sets the optional parameter "claimId": The claimId parameter
// specifies the YouTube claim ID of an existing claim from which a
// reference should be created. (The claimed video is used as the
// reference content.)
func (c *ReferencesInsertCall) ClaimId(claimId string) *ReferencesInsertCall {
	c.urlParams_.Set("claimId", claimId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferencesInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferencesInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Media specifies the media to upload in one or more chunks. The chunk
// size may be controlled by supplying a MediaOption generated by
// googleapi.ChunkSize. The chunk size defaults to
// googleapi.DefaultUploadChunkSize.The Content-Type header used in the
// upload request will be determined by sniffing the contents of r,
// unless a MediaOption generated by googleapi.ContentType is
// supplied.
// At most one of Media and ResumableMedia may be set.
func (c *ReferencesInsertCall) Media(r io.Reader, options ...googleapi.MediaOption) *ReferencesInsertCall {
	if ct := c.reference.ContentType; ct != "" {
		options = append([]googleapi.MediaOption{googleapi.ContentType(ct)}, options...)
	}
	opts := googleapi.ProcessMediaOptions(options)
	chunkSize := opts.ChunkSize
	if !opts.ForceEmptyContentType {
		r, c.mediaType_ = gensupport.DetermineContentType(r, opts.ContentType)
	}
	c.media_, c.mediaBuffer_ = gensupport.PrepareUpload(r, chunkSize)
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx.
//
// Deprecated: use Media instead.
//
// At most one of Media and ResumableMedia may be set. mediaType
// identifies the MIME media type of the upload, such as "image/png". If
// mediaType is "", it will be auto-detected. The provided ctx will
// supersede any context previously provided to the Context method.
func (c *ReferencesInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ReferencesInsertCall {
	c.ctx_ = ctx
	rdr := gensupport.ReaderAtToReader(r, size)
	rdr, c.mediaType_ = gensupport.DetermineContentType(rdr, mediaType)
	c.mediaBuffer_ = gensupport.NewMediaBuffer(rdr, googleapi.DefaultUploadChunkSize)
	c.media_ = nil
	c.mediaSize_ = size
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *ReferencesInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ReferencesInsertCall {
	c.progressUpdater_ = pu
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesInsertCall) Fields(s ...googleapi.Field) *ReferencesInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *ReferencesInsertCall) Context(ctx context.Context) *ReferencesInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferencesInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferencesInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reference)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "references")
	if c.media_ != nil || c.mediaBuffer_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		protocol := "multipart"
		if c.mediaBuffer_ != nil {
			protocol = "resumable"
		}
		c.urlParams_.Set("uploadType", protocol)
	}
	if body == nil {
		body = new(bytes.Buffer)
		reqHeaders.Set("Content-Type", "application/json")
	}
	if c.media_ != nil {
		combined, ctype := gensupport.CombineBodyMedia(body, "application/json", c.media_, c.mediaType_)
		defer combined.Close()
		reqHeaders.Set("Content-Type", ctype)
		body = combined
	}
	if c.mediaBuffer_ != nil && c.mediaType_ != "" {
		reqHeaders.Set("X-Upload-Content-Type", c.mediaType_)
	}
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.references.insert" call.
// Exactly one of *Reference or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Reference.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ReferencesInsertCall) Do(opts ...googleapi.CallOption) (*Reference, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	if c.mediaBuffer_ != nil {
		loc := res.Header.Get("Location")
		rx := &gensupport.ResumableUpload{
			Client:    c.s.client,
			UserAgent: c.s.userAgent(),
			URI:       loc,
			Media:     c.mediaBuffer_,
			MediaType: c.mediaType_,
			Callback: func(curr int64) {
				if c.progressUpdater_ != nil {
					c.progressUpdater_(curr, c.mediaSize_)
				}
			},
		}
		ctx := c.ctx_
		if ctx == nil {
			ctx = context.TODO()
		}
		res, err = rx.Upload(ctx)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if err := googleapi.CheckResponse(res); err != nil {
			return nil, err
		}
	}
	ret := &Reference{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a reference in one of the following ways:\n- If your request is uploading a reference file, YouTube creates the reference from the provided content. You can provide either a video/audio file or a pre-generated fingerprint. If you are providing a pre-generated fingerprint, set the reference resource's fpDirect property to true in the request body. In this flow, you can use either the multipart or resumable upload flows to provide the reference content.\n- If you want to create a reference using a claimed video as the reference content, use the claimId parameter to identify the claim.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.references.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "20GB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/youtube/partner/v1/references"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/youtube/partner/v1/references"
	//       }
	//     }
	//   },
	//   "parameters": {
	//     "claimId": {
	//       "description": "The claimId parameter specifies the YouTube claim ID of an existing claim from which a reference should be created. (The claimed video is used as the reference content.)",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "references",
	//   "request": {
	//     "$ref": "Reference"
	//   },
	//   "response": {
	//     "$ref": "Reference"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "youtubePartner.references.list":

type ReferencesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of references by ID or the list of references
// for the specified asset.
func (r *ReferencesService) List() *ReferencesListCall {
	c := &ReferencesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// AssetId sets the optional parameter "assetId": The assetId parameter
// specifies the YouTube asset ID of the asset for which you are
// retrieving references.
func (c *ReferencesListCall) AssetId(assetId string) *ReferencesListCall {
	c.urlParams_.Set("assetId", assetId)
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of YouTube reference IDs to retrieve.
func (c *ReferencesListCall) Id(id string) *ReferencesListCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferencesListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferencesListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of search results.
func (c *ReferencesListCall) PageToken(pageToken string) *ReferencesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesListCall) Fields(s ...googleapi.Field) *ReferencesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReferencesListCall) IfNoneMatch(entityTag string) *ReferencesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReferencesListCall) Context(ctx context.Context) *ReferencesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferencesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferencesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "references")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.references.list" call.
// Exactly one of *ReferenceListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ReferenceListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ReferencesListCall) Do(opts ...googleapi.CallOption) (*ReferenceListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ReferenceListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of references by ID or the list of references for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.references.list",
	//   "parameters": {
	//     "assetId": {
	//       "description": "The assetId parameter specifies the YouTube asset ID of the asset for which you are retrieving references.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of YouTube reference IDs to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of search results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "references",
	//   "response": {
	//     "$ref": "ReferenceListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ReferencesListCall) Pages(ctx context.Context, f func(*ReferenceListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubePartner.references.patch":

type ReferencesPatchCall struct {
	s           *Service
	referenceId string
	reference   *Reference
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Updates a reference. This method supports patch semantics.
func (r *ReferencesService) Patch(referenceId string, reference *Reference) *ReferencesPatchCall {
	c := &ReferencesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.referenceId = referenceId
	c.reference = reference
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferencesPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferencesPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// ReleaseClaims sets the optional parameter "releaseClaims": The
// releaseClaims parameter indicates that you want to release all match
// claims associated with this reference. This parameter only works when
// the claim's status is being updated to 'inactive' - you can then set
// the parameter's value to true to release all match claims produced by
// this reference.
func (c *ReferencesPatchCall) ReleaseClaims(releaseClaims bool) *ReferencesPatchCall {
	c.urlParams_.Set("releaseClaims", fmt.Sprint(releaseClaims))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesPatchCall) Fields(s ...googleapi.Field) *ReferencesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReferencesPatchCall) Context(ctx context.Context) *ReferencesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferencesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferencesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reference)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "references/{referenceId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"referenceId": c.referenceId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.references.patch" call.
// Exactly one of *Reference or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Reference.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ReferencesPatchCall) Do(opts ...googleapi.CallOption) (*Reference, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Reference{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a reference. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.references.patch",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceId": {
	//       "description": "The referenceId parameter specifies the YouTube reference ID of the reference being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "releaseClaims": {
	//       "default": "false",
	//       "description": "The releaseClaims parameter indicates that you want to release all match claims associated with this reference. This parameter only works when the claim's status is being updated to 'inactive' - you can then set the parameter's value to true to release all match claims produced by this reference.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "references/{referenceId}",
	//   "request": {
	//     "$ref": "Reference"
	//   },
	//   "response": {
	//     "$ref": "Reference"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.references.update":

type ReferencesUpdateCall struct {
	s           *Service
	referenceId string
	reference   *Reference
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Update: Updates a reference.
func (r *ReferencesService) Update(referenceId string, reference *Reference) *ReferencesUpdateCall {
	c := &ReferencesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.referenceId = referenceId
	c.reference = reference
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ReferencesUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ReferencesUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// ReleaseClaims sets the optional parameter "releaseClaims": The
// releaseClaims parameter indicates that you want to release all match
// claims associated with this reference. This parameter only works when
// the claim's status is being updated to 'inactive' - you can then set
// the parameter's value to true to release all match claims produced by
// this reference.
func (c *ReferencesUpdateCall) ReleaseClaims(releaseClaims bool) *ReferencesUpdateCall {
	c.urlParams_.Set("releaseClaims", fmt.Sprint(releaseClaims))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesUpdateCall) Fields(s ...googleapi.Field) *ReferencesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReferencesUpdateCall) Context(ctx context.Context) *ReferencesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReferencesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReferencesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reference)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "references/{referenceId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"referenceId": c.referenceId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.references.update" call.
// Exactly one of *Reference or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Reference.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ReferencesUpdateCall) Do(opts ...googleapi.CallOption) (*Reference, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Reference{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a reference.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.references.update",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceId": {
	//       "description": "The referenceId parameter specifies the YouTube reference ID of the reference being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "releaseClaims": {
	//       "default": "false",
	//       "description": "The releaseClaims parameter indicates that you want to release all match claims associated with this reference. This parameter only works when the claim's status is being updated to 'inactive' - you can then set the parameter's value to true to release all match claims produced by this reference.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "references/{referenceId}",
	//   "request": {
	//     "$ref": "Reference"
	//   },
	//   "response": {
	//     "$ref": "Reference"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.validator.validate":

type ValidatorValidateCall struct {
	s               *Service
	validaterequest *ValidateRequest
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Validate: Validate a metadata file.
func (r *ValidatorService) Validate(validaterequest *ValidateRequest) *ValidatorValidateCall {
	c := &ValidatorValidateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.validaterequest = validaterequest
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *ValidatorValidateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *ValidatorValidateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ValidatorValidateCall) Fields(s ...googleapi.Field) *ValidatorValidateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ValidatorValidateCall) Context(ctx context.Context) *ValidatorValidateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ValidatorValidateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ValidatorValidateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.validaterequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "validator")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.validator.validate" call.
// Exactly one of *ValidateResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ValidateResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ValidatorValidateCall) Do(opts ...googleapi.CallOption) (*ValidateResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ValidateResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Validate a metadata file.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.validator.validate",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "validator",
	//   "request": {
	//     "$ref": "ValidateRequest"
	//   },
	//   "response": {
	//     "$ref": "ValidateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.videoAdvertisingOptions.get":

type VideoAdvertisingOptionsGetCall struct {
	s            *Service
	videoId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves advertising settings for the specified video.
func (r *VideoAdvertisingOptionsService) Get(videoId string) *VideoAdvertisingOptionsGetCall {
	c := &VideoAdvertisingOptionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.videoId = videoId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *VideoAdvertisingOptionsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideoAdvertisingOptionsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VideoAdvertisingOptionsGetCall) Fields(s ...googleapi.Field) *VideoAdvertisingOptionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *VideoAdvertisingOptionsGetCall) IfNoneMatch(entityTag string) *VideoAdvertisingOptionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *VideoAdvertisingOptionsGetCall) Context(ctx context.Context) *VideoAdvertisingOptionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *VideoAdvertisingOptionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *VideoAdvertisingOptionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "videoAdvertisingOptions/{videoId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"videoId": c.videoId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.videoAdvertisingOptions.get" call.
// Exactly one of *VideoAdvertisingOption or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *VideoAdvertisingOption.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *VideoAdvertisingOptionsGetCall) Do(opts ...googleapi.CallOption) (*VideoAdvertisingOption, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &VideoAdvertisingOption{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves advertising settings for the specified video.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.videoAdvertisingOptions.get",
	//   "parameterOrder": [
	//     "videoId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "The videoId parameter specifies the YouTube video ID of the video for which you are retrieving advertising settings.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videoAdvertisingOptions/{videoId}",
	//   "response": {
	//     "$ref": "VideoAdvertisingOption"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.videoAdvertisingOptions.getEnabledAds":

type VideoAdvertisingOptionsGetEnabledAdsCall struct {
	s            *Service
	videoId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetEnabledAds: Retrieves details about the types of allowed ads for a
// specified partner- or user-uploaded video.
func (r *VideoAdvertisingOptionsService) GetEnabledAds(videoId string) *VideoAdvertisingOptionsGetEnabledAdsCall {
	c := &VideoAdvertisingOptionsGetEnabledAdsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.videoId = videoId
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *VideoAdvertisingOptionsGetEnabledAdsCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideoAdvertisingOptionsGetEnabledAdsCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VideoAdvertisingOptionsGetEnabledAdsCall) Fields(s ...googleapi.Field) *VideoAdvertisingOptionsGetEnabledAdsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *VideoAdvertisingOptionsGetEnabledAdsCall) IfNoneMatch(entityTag string) *VideoAdvertisingOptionsGetEnabledAdsCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *VideoAdvertisingOptionsGetEnabledAdsCall) Context(ctx context.Context) *VideoAdvertisingOptionsGetEnabledAdsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *VideoAdvertisingOptionsGetEnabledAdsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *VideoAdvertisingOptionsGetEnabledAdsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "videoAdvertisingOptions/{videoId}/getEnabledAds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"videoId": c.videoId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.videoAdvertisingOptions.getEnabledAds" call.
// Exactly one of *VideoAdvertisingOptionGetEnabledAdsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *VideoAdvertisingOptionGetEnabledAdsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *VideoAdvertisingOptionsGetEnabledAdsCall) Do(opts ...googleapi.CallOption) (*VideoAdvertisingOptionGetEnabledAdsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &VideoAdvertisingOptionGetEnabledAdsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves details about the types of allowed ads for a specified partner- or user-uploaded video.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.videoAdvertisingOptions.getEnabledAds",
	//   "parameterOrder": [
	//     "videoId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "The videoId parameter specifies the YouTube video ID of the video for which you are retrieving advertising options.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videoAdvertisingOptions/{videoId}/getEnabledAds",
	//   "response": {
	//     "$ref": "VideoAdvertisingOptionGetEnabledAdsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.videoAdvertisingOptions.patch":

type VideoAdvertisingOptionsPatchCall struct {
	s                      *Service
	videoId                string
	videoadvertisingoption *VideoAdvertisingOption
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Patch: Updates the advertising settings for the specified video. This
// method supports patch semantics.
func (r *VideoAdvertisingOptionsService) Patch(videoId string, videoadvertisingoption *VideoAdvertisingOption) *VideoAdvertisingOptionsPatchCall {
	c := &VideoAdvertisingOptionsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.videoId = videoId
	c.videoadvertisingoption = videoadvertisingoption
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *VideoAdvertisingOptionsPatchCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideoAdvertisingOptionsPatchCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VideoAdvertisingOptionsPatchCall) Fields(s ...googleapi.Field) *VideoAdvertisingOptionsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *VideoAdvertisingOptionsPatchCall) Context(ctx context.Context) *VideoAdvertisingOptionsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *VideoAdvertisingOptionsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *VideoAdvertisingOptionsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.videoadvertisingoption)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "videoAdvertisingOptions/{videoId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"videoId": c.videoId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.videoAdvertisingOptions.patch" call.
// Exactly one of *VideoAdvertisingOption or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *VideoAdvertisingOption.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *VideoAdvertisingOptionsPatchCall) Do(opts ...googleapi.CallOption) (*VideoAdvertisingOption, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &VideoAdvertisingOption{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the advertising settings for the specified video. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "youtubePartner.videoAdvertisingOptions.patch",
	//   "parameterOrder": [
	//     "videoId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "The videoId parameter specifies the YouTube video ID of the video for which you are updating advertising settings.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videoAdvertisingOptions/{videoId}",
	//   "request": {
	//     "$ref": "VideoAdvertisingOption"
	//   },
	//   "response": {
	//     "$ref": "VideoAdvertisingOption"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.videoAdvertisingOptions.update":

type VideoAdvertisingOptionsUpdateCall struct {
	s                      *Service
	videoId                string
	videoadvertisingoption *VideoAdvertisingOption
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Update: Updates the advertising settings for the specified video.
func (r *VideoAdvertisingOptionsService) Update(videoId string, videoadvertisingoption *VideoAdvertisingOption) *VideoAdvertisingOptionsUpdateCall {
	c := &VideoAdvertisingOptionsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.videoId = videoId
	c.videoadvertisingoption = videoadvertisingoption
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *VideoAdvertisingOptionsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *VideoAdvertisingOptionsUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VideoAdvertisingOptionsUpdateCall) Fields(s ...googleapi.Field) *VideoAdvertisingOptionsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *VideoAdvertisingOptionsUpdateCall) Context(ctx context.Context) *VideoAdvertisingOptionsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *VideoAdvertisingOptionsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *VideoAdvertisingOptionsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.videoadvertisingoption)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "videoAdvertisingOptions/{videoId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"videoId": c.videoId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.videoAdvertisingOptions.update" call.
// Exactly one of *VideoAdvertisingOption or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *VideoAdvertisingOption.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *VideoAdvertisingOptionsUpdateCall) Do(opts ...googleapi.CallOption) (*VideoAdvertisingOption, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &VideoAdvertisingOption{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the advertising settings for the specified video.",
	//   "httpMethod": "PUT",
	//   "id": "youtubePartner.videoAdvertisingOptions.update",
	//   "parameterOrder": [
	//     "videoId"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "videoId": {
	//       "description": "The videoId parameter specifies the YouTube video ID of the video for which you are updating advertising settings.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "videoAdvertisingOptions/{videoId}",
	//   "request": {
	//     "$ref": "VideoAdvertisingOption"
	//   },
	//   "response": {
	//     "$ref": "VideoAdvertisingOption"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.whitelists.delete":

type WhitelistsDeleteCall struct {
	s          *Service
	id         string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Removes a whitelisted channel for a content owner.
func (r *WhitelistsService) Delete(id string) *WhitelistsDeleteCall {
	c := &WhitelistsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *WhitelistsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *WhitelistsDeleteCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WhitelistsDeleteCall) Fields(s ...googleapi.Field) *WhitelistsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *WhitelistsDeleteCall) Context(ctx context.Context) *WhitelistsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *WhitelistsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *WhitelistsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "whitelists/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.whitelists.delete" call.
func (c *WhitelistsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Removes a whitelisted channel for a content owner.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubePartner.whitelists.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube channel ID of the channel being removed from whitelist.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "whitelists/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.whitelists.get":

type WhitelistsGetCall struct {
	s            *Service
	id           string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves a specific whitelisted channel by ID.
func (r *WhitelistsService) Get(id string) *WhitelistsGetCall {
	c := &WhitelistsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *WhitelistsGetCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *WhitelistsGetCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WhitelistsGetCall) Fields(s ...googleapi.Field) *WhitelistsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *WhitelistsGetCall) IfNoneMatch(entityTag string) *WhitelistsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *WhitelistsGetCall) Context(ctx context.Context) *WhitelistsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *WhitelistsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *WhitelistsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "whitelists/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.whitelists.get" call.
// Exactly one of *Whitelist or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Whitelist.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *WhitelistsGetCall) Do(opts ...googleapi.CallOption) (*Whitelist, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Whitelist{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a specific whitelisted channel by ID.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.whitelists.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube channel ID of the whitelisted channel being retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "whitelists/{id}",
	//   "response": {
	//     "$ref": "Whitelist"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.whitelists.insert":

type WhitelistsInsertCall struct {
	s          *Service
	whitelist  *Whitelist
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Whitelist a YouTube channel for your content owner.
// Whitelisted channels are channels that are not owned or managed by
// you, but you would like to whitelist so that no claims from your
// assets are placed on videos uploaded to these channels.
func (r *WhitelistsService) Insert(whitelist *Whitelist) *WhitelistsInsertCall {
	c := &WhitelistsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.whitelist = whitelist
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *WhitelistsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *WhitelistsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WhitelistsInsertCall) Fields(s ...googleapi.Field) *WhitelistsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *WhitelistsInsertCall) Context(ctx context.Context) *WhitelistsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *WhitelistsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *WhitelistsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.whitelist)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "whitelists")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.whitelists.insert" call.
// Exactly one of *Whitelist or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Whitelist.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *WhitelistsInsertCall) Do(opts ...googleapi.CallOption) (*Whitelist, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Whitelist{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Whitelist a YouTube channel for your content owner. Whitelisted channels are channels that are not owned or managed by you, but you would like to whitelist so that no claims from your assets are placed on videos uploaded to these channels.",
	//   "httpMethod": "POST",
	//   "id": "youtubePartner.whitelists.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "whitelists",
	//   "request": {
	//     "$ref": "Whitelist"
	//   },
	//   "response": {
	//     "$ref": "Whitelist"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubePartner.whitelists.list":

type WhitelistsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of whitelisted channels for a content owner.
func (r *WhitelistsService) List() *WhitelistsListCall {
	c := &WhitelistsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of YouTube channel IDs that identify the
// whitelisted channels you want to retrieve.
func (c *WhitelistsListCall) Id(id string) *WhitelistsListCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": The onBehalfOfContentOwner parameter
// identifies the content owner that the user is acting on behalf of.
// This parameter supports users whose accounts are associated with
// multiple content owners.
func (c *WhitelistsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *WhitelistsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The pageToken
// parameter specifies a token that identifies a particular page of
// results to return. Set this parameter to the value of the
// nextPageToken value from the previous API response to retrieve the
// next page of results.
func (c *WhitelistsListCall) PageToken(pageToken string) *WhitelistsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WhitelistsListCall) Fields(s ...googleapi.Field) *WhitelistsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *WhitelistsListCall) IfNoneMatch(entityTag string) *WhitelistsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *WhitelistsListCall) Context(ctx context.Context) *WhitelistsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *WhitelistsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *WhitelistsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "whitelists")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubePartner.whitelists.list" call.
// Exactly one of *WhitelistListResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WhitelistListResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *WhitelistsListCall) Do(opts ...googleapi.CallOption) (*WhitelistListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WhitelistListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of whitelisted channels for a content owner.",
	//   "httpMethod": "GET",
	//   "id": "youtubePartner.whitelists.list",
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of YouTube channel IDs that identify the whitelisted channels you want to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This parameter supports users whose accounts are associated with multiple content owners.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pageToken parameter specifies a token that identifies a particular page of results to return. Set this parameter to the value of the nextPageToken value from the previous API response to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "whitelists",
	//   "response": {
	//     "$ref": "WhitelistListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *WhitelistsListCall) Pages(ctx context.Context, f func(*WhitelistListResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
