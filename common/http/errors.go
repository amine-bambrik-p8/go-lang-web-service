package http

import "errors"

// TODO All HTTP related messages should be moved here
var (
	ErrMessageUnauthorized = errors.New("401 - Unauthorized")
)
