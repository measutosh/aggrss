package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/measutosh/aggrss/internal/database"
)

// the function signature stays the same
// but it some additional config attached to the struct to which access can be gained
// this handler is attached in aggrss.go
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// take input as a json body
	type parameters struct {
		Name string `json:"name"`
	}

	// the request body needs to be passed through the struct
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// if something goes wrong then the error is from client side
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	// if the above code works then create the new user
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		// this will create a new random uuid
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create the user: %s", err))
		return
	}

	// return the custom made user made in models.go
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
