package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"piwi-backend-clean/profiles/core"
	"piwi-backend-clean/profiles/core/domains/profiles"
)

type HttpController struct {
	users *core.Module
}
type ReqUser struct {
	AccountID string
}

func NewHttpController(users *core.Module) *HttpController {
	return &HttpController{users: users}
}

func (a *HttpController) Me(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value("user"); user != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("accountID %v", user)
		profile, err := a.users.GetAccountProfile(r.Context(), user.(*ReqUser).AccountID)
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
	var profile profiles.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ID,err := a.users.CreateNewUserProfile(r.Context(),&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile.ID = ID
	json.NewEncoder(w).Encode(profile)
}
