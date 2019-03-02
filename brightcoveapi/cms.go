package brightcoveapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

/////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	// URL retrieve tokens
	CMS_BASE_URL = "https://cms.api.brightcove.com/v1/"
)

/////////////////////////////////////////////////////////////////////
// STRUCTS

type cms struct {
	client  *Client
	baseurl *url.URL
}

type Video struct {
	Id                string                 `json:"id"`
	AccountId         string                 `json:"account_id"`
	AdKeys            string                 `json:"ad_keys"`
	Category          string                 `json:"category"`
	Live              bool                   `json:"live"`
	ClipSourceVideoId string                 `json:"clip_source_video_id"`
	Complete          bool                   `json:"complete"`
	CreatedAt         string                 `json:"created_at"` // 2015-09-17T16:08:37.108Z
	CuePoints         []*VideoCuePoint       `json:"cue_points"`
	CustomFields      map[string]string      `json:"custom_fields"`
	DeliveryType      string                 `json:"delivery_type"`
	Description       string                 `json:"description"`
	DigitalMasterId   string                 `json:"digital_master_id"`
	DrmDisabled       bool                   `json:"drm_disabled"`
	Duration          uint32                 `json:"duration"`
	Economics         string                 `json:"economics"`
	FolderId          string                 `json:"folder_id"`
	Geo               *VideoGeo              `json:"geo"`
	HasDigitalMaster  bool                   `json:"has_digital_master"`
	Images            map[string]*VideoImage `json:"images"`
	LongDescription   string                 `json:"long_description"`
	Name              string                 `json:"name"`
	OfflineEnabled    bool                   `json:"offline_enabled"`
	OriginalFilename  string                 `json:"original_filename"`
	Projection        string                 `json:"projection"`
	PublishedAt       string                 `json:"published_at"`
	ReferenceId       string                 `json:"reference_id"`
	Schedule          *VideoSchedule         `json:"schedule"`
	State             string                 `json:"state"`
	Tags              []string               `json:"tags"`
	TextTracks        []*VideoTextTrack      `json:"text_tracks"`
	UpdatedAt         string                 `json:"updated_at"`
}

type VideoTextTrack struct {
	Id       string `json:"id"`
	Src      string `json:"src"`
	SrcLang  string `json:"srclang"`
	Default  bool   `json:"default"`
	Kind     string `json:"kind"`
	Label    string `json:"label"`
	MimeType string `json:"mime_type"`
}

type VideoCuePoint struct {
	Id        string  `json:"id"`
	Time      float32 `json:"time"`
	Type      string  `json:"type"`
	ForceStop bool    `json:"force-stop"`
	Metadata  string  `json:"metadata"`
	Name      string  `json:"name"`
}

type VideoGeo struct {
	Countries        string `json:"countries"`
	ExcludeCountries bool   `json:"exclude_countries"`
	Restricted       bool   `json:"restricted"`
}

type VideoImage struct {
	Sources []*VideoImageSource `json:"sources"`
	Src     string              `json:"src"`
	AssetId string              `json:"asset_id"`
}

type VideoImageSource struct {
	Src string `json:"src"`
}

type VideoLink struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type VideoSchedule struct {
	EndsAt   string `json:"ends_at"`
	StartsAt string `json:"starts_at"`
}

type VideoSharing struct {
	ByExternalAcct bool   `json:"by_external_acct"`
	ById           string `json:"by_id"`
	ByReference    bool   `json:"by_reference"`
	SourceId       string `json:"source_id"`
	ToExternalAcct bool   `json:"to_external_acct"`
}

type VideoCount struct {
	Count uint `json:"count"`
}

/////////////////////////////////////////////////////////////////////
// CONSTRUCTOR

func (this *Client) NewCMS() *cms {
	if url, err := url.Parse(CMS_BASE_URL); err != nil {
		return nil
	} else {
		return &cms{
			client:  this,
			baseurl: url,
		}
	}
}

/////////////////////////////////////////////////////////////////////
// GET VIDEOS

// GetVideos returns a list of videos using the "Get_Videos" API call
// https://docs.brightcove.com/cms-api/v1/doc/index.html#operation/GetVideos
func (this *cms) GetVideos(options ...ClientOption) ([]*Video, error) {
	var videos []*Video

	if err := this.client.SetOptions(options); err != nil {
		return nil, err
	} else if absurl, err := URLJoin(this.baseurl, "accounts/"+this.client.AccountId()+"/videos", this.client.options); err != nil {
		return nil, err
	} else if req, err := this.client.NewRequest("GET", absurl.String()); err != nil {
		return nil, err
	} else if resp, err := this.client.Do(req); err != nil {
		return nil, err
	} else {
		decoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()
		if err := decoder.Decode(&videos); err != nil {
			return nil, err
		}
		return videos, nil
	}
}

/////////////////////////////////////////////////////////////////////
// GET VIDEO COUNT

// GetVideoCount returns the number of videos for the account or a search
// https://docs.brightcove.com/cms-api/v1/doc/index.html#operation/GetVideoCount
func (this *cms) GetVideoCount(options ...ClientOption) (uint, error) {
	var count VideoCount
	if err := this.client.SetOptions(options); err != nil {
		return 0, err
	} else if absurl, err := URLJoin(this.baseurl, "accounts/"+this.client.AccountId()+"/counts/videos", this.client.options); err != nil {
		return 0, err
	} else if req, err := this.client.NewRequest("GET", absurl.String()); err != nil {
		return 0, err
	} else if resp, err := this.client.Do(req); err != nil {
		return 0, err
	} else {
		decoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()
		if err := decoder.Decode(&count); err != nil {
			return 0, err
		}
		return count.Count, nil
	}
}

/////////////////////////////////////////////////////////////////////
// GET VIDEO BY ID

// GetVideoById gets a video object - you can include up to 10 video ids
// https://docs.brightcove.com/cms-api/v1/doc/index.html#operation/GetVideoByIdOrReferenceId
func (this *cms) GetVideoById(video_id []string, options ...ClientOption) ([]*Video, error) {
	var video Video
	var videos []*Video

	// Return ErrBadParameter if no videos are specified
	if len(video_id) == 0 {
		return nil, ErrBadParameter
	}

	// We don't count the number of videos, we expect the API to return an error
	// when the maximum number of parameters is reached
	video_list := url.PathEscape(strings.Join(video_id, ","))
	if err := this.client.SetOptions(options); err != nil {
		return nil, err
	} else if absurl, err := URLJoin(this.baseurl, "accounts/"+this.client.AccountId()+"/videos/"+video_list, this.client.options); err != nil {
		return nil, err
	} else if req, err := this.client.NewRequest("GET", absurl.String()); err != nil {
		return nil, err
	} else if resp, err := this.client.Do(req); err != nil {
		return nil, err
	} else {
		// The Brightcove CMS returns either a single JSON object or an array
		// of JSON objects depending on the number of requested items
		if len(video_id) == 1 {
			decoder := json.NewDecoder(resp.Body)
			defer resp.Body.Close()
			if err := decoder.Decode(&video); err != nil {
				return nil, err
			}
			return []*Video{&video}, nil
		} else {
			decoder := json.NewDecoder(resp.Body)
			defer resp.Body.Close()
			if err := decoder.Decode(&videos); err != nil {
				return nil, err
			}
			return videos, nil
		}
	}
}

/////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *VideoImage) String() string {
	fields := make([]string, 0)
	if v.AssetId != "" {
		fields = append(fields, fmt.Sprintf("asset_id=%v", v.AssetId))
	}
	if v.Src != "" {
		fields = append(fields, fmt.Sprintf("src=%v", v.Src))
	}
	if len(v.Sources) > 0 {
		fields = append(fields, fmt.Sprintf("sources=%v", v.Sources))
	}
	return fmt.Sprintf("<VideoImage>{ %v }", strings.Join(fields, " "))
}

func (v *VideoImageSource) String() string {
	return fmt.Sprintf("<VideoImageSource>{ src='%v' }", v.Src)
}

func (v *Video) String() string {
	fields := make([]string, 0)
	if v.AccountId != "" {
		fields = append(fields, fmt.Sprintf("account_id=%v", v.AccountId))
	}
	if v.AdKeys != "" {
		fields = append(fields, fmt.Sprintf("ad_keys='%v'", v.AdKeys))
	}
	if v.Category != "" {
		fields = append(fields, fmt.Sprintf("category='%v'", v.Category))
	}
	if v.Live {
		fields = append(fields, fmt.Sprintf("live=%v", v.Live))
	}
	if v.ClipSourceVideoId != "" {
		fields = append(fields, fmt.Sprintf("clip_source_video_id=%v", v.ClipSourceVideoId))
	}
	if v.Complete {
		fields = append(fields, fmt.Sprintf("complete=%v", v.Complete))
	}
	if v.CreatedAt != "" {
		fields = append(fields, fmt.Sprintf("created_at=%v", v.CreatedAt))
	}
	if len(v.CuePoints) > 0 {
		fields = append(fields, fmt.Sprintf("cue_points=%v", v.CuePoints))
	}
	if len(v.CustomFields) > 0 {
		fields = append(fields, fmt.Sprintf("custom_fields=%v", v.CustomFields))
	}
	if v.DeliveryType != "" {
		fields = append(fields, fmt.Sprintf("delivery_type='%v'", v.DeliveryType))
	}
	if v.Description != "" {
		fields = append(fields, fmt.Sprintf("description='%v'", v.Description))
	}
	if v.DigitalMasterId != "" {
		fields = append(fields, fmt.Sprintf("digital_master_id=%v", v.DigitalMasterId))
	}
	if v.DrmDisabled {
		fields = append(fields, fmt.Sprintf("drm_disabled=%v", v.DrmDisabled))
	}
	if v.Duration > 0 {
		fields = append(fields, fmt.Sprintf("duration=%v", v.Duration))
	}
	if v.Economics != "" {
		fields = append(fields, fmt.Sprintf("economics='%v'", v.Economics))
	}
	if v.FolderId != "" {
		fields = append(fields, fmt.Sprintf("folder_id=%v", v.FolderId))
	}
	if v.Geo != nil {
		fields = append(fields, fmt.Sprintf("geo=%v", v.Geo))
	}
	if v.HasDigitalMaster {
		fields = append(fields, fmt.Sprintf("has_digital_master=%v", v.HasDigitalMaster))
	}
	if len(v.Images) > 0 {
		fields = append(fields, fmt.Sprintf("images=%v", v.Images))
	}
	if v.LongDescription != "" {
		fields = append(fields, fmt.Sprintf("long_description='%v'", v.LongDescription))
	}
	if v.Name != "" {
		fields = append(fields, fmt.Sprintf("name='%v'", v.Name))
	}
	if v.OfflineEnabled {
		fields = append(fields, fmt.Sprintf("offline_enabled=%v", v.OfflineEnabled))
	}
	if v.OriginalFilename != "" {
		fields = append(fields, fmt.Sprintf("original_filename='%v'", v.OriginalFilename))
	}
	if v.Projection != "" {
		fields = append(fields, fmt.Sprintf("projection='%v'", v.Projection))
	}
	if v.PublishedAt != "" {
		fields = append(fields, fmt.Sprintf("published_at=%v", v.PublishedAt))
	}
	if v.ReferenceId != "" {
		fields = append(fields, fmt.Sprintf("reference_id=%v", v.ReferenceId))
	}
	if v.Schedule != nil {
		fields = append(fields, fmt.Sprintf("schedule=%v", v.Schedule))
	}
	if v.State != "" {
		fields = append(fields, fmt.Sprintf("state='%v'", v.State))
	}
	if len(v.Tags) > 0 {
		fields = append(fields, fmt.Sprintf("tags=%v", v.Tags))
	}
	if len(v.TextTracks) > 0 {
		fields = append(fields, fmt.Sprintf("text_tracks=%v", v.TextTracks))
	}
	if v.UpdatedAt != "" {
		fields = append(fields, fmt.Sprintf("updated_at=%v", v.UpdatedAt))
	}
	return fmt.Sprintf("<Video>{ id=%v %v }", v.Id, strings.Join(fields, " "))
}
