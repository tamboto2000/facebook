package raw

type Extensions struct {
	IsFinal   bool       `json:"is_final,omitempty"`
	SrPayload *SrPayload `json:"sr_payload,omitempty"`
}

type SrPayload struct {
	Ddd Ddd `json:"ddd,omitempty"`
}

type Ddd struct {
	Hsrp         Hsrp     `json:"hsrp,omitempty"`
	Jsmods       Jsmods   `json:"jsmods,omitempty"`
	AllResources []string `json:"allResources,omitempty"`
}

type Hsrp struct {
	Hsdp Hsdp `json:"hsdp,omitempty"`
	Hblp Hblp `json:"hblp,omitempty"`
}

type Hblp struct {
	SrRevision  int                `json:"sr_revision,omitempty"`
	Consistency Consistency        `json:"consistency,omitempty"`
	RsrcMap     map[string]RsrcMap `json:"rsrcMap,omitempty"`
	CompMap     CompMap            `json:"compMap,omitempty"`
}
