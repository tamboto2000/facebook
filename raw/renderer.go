package raw

type Renderer struct {
	Typename             string                        `json:"__typename"`
	Field                *Field                        `json:"field"`
	ProfileFieldSections []ProfileFieldSection         `json:"profile_field_sections"`
	Collection           *Collection                   `json:"collection"`
	User                 *User                         `json:"user"`
	IsProdEligible       bool                          `json:"is_prod_eligible"`
	Attachment           *Attachment                   `json:"attachment"`
	Feedback             *Feedback                     `json:"feedback"`
	Text                 string                        `json:"text"`
	TranslationType      *string                       `json:"translation_type,omitempty"`
	ProfileAction        *ActionsRendererProfileAction `json:"profile_action"`
}
