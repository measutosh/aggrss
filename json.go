package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// way to handle errors
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx Error ", msg)
	}

	// defining the specific json response to show for error message
	type errResponse struct {
		Error string `json:error`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})

}

// way to send JSON data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
