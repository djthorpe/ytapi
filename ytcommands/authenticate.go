/*
Copyright David Thorpe 2015 All Rights Reserved
Please see file LICENSE for information on distribution, etc
*/
package ytcommands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
    "github.com/djthorpe/ytapi/util"
)

////////////////////////////////////////////////////////////////////////////////
// File methods

// TODO: Make these relative to home directory, not absolute

func GetCredentialsPath(values *ytapi.Values) string {
	return filepath.Join(util.UserDir(),values.GetString(&ytapi.FlagCredentials))
}

func GetOAuthTokenPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values),values.GetString(&ytapi.FlagAuthToken))
}

func GetServiceAccountPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values),values.GetString(&ytapi.FlagServiceAccount))
}

func GetClientSecretPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values),values.GetString(&ytapi.FlagClientSecret))
}

func GetDefaultsPath(values *ytapi.Values) string {
	return filepath.Join(GetCredentialsPath(values),values.GetString(&ytapi.FlagDefaults))
}

func GetCredentialsFolder(values *ytapi.Values) (string, error) {
	credentialsPath := GetCredentialsPath(values)
	return credentialsPath, nil
}

////////////////////////////////////////////////////////////////////////////////
// Register search output format

func RegisterAuthenticateCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:        "Authenticate",
			Description: "Authenticate against service account or channel",
			Setup:       AuthenticateSetup,
			Execute:     AuthenticateExecute,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Perform authentication

func AuthenticateSetup(values *ytapi.Values, table *ytapi.Table) error {

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

func AuthenticateExecute(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	fmt.Println("TODO: IMPLEMENT values.WriteDefaultsToFile")
/*
	// Write defaults to file
	err := values.WriteDefaultsToFile(GetDefaultsPath(values), credentialsFileMode)
	if err != nil {
		return err
	}
*/
	// success
	return nil
}
