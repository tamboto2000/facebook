package facebook

import "errors"

// ErrInvalidSession occured when provided session cookies is invalid
var ErrInvalidSession = errors.New("Invalid session cookie")

// ErrUserNotFound occured when user with username is not found
var ErrUserNotFound = errors.New("user not found")
