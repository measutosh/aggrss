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
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// take input as a json body
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
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
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		// this will create a new random uuid
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create the feed: %v", err))
		return
	}

	// return the custom made feed is made in models.go
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}


// this one is not authenticated, so no uesr is passed
// this handler has been hooked up in the aggrss.go
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	// this feeds needs to be converted, that has been done models
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get all the feeds: %v", err))
		return
	}

	// return the custom made feed is made in models.go
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
