package api

import (
	"encoding/json"
	"net/http"
)

func Index(rw http.ResponseWriter, req *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(rw, http.StatusOK, response)
}

func New(rw http.ResponseWriter, req *http.Request) {

}

func Add(rw http.ResponseWriter, req *http.Request) {

}

func Remove(rw http.ResponseWriter, req *http.Request) {

}

func Clear(rw http.ResponseWriter, req *http.Request) {

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
