package token

import "errors"

const (
	NAME = "TOKEN"
)

var (
	errInvalidTokenConnection = errors.New("Invalid token connection.")
)
