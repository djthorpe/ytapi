
# ytapi: YouTube API Command-Line Interface

## Introduction

This command-line utility operates on the YouTube API's in 
order to list, update, delete and search various YouTube objects (for example,
channels, videos, playlists, live streams, analytics reports, claims, assets,
policies....)

It implements most of the API calls and allows you to tailor the resulting data 
in a number of forms (text, CSV). You can distribute the binary widely and it 
will use the authentication credentials you specify when building your binary 
package.

Future versions of the tool may also:

  * Output SQL & JSON as well as CSV
  * Allow batch input of information

## Usage

Usage:
  * `ytapi <flags> <call>` Execute an API call

In order to get help:

  * `ytapi -help` will display a list of commands that can be used
  * `ytapi -help <call>` will display flags and output fields for a particular API call

Here is a list of the main flags which can be used on the command line:

  * `-credentials=<foldername>` Folder where credentials are stored
  * `-channel=<channel>` Set the channel to act on behalf of
  * `-contentowner=<contentowner>` Set the content owner
  * `-debug` Debug API calls, output API requests and responses to stderr
  
The following commands affect the input and output of data:

  * `-out=(<filename>.)<csv|ascii>` Output format for displaying results 
  * `-fields=(+|-)(<field|part>),...` Which fields to include or exclude in output

## A note on authentication credentials

The software allows you to authenticate in two ways against YouTube to read,
create, modify and delete information on YouTube:

 * *Authentication using OAuth2* This method ties the tool to running against
   a specific channel, which you specify by using your web browser. It requires
   the installation of credentials called a "Client Secret" which can be
   downloaded from the Google Developers Console
 * *Authentication using a Service Account* This method ties the tool to a
   service account (and associated email address) which allows you to access
   data without having to authenticate before use. You can use this 
   authentication technique to use the YouTube Partner API in addition to the
   YouTube Data API. It requires installation of service account credentials
   which can be downloaded from the Google Developers Console
   
In either case, you will need to enable the YouTube API's in your project, 
also accessible through the Google Developers Console. If using service
accounts, you will also need to add the service account email to your YouTube
Content ID user access list.

## Building the tool

In order to build the tool, please download the latest version from the
releases page here:

  https://github.com/djthorpe/ytapi/releases

The resulting 'ytapi' folder will have a build script under the build folder.
If you have git and go installed, you can then build the software as follows:

```
  ytapi% build/build.sh
```

In most cases, you might want to distribute your binary tool with credentials
"baked in" so that your users can start using the tool without going through
additional installation steps. In order to "bake in" credentials for client 
secrets and service accounts, you can specify the location of those files on 
the build command line. You can download these credentials from your Google 
Developers Console:

  https://developers.google.com/youtube/analytics/registering_an_application
  
You can use the `-c` and `-s` flags in order to do this. For example,

```
  ytapi% build/build.sh -c ~/.ytapi/client_secret.json -s ~/.ytapi/service_account.json
```

Be aware that if you bake in the credentials, anyone making API calls will impact
your API quota, so you should keep an eye on your Google Cloud Console to ensure
the tool isn't being mis-used. In addition, anyone using your service account
credentials within their YouTube Content Management system gives you access to
their CMS system, so be careful when distributing the tool when there are
service account credentials baked into the binary.

## Distribution of your tool

Once you have a binary, you can distribute your tool and people can start using it
by putting the binary in their search path. When they run the tool for the first
time, they can use the "Install" command to create the local credentials. For example,

```
  home% ytapi Install
  home% ytapi ListVideos
```

If you wish to have several different sets of credentials, you can use the
`-credentials` flag to indicate which set. The default value is `.ytapi` but
you can select a different set of credentials as follows:

```
  home% ytapi -credentials=.ytapi2 Install
  home% ytapi -credentials=.ytapi2 ListVideos
```

The `Authenticate` command will refresh the OAuth or Service Account tokens
and allow you to set required scopes of operation. In order to authenticate
against a YouTube channel, use the following form:

```
  home% ytapi Authenticate
```

To authenticate using service account credentials against a named content
owner, use the following form:

```
  home% ytapi -contentowner=<contentowner> -serviceaccount Authenticate
```

This will display a list of channels which the content owner manages. Subsequent
calls will use the same content owner information, so no need to append the
`-contentowner` flag unless you want to reauthenticate. Finally, you can specify
a default channel when using service accounts:

```
  home% ytapi -contentowner=<contentowner> -channel=<channel> Authenticate
```

This sets the default channel on which to operate YouTube API calls. By default,
authentication will set the scopes (or "permissions") of operation to the Data API.
In order to gain other permissions, use the `-scope` flag. For example,

```
  home% ytapi -contentowner=<contentowner> -scope=data,partner Authenticate
```

Here is the list of scopes of operation:

data | Manage YouTube channel account
dataread | Read-only access to YouTube channel account
upload | Ability to upload to YouTube channel account
partner | Manage YouTube partner account
audit | View private information of your YouTube channel relevant during the audit process with a YouTube partner
analytics | Access YouTube analytics
revenue | Access YouTube analytics revenue data
all | All YouTube permissions

By default, the scope set is only "data"

## Parameters to make API calls

## Fields and Parts on output

## Batch input

## API Call Reference

The following sections list the calls that can be made to the YouTube Data and
Partner API's.

### Installation and Authentication operations

	Authenticate
		Authenticate against service account or channel

### Operations on Channels

	ListChannels
		List channels

### Channel Section operations

	ListChannelSections
		List channel sections

### Operations on videos

	ListVideos
		List videos
	GetVideo
		Get single video

### Operations on Broadcasts

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

### Operations on Streams

	ListStreams
		List streams
	DeleteStream
		Delete stream

### Operations on video captions

	ListCaptions
		List captions

### Operations on Playlists

	ListPlaylists
		List playlists for channel
	NewPlaylist
		Create a new playlist
	DeletePlaylist
		Delete an existing playlist
	UpdatePlaylist
		Update playlist metadata

### Operations on PlaylistItems

	ListPlaylistItems
		List playlist items for a playlist
	InsertVideoIntoPlaylist
		Inserts a video into a playlist
	DeleteVideoFromPlaylist
		Deletes a video from a playlist

### Language and Region operations

	ListLanguages
		List languages
	ListRegions
		List regions

### Search operations

	SearchVideos
		Search for videos based on text query or related video
	SearchBroadcasts
		Search for live broadcasts
	SearchPlaylists
		Search for playlists based on text query
	SearchChannels
		Search for channels

### Content owner operations

	ListContentOwners
		List content owners

### Policy operations

	ListPolicies
		List policies

### Claim operations

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

### Asset operations

	GetAsset
		Get a single asset

### Reference operations

	ListReferences
		Get a list of references for an asset
		

		
