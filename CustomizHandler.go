package main

import "net/http"

func HandlerRead(w http.ResponseWriter, r *http.Request) {
	// send an empty JSON object {} with status 200

	respondWithJSON(w, http.StatusInternalServerError, struct{}{})
}
