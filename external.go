package token

import (
	"time"
)

// Sign 签名
func Sign(token *Token, expiries ...time.Duration) (string, error) {
	return module.Sign(token, expiries...)
}

// Validate 验签
func Validate(token string) (*Token, error) {
	return module.Validate(token)
}
