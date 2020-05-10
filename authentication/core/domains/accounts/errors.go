package accounts

import (
	"errors"
	"fmt"
)

var (
	EmptyPasswordError           error = simpleErr("Password cant be empty")
	EmptyUsernameError           error = simpleErr("Username cant be empty")
	InvalidAccountsError               = simpleErr("Invalid accounts")
	InvalidVerificationCodeError       = simpleErr("Invalid verification code")
	AccountAlreadyVerifiedError        = simpleErr("The account its already verified")
	AccountBlockedError                = simpleErr("The accounts its currently blocked")
)

// type InvalidAccounts struct {
// }

// func (e InvalidAccounts) Error() string {
// 	return "Invalid accounts"
// }

// simpleErr create simple error with flat msg
func simpleErr(msg string) error {
	return errors.New(msg)
}

type ErrAccountDontExist struct{}
func (e ErrAccountDontExist) Error() string {
	return fmt.Sprintf("account doesn't exist.")
}

type AlreadyExistUsernameError struct{}
func (e AlreadyExistUsernameError) Error() string {
	return fmt.Sprintf("account usrname already exist.")
}

type UnverifiedAccountError struct{}
func (e UnverifiedAccountError) Error() string {
	return fmt.Sprintf("Unverified Account.")
}

type ErrMissingEmail struct{}
func (e ErrMissingEmail) Error() string {
	return fmt.Sprintf("account doesn't exist.")
}