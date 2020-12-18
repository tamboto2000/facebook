package raw

type Bbox struct {
	Complete       bool        `json:"complete,omitempty"`
	Result         *Result     `json:"result,omitempty"`
	SequenceNumber int         `json:"sequence_number,omitempty"`
	Variables      *Variables  `json:"variables,omitempty"`
	ExtraContext   interface{} `json:"extra_context,omitempty"`
}
