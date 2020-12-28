// Package raw contains raw Facebook JSON data
package raw

type Edge struct {
	Node *Node `json:"node,omitempty"`
}

type Collection struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type WemPrivateSharingBundle struct {
	PrivateSharingControlModelForUser *PrivateSharingControl `json:"private_sharing_control_model_for_user,omitempty"`
}

type PrivateSharingControl struct {
	PrivateSharingEnabled bool `json:"private_sharing_enabled,omitempty"`
}

type StoryBucket struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type StyleRenderer struct {
	Typename                                                    string                                                      `json:"__typename,omitempty"`
	ProfileFieldSections                                        []ProfileFieldSection                                       `json:"profile_field_sections,omitempty"`
	ModuleOperationProfileCometAboutAppSectionContentAppSection ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_operation_ProfileCometAboutAppSectionContent_appSection,omitempty"`
	ModuleComponentProfileCometAboutAppSectionContentAppSection ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_component_ProfileCometAboutAppSectionContent_appSection,omitempty"`
}

type ModuleComponentProfileCometAboutAppSectionContentAppSection struct {
	DR string `json:"__dr,omitempty"`
}

type ProfileFieldSection struct {
	ID               string                    `json:"id,omitempty"`
	Title            *ProfileFieldSectionTitle `json:"title,omitempty"`
	Subtitle         interface{}               `json:"subtitle,omitempty"`
	FieldSectionType string                    `json:"field_section_type,omitempty"`
	ProfileFields    *ProfileFields            `json:"profile_fields,omitempty"`
}

type ProfileFields struct {
	Nodes []ProfileFieldsNode `json:"nodes,omitempty"`
}

type ProfileFieldsNode struct {
	Title                 TextClass        `json:"title,omitempty"`
	FieldType             string           `json:"field_type,omitempty"`
	ListItemGroups        []ListItemGroup  `json:"list_item_groups,omitempty"`
	LinkURL               interface{}      `json:"link_url,omitempty"`
	GroupKey              interface{}      `json:"group_key,omitempty"`
	Icon                  *Icon            `json:"icon,omitempty"`
	PrivacyScope          NodePrivacyScope `json:"privacy_scope,omitempty"`
	Renderer              *Renderer        `json:"renderer,omitempty"`
	EditRenderer          EditRenderer     `json:"edit_renderer,omitempty"`
	PublishOptionRenderer interface{}      `json:"publish_option_renderer,omitempty"`
}

type EditRenderer struct {
	Typename                                                               string                                                      `json:"__typename,omitempty"`
	DeleteLabel                                                            string                                                      `json:"delete_label,omitempty"`
	EditLabel                                                              string                                                      `json:"edit_label,omitempty"`
	WorkExperience                                                         *Experience                                                 `json:"work_experience,omitempty"`
	PrivacyScope                                                           EditRendererPrivacyScope                                    `json:"privacy_scope,omitempty"`
	ModuleOperationProfileCometAboutProfileFieldSectionSectionEditRenderer ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_operation_ProfileCometAboutProfileFieldSection_section_edit_renderer,omitempty"`
	ModuleComponentProfileCometAboutProfileFieldSectionSectionEditRenderer ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_component_ProfileCometAboutProfileFieldSection_section_edit_renderer,omitempty"`
	EducationExperience                                                    *Experience                                                 `json:"education_experience,omitempty"`
}

type Experience struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

type EditRendererPrivacyScope struct {
	PrivacyScopeRenderer PurplePrivacyScopeRenderer `json:"privacy_scope_renderer,omitempty"`
}

type PurplePrivacyScopeRenderer struct {
	Typename                  string                    `json:"__typename,omitempty"`
	IsPrivacySelectorRenderer string                    `json:"__isPrivacySelectorRenderer,omitempty"`
	PrivacyRowInput           PrivacyRowInput           `json:"privacy_row_input,omitempty"`
	Scope                     PrivacyScopeRendererScope `json:"scope,omitempty"`
	ID                        string                    `json:"id,omitempty"`
	EntryPointRenderer        PurpleEntryPointRenderer  `json:"entry_point_renderer,omitempty"`
}

type PurpleEntryPointRenderer struct {
	Typename string       `json:"__typename,omitempty"`
	Source   PurpleSource `json:"source,omitempty"`
}

type PurpleSource struct {
	Typename string      `json:"__typename,omitempty"`
	Scope    PurpleScope `json:"scope,omitempty"`
	ID       string      `json:"id,omitempty"`
}

type PurpleScope struct {
	Label                   string               `json:"label,omitempty"`
	IconImage               Icon                 `json:"icon_image,omitempty"`
	SelectedOption          PurpleSelectedOption `json:"selected_option,omitempty"`
	ShowTagExpansionOptions bool                 `json:"show_tag_expansion_options,omitempty"`
	CanViewerEdit           bool                 `json:"can_viewer_edit,omitempty"`
	Description             string               `json:"description,omitempty"`
	ExtendedDescription     string               `json:"extended_description,omitempty"`
}

type Icon struct {
	Height int     `json:"height,omitempty"`
	Scale  float64 `json:"scale,omitempty"`
	URI    string  `json:"uri,omitempty"`
	Width  int     `json:"width,omitempty"`
}

type PurpleSelectedOption struct {
	CurrentTagExpansion string `json:"current_tag_expansion,omitempty"`
	ID                  string `json:"id,omitempty"`
}

type PrivacyRowInput struct {
	Allow             []interface{} `json:"allow,omitempty"`
	BaseState         string        `json:"base_state,omitempty"`
	Deny              []interface{} `json:"deny,omitempty"`
	PrivacyTargeting  interface{}   `json:"privacy_targeting,omitempty"`
	TagExpansionState string        `json:"tag_expansion_state,omitempty"`
}

type PrivacyScopeRendererScope struct {
	SelectedRowOverride interface{}          `json:"selected_row_override,omitempty"`
	SelectedOption      FluffySelectedOption `json:"selected_option,omitempty"`
	CanViewerEdit       bool                 `json:"can_viewer_edit,omitempty"`
	PrivacyWriteID      string               `json:"privacy_write_id,omitempty"`
}

type FluffySelectedOption struct {
	PrivacyRowInput PrivacyRowInput `json:"privacy_row_input,omitempty"`
	ID              string          `json:"id,omitempty"`
}

type ListItemGroup struct {
	ListItems []ListItem `json:"list_items,omitempty"`
}

type ListItem struct {
	HeadingType string    `json:"heading_type,omitempty"`
	Text        TextClass `json:"text,omitempty"`
}

type TextClass struct {
	DelightRanges     []interface{} `json:"delight_ranges,omitempty"`
	ImageRanges       []interface{} `json:"image_ranges,omitempty"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges,omitempty"`
	AggregatedRanges  []interface{} `json:"aggregated_ranges,omitempty"`
	Ranges            []Range       `json:"ranges,omitempty"`
	ColorRanges       []interface{} `json:"color_ranges,omitempty"`
	Text              string        `json:"text,omitempty"`
}

type Range struct {
	Entity                Entity `json:"entity,omitempty"`
	EntityIsWeakReference bool   `json:"entity_is_weak_reference,omitempty"`
	Length                int    `json:"length,omitempty"`
	Offset                int    `json:"offset,omitempty"`
}

type Entity struct {
	Typename                                                  string                                                      `json:"__typename,omitempty"`
	IsEntity                                                  string                                                      `json:"__isEntity,omitempty"`
	IsActor                                                   string                                                      `json:"__isActor,omitempty"`
	ID                                                        string                                                      `json:"id,omitempty"`
	URL                                                       string                                                      `json:"url,omitempty"`
	CometURL                                                  string                                                      `json:"comet_url,omitempty"`
	ModuleOperationCometTextWithEntitiesRelayTextWithEntities ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_operation_CometTextWithEntitiesRelay_textWithEntities,omitempty"`
	ModuleComponentCometTextWithEntitiesRelayTextWithEntities ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_component_CometTextWithEntitiesRelay_textWithEntities,omitempty"`
	CategoryType                                              string                                                      `json:"category_type,omitempty"`
	VerificationStatus                                        string                                                      `json:"verification_status,omitempty"`
	IsVerified                                                bool                                                        `json:"is_verified,omitempty"`
	ProfileURL                                                string                                                      `json:"profile_url,omitempty"`
	MobileURL                                                 string                                                      `json:"mobileUrl,omitempty"`
	IsNode                                                    string                                                      `json:"__isNode,omitempty"`
}

type NodePrivacyScope struct {
	PrivacyScopeRenderer FluffyPrivacyScopeRenderer `json:"privacy_scope_renderer,omitempty"`
}

type FluffyPrivacyScopeRenderer struct {
	Typename                  string                    `json:"__typename,omitempty"`
	IsPrivacySelectorRenderer string                    `json:"__isPrivacySelectorRenderer,omitempty"`
	PrivacyRowInput           PrivacyRowInput           `json:"privacy_row_input,omitempty"`
	Scope                     PrivacyScopeRendererScope `json:"scope,omitempty"`
	ID                        string                    `json:"id,omitempty"`
	EntryPointRenderer        FluffyEntryPointRenderer  `json:"entry_point_renderer,omitempty"`
}

type FluffyEntryPointRenderer struct {
	Typename string       `json:"__typename,omitempty"`
	Source   FluffySource `json:"source,omitempty"`
}

type FluffySource struct {
	Typename string      `json:"__typename,omitempty"`
	Scope    FluffyScope `json:"scope,omitempty"`
	ID       string      `json:"id,omitempty"`
}

type FluffyScope struct {
	Label               string `json:"label,omitempty"`
	IconImage           Icon   `json:"icon_image,omitempty"`
	CanViewerEdit       bool   `json:"can_viewer_edit,omitempty"`
	Description         string `json:"description,omitempty"`
	ExtendedDescription string `json:"extended_description,omitempty"`
}

type Renderer struct {
	Typename                                                   string                                                      `json:"__typename,omitempty"`
	TargetFieldType                                            string                                                      `json:"target_field_type,omitempty"`
	Title                                                      *TextClass                                                  `json:"title,omitempty"`
	ModuleOperationProfileCometAboutProfileFieldSectionSection ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_operation_ProfileCometAboutProfileFieldSection_section,omitempty"`
	ModuleComponentProfileCometAboutProfileFieldSectionSection ModuleComponentProfileCometAboutAppSectionContentAppSection `json:"__module_component_ProfileCometAboutProfileFieldSection_section,omitempty"`
	Field                                                      *Field                                                      `json:"field,omitempty"`
}

type Field struct {
	Title          TextClass        `json:"title,omitempty"`
	TextContent    interface{}      `json:"text_content,omitempty"`
	ListItemGroups []ListItemGroup  `json:"list_item_groups,omitempty"`
	PrivacyScope   NodePrivacyScope `json:"privacy_scope,omitempty"`
	Icon           Icon             `json:"icon,omitempty"`
}

type ProfileFieldSectionTitle struct {
	Text string `json:"text,omitempty"`
}

type BanzaiScubaDEPRECATED struct {
	R   []string                 `json:"r,omitempty"`
	RDS BanzaiScubaDEPRECATEDRDS `json:"rds,omitempty"`
}

type BanzaiScubaDEPRECATEDRDS struct {
	M []string `json:"m,omitempty"`
}

type BladeRunnerClient struct {
	R   []string             `json:"r,omitempty"`
	RDS BladeRunnerClientRDS `json:"rds,omitempty"`
}

type BladeRunnerClientRDS struct {
	M []string `json:"m,omitempty"`
	R []string `json:"r,omitempty"`
}

type BladeRunnerStreamHandler struct {
	R  []string `json:"r,omitempty"`
	Be int      `json:"be,omitempty"`
}

type CometCompatModalReact struct {
	R   []string             `json:"r,omitempty"`
	RDS BladeRunnerClientRDS `json:"rds,omitempty"`
	Be  int                  `json:"be,omitempty"`
}

type FbtLogging struct {
	R []string `json:"r,omitempty"`
}

type Consistency struct {
	Rev int `json:"rev,omitempty"`
}

type RsrcMap struct {
	Type    string `json:"type,omitempty"`
	Src     string `json:"src,omitempty"`
	Nc      int    `json:"nc,omitempty"`
	P       string `json:"p,omitempty"`
	Prelude int    `json:"prelude,omitempty"`
}

type Hsdp struct {
	IxData  map[string]IxDatum  `json:"ixData,omitempty"`
	ClpData map[string]ClpDatum `json:"clpData,omitempty"`
	GkxData map[string]GkxDatum `json:"gkxData,omitempty"`
	QexData map[string]QexDatum `json:"qexData,omitempty"`
}

type ClpDatum struct {
	R int `json:"r,omitempty"`
	S int `json:"s,omitempty"`
}

type GkxDatum struct {
	Result bool   `json:"result,omitempty"`
	Hash   string `json:"hash,omitempty"`
}

type IxDatum struct {
	Sprited           int    `json:"sprited,omitempty"`
	SpriteCSSClass    string `json:"spriteCssClass,omitempty"`
	SpriteMapCSSClass string `json:"spriteMapCssClass,omitempty"`
	SPI               string `json:"_spi,omitempty"`
}

type QexDatum struct {
	R *RUnion `json:"r,omitempty"`
}

type Jsmods struct {
	Define  [][]DefineElement `json:"define,omitempty"`
	Require [][]JsmodsRequire `json:"require,omitempty"`
}

type DefineClass struct {
	RC                                    []string                    `json:"__rc,omitempty"`
	DeferBootloads                        bool                        `json:"deferBootloads,omitempty"`
	HighPriBootloads                      bool                        `json:"highPriBootloads,omitempty"`
	JSRetries                             []int                       `json:"jsRetries,omitempty"`
	JSRetryAbortNum                       int                         `json:"jsRetryAbortNum,omitempty"`
	JSRetryAbortTime                      int                         `json:"jsRetryAbortTime,omitempty"`
	RetryQueuedBootloads                  bool                        `json:"retryQueuedBootloads,omitempty"`
	SilentDups                            bool                        `json:"silentDups,omitempty"`
	Timeout                               int                         `json:"timeout,omitempty"`
	ModulePrefix                          string                      `json:"modulePrefix,omitempty"`
	Facebookdotcom                        bool                        `json:"facebookdotcom,omitempty"`
	Messengerdotcom                       bool                        `json:"messengerdotcom,omitempty"`
	Workplacedotcom                       bool                        `json:"workplacedotcom,omitempty"`
	Token                                 string                      `json:"token,omitempty"`
	ServerNonce                           string                      `json:"ServerNonce,omitempty"`
	ServerRevision                        int                         `json:"server_revision,omitempty"`
	ClientRevision                        int                         `json:"client_revision,omitempty"`
	Tier                                  string                      `json:"tier,omitempty"`
	PushPhase                             string                      `json:"push_phase,omitempty"`
	PkgCohort                             string                      `json:"pkg_cohort,omitempty"`
	PR                                    float64                     `json:"pr,omitempty"`
	HasteSite                             string                      `json:"haste_site,omitempty"`
	BeOneAhead                            bool                        `json:"be_one_ahead,omitempty"`
	IROn                                  bool                        `json:"ir_on,omitempty"`
	IsRTL                                 bool                        `json:"is_rtl,omitempty"`
	IsComet                               bool                        `json:"is_comet,omitempty"`
	IsExperimentalTier                    bool                        `json:"is_experimental_tier,omitempty"`
	IsJITWarmedUp                         bool                        `json:"is_jit_warmed_up,omitempty"`
	Hsi                                   string                      `json:"hsi,omitempty"`
	SemrHostBucket                        string                      `json:"semr_host_bucket,omitempty"`
	Spin                                  int                         `json:"spin,omitempty"`
	SpinR                                 int                         `json:"__spin_r,omitempty"`
	SpinB                                 string                      `json:"__spin_b,omitempty"`
	SpinT                                 int                         `json:"__spin_t,omitempty"`
	CometEnv                              string                      `json:"comet_env,omitempty"`
	Vip                                   string                      `json:"vip,omitempty"`
	ParamName                             string                      `json:"param_name,omitempty"`
	Version                               int                         `json:"version,omitempty"`
	ShouldRandomize                       bool                        `json:"should_randomize,omitempty"`
	WWWAlwaysUsePolyfillSetimmediate      bool                        `json:"www_always_use_polyfill_setimmediate,omitempty"`
	Killed                                *Killed                     `json:"killed,omitempty"`
	Ko                                    *Killed                     `json:"ko,omitempty"`
	PurpleAppID                           int                         `json:"appId,omitempty"`
	Extra                                 []interface{}               `json:"extra,omitempty"`
	ReportInterval                        int                         `json:"reportInterval,omitempty"`
	SampleWeight                          interface{}                 `json:"sampleWeight,omitempty"`
	SampleWeightKey                       string                      `json:"sampleWeightKey,omitempty"`
	PreferMessageChannel                  bool                        `json:"prefer_message_channel,omitempty"`
	AsyncGetToken                         string                      `json:"async_get_token,omitempty"`
	Uris                                  []string                    `json:"uris,omitempty"`
	DeferCookies                          bool                        `json:"deferCookies,omitempty"`
	InitialConsent                        *InitialConsent             `json:"initialConsent,omitempty"`
	NoCookies                             bool                        `json:"noCookies,omitempty"`
	ShouldShowCookieBanner                bool                        `json:"shouldShowCookieBanner,omitempty"`
	UseTrustedTypes                       bool                        `json:"useTrustedTypes,omitempty"`
	ReportOnly                            bool                        `json:"reportOnly,omitempty"`
	ConnectionClass                       string                      `json:"connectionClass,omitempty"`
	AllowedDomains                        []string                    `json:"allowed_domains,omitempty"`
	DebugNoBatching                       bool                        `json:"debugNoBatching,omitempty"`
	EndpointURI                           string                      `json:"endpointURI,omitempty"`
	A11Y                                  *A11Y                       `json:"a11y,omitempty"`
	CUser                                 *A11Y                       `json:"c_user,omitempty"`
	Cppo                                  *Cppo                       `json:"cppo,omitempty"`
	Dpr                                   *Cppo                       `json:"dpr,omitempty"`
	FblCi                                 *Cppo                       `json:"fbl_ci,omitempty"`
	FblCS                                 *Cppo                       `json:"fbl_cs,omitempty"`
	FblSt                                 *Cppo                       `json:"fbl_st,omitempty"`
	IUser                                 *A11Y                       `json:"i_user,omitempty"`
	JSVer                                 *Cppo                       `json:"js_ver,omitempty"`
	Locale                                *Locale                     `json:"locale,omitempty"`
	MPixelRatio                           *Cppo                       `json:"m_pixel_ratio,omitempty"`
	Noscript                              *A11Y                       `json:"noscript,omitempty"`
	Presence                              *Cppo                       `json:"presence,omitempty"`
	Sfau                                  *A11Y                       `json:"sfau,omitempty"`
	Vpd                                   *Cppo                       `json:"vpd,omitempty"`
	Wd                                    *Cppo                       `json:"wd,omitempty"`
	XReferer                              *A11Y                       `json:"x-referer,omitempty"`
	XSrc                                  *Cppo                       `json:"x-src,omitempty"`
	ShouldReturnFbtResult                 bool                        `json:"shouldReturnFbtResult,omitempty"`
	InlineMode                            string                      `json:"inlineMode,omitempty"`
	Meta                                  *Meta                       `json:"meta,omitempty"`
	Patterns                              *Patterns                   `json:"patterns,omitempty"`
	Gender                                int                         `json:"GENDER,omitempty"`
	Language                              string                      `json:"language,omitempty"`
	DecimalSeparator                      string                      `json:"decimalSeparator,omitempty"`
	NumberDelimiter                       string                      `json:"numberDelimiter,omitempty"`
	MinDigitsForThousandsSeparator        int                         `json:"minDigitsForThousandsSeparator,omitempty"`
	StandardDecimalPatternInfo            *StandardDecimalPatternInfo `json:"standardDecimalPatternInfo,omitempty"`
	NumberingSystemData                   interface{}                 `json:"numberingSystemData,omitempty"`
	AccessToken                           string                      `json:"accessToken,omitempty"`
	ActorID                               string                      `json:"actorID,omitempty"`
	CustomHeaders                         *A11Y                       `json:"customHeaders,omitempty"`
	EnableNetworkLogger                   bool                        `json:"enableNetworkLogger,omitempty"`
	FetchTimeout                          int                         `json:"fetchTimeout,omitempty"`
	GraphBatchURI                         string                      `json:"graphBatchURI,omitempty"`
	GraphURI                              string                      `json:"graphURI,omitempty"`
	RetryDelays                           []int                       `json:"retryDelays,omitempty"`
	UseXController                        bool                        `json:"useXController,omitempty"`
	XhrEncoding                           interface{}                 `json:"xhrEncoding,omitempty"`
	SubscriptionTopicURI                  string                      `json:"subscriptionTopicURI,omitempty"`
	WithCredentials                       bool                        `json:"withCredentials,omitempty"`
	IsProductionEndpoint                  bool                        `json:"isProductionEndpoint,omitempty"`
	BrowserArchitecture                   string                      `json:"browserArchitecture,omitempty"`
	BrowserFullVersion                    interface{}                 `json:"browserFullVersion,omitempty"`
	BrowserMinorVersion                   interface{}                 `json:"browserMinorVersion,omitempty"`
	BrowserName                           string                      `json:"browserName,omitempty"`
	BrowserVersion                        interface{}                 `json:"browserVersion,omitempty"`
	DeviceName                            string                      `json:"deviceName,omitempty"`
	EngineName                            string                      `json:"engineName,omitempty"`
	EngineVersion                         interface{}                 `json:"engineVersion,omitempty"`
	PlatformArchitecture                  string                      `json:"platformArchitecture,omitempty"`
	PlatformName                          string                      `json:"platformName,omitempty"`
	PlatformVersion                       interface{}                 `json:"platformVersion,omitempty"`
	PlatformFullVersion                   interface{}                 `json:"platformFullVersion,omitempty"`
	OnDemandReferenceCounting             bool                        `json:"on_demand_reference_counting,omitempty"`
	OnDemandProfilingCounters             bool                        `json:"on_demand_profiling_counters,omitempty"`
	DefaultRate                           int                         `json:"default_rate,omitempty"`
	LiteDefaultRate                       int                         `json:"lite_default_rate,omitempty"`
	InteractionToLiteCoinflip             *InteractionToCoinflip      `json:"interaction_to_lite_coinflip,omitempty"`
	InteractionToCoinflip                 *InteractionToCoinflip      `json:"interaction_to_coinflip,omitempty"`
	EnableHeartbeat                       bool                        `json:"enable_heartbeat,omitempty"`
	MaxBlockMergeDuration                 int                         `json:"maxBlockMergeDuration,omitempty"`
	MaxBlockMergeDistance                 int                         `json:"maxBlockMergeDistance,omitempty"`
	EnableBanzaiStream                    bool                        `json:"enable_banzai_stream,omitempty"`
	UserTimingCoinflip                    int                         `json:"user_timing_coinflip,omitempty"`
	BanzaiStreamCoinflip                  int                         `json:"banzai_stream_coinflip,omitempty"`
	CompressionEnabled                    bool                        `json:"compression_enabled,omitempty"`
	RefCountingFix                        bool                        `json:"ref_counting_fix,omitempty"`
	RefCountingContFix                    bool                        `json:"ref_counting_cont_fix,omitempty"`
	AlsoRecordNewTimesliceFormat          bool                        `json:"also_record_new_timeslice_format,omitempty"`
	ForceAsyncRequestTracingOn            bool                        `json:"force_async_request_tracing_on,omitempty"`
	MaximumIgnorableStallMS               float64                     `json:"maximumIgnorableStallMs,omitempty"`
	SampleRate                            float64                     `json:"sampleRate,omitempty"`
	SampleRateClassic                     float64                     `json:"sampleRateClassic,omitempty"`
	SampleRateFastStale                   float64                     `json:"sampleRateFastStale,omitempty"`
	QplEvents                             map[string]QplEvent         `json:"qpl_events,omitempty"`
	Killswitch                            bool                        `json:"killswitch,omitempty"`
	Domains                               *DomainsUnion               `json:"domains,omitempty"`
	AdaptiveConfig                        *AdaptiveConfig             `json:"adaptive_config,omitempty"`
	Map                                   map[string]ClpDatum         `json:"map,omitempty"`
	Config                                *Config                     `json:"config,omitempty"`
	Autobot                               *A11Y                       `json:"autobot,omitempty"`
	Assimilator                           *A11Y                       `json:"assimilator,omitempty"`
	UnsubscribeRelease                    bool                        `json:"unsubscribe_release,omitempty"`
	BladerunnerWWWSandbox                 interface{}                 `json:"bladerunner_www_sandbox,omitempty"`
	DefineIsIntern                        bool                        `json:"is_intern,omitempty"`
	RewriteRules                          *A11Y                       `json:"rewrite_rules,omitempty"`
	Whitelist                             map[string]int              `json:"whitelist,omitempty"`
	HeartbeatInterval                     int                         `json:"heartbeat_interval,omitempty"`
	SandboxDomain                         interface{}                 `json:"sandboxDomain,omitempty"`
	IsIntern                              bool                        `json:"isIntern,omitempty"`
	EnableRequestStreamLogging            bool                        `json:"enableRequestStreamLogging,omitempty"`
	GcReleaseBufferSize                   int                         `json:"gc_release_buffer_size,omitempty"`
	Max                                   int                         `json:"max,omitempty"`
	ClientID                              string                      `json:"clientID,omitempty"`
	LiveQueryWebRelayKillList             []string                    `json:"liveQueryWebRelayKillList,omitempty"`
	FantailLogQueue                       interface{}                 `json:"FantailLogQueue,omitempty"`
	MqttPublishTimeoutMS                  int                         `json:"mqttPublishTimeoutMs,omitempty"`
	MethodToSamplingMultiplier            *A11Y                       `json:"methodToSamplingMultiplier,omitempty"`
	IsTestRunning                         bool                        `json:"isTestRunning,omitempty"`
	AuxiliaryServiceInfo                  *A11Y                       `json:"auxiliaryServiceInfo,omitempty"`
	TestPath                              interface{}                 `json:"testPath,omitempty"`
	ServerTime                            int                         `json:"serverTime,omitempty"`
	MaxSize                               int                         `json:"MAX_SIZE,omitempty"`
	MaxWait                               int                         `json:"MAX_WAIT,omitempty"`
	RestoreWait                           int                         `json:"RESTORE_WAIT,omitempty"`
	Blacklist                             []string                    `json:"blacklist,omitempty"`
	Gks                                   *Gks                        `json:"gks,omitempty"`
	RouteLimit                            int                         `json:"ROUTE_LIMIT,omitempty"`
	DataLimit                             int                         `json:"DATA_LIMIT,omitempty"`
	MaxPrefetchVideosNum                  int                         `json:"maxPrefetchVideosNum,omitempty"`
	ConsolidateFragmentedPrefetchRequest  bool                        `json:"consolidateFragmentedPrefetchRequest,omitempty"`
	FixPrefetchCacheAbort                 bool                        `json:"fixPrefetchCacheAbort,omitempty"`
	DisablePrefetchCache                  bool                        `json:"disablePrefetchCache,omitempty"`
	EnableGlobalSchedulerForPrefetch      bool                        `json:"enableGlobalSchedulerForPrefetch,omitempty"`
	PrefetchPriority                      int                         `json:"prefetchPriority,omitempty"`
	DisableShakaBandwidthEstimator        bool                        `json:"disableShakaBandwidthEstimator,omitempty"`
	SwitchPrefetchTaskToHighPriWhenPlayed bool                        `json:"switchPrefetchTaskToHighPriWhenPlayed,omitempty"`
	UseFetch                              bool                        `json:"useFetch,omitempty"`
	UserID                                string                      `json:"USER_ID,omitempty"`
	AccountID                             string                      `json:"ACCOUNT_ID,omitempty"`
	Name                                  string                      `json:"NAME,omitempty"`
	ShortName                             string                      `json:"SHORT_NAME,omitempty"`
	IsBusinessPersonAccount               bool                        `json:"IS_BUSINESS_PERSON_ACCOUNT,omitempty"`
	HasSecondaryBusinessPerson            bool                        `json:"HAS_SECONDARY_BUSINESS_PERSON,omitempty"`
	IsMessengerOnlyUser                   bool                        `json:"IS_MESSENGER_ONLY_USER,omitempty"`
	IsDeactivatedAllowedOnMessenger       bool                        `json:"IS_DEACTIVATED_ALLOWED_ON_MESSENGER,omitempty"`
	IsMessengerCallGuestUser              bool                        `json:"IS_MESSENGER_CALL_GUEST_USER,omitempty"`
	IsWorkMessengerCallGuestUser          bool                        `json:"IS_WORK_MESSENGER_CALL_GUEST_USER,omitempty"`
	AppID                                 string                      `json:"APP_ID,omitempty"`
	Overrides                             *Overrides                  `json:"overrides,omitempty"`
	Impl                                  string                      `json:"impl,omitempty"`
	DeviceID                              string                      `json:"device_id,omitempty"`
	FluffyAppID                           string                      `json:"app_id,omitempty"`
	EnableBladerunner                     bool                        `json:"enable_bladerunner,omitempty"`
	EnableQueue                           bool                        `json:"enable_queue,omitempty"`
	EnableLocalstorage                    bool                        `json:"enable_localstorage,omitempty"`
	EnableACK                             bool                        `json:"enable_ack,omitempty"`
	Fbid                                  string                      `json:"fbid,omitempty"`
	DefineAppID                           int                         `json:"appID,omitempty"`
	Endpoint                              string                      `json:"endpoint,omitempty"`
	PollingEndpoint                       string                      `json:"pollingEndpoint,omitempty"`
	SubscribedTopics                      []interface{}               `json:"subscribedTopics,omitempty"`
	Capabilities                          int                         `json:"capabilities,omitempty"`
	ClientCapabilities                    int                         `json:"clientCapabilities,omitempty"`
	ChatVisibility                        bool                        `json:"chatVisibility,omitempty"`
	InitialSetting                        string                      `json:"initialSetting,omitempty"`
	SupportsMetaReferrer                  bool                        `json:"supports_meta_referrer,omitempty"`
	DefaultMetaReferrerPolicy             string                      `json:"default_meta_referrer_policy,omitempty"`
	SwitchedMetaReferrerPolicy            string                      `json:"switched_meta_referrer_policy,omitempty"`
	NonLinkshimLnfbMode                   string                      `json:"non_linkshim_lnfb_mode,omitempty"`
	LinkReactDefaultHash                  string                      `json:"link_react_default_hash,omitempty"`
	UntrustedLinkDefaultHash              string                      `json:"untrusted_link_default_hash,omitempty"`
	LinkshimHost                          string                      `json:"linkshim_host,omitempty"`
	UseRelNoOpener                        bool                        `json:"use_rel_no_opener,omitempty"`
	AlwaysUseHTTPS                        bool                        `json:"always_use_https,omitempty"`
	OnionAlwaysShim                       bool                        `json:"onion_always_shim,omitempty"`
	MiddleClickRequiresEvent              bool                        `json:"middle_click_requires_event,omitempty"`
	WWWSafeJSMode                         string                      `json:"www_safe_js_mode,omitempty"`
	MSafeJSMode                           interface{}                 `json:"m_safe_js_mode,omitempty"`
	GhlParamLinkShim                      bool                        `json:"ghl_param_link_shim,omitempty"`
	ClickIDS                              []string                    `json:"click_ids,omitempty"`
	IsLinkshimSupported                   bool                        `json:"is_linkshim_supported,omitempty"`
	CurrentDomain                         string                      `json:"current_domain,omitempty"`
}

type A11Y struct {
}

type AdaptiveConfig struct {
	Interactions Interactions `json:"interactions,omitempty"`
	Qpl          Qpl          `json:"qpl,omitempty"`
	Modules      interface{}  `json:"modules,omitempty"`
	Events       interface{}  `json:"events,omitempty"`
}

type Interactions struct {
	Modules map[string]int     `json:"modules,omitempty"`
	Events  map[string]float64 `json:"events,omitempty"`
}

type Qpl struct {
	Modules A11Y `json:"modules,omitempty"`
	Events  A11Y `json:"events,omitempty"`
}

type Config struct {
	MaxSubscriptions              int                       `json:"max_subscriptions,omitempty"`
	WWWIdleUnsubscribeMinTimeMS   int                       `json:"www_idle_unsubscribe_min_time_ms,omitempty"`
	WWWIdleUnsubscribeTimesMS     WWWIdleUnsubscribeTimesMS `json:"www_idle_unsubscribe_times_ms,omitempty"`
	WWWUnevictableTopicRegexes    []string                  `json:"www_unevictable_topic_regexes,omitempty"`
	AutobotTiers                  AutobotTiers              `json:"autobot_tiers,omitempty"`
	MaxSubscriptionFlushBatchSize int                       `json:"max_subscription_flush_batch_size,omitempty"`
}

type AutobotTiers struct {
	Latest string `json:"latest,omitempty"`
	Intern string `json:"intern,omitempty"`
	Sb     string `json:"sb,omitempty"`
}

type WWWIdleUnsubscribeTimesMS struct {
	FeedbackLikeSubscribe            int `json:"feedback_like_subscribe,omitempty"`
	CommentLikeSubscribe             int `json:"comment_like_subscribe,omitempty"`
	FeedbackTypingSubscribe          int `json:"feedback_typing_subscribe,omitempty"`
	CommentCreateSubscribe           int `json:"comment_create_subscribe,omitempty"`
	VideoTipJarPaymentEventSubscribe int `json:"video_tip_jar_payment_event_subscribe,omitempty"`
}

type Cppo struct {
	T int `json:"t,omitempty"`
}

type DomainsClass struct {
	Map [][]Path `json:"__map,omitempty"`
}

type Gks struct {
	BoostedPagelikes          bool `json:"boosted_pagelikes,omitempty"`
	MercurySendErrorLogging   bool `json:"mercury_send_error_logging,omitempty"`
	PlatformOauthClientEvents bool `json:"platform_oauth_client_events,omitempty"`
	VisibilityTracking        bool `json:"visibility_tracking,omitempty"`
	Graphexplorer             bool `json:"graphexplorer,omitempty"`
	StickerSearchRanking      bool `json:"sticker_search_ranking,omitempty"`
}

type InitialConsent struct {
	Set []int `json:"__set,omitempty"`
}

type InteractionToCoinflip struct {
	AdsInterfacesInteraction int `json:"ADS_INTERFACES_INTERACTION,omitempty"`
	AdsPerfScenario          int `json:"ads_perf_scenario,omitempty"`
	AdsWaitTime              int `json:"ads_wait_time,omitempty"`
	Event                    int `json:"Event,omitempty"`
}

type Killed struct {
	Set []string `json:"__set,omitempty"`
}

type Overrides struct {
	The1_Abf9D2B57Eeb10432B6C440B381D95DB string `json:"1_abf9d2b57eeb10432b6c440b381d95db,omitempty"`
}

type Patterns struct {
	The039S039S string `json:"/\u0001(.*)('|&#039;)s\u0001(?:'|&#039;)s(.*)/,omitempty"`
	Empty       string `json:"/_\u0001([^\u0001]*)\u0001/,omitempty"`
}

type QplEvent struct {
	SampleRate int `json:"sampleRate,omitempty"`
}

type StandardDecimalPatternInfo struct {
	PrimaryGroupSize   int `json:"primaryGroupSize,omitempty"`
	SecondaryGroupSize int `json:"secondaryGroupSize,omitempty"`
}

type RUnion struct {
	Bool    bool
	Integer int
}

type DefineElement struct {
	DefineClass *DefineClass
	Integer     int
	String      string
	StringArray []string
}

type DomainsUnion struct {
	DomainsClass *DomainsClass
	StringArray  []string
}

type Path struct {
	Integer int
	String  string
}

type Locale struct {
	Cppo   *Cppo
	String string
}

type JsmodsRequire struct {
	String     string
	UnionArray []PurpleRequire
}

type PurpleRequire struct {
	String     string
	UnionArray []FluffyRequire
}

type FluffyRequire struct {
	ModuleComponentProfileCometAboutAppSectionContentAppSection *ModuleComponentProfileCometAboutAppSectionContentAppSection
	String                                                      string
}
