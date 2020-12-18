package raw

type Result struct {
	Data       *Data       `json:"data"`
	Extensions *Extensions `json:"extensions"`
	Label      string      `json:"label"`
	Path       []string    `json:"path"`	
}
