package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cfg.handlerUsersCreate(w, r)
	case http.MethodPut:
		cfg.handlerUsersUpdate(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}
}
