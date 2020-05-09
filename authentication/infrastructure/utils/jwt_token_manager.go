package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"piwi-backend-clean/authentication/core"
	"piwi-backend-clean/authentication/core/domains/accounts"
)

var secretSign = []byte("secretclae")

type JWTTokenManager struct{}

func (t JWTTokenManager) GenerateToken(account *accounts.Account, profileID string) (token string, err error) {

	claims := core.TokenClaims{
		AccountID: account.ID,
		ProfileID: profileID,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString(secretSign)
	if err != nil {
		return "nil", fmt.Errorf("Error signining")
	}
	return token, nil
}

// ValidateToken use to validate json token ang get claims data
func (t JWTTokenManager) ValidateToken(tokenString string) (claims *core.TokenClaims, err error) {

	// Parse the token
	claims = &core.TokenClaims{}
	tk, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify

		return secretSign, nil
	})
	if err != nil {
		return nil, err
	}

	err = tk.Claims.Valid()
	if err != nil {
		return nil, err
	}
	fmt.Println(claims)
	return claims, nil
}
