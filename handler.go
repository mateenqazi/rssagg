package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mateenqazi/rssagg/internal/database"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

// func (cfg *apiConfig) Handler(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		Name string
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
// 		return
// 	}

// 	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
// 		ID:        uuid.New(),
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 		Name:      params.Name,
// 	})
// 	if err != nil {
// 		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
// 		return
// 	}

// 	RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
// }

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
