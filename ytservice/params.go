/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

////////////////////////////////////////////////////////////////////////////////

// Params object stores all the parameters used for making API requests
type Params struct {
	ContentOwner    *string `json:"contentowner,omitempty"`
	Channel         *string `json:"channel,omitempty"`
	Video           *string `json:"video,omitempty"`
	MaxResults      int64   `json:"maxresults,omitempty"`
	Query           *string `json:"q,omitempty"`
	BroadcastStatus *string `json:"broadcaststatus,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////

// Returns a new Params object
func NewParams() *Params {
	this := new(Params)
	this.MaxResults = 0
	this.ContentOwner = nil
	this.Channel = nil
	this.Video = nil
	this.BroadcastStatus = nil
	this.Query = nil
	return this
}

// Returns a params object from a JSON file
func NewParamsFromJSON(filename string) (*Params, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, NewError(ErrorInvalidDefaults, err)
	}
	this := NewParams()
	err = json.Unmarshal(bytes, this)
	if err != nil {
		return nil, NewError(ErrorInvalidDefaults, err)
	}
	return this, nil
}

////////////////////////////////////////////////////////////////////////////////

// Makes a copy of the object
func (this *Params) Copy() *Params {
	copy := NewParams()
	copy.MaxResults = this.MaxResults
	copy.ContentOwner = this.ContentOwner // TODO copy pointer?
	copy.Channel = this.Channel           // TODO copy pointer?
	copy.Video = this.Video               // TODO copy pointer?
	copy.Query = this.Query               // TODO copy pointer?
	copy.BroadcastStatus = this.BroadcastStatus
	return copy
}

// Save params object
func (this *Params) Save(filename string, perm os.FileMode) error {
	json, err := json.MarshalIndent(this, "", "  ")
	if err != nil {
		return NewError(ErrorInvalidDefaults, err)
	}
	err = ioutil.WriteFile(filename, json, perm)
	if err != nil {
		return NewError(ErrorInvalidDefaults, err)
	}
	return nil
}

// Return boolean value which indicates if a content owner parameter is missing
func (this *Params) IsEmptyContentOwner() bool {
	if this.ContentOwner == nil {
		return true
	}
	if len(*this.ContentOwner) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid content owner setting
func (this *Params) IsValidContentOwner() bool {
	if this.IsEmptyContentOwner() {
		return false
	}
	matched, _ := regexp.MatchString("^([a-zA-Z0-9]{22})$", *this.ContentOwner)
	return matched
}

// Return boolean value which indicates an empty channel
func (this *Params) IsEmptyChannel() bool {
	if this.Channel == nil {
		return true
	}
	if len(*this.Channel) == 0 {
		return true
	}
	return false
}

// Return boolean value whichindicates a valid content owner setting
func (this *Params) IsValidChannel() bool {
	if this.IsEmptyChannel() {
		return false
	}
	matched, _ := regexp.MatchString("^UC([a-zA-Z0-9\\-]{22})$", *this.Channel)
	return matched
}

// Return boolean value which indicates an empty query
func (this *Params) IsEmptyQuery() bool {
	if this.Query == nil {
		return true
	}
	if len(*this.Query) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates an empty video parameter
func (this *Params) IsEmptyVideo() bool {
	if this.Video == nil {
		return true
	}
	if len(*this.Video) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid video parameter
func (this *Params) IsValidVideo() bool {
	if this.IsEmptyVideo() {
		return false
	}
	matched, _ := regexp.MatchString("^([a-zA-Z0-9\\-]{11})$", *this.Video)
	return matched
}

// Return boolean value which indicates an empty video parameter
func (this *Params) IsEmptyBroadcastStatus() bool {
	if this.BroadcastStatus == nil {
		return true
	}
	if len(*this.BroadcastStatus) == 0 {
		return true
	}
	return false
}

// Return boolean value which indicates a valid video parameter
func (this *Params) IsValidBroadcastStatus() bool {
	if this.IsEmptyBroadcastStatus() {
		return false
	}
	matched, _ := regexp.MatchString("^(all|upcoming|completed|active)$", *this.BroadcastStatus)
	return matched
}
