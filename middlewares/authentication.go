package middlewares

import (
	"context"
	"fmt"
	"net/http"

	auth "piwi-backend-clean/authentication/core"
)



var authModule *auth.Module

func SetAuthModule(module *auth.Module) {
	authModule = module
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


