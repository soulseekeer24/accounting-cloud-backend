package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"piwi-backend-clean/authentication/core"
	"piwi-backend-clean/authentication/core/dto"
)

type AuthenticationHTTP struct {
	auth *core.Module
}

func NewAuthHTTP(auth *core.Module) *AuthenticationHTTP {
	return &AuthenticationHTTP{auth: auth}
}

func (a *AuthenticationHTTP) Signin(w http.ResponseWriter, r *http.Request) {
	var dto dto.LoginAccount

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.auth.Authenticate(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, token)
	// json.NewEncoder(w).Encode(&account)

}

func (a *AuthenticationHTTP) SignUp(w http.ResponseWriter, r *http.Request) {

	var dto dto.RegisterUser

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success, err := a.auth.RegisterAccounts(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !success {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Check you email for confirmation code")
}

func (a *AuthenticationHTTP) ValidateAccount(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["validation_code"]

	success, err := a.auth.ValidateAccount(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !success {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Check you email for confirmation code")
}
