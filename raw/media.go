package raw

type Media struct {
	Typename                            string                              `json:"__typename"`
	IsPlayable                          bool                                `json:"is_playable"`
	Image                               *Image                              `json:"image"`
	ID                                  string                              `json:"id"`
	AccessibilityCaption                string                              `json:"accessibility_caption"`
	PhotoImage                          *Image                              `json:"photo_image"`
	URL                                 string                              `json:"url"`
	IsLiveStreaming                     bool                                `json:"is_live_streaming"`
	IsLiveTraceEnabled                  bool                                `json:"is_live_trace_enabled"`
	IsLooping                           bool                                `json:"is_looping"`
	LoopCount                           int64                               `json:"loop_count"`
	IsSpherical                         bool                                `json:"is_spherical"`
	PermalinkURL                        string                              `json:"permalink_url"`
	PlayableURL                         string                              `json:"playable_url"`
	VideoID                             string                              `json:"videoId"`
	IsPremiere                          bool                                `json:"isPremiere"`
	LiveViewerCount                     int64                               `json:"liveViewerCount"`
	IsGamingVideo                       bool                                `json:"is_gaming_video"`
	PublishTime                         int64                               `json:"publish_time"`
	OriginalRotation                    string                              `json:"original_rotation"`
	CanViewerShare                      bool                                `json:"can_viewer_share"`
	CreationStory                       *CreationStory                      `json:"creation_story"`
	OwnerVideoChannel                   *Owner                              `json:"owner_video_channel"`
	Height                              int64                               `json:"height"`
	Width                               int64                               `json:"width"`
	Owner                               *Owner                              `json:"owner"`
	CreatedTime                         int64                               `json:"created_time"`
	PrivacyScope                        *WelcomePrivacyScope                `json:"privacy_scope"`
	Feedback                            *Feedback                           `json:"feedback"`
	CanViewerAddTags                    bool                                `json:"can_viewer_add_tags"`
	Tags                                *Tags                               `json:"tags"`
	ContainerStory                      *ContainerStory                     `json:"container_story"`
	DefaultMediaset                     *DefaultMediaset                    `json:"default_mediaset"`
	OriginalWidth                       int64                               `json:"original_width"`
	OriginalHeight                      int64                               `json:"original_height"`
	PlayableURLQualityHD                string                              `json:"playable_url_quality_hd"`
	DashPrefetchExperimental            []string                            `json:"dash_prefetch_experimental"`
	CanUseOz                            bool                                `json:"can_use_oz"`
	PlayableURLDash                     string                              `json:"playable_url_dash"`
	AutoplayGatingResult                string                              `json:"autoplay_gating_result"`
	ViewerAutoplaySetting               string                              `json:"viewer_autoplay_setting"`
	CanAutoplay                         bool                                `json:"can_autoplay"`
	PreferredThumbnail                  PreferredThumbnail                  `json:"preferred_thumbnail"`
	VideoCards                          VideoCards                          `json:"video_cards"`
	LiveRewindEnabled                   bool                                `json:"live_rewind_enabled"`
	ScrubberPreviewThumbnailInformation ScrubberPreviewThumbnailInformation `json:"scrubber_preview_thumbnail_information"`
	BreakingStatus                      bool                                `json:"breakingStatus"`
	Edges                               []struct {
		Node *Media `json:"node"`
	} `json:"edges"`
}

type CreationStory struct {
	ID                    string               `json:"id"`
	Actors                []CreationStoryActor `json:"actors"`
	CometSections         CometSections        `json:"comet_sections"`
	EncryptedTracking     string               `json:"encrypted_tracking"`
	CanViewerEdit         bool                 `json:"can_viewer_edit"`
	LegacyStoryHideableID string               `json:"legacy_story_hideable_id"`
	Shareable             Shareable            `json:"shareable"`
}

type Owner struct {
	Typename       string         `json:"__typename"`
	ID             string         `json:"id"`
	ProfilePicture ProfilePicture `json:"profile_picture"`
	Name           string         `json:"name"`
}

type PreferredThumbnail struct {
	Image               Image       `json:"image"`
	ImagePreviewPayload interface{} `json:"image_preview_payload"`
	ID                  string      `json:"id"`
}

type ScrubberPreviewThumbnailInformation struct {
	SpriteUris                 []interface{} `json:"sprite_uris"`
	ThumbnailWidth             int64         `json:"thumbnail_width"`
	ThumbnailHeight            int64         `json:"thumbnail_height"`
	HasPreviewThumbnails       bool          `json:"has_preview_thumbnails"`
	NumImagesPerRow            int64         `json:"num_images_per_row"`
	MaxNumberOfImagesPerSprite int64         `json:"max_number_of_images_per_sprite"`
	TimeIntervalBetweenImage   int64         `json:"time_interval_between_image"`
}

type Mediaset struct {
	Typename          string         `json:"__typename"`
	PrevMediaNoCursor *MediaNoCursor `json:"prevMediaNoCursor"`
	NextMediaNoCursor *MediaNoCursor `json:"nextMediaNoCursor"`
	CurrMedia         *Media         `json:"currMedia"`
	ID                string         `json:"id"`
}

type MediaNoCursor struct {
	Typename   string `json:"__typename"`
	IsPlayable bool   `json:"is_playable"`
	ID         string `json:"id"`
}

type Shareable struct {
	Typename string `json:"__typename"`
	WWWURL   string `json:"wwwUrl"`
	ID       string `json:"id"`
}

type Image struct {
	URI    string `json:"uri"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type VideoCards struct {
	Nodes []interface{} `json:"nodes"`
}
