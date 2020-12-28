package facebook

import "errors"

// ErrInvalidSession occured when provided session cookies is invalid
var ErrInvalidSession = errors.New("Invalid session cookie")
