
Creating a release
git tag alpha-002 -a -m "Added new alpha release"
build/build.sh -c ~/.ytapi/client_secrets.json -s ~/.ytapi/service_account.json
git push origin alpha-002

On github, then create a release from the tags


Here are some methods that need implemented:

Activities
-> List
-> Insert

Captions
-> Update

Channels
-> InsertChannelBanner
-> UpdateChannelMetadata

-> GetChannelLocalizedMetadata
-> AddChannelLocalizedMetadata
-> DeleteChannelLocalizedMetadata

-> ListChannelSections
-> GetChannelSectionMetadata
-> UpdateChannelSectionMetadata
-> DeleteChannelSection
-> GetChannelSectionLocalizationMetadata
-> AddChannelSectionLocalizationMetadata
-> DeleteChannelSectionLocalizationMetadata

-> SetChannelWatermark
-> DeleteChannelWatermark

Videos
-> DownloadVideoThumbnail
-> UploadVideo
-> DeleteVideo
-> SetVideoRating
-> GetVideoRating
-> UpdateVideoMetadata
-> GetVideoLocalizationMetadata
-> AddVideoLocalizationMetadata
-> DeleteVideoLocalizationMetadata

Here are some other future enhancements:

- SQL & JSON output
- Batch input with -in flag
- Time output format, fix parsing of input time formats
- Local timezone flag
- Analytics API support
