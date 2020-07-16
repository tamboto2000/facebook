package raw

type Attachment struct {
	StyleTypeRenderer *Renderer          `json:"style_type_renderer"`
	AllSubattachments *AllSubattachments `json:"all_subattachments"`
	MediasetToken     string             `json:"mediaset_token"`
	URL               string             `json:"url"`
	ID                string             `json:"id"`
	Media             *Media             `json:"media"`
}
