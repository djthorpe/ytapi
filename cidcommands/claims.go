/*
  Copyright David Thorpe 2015 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package cidcommands

import (
	"errors"
	
	"github.com/djthorpe/ytapi/ytapi"
	"github.com/djthorpe/ytapi/ytservice"
	"github.com/djthorpe/ytapi/youtubepartner/v1"
)

////////////////////////////////////////////////////////////////////////////////
// Register commands

func RegisterClaimCommands() []ytapi.Command {
	return []ytapi.Command{
		ytapi.Command{
			Name:        "Claim",
			Description: "Create a claim between a video and asset with a defined policy",
            ServiceAccount: true,
			Required:    []*ytapi.Flag{ &ytapi.FlagVideo,&ytapi.FlagAsset,&ytapi.FlagPolicy },
			Optional:    []*ytapi.Flag{ &ytapi.FlagClaimType, &ytapi.FlagClaimBlockOutsideOwnership },
			Setup:       RegisterClaimFormat,
			Execute:     InsertClaim,
		},
		ytapi.Command{
			Name:        "GetClaim",
			Description: "Get Existing claim",
            ServiceAccount: true,
			Required:    []*ytapi.Flag{ &ytapi.FlagClaim },
			Setup:       RegisterClaimFormat,
			Execute:     GetClaim,
		},
		ytapi.Command{
			Name:        "ListClaims",
			Description: "List all claims",
            ServiceAccount: true,
			Optional:    []*ytapi.Flag{ &ytapi.FlagMaxResults },
			Setup:       RegisterClaimFormat,
			Execute:     ListClaims,
		},
		ytapi.Command{
			Name:        "ClaimHistory",
			Description: "List history for a claim",
            ServiceAccount: true,
			Required:    []*ytapi.Flag{ &ytapi.FlagClaim },
			Setup:       RegisterClaimHistoryFormat,
			Execute:     GetClaimHistory,
		},
		ytapi.Command{
			Name:        "UpdateClaim",
			Description: "Update an existing claim",
            ServiceAccount: true,
            Required:    []*ytapi.Flag{ &ytapi.FlagClaim },
			Optional:    []*ytapi.Flag{ &ytapi.FlagClaimStatus,&ytapi.FlagPolicy,&ytapi.FlagClaimBlockOutsideOwnership },
			Setup:       RegisterClaimFormat,
			Execute:     PatchClaim,
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
	table.SetColumns([]string{"claim", "asset", "video","status","origin","type","timeCreated","policyName" })

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
	table.SetColumns([]string{"type", "time" })

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// List Claims

func ListClaims(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	call := service.PAPI.Claims.List()
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	return ytapi.DoClaimsList(call,table,values.GetInt(&ytapi.FlagMaxResults))
}

////////////////////////////////////////////////////////////////////////////////
// Get Claim

func GetClaim(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	call := service.PAPI.Claims.Get(values.GetString(&ytapi.FlagClaim))
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response,err := call.Do()
	if err != nil {
		return err
	}
	if err = table.Append([]*youtubepartner.Claim{ response }); err != nil {
		return err
	}
	return nil
}


////////////////////////////////////////////////////////////////////////////////
// Claim History

func GetClaimHistory(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
	// create call and set parameters
	call := service.PAPI.ClaimHistory.Get(values.GetString(&ytapi.FlagClaim))
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response,err := call.Do()
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
	// create call and set parameters
	call := service.PAPI.Claims.Insert(&youtubepartner.Claim{
		AssetId: values.GetString(&ytapi.FlagAsset),
		VideoId: values.GetString(&ytapi.FlagVideo),
		ContentType: values.GetString(&ytapi.FlagClaimType),
		Policy: &youtubepartner.Policy{
			Id: values.GetString(&ytapi.FlagPolicy),
		},
	})
	call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))

	// Execute
	response,err := call.Do()
	if err != nil {
		return err
	}

	// Handle block outside ownership
	// TODO: always force sending the flag
	if values.IsSet(&ytapi.FlagClaimBlockOutsideOwnership) {
		call := service.PAPI.Claims.Patch(response.Id,&youtubepartner.Claim{
			BlockOutsideOwnership: values.GetBool(&ytapi.FlagClaimBlockOutsideOwnership),
		})
		call = call.OnBehalfOfContentOwner(values.GetString(&ytapi.FlagContentOwner))
		// Execute
		response,err = call.Do()
		if err != nil {
			return err
		}
	}

	if err = table.Append([]*youtubepartner.Claim{ response }); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Patch Claim

func PatchClaim(service *ytservice.Service, values *ytapi.Values, table *ytapi.Table) error {
    // TODO
    return nil
}


