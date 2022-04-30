package api

import (
	"encoding/json"
	"net/http"
)

// Common JSON response struct
type Response struct {
	Response string
}

// Common JSON error struct
type ErrorResponse struct {
	Error string
}

func Index(rw http.ResponseWriter, req *http.Request) {
	respond(rw, http.StatusOK, Response{"OK"})
}

func respond(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
