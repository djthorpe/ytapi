
# Using "ytapi" for YouTube Channel Localization

## Introduction

This guide provides information on using the `ytapi` command-line tool for
localization of your channel, playlists and videos, including adding localized
metadata and captions to your videos. In order to use the tool, you will need
to have a Windows, Macintosh or Linux operating system and be familiar with
running command-line tools.

Localization of metadata is providing alternate information (usually titles and descriptions)
in alternate languages, so that if a viewer of your channel, playlist or video who has an alternate
language selected, they may receive localized metadata rather than your default metadata.
Clearly if you don't provide localized metadata in your viewers language, your default
metadata is displayed instead.

You are free to upload one or more caption tracks to your videos, which provided timed subtitles
overlayed onto your videos. YouTube can also support automatic "rough translation" of your
caption tracks into alternate languages when you download existing caption tracks.

Finally, you can alter the display of your channel based on targeting channel sections
to regions and languages. For example, if you want to only display a certain playlist
of videos on your channel to viewers with France set as their chosen country, you can
target that section only to that country.

Many of these functions are provided through the YouTube Creator Studio. Only
section targetting is only available through the `ytapi` tool. Here is a list of
the operations you might want to perform to localize your channel:

| Operation                             |  Web  | YT API Command                         |
| ------------------------------------- | :---: | -------------------------------------- |
| Set Default Channel Language          |   Y   | ytapi UpdateChannelMetadata            |
| Localize Channel Metadata             |   Y   | ytapi UpdateChannelLocalizedMetadata   |
| Set Default Video Language            |   Y   | ytapi UpdateVideoMetadata              |
| Localize Video Metadata               |   Y   | ytapi UpdateVideoLocalizedMetadata     |
| Set Default Playlist Language         |   Y   | ytapi UpdatePlaylistMetadata           |
| Add Country-Specific Channel Section  |   N   | ytapi UpdateSectionLocalizedMetadata   |
| Add Language-Specific Channel Section |   N   | ytapi UpdateSectionLocalizedMetadata   |
| List Video Caption Track              |   Y   | ytapi ListCaptionTrack                 |
| Upload Video Caption Track            |   Y   | ytapi UploadCaptionTrack               |
| Download Video Caption Track          |   Y   | ytapi DownloadCaptionTrack             |
| Delete Video Caption Track            |   Y   | ytapi DeleteCaptionTrack               |

## Installation and Authentication

## Using the Tool

## Localizing Metadata

## Channel Section targetting

## Video Caption tracks




