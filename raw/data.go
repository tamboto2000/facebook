package raw

type Data struct {
	Name                 string            `json:"name"`
	Gender               string            `json:"gender"`
	User                 *User             `json:"user"`
	Viewer               *Viewer           `json:"viewer"`
	Nux                  interface{}       `json:"nux"`
	IsProfile            string            `json:"__isProfile"`
	ProfilePhoto         *Photo            `json:"profilePhoto"`
	ProfilePicNormal     *Image            `json:"profilePicNormal"`
	ProfilePicSmall      *Image            `json:"profilePicSmall"`
	ProfileVideo         interface{}       `json:"profile_video"`
	PrefillContactpoint  interface{}       `json:"prefill_contactpoint"`
	LoginPostURI         string            `json:"login_post_uri"`
	AbTestingEnabled     bool              `json:"ab_testing_enabled"`
	ResetURI             string            `json:"reset_uri"`
	SketchSeed1          interface{}       `json:"sketch_seed1"`
	SketchSeed2          interface{}       `json:"sketch_seed2"`
	Rounds               interface{}       `json:"rounds"`
	PublicKey            string            `json:"public_key"`
	KeyID                int64             `json:"key_id"`
	PrefillSource        interface{}       `json:"prefill_source"`
	IddUserCryptedUID    interface{}       `json:"idd_user_crypted_uid"`
	Locale               string            `json:"locale"`
	Lsd                  *KeyVal           `json:"lsd"`
	Jazoest              *KeyVal           `json:"jazoest"`
	LoginSource          string            `json:"login_source"`
	Timestamp            int64             `json:"timestamp"`
	Lgnrnd               string            `json:"lgnrnd"`
	SendScreenDimensions bool              `json:"send_screen_dimensions"`
	LoginData            *Data             `json:"login_data"`
	ProfileIntroCard     *ProfileIntroCard `json:"profile_intro_card"`
	ShouldUsePageRename  bool              `json:"should_use_page_rename"`
}

type Viewer struct {
}

type KeyVal struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProfileIntroCard struct {
	Bio *Bio   `json:"bio"`
	ID  string `json:"id"`
}

type Bio struct {
	DelightRanges     []interface{} `json:"delight_ranges"`
	ImageRanges       []interface{} `json:"image_ranges"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges"`
	AggregatedRanges  []interface{} `json:"aggregated_ranges"`
	Ranges            []interface{} `json:"ranges"`
	ColorRanges       []interface{} `json:"color_ranges"`
	Text              string        `json:"text"`
}
