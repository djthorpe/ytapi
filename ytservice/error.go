/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"fmt"
	"errors"
)

// Error type
type Error struct {
	base   error
	detail string
}

// Enumeration of Errors
var (
	ErrorMissingContentOwner   = errors.New("Missing content owner parameter")
	ErrorInvalidServiceAccount = errors.New("Invalid service account")
	ErrorInvalidClientSecrets  = errors.New("Invalid client secrets configuration")
	ErrorInvalidDefaults       = errors.New("Invalid defaults file")
	ErrorCacheTokenRead        = errors.New("Invalid Cache Token")
	ErrorCacheTokenWrite       = errors.New("Unable to create cache token")
	ErrorTokenExchange         = errors.New("Token Exchange Error")
	ErrorResponse              = errors.New("Bad Response")
	ErrorBadParameter          = errors.New("Invalid Parameter")
)

// return an error as a string
func (this *Error) Error() string {
	if len(this.detail) > 0 {
		return fmt.Sprintf("%s: %s",this.base.Error(),this.detail)
	} else {
		return fmt.Sprintf("%s",this.base.Error())
	}
}

func NewError(base error,detail error) *Error {
	this := new(Error)
	this.base = base
	if detail != nil {
		this.detail = detail.Error()
	}
	return this
}
