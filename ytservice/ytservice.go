/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

////////////////////////////////////////////////////////////////////////////////

// Service object which contains the main context for calling the YouTube API
type Service struct {
	API            *youtube.Service
	ServiceAccount bool
	token          *oauth2.Token
}

////////////////////////////////////////////////////////////////////////////////

// Constants
const (
	YouTubeMaxPagingResults = 50
)

////////////////////////////////////////////////////////////////////////////////

// Returns a service object given service account details
func NewYouTubeServiceFromServiceAccountJSON(filename string,params *Params,debug bool) (*Service, error) {
	if params.IsValidContentOwner() == false {
		return nil, ErrorMissingContentOwner
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount,err)
	}
	sa_config, err := google.JWTConfigFromJSON(bytes,youtube.YoutubeScope,youtube.YoutubepartnerScope)
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount,err)
	}
	ctx := getContext(debug)
	service, err := youtube.New(sa_config.Client(ctx))
	if err != nil {
		return nil, NewError(ErrorInvalidServiceAccount,err)
	}
	this := new(Service)
	this.API = service
	this.ServiceAccount = true
	return this, nil
}

// Returns a service object given client secrets details
func NewYouTubeServiceFromClientSecretsJSON(clientsecrets string, tokencache string,params *Params,debug bool) (*Service, error) {
	bytes, err := ioutil.ReadFile(clientsecrets)
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets,err)
	}
	config, err := google.ConfigFromJSON(bytes, youtube.YoutubeScope)
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets,err)
	}
	ctx := getContext(debug)

	// Attempt to get token from cache
	token, err := tokenFromFile(tokencache)
	if err != nil {
		token, err = tokenFromWeb(config, ctx)
		saveToken(tokencache, token)
	}
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets,err)
	}

	// create client
	service, err := youtube.New(config.Client(ctx, token))
	if err != nil {
		return nil, NewError(ErrorInvalidClientSecrets,err)
	}

	this := new(Service)
	this.API = service
	this.ServiceAccount = false
	this.token = token
	return this, nil
}

////////////////////////////////////////////////////////////////////////////////


