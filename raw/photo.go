package raw

type Photo struct {
	ID           string `json:"id,omitempty"`
	Image        *Image `json:"image,omitempty"`
	ViewerImage  *Image `json:"viewer_image,omitempty"`
	BlurredImage *Image `json:"blurred_image,omitempty"`
	URL          string `json:"url,omitempty"`
}

type Image struct {
	URI    string `json:"uri,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}
