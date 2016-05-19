
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

This command-line utility operates on the YouTube API's' in order to list,
update, delete and search various YouTube objects such as videos, channels,
broadcasts, streams and playlists. It implements most of the API calls for
both the YouTube Data API and the YouTube Partner API, and allows you to
tailor the resulting data in a number of forms (text, CSV). You can distribute
the binary widely and it will use the authentication credentials you specify
when building your binary package.

## Authentication credentials

The software allows you to authenticate in two ways against YouTube to read,
create, modify and delete information on YouTube:

 * _Authentication using OAuth2_ This method ties the tool to running against
   a specific channel, which you specify by using your web browser. It requires
   the installation of credentials called a "Client Secret" which can be
   downloaded from the Google Developers Console
 * _Authentication using a Service Account_ This method ties the tool to a
   service account (and associated email address) which allows you to access
   data without having to authenticate before use. You can use this 
   authentication technique to use the YouTube Partner API in addition to the
   YouTube Data API. It requires installation of service account credentials
   which can be downloaded from the Google Developers Console
   
In either case, you will need to enable the YouTube Data API and YouTube Partner
API in your project, also accessible through the Google Developers Console.

## Building the tool

In order to build the tool, please download the latest version from the
releases page here:

  https://github.com/djthorpe/ytapi/releases

The resulting 'ytapi' folder will have a build script under the build folder.
If you have git and go installed, you can then build the software as follows:

```
  ytapi% build/build.sh
```

In order to "bake in" credentials for client secrets and service accounts,
you can specify the location of those files on the build command line. You
can download these credentials from your Google Cloud Console:

```
  ytapi% build/build.sh -c ~/.ytapi/client_secret.json -s ~/.ytapi/service_account.json
```

Be aware that if you bake in the credentials, anyone making API calls will impact
your API quota, so you should keep an eye on your Google Cloud Console to ensure
the tool isn't being mis-used. In addition, anyone using your service account
credentials within their YouTube Content Management system gives you access to
their CMS system, so be careful when distributing the tool when there are
service account credentials baked into the binary.

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
		
