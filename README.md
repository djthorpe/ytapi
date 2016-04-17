# ytapi
YouTube API Command-Line Interface

Usage: `ytapi <flags> <command>`

Or: `ytapi --help`

Flags:

  * `--help`              Display usage information
  * `--channel=<id>`      Set the channel
  * `--contentowner=<id>` Set the content owner
  * `--debug`             Debug API calls
  * `--maxresults=<int>`  The maximum number of results to return or 0 for unlimited

== Authentication ==

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

== Channels ==

Commands for accessing channels:

  * `ListChannels` Use this for listing the channel or channels that you have
    access for.

== Installation ==

In order to use this package, you'll need to create a ".credentials" folder
into which you need to place two files:

  * `client_secrets.json`
  * `service_account.json`
  
There is also a file called `defaults.json` into which you can place any
default parameters you wish to use without putting them onto the command
line.

