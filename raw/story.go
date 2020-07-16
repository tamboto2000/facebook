package raw

type Story struct {
	ID                        string          `json:"id"`
	CometSections             *CometSections  `json:"comet_sections"`
	EncryptedTracking         string          `json:"encrypted_tracking"`
	DebugInfo                 interface{}     `json:"debug_info"`
	SerializedFrtpIdentifiers interface{}     `json:"serialized_frtp_identifiers"`
	CanViewerSeeMenu          bool            `json:"can_viewer_see_menu"`
	CreationTime              int64           `json:"creation_time"`
	URL                       string          `json:"url"`
	GhlLabel                  interface{}     `json:"ghl_label"`
	Feedback                  Feedback        `json:"feedback"`
	Attachments               []Attachment    `json:"attachments"`
	SponsoredData             interface{}     `json:"sponsored_data"`
	AttachedStory             *Story          `json:"attached_story"`
	Typename                  string          `json:"__typename"`
	IsProdEligible            bool            `json:"is_prod_eligible"`
	Story                     *Story          `json:"story"`
	IsTextOnlyStory           bool            `json:"is_text_only_story"`
	FeedbackContext           FeedbackContext `json:"feedback_context"`
	Actors                    []Actor         `json:"actors"`
	Message                   Message         `json:"message"`
}
