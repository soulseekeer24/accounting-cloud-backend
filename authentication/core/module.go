package core

import (
	"context"
	"fmt"
	"piwi-backend-clean/authentication/core/domains/accounts"
	"piwi-backend-clean/authentication/core/dto"
)

type Module struct {
	AccountsService *accounts.Service
	tokenManager    TokenManager
}

func NewAuthentication(accountRepository accounts.Repository, encrypter accounts.Encrypter, tokenManager TokenManager) *Module {
	credService := accounts.NewService(accountRepository, encrypter)

	auth := Module{
		AccountsService: credService,
		tokenManager:    tokenManager}
	return &auth
}

//RegisterAccounts register a account and sent te profile data to the profiles data.
func (m *Module) RegisterAccounts(ctx context.Context, acc *accounts.Account) (success bool, err error) {

	keys, err := m.AccountsService.CreateAccount(ctx, acc)
	if err != nil {
		return false, err
	}

	if keys != nil {
		success = true
	}

	fmt.Printf("enviar a comunication %v \n", keys.VerificationHash)

	return success, err
}

func (m *Module) Authenticate(ctx context.Context, loginAccount *dto.LoginAccount) (token string, err error) {
	account, err := m.AccountsService.Authenticate(ctx, loginAccount.Username, loginAccount.Password)
	if err != nil {
		return "", err
	}

	token, err = m.tokenManager.GenerateToken(account)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Module) ValidateAccount(ctx context.Context, hash string) (success bool, err error) {
	_, err = m.AccountsService.ValidateAccountWithHash(ctx, hash)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *Module) ValidateToken(ctx context.Context, token string) (claims *TokenClaims, err error) {
	return m.tokenManager.ValidateToken(token)
}


