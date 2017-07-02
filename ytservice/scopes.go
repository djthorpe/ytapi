package ytservice

/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/youtubeanalytics/v1"
)

////////////////////////////////////////////////////////////////////////////////

const (
	SCOPE_DATASSL   = youtube.YoutubeForceSslScope
	SCOPE_DATA      = youtube.YoutubeScope
	SCOPE_DATAREAD  = youtube.YoutubeReadonlyScope
	SCOPE_UPLOAD    = youtube.YoutubeUploadScope
	SCOPE_PARTNER   = youtube.YoutubepartnerScope
	SCOPE_AUDIT     = youtube.YoutubepartnerChannelAuditScope
	SCOPE_ANALYTICS = youtubeanalytics.YtAnalyticsReadonlyScope
	SCOPE_REVENUE   = youtubeanalytics.YtAnalyticsMonetaryReadonlyScope
)

var (
	scopeMap = map[string][]string{
		"default":   []string{SCOPE_DATASSL},
		"all":       []string{SCOPE_DATASSL, SCOPE_UPLOAD, SCOPE_PARTNER, SCOPE_AUDIT, SCOPE_ANALYTICS, SCOPE_REVENUE},
		"data":      []string{SCOPE_DATASSL},
		"datanossl": []string{SCOPE_DATA},
		"partner":   []string{SCOPE_PARTNER},
		"upload":    []string{SCOPE_UPLOAD},
		"audit":     []string{SCOPE_AUDIT},
		"analytics": []string{SCOPE_ANALYTICS},
		"revenue":   []string{SCOPE_ANALYTICS, SCOPE_REVENUE},
	}
)

////////////////////////////////////////////////////////////////////////////////

// scopesForFlags returns the scope names for the current set of flags
func scopesForFlags(flags []string) ([]string, error) {
	scope_map := make(map[string]bool, 8)
	for _, flag := range flags {
		if flag_scopes, exists := scopeMap[flag]; exists {
			for _, scope := range flag_scopes {
				scope_map[scope] = true
			}
		} else {
			return nil, ErrInvalidScopeFlag
		}
	}
	scope_values := make([]string, 0, len(scope_map))
	for k := range scope_map {
		scope_values = append(scope_values, k)
	}
	return scope_values, nil
}
