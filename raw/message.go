package raw

type Message struct {
	ID                string        `json:"id"`
	URL               string        `json:"url,omitempty"`
	Name              string        `json:"name,omitempty"`
	Typename          string        `json:"__typename"`
	IsProdEligible    bool          `json:"is_prod_eligible"`
	Story             *Story        `json:"story"`
	Text              string        `json:"text"`
	ImageRanges       []interface{} `json:"image_ranges"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges"`
}
