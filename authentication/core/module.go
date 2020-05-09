package core

import (
	"context"
	"errors"
	"fmt"
	"piwi-backend-clean/authentication/core/domains/accounts"
	"piwi-backend-clean/authentication/core/dto"
	"piwi-backend-clean/authentication/core/external"
)



type Module struct {
	AccountsService    *accounts.Service
	tokenManager  TokenManager
	profileModule external.ProfileModule
}

func NewAuthentication(accountRepository accounts.Repository, encrypter accounts.Encrypter, tokenManager TokenManager) *Module {
	credService := accounts.NewService(accountRepository, encrypter)

	auth := Module{
		AccountsService: credService,
		tokenManager:    tokenManager}
	return &auth
}

//RegisterAccounts register a account and sent te profile data to the profiles data.
func (m *Module) RegisterAccounts(ctx context.Context, register *dto.RegisterUser) (success bool, err error) {

	if register.FirstName == "" {
		return false, errors.New("firstname missing.")
	}
	if register.FirstName == "" {
		return false, errors.New("lastname missing.")
	}
	if register.Email == "" {
		return false, errors.New("Email missing.")
	}

	if register.Role == "" {
		return false, errors.New("Missing Role.")

	}

	keys, err := m.AccountsService.CreateAccount(ctx, register.Username, register.Password)
	if err != nil {
		return false, err
	}

	fmt.Printf("enviar a comunication %v", keys.VerificationHash)

	// Create a profile information and then pass to the profiles module to
	// create the profile linked to this account
	newProfile := dto.Profile{
		AccountID: keys.AccountID,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Role:      register.Role,
	}

	// Comunication between the profile and authentication module to create profile
	success, err = m.profileModule.CreateProfile(&newProfile)
	if err != nil {
		//handle what kind of error cud happend and retry probably
		// panic(err)
	}

	return success, err
}

func (m *Module) Authenticate(ctx context.Context, loginAccount *dto.LoginAccount) (token string, err error) {
	account, err := m.AccountsService.Authenticate(ctx, loginAccount.Username, loginAccount.Password)
	if err != nil {
		return "", err
	}

	profile, err := m.profileModule.GetProfileByAccountID(account.ID)
	if err != nil {
		return "", err
	}

	token, err = m.tokenManager.GenerateToken(account, profile.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Module) ValidateAccount(ctx context.Context, hash string) (success bool, err error) {
	account, err := m.AccountsService.ValidateAccountWithHash(ctx, hash)
	if err != nil {
		return false, err
	}
	//Once the acccount its validated we procced to mark the email as valid
	m.profileModule.ValidateEmail(account.ID)

	return true, nil
}

func (m *Module) ValidateToken(ctx context.Context, token string) (claims *TokenClaims, err error) {
	return m.tokenManager.ValidateToken(token)
}

func (m *Module) ConnectToProfiles(pm external.ProfileModule) {
	m.profileModule = pm
}
