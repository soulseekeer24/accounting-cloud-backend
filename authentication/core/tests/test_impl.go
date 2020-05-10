package tests

import (
	"piwi-backend-clean/authentication/core"
	"piwi-backend-clean/authentication/core/domains/accounts"
)

type TestEncripter struct{}

func (e *TestEncripter) ValidateHash(original string, underTest string) (success bool, err error) {

	success = original == underTest

	return success, nil
}

func (e *TestEncripter) GenerateValidationHash(key string, seed string) (hast string, err error) {
	return "key-has", nil
}
func (e *TestEncripter) HashPassword(password string) (hash string, err error) {
	return password, nil
}

type MuckTokenManager struct{}

func (t MuckTokenManager) GenerateToken(account *accounts.Account, profileID string) (token string, err error) {
	return "token_lol", nil
}
func (t MuckTokenManager) ValidateToken(token string) (claims *core.TokenClaims, err error) {

	return
}
