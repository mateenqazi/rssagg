package main

import (
	"net/http"

	"github.com/mateenqazi/rssagg/internal/auth"
	"github.com/mateenqazi/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUserApiKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
