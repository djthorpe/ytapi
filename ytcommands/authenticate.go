/*
Copyright David Thorpe 2015 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////

const (
	credentialsPathMode = 0700
	credentialsFileMode = 0644
)

////////////////////////////////////////////////////////////////////////////////
// File methods

func userDir() string {
	currentUser, _ := user.Current()
	return currentUser.HomeDir
}

func GetCredentialsPath(values *ytapi.Values) string {
	return filepath.Join(userDir(), values.GetString(&ytapi.FlagCredentials))
}

func GetOAuthTokenPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values), values.GetString(&ytapi.FlagAuthToken))
}

func GetServiceAccountPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values), values.GetString(&ytapi.FlagServiceAccount))
}

func GetClientSecretPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values), values.GetString(&ytapi.FlagClientSecret))
}

func GetDefaultsPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values), values.GetString(&ytapi.FlagDefaults))
}

func GetCredentialsFolder(values *ytapi.Values) (string, error) {
	// Obtain path for credentials
	credentialsPath := GetCredentialsPath(values)

	// Create if not already made
	if credentialsPathInfo, err := os.Stat(credentialsPath); err != nil || !credentialsPathInfo.IsDir() {
		// if path is missing, try and create the folder
		if err := os.Mkdir(credentialsPath, credentialsPathMode); err != nil {
			return "", err
		}
	}
	return credentialsPath, nil
}

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterAuthenticateCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "Authenticate",
			Description: "Authenticate against service account or channel",
			Setup:       AuthenticateSetup,
			Execute:     AuthenticateExecute,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Perform authentication

func AuthenticateSetup(values *ytapi.Values, table *ytservice.Table) error {

	// remove existing oauth token
	tokenPath := GetOAuthTokenPath(values)
	if _, err := os.Stat(tokenPath); os.IsNotExist(err) == false {
		err = os.Remove(tokenPath)
		if err != nil {
			return err
		}
	}

	// success
	return nil
}

func AuthenticateExecute(service *ytservice.Service, values *ytapi.Values, table *ytservice.Table) error {

	// Write defaults to file
	err := values.WriteDefaultsToFile(GetDefaultsPath(values), credentialsFileMode)
	if err != nil {
		return err
	}

	// success
	return nil
}
