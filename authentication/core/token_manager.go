package core

import "piwi-backend-clean/authentication/core/domains/accounts"

type TokenManager interface {
	GenerateToken(account *accounts.Account) (token string, err error)
	ValidateToken(token string) (claims *TokenClaims, err error)
}
