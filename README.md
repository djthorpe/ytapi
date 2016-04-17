
# ytapi: YouTube API Command-Line Interface

Usage: `ytapi <flags> <command>`

Or: `ytapi --help`

Flags:

  * `-help`               Display usage information
  * `-channel=<id>`       Set the channel
  * `-contentowner=<id>`  Set the content owner
  * `-q=<string>`         Query text used when searching
  * `-debug`              Debug API calls
  * `-output=<csv|ascii>` Output format for displaying results 
  * `-part=<+part,-part,...>` Add and/or remote parts from the output
  * `-maxresults=<int>`   The maximum number of results to return or 0 for unlimited

## Introduction

This command-line utility operates on the YouTube Data API in order to list,
update, delete and search various YouTube objects such as videos, channels,
broadcasts, streams and playlists.


## Authentication

Commands for authenticating for using the YouTube API:

  * `Authenticate` Use this form when authenticating against a channel. It will open
    your web browser where you can give permission to access your YouTube
	account. The channel you specify will be used as a default to operate on
	for subsequent operations.

  * `--contentowner=<id> Authenticate` Use this form when authenticating against a
    service account. The content owner you specify will be used as a default to 
	operate on for subsequent operations.

  * `--contentowner=<id> --channel=<id> Authenticate` Use this form when 
    authenticating against a service account. The content owner and channel you 
	specify will be used as defaults for subsequent operations.

## Searching YouTube

Commands for performing searches of YouTube:

  * `--q=<search> Search` Search videos, channels and playlists.

## Operations on Channels

Commands for accessing channels:

  * `ListChannels` Use this for listing the channel or channels that you have
    access for

## Operations for Live Streams

 * `ListBroadcasts` Use this for listing your broadcasts
 
 * `--video=<id> DeleteBroadcast` Delete a single broadcast from your channel
 
 * `ListStreams` Use this for listing your streams


## Installation

In order to use this package, you'll need to create a ".credentials" folder
into which you need to place one or two files:

  * `client_secrets.json` is required to be placed in the folder
  * `service_account.json` is optional where you wish to operate on the API on 
    behalf of a YouTube content owner.
  
These files can be downloaded from your Google Developer console.

