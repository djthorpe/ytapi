// Package ytservice encapsulates all the YouTube API services
/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"runtime"
	"time"
)

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

////////////////////////////////////////////////////////////////////////////////

// Returns context
func getContext(debug bool) context.Context {
	ctx := context.Background()
	if debug {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: &logTransport{http.DefaultTransport},
		})
	}
	return ctx
}

// Returns token from cache
func tokenFromFile(filename string) (*oauth2.Token, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, ErrorCacheTokenRead
	}
	t := new(oauth2.Token)
	err = gob.NewDecoder(f).Decode(t)
	return t, err
}

// Saves token
func saveToken(filename string, token *oauth2.Token) error {
	f, err := os.Create(filename)
	if err != nil {
		return ErrorCacheTokenWrite
	}
	defer f.Close()
	gob.NewEncoder(f).Encode(token)
	return nil
}

// Creates a webserver for user interaction with Google
func tokenFromWeb(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	ch := make(chan string)
	randState := fmt.Sprintf("ts%d", time.Now().UnixNano())
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			http.Error(rw, "State doesn't match", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "<h1>Success</h1>Authorized - You can now close this window")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}
		http.Error(rw, "Denied - You can now close this window", 500)
		ch <- ""
	}))
	defer ts.Close()
	config.RedirectURL = ts.URL
	authURL := config.AuthCodeURL(randState)
	go openURL(authURL)
	code := <-ch
	if code == "" {
		return nil, ErrorDenied
	}
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, ErrorTokenExchange
	}
	return token, nil
}

// Attempt to open a URL using a browser
func openURL(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
	/*
		try := []string{"xdg-open", "google-chrome", "open"}
		for _, bin := range try {
			err := exec.Command(bin, url).Run()
			if err == nil {
				return
			}
		}
	*/
}
