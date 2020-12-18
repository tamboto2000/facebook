package facebook

// Photo contains image type media
type Photo struct {
	ID     string `json:"id"`
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}
