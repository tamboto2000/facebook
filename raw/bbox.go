package raw

type Bbox struct {
	Complete       bool        `json:"complete"`
	Result         *Result     `json:"result"`
	SequenceNumber int         `json:"sequence_number"`
	Variables      *Variables  `json:"variables"`
	ExtraContext   interface{} `json:"extra_context"`	
}
