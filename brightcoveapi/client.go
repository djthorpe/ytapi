package brightcoveapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

/////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	// URL retrieve tokens
	OAUTH_TOKEN_URL = "https://oauth.brightcove.com/v3/access_token"

	// Default size of buffer
	BUFFER_SIZE = 1024
)

/////////////////////////////////////////////////////////////////////
// STRUCTS

// Credentials define the Brightcove credentials and defaults used
// for accessing the API's
type Credentials struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccountId    string `json:"account_id"`
}

// Token is the returned access token from tbe Brightcove OAuth
type Token struct {
	timestamp   time.Time
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresSecs uint   `json:"expires_in"`
}

// Client defines a Brightcove client
type Client struct {
	// Private members
	http        *http.Client
	credentials *Credentials
	token       *Token
	options     url.Values
	debug       bool

	// Public members
	CMS *cms
}

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	current *http.Request
}

////////////////////////////////////////////////////////////////////

func NewClientWithCredentials(credentials *Credentials, options ...ClientOption) (*Client, error) {
	// Sanity check
	if credentials == nil {
		return nil, ErrBadParameter
	}
	// Create a client and apply options
	this := &Client{
		http:        &http.Client{},
		credentials: credentials,
		token:       nil,
	}
	for _, option := range options {
		if err := option.apply(this); err != nil {
			return nil, err
		}
	}
	// Create subsystems
	this.CMS = this.NewCMS()
	// Return success
	return this, nil
}

func NewClient(path string, options ...ClientOption) (*Client, error) {
	if credentials, err := ReadCredentials(path); err != nil {
		return nil, err
	} else if credentials.AccountId == "" {
		return nil, ErrAccountId
	} else {
		return NewClientWithCredentials(credentials, options...)
	}
}

// NewRequest returns a http request with access token embedded
func (this *Client) NewRequest(method, url string) (*http.Request, error) {
	if _, err := this.AccessToken(); err != nil {
		return nil, err
	} else if req, err := http.NewRequest(method, url, nil); err != nil {
		return nil, err
	} else {
		req.Header.Add("Authorization", "Bearer "+this.token.AccessToken)
		return req, nil
	}
}

// Copy will download data from a URL and write to a stream
func (this *Client) Copy(method, url string, w io.Writer) error {
	if req, err := this.NewRequest(method, url); err != nil {
		return err
	} else if response, err := this.Do(req); err != nil {
		return err
	} else {
		defer response.Body.Close()
		if _, err = io.Copy(w, response.Body); err != nil {
			return err
		} else {
			return nil
		}
	}
}

// AccountId returns the brightcove AccountId field
func (this *Client) AccountId() string {
	return this.credentials.AccountId
}

// AccessToken returns the access token or requests a new one if it has expired
func (this *Client) AccessToken() (*Token, error) {
	if this.token == nil || this.token.Expired() {
		if token, err := this.getOAuthToken(this.credentials); err != nil {
			return nil, err
		} else {
			this.token = token
		}
	}
	return this.token, nil
}

// Do performs a http request and returns the raw http response
func (this *Client) Do(req *http.Request) (*http.Response, error) {
	if this.debug {
		log.Println("=>", req.URL.String())
	}
	if response, err := this.http.Do(req); err != nil {
		return nil, err
	} else {
		if this.debug {
			dump, err := httputil.DumpResponse(response, true)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%q", dump)
		}
		if response.StatusCode != http.StatusOK {
			if this.debug {
				log.Println("<=", response.Status)
			}
			return nil, fmt.Errorf("%v [%v]", response.Status, response.StatusCode)
		} else {
			return response, nil
		}
	}
}

// SetOptions resets the set of http query parameters and in turn
// sets them from handler functions
func (this *Client) SetOptions(options []ClientOption) error {
	this.options = url.Values{}
	for _, option := range options {
		if err := option.apply(this); err != nil {
			return err
		}
	}
	return nil
}

/////////////////////////////////////////////////////////////////////

// ReadCredentials returns brightcove credentials, or if they don't
// exist yet will also write an empty credentials file that can be
// filled out with a text editor
func ReadCredentials(path string) (*Credentials, error) {
	var credentials Credentials
	if abspath, err := AbsolutePath(path); err != nil {
		return nil, err
	} else if stat, err := os.Stat(abspath); os.IsNotExist(err) {
		// Write empty credentials file
		if fw, err := os.Create(abspath); err != nil {
			return nil, err
		} else {
			defer fw.Close()
			encoder := json.NewEncoder(fw)
			if err := encoder.Encode(&credentials); err != nil {
				return nil, err
			}
			return &credentials, nil
		}
	} else if err != nil {
		return nil, err
	} else if stat.Mode().IsRegular() == false {
		return nil, fmt.Errorf("Invalid credentials file")
	} else if fh, err := os.Open(abspath); err != nil {
		return nil, err
	} else {
		defer fh.Close()
		decoder := json.NewDecoder(fh)
		if err := decoder.Decode(&credentials); err != nil {
			return nil, err
		} else {
			return &credentials, nil
		}
	}
}

// getOAuthRequest returns a http.Request object which can be used to
// request an access token from brightcove oauth api
func getOAuthRequest(credentials *Credentials) (*http.Request, error) {
	data := url.Values{
		"grant_type": {"client_credentials"},
	}
	body := strings.NewReader(data.Encode())
	if req, err := http.NewRequest("POST", OAUTH_TOKEN_URL, body); err != nil {
		return nil, err
	} else {
		req.SetBasicAuth(credentials.ClientId, credentials.ClientSecret)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		return req, nil
	}
}

// getOAuthToken returns a Token object which can be used to
// make subsequent requests to the brightcove api. The expiry
// time for the access token can be determined using the
// Expiry() method
func (this *Client) getOAuthToken(credentials *Credentials) (*Token, error) {
	var token Token
	if req, err := getOAuthRequest(credentials); err != nil {
		return nil, err
	} else if resp, err := this.Do(req); err != nil {
		return nil, err
	} else {
		decoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()
		if err := decoder.Decode(&token); err != nil {
			return nil, err
		}
		token.timestamp = time.Now()
		return &token, nil
	}
}

/////////////////////////////////////////////////////////////////////
// Stringify

func (this *Credentials) String() string {
	return fmt.Sprintf("<Credentials>{ client_id='%v' }", this.ClientId)
}

func (this *Token) String() string {
	return fmt.Sprintf("<Token>{ access_token='%v' token_type='%v' expired=%v }", this.AccessToken, this.TokenType, this.Expired())
}

func (this *Client) String() string {
	return fmt.Sprintf("<Client>{ credentials=%v token=%v client_timeout=%v }", this.credentials, this.token, this.http.Timeout)
}

/////////////////////////////////////////////////////////////////////
// Credentials & Tokens

func (this *Token) Expiry() time.Time {
	if this.timestamp.IsZero() || this.ExpiresSecs == 0 {
		return time.Time{}
	} else {
		return this.timestamp.Add(time.Duration(time.Second * time.Duration(this.ExpiresSecs)))
	}
}

func (this *Token) Expired() bool {
	return this.Expiry().IsZero() || this.Expiry().Before(time.Now())
}

/////////////////////////////////////////////////////////////////////
// Transport

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	fmt.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}
