package raw

type Data struct {
	Name                 string            `json:"name,omitempty"`
	Gender               string            `json:"gender,omitempty"`
	User                 *User             `json:"user,omitempty"`
	Viewer               *Viewer           `json:"viewer,omitempty"`
	Nux                  interface{}       `json:"nux,omitempty"`
	IsProfile            string            `json:"__isProfile,omitempty"`
	ProfilePhoto         *Photo            `json:"profilePhoto,omitempty"`
	ProfilePicNormal     *Image            `json:"profilePicNormal,omitempty"`
	ProfilePicSmall      *Image            `json:"profilePicSmall,omitempty"`
	ProfileVideo         interface{}       `json:"profile_video,omitempty"`
	PrefillContactpoint  interface{}       `json:"prefill_contactpoint,omitempty"`
	LoginPostURI         string            `json:"login_post_uri,omitempty"`
	AbTestingEnabled     bool              `json:"ab_testing_enabled,omitempty"`
	ResetURI             string            `json:"reset_uri,omitempty"`
	SketchSeed1          interface{}       `json:"sketch_seed1,omitempty"`
	SketchSeed2          interface{}       `json:"sketch_seed2,omitempty"`
	Rounds               interface{}       `json:"rounds,omitempty"`
	PublicKey            string            `json:"public_key,omitempty"`
	KeyID                int               `json:"key_id,omitempty"`
	PrefillSource        interface{}       `json:"prefill_source,omitempty"`
	IddUserCryptedUID    interface{}       `json:"idd_user_crypted_uid,omitempty"`
	Locale               string            `json:"locale,omitempty"`
	Lsd                  *KeyVal           `json:"lsd,omitempty"`
	Jazoest              *KeyVal           `json:"jazoest,omitempty"`
	LoginSource          string            `json:"login_source,omitempty"`
	Timestamp            int               `json:"timestamp,omitempty"`
	Lgnrnd               string            `json:"lgnrnd,omitempty"`
	SendScreenDimensions bool              `json:"send_screen_dimensions,omitempty"`
	LoginData            *Data             `json:"login_data,omitempty"`
	ProfileIntroCard     *ProfileIntroCard `json:"profile_intro_card,omitempty"`
	ShouldUsePageRename  bool              `json:"should_use_page_rename,omitempty"`
	ActiveCollections    *Collection       `json:"activeCollections,omitempty"`
}

type Viewer struct {
}

type KeyVal struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ProfileIntroCard struct {
	Bio *Bio   `json:"bio,omitempty"`
	ID  string `json:"id,omitempty"`
}

type Bio struct {
	DelightRanges     []interface{} `json:"delight_ranges,omitempty"`
	ImageRanges       []interface{} `json:"image_ranges,omitempty"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges,omitempty"`
	AggregatedRanges  []interface{} `json:"aggregated_ranges,omitempty"`
	Ranges            []interface{} `json:"ranges,omitempty"`
	ColorRanges       []interface{} `json:"color_ranges,omitempty"`
	Text              string        `json:"text,omitempty"`
}
