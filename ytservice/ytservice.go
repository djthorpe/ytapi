/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"io/ioutil"

	"github.com/djthorpe/ytapi/youtubepartner/v1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////

// Service object which contains the main context for calling the YouTube API
type Service struct {
	API                 *youtube.Service
	PAPI                *youtubepartner.Service
	ServiceAccount      bool
	ServiceAccountEmail string
	token               *oauth2.Token
}

////////////////////////////////////////////////////////////////////////////////

// Returns a service object given service account details
func NewYouTubeServiceFromServiceAccountJSON(filename string, debug bool) (*Service, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount, err)
	}
	sa_config, err := google.JWTConfigFromJSON(bytes, youtube.YoutubeForceSslScope, youtube.YoutubepartnerScope)
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount, err)
	}
	ctx := getContext(debug)
	service, err := youtube.New(sa_config.Client(ctx))
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount, err)
	}
	partnerservice, err := youtubepartner.New(sa_config.Client(ctx))
	this := new(Service)
	this.API = service
	this.PAPI = partnerservice
	this.ServiceAccount = true
	this.ServiceAccountEmail = sa_config.Email
	return this, nil
}

// Returns a service object given client secrets details
func NewYouTubeServiceFromClientSecretsJSON(clientsecrets string, tokencache string, debug bool) (*Service, error) {
	bytes, err := ioutil.ReadFile(clientsecrets)
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}
	config, err := google.ConfigFromJSON(bytes, youtube.YoutubeForceSslScope)
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}
	ctx := getContext(debug)

	// Attempt to get token from cache
	token, err := tokenFromFile(tokencache)
	if err != nil {
		token, err = tokenFromWeb(config, ctx)
		if err != nil {
			return nil, err
		}
		saveToken(tokencache, token)
	}
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}

	// create client
	service, err := youtube.New(config.Client(ctx, token))
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets, err)
	}

	this := new(Service)
	this.API = service
	this.ServiceAccount = false
	this.token = token
	return this, nil
}
