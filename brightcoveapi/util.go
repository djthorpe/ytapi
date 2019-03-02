package brightcoveapi

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

var (
	ErrNotImplemented = errors.New("Not implemented")
	ErrBadParameter   = errors.New("Bad parameter")
	ErrAccountId      = errors.New("Invalid account id ")
)

// HomeDirPath returns home directory
func HomeDirPath() (string, error) {
	if user, err := user.Current(); err != nil {
		return "", err
	} else if stat, err := os.Stat(user.HomeDir); os.IsNotExist(err) {
		return "", err
	} else if stat.IsDir() == false {
		return "", fmt.Errorf("Invalid home directory")
	} else {
		return filepath.Clean(user.HomeDir), nil
	}
}

// AbsolutePath resolves path to include home directory if the
// path is not already absolute
func AbsolutePath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	} else if homedir, err := HomeDirPath(); err != nil {
		return "", err
	} else {
		return filepath.Join(homedir, path), nil
	}
}

// URLJoin returns a URL object which is composed of an absolute base
// URL and a path, which is appended
func URLJoin(baseurl *url.URL, path string, query url.Values) (*url.URL, error) {
	if relurl, err := url.Parse(path); err != nil {
		return nil, err
	} else if url := baseurl.ResolveReference(relurl); url == nil {
		return nil, errors.New("URL parse error")
	} else {
		url.RawQuery = query.Encode()
		return url, nil
	}
}
