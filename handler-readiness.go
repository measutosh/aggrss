package main

import "net/http"

// http handler that will trigger the json sending process
// this handler is hooked in the aggrss file
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
