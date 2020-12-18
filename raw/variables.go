package raw

type Variables struct {
	NotInViewAs bool    `json:"notInViewAs,omitempty"`
	Scale       float64 `json:"scale,omitempty"`
	UserID      string  `json:"userID,omitempty"`
}
