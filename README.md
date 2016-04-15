# ytapi
YouTube API Command-Line Interface

Usage:
  ytapi <flags> <command>

Flags:

  * `--help`              Display usage information
  * `--channel=<id>`      Set the channel
  * `--contentowner=<id>` Set the content owner
  * `--debug`             Debug API calls
  * `--max-results=<int>` The maximum number of results to return or 0 for unlimited

Commands for authenticating for using the YouTube API:

  * `auth` Use this form when authenticating against a channel. It will open
    your web browser where you can give permission to access your YouTube
	account.

  * `--contentowner=<id> auth` Use this form when authenticating against a
    service account.

  * `--contentowner=<id> --channel=<id> auth` Use this form when authenticating 
    against a service account, and setting a default channel to operate on.

Commands for accessing channels:

  * `channels` Use this for listing all channels that you have Read & Write 
    access for.
 
In order to use this package, you'll need to create a ".credentials" folder
into which you need to place two files:

  * client_secrets.json
  * service_account.json
  
There is also a file called `defaults.json` into which you can place any
default parameters you wish to use without putting them onto the command
line.

