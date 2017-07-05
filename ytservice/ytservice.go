package ytservice

/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"io/ioutil"

	"github.com/djthorpe/ytapi/youtubepartner/v1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/youtubeanalytics/v1"
	"google.golang.org/api/youtubereporting/v1"
)

////////////////////////////////////////////////////////////////////////////////

// Service object which contains the main context for calling the YouTube API
type Service struct {
	// Data API
	API *youtube.Service
	// Partner API
	PAPI *youtubepartner.Service
	// Analytics API
	AAPI *youtubeanalytics.Service
	// Bulk Reporting API
	RAPI *youtubereporting.Service
	// Whether this is a service account
	ServiceAccount bool
	// The email address of the service account
	ServiceAccountEmail string
	// the private OAuth token
	token *oauth2.Token
	// call options
	callopts []googleapi.CallOption
}

////////////////////////////////////////////////////////////////////////////////

// NewYouTubeServiceFromServiceAccountJSON returns a service object given
// service account filename, an array of scopes and a debug flag
func NewYouTubeServiceFromServiceAccountJSON(filename string, scope_flags []string, debug bool) (*Service, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount, err)
	}

	scopes, err := scopesForFlags(scope_flags)
	if err != nil {
		return nil, err
	}
	saConfig, err := google.JWTConfigFromJSON(bytes, scopes...)
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount, err)
	}
	ctx := getContext(debug)
	service, err := youtube.New(saConfig.Client(ctx))
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount, err)
	}
	partnerservice, err := youtubepartner.New(saConfig.Client(ctx))
	if err != nil {
		return nil, err
	}
	analyticsservice, err := youtubeanalytics.New(saConfig.Client(ctx))
	if err != nil {
		return nil, err
	}
	reportingservice, err := youtubereporting.New(saConfig.Client(ctx))
	if err != nil {
		return nil, err
	}

	// create the service object
	this := new(Service)
	this.API = service
	this.PAPI = partnerservice
	this.AAPI = analyticsservice
	this.RAPI = reportingservice
	this.ServiceAccount = true
	this.ServiceAccountEmail = saConfig.Email
	this.callopts = make([]googleapi.CallOption, 0)

	// Success
	return this, nil
}

// NewYouTubeServiceFromClientSecretsJSON returns a service object given client secrets details
func NewYouTubeServiceFromClientSecretsJSON(clientsecrets string, tokencache string, debug bool) (*Service, error) {
	bytes, err := ioutil.ReadFile(clientsecrets)
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}
	config, err := google.ConfigFromJSON(bytes, youtube.YoutubeForceSslScope, youtubeanalytics.YtAnalyticsReadonlyScope)
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}
	ctx := getContext(debug)

	// Attempt to get token from cache
	token, err := tokenFromFile(tokencache)
	if err != nil {
		token, err = tokenFromWeb(ctx, config)
		if err != nil {
			return nil, err
		}
		saveToken(tokencache, token)
	}
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}

	// create data client
	service, err := youtube.New(config.Client(ctx, token))
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}
	// create analytics client
	analyticsservice, err := youtubeanalytics.New(config.Client(ctx, token))
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}
	// create reporting client
	reportingservice, err := youtubereporting.New(config.Client(ctx, token))
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}

	// create the service object
	this := new(Service)
	this.API = service
	this.AAPI = analyticsservice
	this.RAPI = reportingservice
	this.ServiceAccount = false
	this.token = token
	this.callopts = make([]googleapi.CallOption, 0)
	return this, nil
}

// SetQuotaUser sets the quota user parameter for all API requests
func (this *Service) SetQuotaUser(value string) {
	this.callopts = append(this.callopts, googleapi.QuotaUser(value))
}

// SetQuotaAddress sets the userid parameter for all API requests
func (this *Service) SetQuotaAddress(value string) {
	this.callopts = append(this.callopts, googleapi.UserIP(value))
}

// SetTraceToken sets the tracetoken parameter for all API requests
func (this *Service) SetTraceToken(value string) {
	this.callopts = append(this.callopts, googleapi.Trace(value))
}

// CallOptions returns the array of call options
func (this *Service) CallOptions() []googleapi.CallOption {
	return this.callopts
}
