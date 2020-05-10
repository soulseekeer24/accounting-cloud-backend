package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"piwi-backend-clean/middlewares"
	"piwi-backend-clean/profiles/core"
	"piwi-backend-clean/profiles/core/domains/profiles"
)

type HttpController struct {
	users *core.Module
}

func NewHttpController(users *core.Module) *HttpController {
	return &HttpController{users: users}
}

func (a *HttpController) Me(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value("user"); user != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("accountID %v", user)
		profile, err := a.users.GetAccountProfile(r.Context(), user.(middlewares.UserLogged).AccountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(&profile)

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}

}

func (a *HttpController) CreateProfile(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value("user"); user != nil {
		var profile profiles.Profile
		err := json.NewDecoder(r.Body).Decode(&profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		accountID := user.(middlewares.UserLogged).AccountID
		ID, err := a.users.CreateNewUserProfile(r.Context(),accountID , &profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		profile.ID = ID
		json.NewEncoder(w).Encode(profile)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
