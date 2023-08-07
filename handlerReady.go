package main

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	// RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	RespondWithJSON(w, 200, struct{}{})
}

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
