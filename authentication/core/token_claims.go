package core

import "errors"

type TokenClaims struct {
	AccountID string `json:"account_id"`
}

func (t TokenClaims) Valid() error {
	if t.AccountID == "" {
		return errors.New("missing account ID")
	}
	return nil
}
