package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	profiles "piwi-backend-clean/profiles/core"

	auth "piwi-backend-clean/authentication/core"
)

var authModule *auth.Module
var profilesModule *profiles.Module

func SetAuthModule(module *auth.Module) {
	authModule = module
}
func SetProfilesModule(module *profiles.Module) {
	profilesModule = module
}

type UserLogged struct {
	AccountID string
}

func IsAuthenticated(callback func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := r.Header["Token"]; ok {
			// Parse the token
			claims, err := authModule.ValidateToken(r.Context(), token[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			//Add data to context
			ctx := context.WithValue(r.Context(), "user", UserLogged{claims.AccountID})
			callback(w, r.WithContext(ctx))
		} else {
			fmt.Fprintf(w, "no aturoize")
		}
	})
}
func IsOwnProfile(callback func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := r.Header["Token"]; ok {
			// Parse the token
			claims, err := authModule.ValidateToken(r.Context(), token[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			user := UserLogged{claims.AccountID}
			profileID := chi.URLParam(r, "profile_id")
			//TODO fix error handling
			profile, err := profilesModule.GetAccountProfile(r.Context(), user.AccountID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if profile.ID != profileID{
				http.Error(w, errors.New("You have not permission to this resource.").Error(), http.StatusUnauthorized)
				return
			}

			//Add data to context
			ctx := context.WithValue(r.Context(), "user", user)
			callback(w, r.WithContext(ctx))
		} else {
			fmt.Fprintf(w, "no aturoize")
		}
	})
}
