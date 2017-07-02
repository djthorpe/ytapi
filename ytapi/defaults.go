package ytapi

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

////////////////////////////////////////////////////////////////////////////////

// Defaults structure defines values read from store
type Defaults struct {
	ContentOwner string `json:"contentowner,omitempty"`
	Channel      string `json:"channel,omitempty"`
	QuotaUser    string `json:"quotauser,omitempty"`
	QuotaAddress string `json:"quotaip,omitempty"`
	TraceToken   string `json:"tracetoken,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// ReadDefaults sets values from the defaults file
func (this *FlagSet) ReadDefaults() error {
	var err error
	// if a file exists, then read it
	if _, err := os.Stat(this.Defaults); os.IsNotExist(err) {
		// file doesn't exist, so just return
		return err
	}
	// read in the file
	bytes, err := ioutil.ReadFile(this.Defaults)
	if err != nil {
		return err
	}
	defaults := &Defaults{}
	err = json.Unmarshal(bytes, defaults)
	if err != nil {
		return err
	}
	// ContentOwner
	if err == nil && defaults.ContentOwner != "" {
		err = this.Values.SetDefault(&FlagContentOwner, string(defaults.ContentOwner))
	}
	// Channel
	if err == nil && defaults.Channel != "" {
		err = this.Values.SetDefault(&FlagChannel, string(defaults.Channel))
	}
	// QuotaUser
	if err == nil && defaults.QuotaUser != "" {
		err = this.Values.SetDefault(&FlagQuotaUser, string(defaults.QuotaUser))
	}
	// QuotaAddress
	if err == nil && defaults.QuotaAddress != "" {
		err = this.Values.SetDefault(&FlagQuotaAddress, string(defaults.QuotaAddress))
	}
	// TraceToken
	if err == nil && defaults.TraceToken != "" {
		err = this.Values.SetDefault(&FlagTraceToken, string(defaults.TraceToken))
	}
	if err != nil {
		return err
	}
	// success
	return nil
}

// WriteDefaults writes defaults file
func (this *FlagSet) WriteDefaults() error {
	defaults := &Defaults{}
	if this.Values.IsSet(&FlagContentOwner) {
		defaults.ContentOwner = this.Values.GetString(&FlagContentOwner)
	}
	if this.Values.IsSet(&FlagChannel) && this.Values.IsSet(&FlagContentOwner) {
		defaults.Channel = this.Values.GetString(&FlagChannel)
	}
	if this.Values.IsSet(&FlagQuotaUser) {
		defaults.QuotaUser = this.Values.GetString(&FlagQuotaUser)
	}
	if this.Values.IsSet(&FlagQuotaAddress) {
		defaults.QuotaAddress = this.Values.GetString(&FlagQuotaAddress)
	}
	if this.Values.IsSet(&FlagTraceToken) {
		defaults.TraceToken = this.Values.GetString(&FlagTraceToken)
	}
	json, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(this.Defaults, json, credentialsFileMode)
	if err != nil {
		return err
	}
	return nil
}
