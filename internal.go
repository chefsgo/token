package token

import "time"

// Sign 签名
func (this *Module) Sign(token *Token, expiries ...time.Duration) (string, error) {
	if this.connect == nil {
		return "", errInvalidTokenConnection
	}

	expiry := this.config.Expiry
	if len(expiries) > 0 {
		expiry = expiries[0]
	}

	if expiry > 0 {
		token.Expiry = time.Now().Add(expiry).Unix()
	}

	return this.connect.Sign(token)
}

// Validate 验签
func (this *Module) Validate(token string) (*Token, error) {
	if this.connect == nil {
		return nil, errInvalidTokenConnection
	}

	return this.connect.Validate(token)
}
