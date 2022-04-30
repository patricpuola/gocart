package api

import (
	"encoding/json"
	"net/http"
)

// Generic Response
type Response struct {
	Response string
}

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
