package ytservice

import (
	"os"
	"io/ioutil"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// Defaults object stores all the defaults for various fields
type Defaults struct {
	Debug        bool     `json:"debug"`
	ContentOwner *string  `json:"content_owner,omitempty"`
	Channel      *string  `json:"channel,omitempty"`
	MaxResults   uint64   `json:"max_results"`
}

// YTService object which contains the main context for calling the YouTube API
type YTService struct {
	API            *youtube.Service
	ServiceAccount bool
	token          *oauth2.Token
}

// Constants
const (
	YouTubeMaxPagingResults = 50
)

// Returns a service object given service account details
func NewYouTubeServiceFromServiceAccountJSON(filename string,defaults *Defaults) (*YTService, error) {
	if len(*defaults.ContentOwner) == 0 {
		return nil, ErrorMissingContentOwner
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, ErrorInvalidServiceAccount
	}
	sa_config, err := google.JWTConfigFromJSON(bytes, youtube.YoutubeScope, youtube.YoutubepartnerScope)
	if err != nil {
		return nil, ErrorInvalidServiceAccount
	}
	ctx := getContext(defaults.Debug)
	service, err := youtube.New(sa_config.Client(ctx))
	if err != nil {
		return nil, ErrorInvalidServiceAccount
	}
	this := new(YTService)
	this.API = service
	this.ServiceAccount = true
	return this, nil
}


// Returns a service object given client secrets details
func NewYouTubeServiceFromClientSecretsJSON(clientsecrets string, tokencache string,defaults *Defaults) (*YTService, error) {
	bytes, err := ioutil.ReadFile(clientsecrets)
	if err != nil {
		return nil, ErrorInvalidClientSecrets
	}
	config, err := google.ConfigFromJSON(bytes, youtube.YoutubeScope)
	if err != nil {
		return nil, ErrorInvalidClientSecrets
	}
	ctx := getContext(defaults.Debug)

	// Attempt to get token from cache
	token, err := tokenFromFile(tokencache)
	if err != nil {
		token, err = tokenFromWeb(config, ctx)
		saveToken(tokencache, token)
	}
	if err != nil {
		return nil, ErrorInvalidClientSecrets
	}

	// create client
	service, err := youtube.New(config.Client(ctx, token))
	if err != nil {
		return nil, ErrorInvalidClientSecrets
	}

	this := new(YTService)
	this.API = service
	this.ServiceAccount = false
	this.token = token
	return this, nil
}


// Returns a defaults object from a JSON file
func NewDefaultsFromJSON(defaults string) (*Defaults, error) {
	bytes, err := ioutil.ReadFile(defaults)
	if err != nil {
		return nil, ErrorInvalidDefaults
	}
	this := NewDefaults()
    err = json.Unmarshal(bytes,this)
	if err != nil {
		return nil, ErrorInvalidDefaults
	}
    return this,nil
}

// Returns a defaults object
func NewDefaults() (*Defaults) {
	this := new(Defaults)
	this.Debug = false
	this.MaxResults = 0
	this.ContentOwner = nil
	this.Channel = nil
	return this
}

// Save defaults object
func (this *Defaults) Save(filename string,perm os.FileMode) error {
	json, err := json.MarshalIndent(this,"","  ")
	if err != nil {
		return err
	}
    err = ioutil.WriteFile(filename,json,perm)
	if err != nil {
		return err
	}
	return nil
}


