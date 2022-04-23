package token

import "errors"

const (
	NAME = "token"
)

var (
	errInvalidTokenConnection = errors.New("Invalid token connection.")
)
