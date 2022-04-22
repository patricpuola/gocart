package api

import (
	"encoding/json"
	"net/http"
	"patricpuola/gocart/cartservice"
)

type Response struct {
	Response string
}

func Index(rw http.ResponseWriter, req *http.Request) {
	respond(rw, http.StatusOK, cartservice.GetAll())
}

func New(rw http.ResponseWriter, req *http.Request) {
	cartOwnerName := req.FormValue("name")
	if cartOwnerName == "" {
		respond(rw, http.StatusBadRequest, Response{"Empty name"})
		return
	}

	cart, err := cartservice.NewCart(cartOwnerName)
	if err != nil {
		respond(rw, http.StatusNotAcceptable, Response{err.Error()})
		return
	}
	respond(rw, http.StatusOK, cart)
}

func Add(rw http.ResponseWriter, req *http.Request) {

}

func Remove(rw http.ResponseWriter, req *http.Request) {

}

func Clear(rw http.ResponseWriter, req *http.Request) {

}

func respond(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
