package raw

type Result struct {
	Data       *Data       `json:"data,omitempty"`
	Extensions *Extensions `json:"extensions,omitempty"`
	Label      string      `json:"label,omitempty"`
	Path       []string    `json:"path,omitempty"`
}
