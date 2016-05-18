
# ytapi: YouTube API Command-Line Interface

Usage: `ytapi <flags> <command>`

Or: `ytapi -help`

Here is a list of the main flags which can be used on the command line which
affect most commands

  * `-help`               Display usage information
  * `-channel=<channel>`       Set the channel to act on behalf of
  * `-contentowner=<contentowner>`  Set the content owner
  * `-debug`              Debug API calls, output API requests and responses to stderr
  * `-output=<csv|ascii>` Output format for displaying results 

## Introduction

This command-line utility operates on the YouTube Data API in order to list,
update, delete and search various YouTube objects such as videos, channels,
broadcasts, streams and playlists.

## Installation

In order to use this package, you'll need to create a `.ytapi` folder
into which you need to place one or two files:

  * `client_secrets.json` is required to be placed in the folder
  * `service_account.json` is optional where you wish to operate on the API on 
    behalf of a YouTube content owner.
  
These files can be downloaded from your Google Developer console.


## Installation and Authentication operations

	Authenticate
		Authenticate against service account or channel

## Operations on Channels

	ListChannels
		List channels

## Channel Section operations

	ListChannelSections
		List channel sections

## Operations on videos

	ListVideos
		List videos
	GetVideo
		Get single video

## Operations on Broadcasts

	ListBroadcasts
		List broadcasts
	DeleteBroadcast
		Delete broadcast
	TransitionBroadcast
		Transition broadcast to another state
	BindBroadcast
		Bind or unbind broadcast to stream
	NewBroadcast
		Create a new broadcast

## Operations on Streams

	ListStreams
		List streams
	DeleteStream
		Delete stream

## Operations on video captions

	ListCaptions
		List captions

## Operations on Playlists

	ListPlaylists
		List playlists for channel
	NewPlaylist
		Create a new playlist
	DeletePlaylist
		Delete an existing playlist
	UpdatePlaylist
		Update playlist metadata

## Operations on PlaylistItems

	ListPlaylistItems
		List playlist items for a playlist
	InsertVideoIntoPlaylist
		Inserts a video into a playlist
	DeleteVideoFromPlaylist
		Deletes a video from a playlist

## Language and Region operations

	ListLanguages
		List languages
	ListRegions
		List regions

## Search operations

	SearchVideos
		Search for videos based on text query or related video
	SearchBroadcasts
		Search for live broadcasts
	SearchPlaylists
		Search for playlists based on text query
	SearchChannels
		Search for channels

## Content owner operations

	ListContentOwners
		List content owners

## Policy operations

	ListPolicies
		List policies

## Claim operations

	Claim
		Create a claim between a video and asset with a defined policy
	GetClaim
		Get Existing claim
	ListClaims
		List all claims
	ClaimHistory
		List history for a claim
	UpdateClaim
		Update an existing claim

## Asset operations

	GetAsset
		Get a single asset

## Reference operations

	ListReferences
		Get a list of references for an asset
		
