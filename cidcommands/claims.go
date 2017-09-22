package cidcommands

/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"errors"
	"strings"

	"github.com/djthorpe/ytapi/youtubepartner/v1"
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterClaimCommands() []*ytapi.Command {
	return []*ytapi.Command{
		&ytapi.Command{
			Name:           "Claim",
			Description:    "Create a claim between a video and asset with a defined policy",
			ServiceAccount: true,
			Required:       []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagAsset, &ytapi.FlagPolicy},
			Optional:       []*ytapi.Flag{&ytapi.FlagClaimType, &ytapi.FlagClaimBlockOutsideOwnership},
			Setup:          RegisterClaimFormat,
			Execute:        InsertClaim,
		},
		&ytapi.Command{
			Name:           "GetClaim",
			Description:    "Get Existing claim",
			ServiceAccount: true,
			Optional:       []*ytapi.Flag{&ytapi.FlagClaim, &ytapi.FlagVideo},
			Setup:          RegisterClaimFormat,
			Execute:        GetClaim,
		},
		&ytapi.Command{
			Name:           "ListClaims",
			Description:    "List all claims",
			ServiceAccount: true,
			Optional:       []*ytapi.Flag{&ytapi.FlagMaxResults},
			Setup:          RegisterClaimFormat,
			Execute:        ListClaims,
		},
		&ytapi.Command{
			Name:           "ClaimHistory",
			Description:    "List history for a claim",
			ServiceAccount: true,
			Optional:       []*ytapi.Flag{&ytapi.FlagClaim, &ytapi.FlagVideo},
			Setup:          RegisterClaimHistoryFormat,
			Execute:        GetClaimHistory,
		},
		&ytapi.Command{
			Name:           "UpdateClaim",
			Description:    "Update an existing claim or video",
			ServiceAccount: true,
			Optional:       []*ytapi.Flag{&ytapi.FlagVideo, &ytapi.FlagClaim, &ytapi.FlagClaimStatus, &ytapi.FlagPolicy, &ytapi.FlagClaimBlockOutsideOwnership},
			Setup:          RegisterClaimFormat,
			Execute:        PatchClaim,
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Register format

func RegisterClaimFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("id", []*ytapi.Flag{
		&ytapi.Flag{Name: "claim", Path: "Id", Type: ytapi.FLAG_STRING},
	})

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "asset", Path: "AssetId", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "video", Path: "VideoId", Type: ytapi.FLAG_VIDEO},
		&ytapi.Flag{Name: "timeCreated", Path: "TimeCreated", Type: ytapi.FLAG_TIME},
	})

	table.RegisterPart("claimDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "status", Path: "Status", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "origin", Path: "Origin/Source", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "type", Path: "ContentType", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "isPartnerUploaded", Path: "IsPartnerUploaded", Type: ytapi.FLAG_BOOL},
		&ytapi.Flag{Name: "blockOutsideOwnership", Path: "BlockOutsideOwnership", Type: ytapi.FLAG_BOOL},
	})

	table.RegisterPart("policyDetails", []*ytapi.Flag{
		&ytapi.Flag{Name: "policy", Path: "Policy/Id", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "policyName", Path: "Policy/Name", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"claim", "asset", "video", "status", "origin", "type", "timeCreated", "policyName"})

	// success
	return nil
}

func RegisterClaimHistoryFormat(values *ytapi.Values, table *ytapi.Table) error {

	table.RegisterPart("snippet", []*ytapi.Flag{
		&ytapi.Flag{Name: "type", Path: "Type", Type: ytapi.FLAG_STRING},
		&ytapi.Flag{Name: "time", Path: "Time", Type: ytapi.FLAG_TIME},
		&ytapi.Flag{Name: "source", Path: "Source/Type", Type: ytapi.FLAG_STRING},
	})

	// set default columns
	table.SetColumns([]string{"type", "time"})

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Return Policy

func policyByValue(value string) (*youtubepartner.Policy, error) {
	value2 := strings.TrimSpace(strings.ToLower(value))
	switch value2 {
	case "block", "track", "monetize", "takedown":
		return policyByWorldwideRule(value2)
	default:
		return policyByIdentifier(value)
	}
}

func policyByIdentifier(name string) (*youtubepartner.Policy, error) {
	return &youtubepartner.Policy{
		Id: name,
	}, nil
}

func policyByWorldwideRule(name string) (*youtubepartner.Policy, error) {
	switch name {
	case "block":
		return &youtubepartner.Policy{
			Rules: []*youtubepartner.PolicyRule{
				&youtubepartner.PolicyRule{
					Action: "block",
				},
			},
		}, nil
	case "track":
		return &youtubepartner.Policy{
			Rules: []*youtubepartner.PolicyRule{
				&youtubepartner.PolicyRule{
					Action: "track",
				},
			},
		}, nil
	case "monetize":
		return &youtubepartner.Policy{
			Rules: []*youtubepartner.PolicyRule{
				&youtubepartner.PolicyRule{
					Action: "monetize",
				},
			},
		}, nil
	case "takedown":
		return &youtubepartner.Policy{
			Rules: []*youtubepartner.PolicyRule{
				&youtubepartner.PolicyRule{
					Action: "takedown",
				},
			},
		}, nil
	default:
		return nil, errors.New("Invalid policy value")
	}
}

////////////////////////////////////////////////////////////////////////////////
// Return Claim

func claimSearchByVideo(service *ytservice.Service, values *ytapi.Values) (string, error) {
	// create call and set parameters
	call := service.PAPI.ClaimSearch.List()
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
	call = call.VideoId(values.GetString(&ytapi.FlagVideo))
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return "", err
	}
	// TODO - We might return more than one claim, only count active claims
	// not inactive ones
	if len(response.Items) != 1 {
		return "", errors.New("Invalid claim")
	}
	return response.Items[0].Id, nil
}

////////////////////////////////////////////////////////////////////////////////
// List Claims

func ListClaims(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	call := service.PAPI.Claims.List()
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	return ytapi.DoClaimsList(call, table, values.GetInt(&ytapi.FlagMaxResults), service.CallOptions()...)
}

////////////////////////////////////////////////////////////////////////////////
// Get Claim

func GetClaim(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// claim or video parameter
	var err error
	var claim string
	if values.IsSet(&ytapi.FlagClaim) {
		claim = values.GetString(&ytapi.FlagClaim)
	} else if values.IsSet(&ytapi.FlagVideo) {
		if claim, err = claimSearchByVideo(service, values); err != nil {
			return err
		}
	} else {
		return errors.New("Expect -claim or -video flag")
	}

	// create call and set parameters
	call := service.PAPI.Claims.Get(claim)
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}
	if err = table.Append([]*youtubepartner.Claim{response}); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Claim History

func GetClaimHistory(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// claim or video parameter
	var err error
	var claim string
	if values.IsSet(&ytapi.FlagClaim) {
		claim = values.GetString(&ytapi.FlagClaim)
	} else if values.IsSet(&ytapi.FlagVideo) {
		if claim, err = claimSearchByVideo(service, values); err != nil {
			return err
		}
	} else {
		return errors.New("Expect -claim or -video flag")
	}

	// create call and set parameters
	call := service.PAPI.ClaimHistory.Get(claim)
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}
	if err = table.Append(response.Event); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Create Claim

func InsertClaim(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {

	// policy parameter
	var err error
	var policy *youtubepartner.Policy
	if values.IsSet(&ytapi.FlagPolicy) {
		if policy, err = policyByValue(values.GetString(&ytapi.FlagPolicy)); err != nil {
			return err
		}
	}

	// create call and set parameters
	call := service.PAPI.Claims.Insert(&youtubepartner.Claim{
		AssetId:     values.GetString(&ytapi.FlagAsset),
		VideoId:     values.GetString(&ytapi.FlagVideo),
		ContentType: values.GetString(&ytapi.FlagClaimType),
		Policy:      policy,
	})
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Handle block outside ownership
	// TODO: always force sending the flag
	if values.IsSet(&ytapi.FlagClaimBlockOutsideOwnership) {
		call := service.PAPI.Claims.Patch(response.Id, &youtubepartner.Claim{
			BlockOutsideOwnership: values.GetBool(&ytapi.FlagClaimBlockOutsideOwnership),
		})
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		// Execute
		response, err = call.Do(service.CallOptions()...)
		if err != nil {
			return err
		}
	}

	if err = table.Append([]*youtubepartner.Claim{response}); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Patch Claim

func PatchClaim(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	var err error

	// claim or video parameter
	var claim string
	if values.IsSet(&ytapi.FlagClaim) {
		claim = values.GetString(&ytapi.FlagClaim)
	} else if values.IsSet(&ytapi.FlagVideo) {
		if claim, err = claimSearchByVideo(service, values); err != nil {
			return err
		}
	} else {
		return errors.New("Expect -claim or -video flag")
	}

	// policy parameter
	var policy *youtubepartner.Policy
	if values.IsSet(&ytapi.FlagPolicy) {
		policy, err = policyByValue(values.GetString(&ytapi.FlagPolicy))
	}

	call := service.PAPI.Claims.Patch(claim, &youtubepartner.Claim{
		BlockOutsideOwnership: values.GetBool(&ytapi.FlagClaimBlockOutsideOwnership),
		Status:                values.GetString(&ytapi.FlagClaimStatus),
		Policy:                policy,
		ForceSendFields: values.SetFields(map[string]*ytapi.Flag{
			"BlockOutsideOwnership": &ytapi.FlagClaimBlockOutsideOwnership,
			"Status":                &ytapi.FlagClaimStatus,
			"Policy":                &ytapi.FlagPolicy,
		}),
	})
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response, err := call.Do(service.CallOptions()...)
	if err != nil {
		return err
	}

	// Append claim to table
	if err = table.Append([]*youtubepartner.Claim{response}); err != nil {
		return err
	}
	return nil
}
