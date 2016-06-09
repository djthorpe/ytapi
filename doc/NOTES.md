
Creating a release
git tag alpha-002 -a -m "Added new alpha release"
build/build.sh -c ~/.ytapi/client_secrets.json -s ~/.ytapi/service_account.json
git push origin alpha-002

On github, then create a release from the tags

In order to generate the YouTube Partner API package, you need to run the
following commands:

cd $GOPATH/src/github.com/djthorpe/ytapi
curl https://www.googleapis.com/discovery/v1/apis/youtubePartner/v1/rest > youtubepartner/v1/youtubepartner.json
google-api-go-generator -api_json_file=youtubepartner/v1/youtubepartner.json

----- old notes -------


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
