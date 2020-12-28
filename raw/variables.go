package raw

type Variables struct {
	NotInViewAs       bool    `json:"notInViewAs,omitempty"`
	Scale             float64 `json:"scale,omitempty"`
	UserID            string  `json:"userID,omitempty"`
	AppSectionFeedKey string  `json:"appSectionFeedKey,omitempty"`
	CollectionToken   string  `json:"collectionToken,omitempty"`
	RawSectionToken   string  `json:"rawSectionToken,omitempty"`
	SectionToken      string  `json:"sectionToken,omitempty"`
}
