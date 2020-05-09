package core

import "piwi-backend-clean/authentication/core/domains/accounts"

type TokenManager interface {
	GenerateToken(account *accounts.Account, profileID string) (token string, err error)
	ValidateToken(token string) (claims *TokenClaims, err error)
}
