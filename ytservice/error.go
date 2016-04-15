package ytservice

import (
	"errors"
)

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
