package brightcoveapi

import (
	"net/url"
	"os"
	"path/filepath"
)

// Return the path to write to or create from src and path
func extractPath(src, path string) (string, error) {
	// Clean up the path argument
	if path == "" || path == "." {
		if wd, err := os.Getwd(); err != nil {
			return "", err
		} else {
			path = wd
		}
	}
	path = filepath.Clean(path)
	basepath := filepath.Dir(path)
	if stat, err := os.Stat(path); os.IsNotExist(err) == false && stat.IsDir() {
		if u, err := url.Parse(src); err != nil {
			return "", err
		} else {
			return filepath.Join(path, filepath.Base(u.Path)), nil
		}
	} else if os.IsNotExist(err) == false && stat.Mode().IsRegular() {
		return path, nil
	} else if os.IsNotExist(err) {
		if stat, err := os.Stat(basepath); os.IsNotExist(err) == false && stat.IsDir() {
			return path, nil
		}
	} else if err != nil {
		return "", err
	}
	return "", ErrBadParameter
}

// Download downloads a URL to a path, which can be either
// an existing folder or a file within an existing folder
func (this *Client) Download(method, src, path string) error {
	if path, err := extractPath(src, path); err != nil {
		return err
	} else if fh, err := os.Create(path); err != nil {
		return err
	} else {
		defer fh.Close()
		if this.Copy("GET", src, fh); err != nil {
			return err
		} else {
			return nil
		}
	}
}
