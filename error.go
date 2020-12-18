package facebook

// ErrInvalidSession occured when provided session cookies is invalid
type ErrInvalidSession string

func (ErrInvalidSession *ErrInvalidSession) Error() string {
	return "Invalid session cookie"
}
