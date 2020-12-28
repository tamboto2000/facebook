package facebook

// Photo contains image type media
type Photo struct {
	ID     string  `json:"id,omitempty"`
	URI    string  `json:"uri,omitempty"`
	Width  int     `json:"width,omitempty"`
	Height int     `json:"height,omitempty"`
	URL    string  `json:"url,omitempty"`
	Scale  float64 `json:"scale,omitempty"`
}
