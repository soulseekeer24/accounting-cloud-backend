package tests

import (
	"piwi-backend-clean/authentication/core/domains/accounts"
	"testing"

)

type ServiceTestFunction func(s *accounts.Service, t *testing.T)

type TestCase struct {
	Name     string
	Callback ServiceTestFunction
}

// func ServiceSuite(s *accounts.Service, t *testing.T) {
// 	tc := []TestCase{
// 		{Name: "create new accounts", Callback: CreateAccountTest},
// 	}

// 	for _, tCase := range tc {
// 		t.Run(tCase.Name, func(t *testing.T) {
// 			tCase.Callback(s, t)
// 		})
// 	}
// }

// func CreateAccountTest(s *accounts.Service, t *testing.T) {
// 	ctx := context.Background()
// 	_, err := s.CreateAccount(ctx, "username", "password")
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestService(t *testing.T) {
// 	repo := &accounts.MuckRepository{
// 		SaveAccountFunc:           SaveAccount,
// 		GetAccountsByUserNameFunc: GetAccountsByUserName,
// 	}

// 	s := accounts.NewService(repo, &TestEncripter{})

// 	ServiceSuite(s, t)
// }
