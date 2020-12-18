package raw

type Photo struct {
	ID           string `json:"id"`
	Image        *Image `json:"image"`
	ViewerImage  *Image `json:"viewer_image"`
	BlurredImage *Image `json:"blurred_image"`
	URL          string `json:"url"`
}

type Image struct {
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
