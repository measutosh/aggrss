package main

import (
	"fmt"
	"net/http"

	"github.com/measutosh/aggrss/internal/auth"
	"github.com/measutosh/aggrss/internal/database"
)

// this middleware exists because any new handler is added to the app that will go through
// the similar auth process, instead of repeating the code everywhere, this middleware
// wil be called

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// the above function doesn't match the function sign of a http.handleFunc
// so an indirect typecasting of the function has been done below
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// collect the apikey
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// pass the key to the db query
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get the user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
