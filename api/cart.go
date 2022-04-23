package api

import (
	"net/http"
	"patricpuola/gocart/cartservice"
	"strconv"

	"github.com/gorilla/mux"
)

func CartIndex(rw http.ResponseWriter, req *http.Request) {
	uuid := mux.Vars(req)["uuid"]
	if uuid != "" {
		respond(rw, http.StatusOK, cartservice.Get(&uuid))
	} else {
		respond(rw, http.StatusOK, cartservice.GetAll())
	}
}

func CartClear(rw http.ResponseWriter, req *http.Request) {
	respond(rw, http.StatusOK, Response{"OK"})
}

func CartNew(rw http.ResponseWriter, req *http.Request) {
	cid := mux.Vars(req)["cid"]
	if cid == "" {
		respond(rw, http.StatusBadRequest, Response{"No customer id given"})
		return
	}

	customerId, _ := strconv.Atoi(cid)
	cart, err := cartservice.New(customerId)
	if err != nil {
		respond(rw, http.StatusBadRequest, Response{err.Error()})
		return
	}
	respond(rw, http.StatusOK, cart)
}
