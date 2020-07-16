package raw

import "encoding/json"

type Exports struct {
	Type         string      `json:"__type"`
	RootView     View        `json:"rootView"`
	TracePolicy  string      `json:"tracePolicy"`
	Meta         Meta        `json:"meta"`
	CanonicalURL interface{} `json:"canonicalUrl"`
}

type View struct {
	Props Props `json:"props"`
}

type Props struct {
	CollectionToken string `json:"collectionToken"`
	RawSectionToken string `json:"rawSectionToken"`
	SectionToken    string `json:"sectionToken"`
	UserID          string `json:"userID"`
	UserVanity      string `json:"userVanity"`
	ViewerID        string `json:"viewerID"`
}

type Meta struct {
	Title string `json:"title"`
}

type GraphQLData struct {
	Data   Data    `json:"data"`
	Errors []Error `json:"errors"`
	Raw    []byte  `json:"-"`
}

type Error struct {
	Message        string      `json:"message"`
	Severity       string      `json:"severity"`
	Code           int64       `json:"code"`
	APIErrorCode   interface{} `json:"api_error_code"`
	Summary        string      `json:"summary"`
	Description    string      `json:"description"`
	IsSilent       bool        `json:"is_silent"`
	IsTransient    bool        `json:"is_transient"`
	RequiresReauth bool        `json:"requires_reauth"`
	AllowUserRetry bool        `json:"allow_user_retry"`
	DebugInfo      interface{} `json:"debug_info"`
	QueryPath      interface{} `json:"query_path"`
	FbtraceID      string      `json:"fbtrace_id"`
	WWWRequestID   string      `json:"www_request_id"`
	Path           []string    `json:"path"`
}

type ActiveCollections struct {
	Nodes []Node `json:"nodes"`
}

type ProfileFieldSection struct {
	SectionType   string        `json:"section_type"`
	ProfileFields ProfileFields `json:"profile_fields"`
}

type ProfileFields struct {
	Nodes []Node `json:"nodes"`
}

type Field struct {
	Title          TextClass       `json:"title"`
	TextContent    TextContent     `json:"text_content"`
	ListItemGroups []ListItemGroup `json:"list_item_groups"`
	PrivacyScope   interface{}     `json:"privacy_scope"`
	Icon           Icon            `json:"icon"`
}

type Icon struct {
	Height int64  `json:"height"`
	Scale  int64  `json:"scale"`
	URI    string `json:"uri"`
	Width  int64  `json:"width"`
}

type ListItemGroup struct {
	ListItems []ListItem `json:"list_items"`
}

type ListItem struct {
	HeadingType string    `json:"heading_type"`
	Text        TextClass `json:"text"`
}

type TextClass struct {
	Ranges []Range `json:"ranges"`
	Text   string  `json:"text"`
}

type Range struct {
	Entity Entity `json:"entity"`
	Length int64  `json:"length"`
	Offset int64  `json:"offset"`
}

type Entity struct {
	Typename     string  `json:"__typename"`
	ID           string  `json:"id"`
	URL          string  `json:"url"`
	ProfileURL   string  `json:"profile_url"`
	CategoryType string  `json:"category_type"`
	ExternalURL  *string `json:"external_url,omitempty"`
}

type CoverPhoto struct {
	Photo Photo `json:"photo"`
}

type Photo struct {
	ID           string `json:"id"`
	Image        Image  `json:"image"`
	BlurredImage Image  `json:"blurred_image"`
	URL          string `json:"url"`
}

type ProfilePic struct {
	URI string `json:"uri"`
}

type TimelineSections struct {
	Nodes []Node `json:"nodes"`
}

type YearOverview struct {
	Items Items `json:"items"`
}

type Title struct {
	Text           string `json:"text"`
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story"`
}

type AggregatedStorySection struct {
	AggregatedStories interface{} `json:"aggregated_stories"`
}

type ActorPhoto struct {
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story"`
}

type LargerProfilePic struct {
	URI string `json:"uri"`
}

type Timestamp struct {
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story"`
}

type Content struct {
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story,omitempty"`
}

type AllSubattachments struct {
	Count int64         `json:"count"`
	Nodes []NodeElement `json:"nodes"`
}

type AttachedStoryLayout struct {
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story"`
}

type CometSectionsMessage struct {
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story"`
}

type MessageContainer struct {
	Typename       string `json:"__typename"`
	IsProdEligible bool   `json:"is_prod_eligible"`
	Story          *Story `json:"story"`
}

type FeedbackContext struct {
	FeedbackTargetWithContext *Feedback `json:"feedback_target_with_context"`
}

type CommentCount struct {
	TotalCount int64 `json:"total_count"`
	Count      int64 `json:"count"`
}

type ReactionCount struct {
	Count int64 `json:"count"`
}

type SupportedReaction struct {
	Key int64 `json:"key"`
}

type TopReactions struct {
	Count   int64     `json:"count"`
	Summary []Summary `json:"summary"`
}

type TimelineListFeedUnits struct {
	Edges []TimelineListFeedUnitsEdge `json:"edges"`
}

type TimelineListFeedUnitsEdge struct {
	Node   Node   `json:"node"`
	Cursor string `json:"cursor"`
}

type Author struct {
	Typename             string           `json:"__typename"`
	ID                   string           `json:"id"`
	Name                 string           `json:"name"`
	ProfilePictureDepth0 LargerProfilePic `json:"profile_picture_depth_0"`
	URL                  string           `json:"url"`
	IsVerified           bool             `json:"is_verified,omitempty"`
	ShortName            string           `json:"short_name,omitempty"`
	UserAvatar           interface{}      `json:"user_avatar"`
}

type Body struct {
	Text   string        `json:"text"`
	Ranges []interface{} `json:"ranges"`
}

type FeedbackDisplayComments struct {
	HighlightedComments              []interface{} `json:"highlighted_comments"`
	CommentOrder                     string        `json:"comment_order"`
	IsInitiallyExpanded              bool          `json:"is_initially_expanded"`
	PageSize                         int64         `json:"page_size"`
	ReplyCommentOrder                string        `json:"reply_comment_order"`
	ShouldRenderComposerPreemptively bool          `json:"should_render_composer_preemptively"`
	AfterCount                       int64         `json:"after_count"`
	BeforeCount                      int64         `json:"before_count"`
	Count                            int64         `json:"count"`
	Edges                            []Edge        `json:"edges"`
	PageInfo                         *PageInfo     `json:"page_info"`
}

type PageInfo struct {
	EndCursor       string `json:"end_cursor"`
	HasNextPage     bool   `json:"has_next_page"`
	HasPreviousPage bool   `json:"has_previous_page"`
	StartCursor     string `json:"start_cursor"`
}

type Reactors struct {
	Count        int64  `json:"count"`
	CountReduced string `json:"count_reduced"`
}

type Result struct {
	Data    Data    `json:"data"`
	Type    string  `json:"type"`
	Exports Exports `json:"exports"`
}

type UserIDInfo struct {
	UserID                          string `json:"USER_ID"`
	AccountID                       string `json:"ACCOUNT_ID"`
	Name                            string `json:"NAME"`
	ShortName                       string `json:"SHORT_NAME"`
	IsMessengerOnlyUser             bool   `json:"IS_MESSENGER_ONLY_USER"`
	IsDeactivatedAllowedOnMessenger bool   `json:"IS_DEACTIVATED_ALLOWED_ON_MESSENGER"`
	AppID                           string `json:"APP_ID"`
}

type RequireBbox struct {
	Bbox Bbox `json:"__bbox"`
}

type Bbox struct {
	Result         Result `json:"result"`
	SequenceNumber int64  `json:"sequence_number"`
}

type Data struct {
	LoginData                         LoginData              `json:"login_data"`
	User                              User                   `json:"user"`
	HasTributesSection                bool                   `json:"has_tributes_section"`
	TimelineNavAppSections            TimelineNavAppSections `json:"timeline_nav_app_sections"`
	IsViewerFriend                    bool                   `json:"is_viewer_friend"`
	ActiveCollections                 ActiveCollections      `json:"activeCollections"`
	Node                              Node                   `json:"node"`
	Feedback                          *Feedback              `json:"feedback"`
	URL                               string                 `json:"url"`
	CurrMedia                         *Media                 `json:"currMedia"`
	Mediaset                          *Mediaset              `json:"mediaset"`
	Cursor                            string                 `json:"cursor"`
	CommentCount                      *CommentCount          `json:"comment_count"`
	CanViewerComment                  bool                   `json:"can_viewer_comment"`
	TopLevelCommentListRenderer       *Renderer              `json:"top_level_comment_list_renderer"`
	CometUFISummaryAndActionsRenderer *Renderer              `json:"comet_ufi_summary_and_actions_renderer"`
}

type User struct {
	ID string `json:"id"`

	Name                   string                 `json:"name"`
	Typename               string                 `json:"__typename"`
	URL                    string                 `json:"url"`
	IsViewerFriend         bool                   `json:"is_viewer_friend"`
	CoverPhoto             CoverPhoto             `json:"cover_photo"`
	ProfilePic160          ProfilePic             `json:"profilePic160"`
	AlternateName          string                 `json:"alternate_name"`
	IsVerified             bool                   `json:"is_verified"`
	IsVisiblyMemorialized  bool                   `json:"is_visibly_memorialized"`
	TimelineSections       TimelineSections       `json:"timeline_sections"`
	TimelineListFeedUnits  TimelineListFeedUnits  `json:"timeline_list_feed_units"`
	TimelineNavAppSections TimelineNavAppSections `json:"timeline_nav_app_sections"`
}

type Ections struct {
	Nodes []Node `json:"nodes"`
}

type TimelineNavAppSections struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
}

type Edge struct {
	I18NReactionCount string `json:"i18n_reaction_count,omitempty"`
	Node              *Node  `json:"node,omitempty"`
	ReactionCount     int64  `json:"reaction_count,omitempty"`
	Cursor            string `json:"cursor,omitempty"`
}

type AllCollections struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type NodeElement struct {
	URL              string      `json:"url"`
	ID               string      `json:"id"`
	DeduplicationKey string      `json:"deduplication_key"`
	Media            Media       `json:"media"`
	StyleRenderer    *Renderer   `json:"style_renderer"`
	FirstCardToShow  interface{} `json:"first_card_to_show"`
}

type Items struct {
	Count    int64    `json:"count"`
	Nodes    []Node   `json:"nodes"`
	Edges    []Edge   `json:"edges"`
	PageInfo PageInfo `json:"page_info"`
}

type TextContent struct {
	ImageRanges       []interface{} `json:"image_ranges"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges"`
	AggregatedRanges  []interface{} `json:"aggregated_ranges"`
	Ranges            []interface{} `json:"ranges"`
	Text              string        `json:"text"`
}

type QueryID struct {
	ActorID     string `json:"actorID"`
	PreloaderID string `json:"preloaderID"`
	QueryID     string `json:"queryID"`
}

type FBToken1 struct {
	ServerRevision int    `json:"server_revision"`
	ClientRevision int    `json:"client_revision"`
	Tier           string `json:"tier"`
	PushPhase      string `json:"push_phase"`
	PkgCohort      string `json:"pkg_cohort"`
	PR             int    `json:"pr"`
	HasteSite      string `json:"haste_site"`
	BeOneAhead     bool   `json:"be_one_ahead"`
	IROn           bool   `json:"ir_on"`
	IsRTL          bool   `json:"is_rtl"`
	IsComet        bool   `json:"is_comet"`
	Hsi            string `json:"hsi"`
	Spin           int    `json:"spin"`
	SpinR          int    `json:"__spin_r"`
	SpinB          string `json:"__spin_b"`
	SpinT          int    `json:"__spin_t"`
	Vip            string `json:"vip"`
}

//FBToken2 contain fb_dtsg
type FBToken2 struct {
	Require [][]json.RawMessage `json:"require"`
}

type LoginData struct {
	PrefillContactpoint  interface{} `json:"prefill_contactpoint"`
	PrefillSource        interface{} `json:"prefill_source"`
	IddUserCryptedUid    interface{} `json:"idd_user_crypted_uid"`
	Locale               string      `json:"locale"`
	Lsd                  Jazoest     `json:"lsd"`
	Jazoest              Jazoest     `json:"jazoest"`
	LoginPostURI         string      `json:"login_post_uri"`
	AbTestingEnabled     bool        `json:"ab_testing_enabled"`
	LoginSource          string      `json:"login_source"`
	Timestamp            int64       `json:"timestamp"`
	Lgnrnd               string      `json:"lgnrnd"`
	SendScreenDimensions bool        `json:"send_screen_dimensions"`
	ResetURI             string      `json:"reset_uri"`
	SketchSeed1          string      `json:"sketch_seed1"`
	SketchSeed2          string      `json:"sketch_seed2"`
	Rounds               int64       `json:"rounds"`
}

type Jazoest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Collection struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	Items        *Items `json:"items"`
	NullStateMsg string `json:"null_state_msg"`
}

type ActionsRendererProfileAction struct {
	Typename      string        `json:"__typename"`
	IsActive      bool          `json:"is_active"`
	ClientHandler ClientHandler `json:"client_handler"`
	ID            string        `json:"id"`
}

type ClientHandler struct {
	Typename      string                     `json:"__typename"`
	ProfileAction ClientHandlerProfileAction `json:"profile_action"`
}

type ClientHandlerProfileAction struct {
	Typename                 string                   `json:"__typename"`
	RestrictableProfileOwner RestrictableProfileOwner `json:"restrictable_profile_owner"`
	ID                       string                   `json:"id"`
}

type RestrictableProfileOwner struct {
	Typename  string `json:"__typename"`
	ID        string `json:"id"`
	Gender    string `json:"gender"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Summary struct {
	ReactionCount        int64    `json:"reaction_count"`
	ReactionCountReduced string   `json:"reaction_count_reduced"`
	Reaction             Reaction `json:"reaction"`
}

type Reaction struct {
	ReactionType  string `json:"reaction_type"`
	LocalizedName string `json:"localized_name"`
	Color         string `json:"color"`
	ID            string `json:"id"`
}

type PhotoImage struct {
	URI    string `json:"uri"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}

type RouteDefinitions struct {
	Payload struct {
		SrPayload struct {
			Bootloadable map[string]struct {
				R     []string            `json:"r"`
				Rdfds map[string][]string `json:"rdfds"`
				Rds   map[string][]string `json:"rds"`
				Be    int                 `json:"be"`
			} `json:"bootloadable"`

			ResourceMap map[string]struct {
				Type string `json:"type"`
				Src  string `json:"src"`
				P    string `json:"p"`
			} `json:"resource_map"`
		} `json:"sr_payload"`
	} `json:"payload"`
}

type ContainerStory struct {
	ID                string           `json:"id"`
	PostID            string           `json:"post_id"`
	EncryptedTracking string           `json:"encrypted_tracking"`
	ViewabilityConfig []int64          `json:"viewability_config"`
	ClientViewConfig  ClientViewConfig `json:"client_view_config"`
}

type ClientViewConfig struct {
	CanDelayLogImpression bool `json:"can_delay_log_impression"`
	UseBanzaiSignalImp    bool `json:"use_banzai_signal_imp"`
	UseBanzaiVitalImp     bool `json:"use_banzai_vital_imp"`
}

type CreationStoryActor struct {
	Typename              string      `json:"__typename"`
	ID                    string      `json:"id"`
	URL                   string      `json:"url"`
	WorkForeignEntityInfo interface{} `json:"work_foreign_entity_info"`
	Name                  string      `json:"name"`
}

type ModuleSectionStory struct {
	DR string `json:"__dr"`
}

type ActorPhotoStory struct {
	Actors []StoryActor `json:"actors"`
	ID     string       `json:"id"`
}

type StoryActor struct {
	Typename       string      `json:"__typename"`
	ID             string      `json:"id"`
	ProfileURL     string      `json:"profile_url"`
	StoryBucket    StoryBucket `json:"story_bucket"`
	URL            string      `json:"url"`
	Name           string      `json:"name"`
	ProfilePicture Image       `json:"profile_picture"`
}

type StoryBucket struct {
	Nodes []NodeElement `json:"nodes"`
}

type Audience struct {
	Typename       string        `json:"__typename"`
	IsProdEligible bool          `json:"is_prod_eligible"`
	Story          AudienceStory `json:"story"`
}

type AudienceStory struct {
	PrivacyScope StoryPrivacyScope `json:"privacy_scope"`
	ID           string            `json:"id"`
}

type StoryPrivacyScope struct {
	IconImage   IconImage `json:"icon_image"`
	Description string    `json:"description"`
}

type IconImage struct {
	Name string `json:"name"`
}

type TimestampStory struct {
	CreationTime int64       `json:"creation_time"`
	URL          string      `json:"url"`
	GhlLabel     interface{} `json:"ghl_label"`
	ID           string      `json:"id"`
}

type DefaultMediaset struct {
	Typename  string `json:"__typename"`
	PrevMedia Tags   `json:"prevMedia"`
	NextMedia Tags   `json:"nextMedia"`
	ID        string `json:"id"`
}

type Tags struct {
	Edges []TagsEdge `json:"edges"`
}

type TagsEdge struct {
	Node PurpleNode `json:"node"`
}

type PurpleNode struct {
	Typename   string `json:"__typename"`
	IsPlayable bool   `json:"is_playable"`
	ID         string `json:"id"`
}

type TopReactionsEdge struct {
	ReactionCount int64      `json:"reaction_count"`
	Node          FluffyNode `json:"node"`
}

type FluffyNode struct {
	Key int64  `json:"key"`
	ID  string `json:"id"`
}

type ViewerActor struct {
	Typename string `json:"__typename"`
	ID       string `json:"id"`
}

type ProfilePicture struct {
	URI string `json:"uri"`
}

type WelcomePrivacyScope struct {
	PrivacyScopeRenderer PrivacyScopeRenderer `json:"privacy_scope_renderer"`
}

type PrivacyScopeRenderer struct {
	Typename        string          `json:"__typename"`
	PrivacyRowInput PrivacyRowInput `json:"privacy_row_input"`
	Scope           Scope           `json:"scope"`
	ID              string          `json:"id"`
}

type PrivacyRowInput struct {
	Allow             []interface{} `json:"allow"`
	BaseState         string        `json:"base_state"`
	Deny              []interface{} `json:"deny"`
	TagExpansionState string        `json:"tag_expansion_state"`
}

type Scope struct {
	Label                   string         `json:"label"`
	IconImage               Image          `json:"icon_image"`
	SelectedOption          SelectedOption `json:"selected_option"`
	ShowTagExpansionOptions bool           `json:"show_tag_expansion_options"`
	CanViewerEdit           bool           `json:"can_viewer_edit"`
	Description             string         `json:"description"`
}

type SelectedOption struct {
	CurrentTagExpansion string `json:"current_tag_expansion"`
	ID                  string `json:"id"`
}
