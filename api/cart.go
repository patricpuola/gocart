package api

import (
	"fmt"
	"net/http"
	"patricpuola/gocart/cartservice"
	"patricpuola/gocart/config"
	"patricpuola/gocart/itemservice"
	"patricpuola/gocart/util"
	"strconv"

	"github.com/gorilla/mux"
)

func CartIndex(rw http.ResponseWriter, req *http.Request) {
	if uuid, uuidProvided := mux.Vars(req)["uuid"]; uuidProvided {
		if cart, cartFound := cartservice.Get(&uuid); cartFound {
			respond(rw, http.StatusOK, cart)
		} else {
			respond(rw, http.StatusNotFound, ErrorResponse{"Cart not found"})
		}
	} else {
		respond(rw, http.StatusOK, cartservice.GetAll())
	}
}

func CartClear(rw http.ResponseWriter, req *http.Request) {
	util.PrintVeryVerbose("Carts cleared")
	respond(rw, http.StatusOK, Response{"OK"})
}

func CartNew(rw http.ResponseWriter, req *http.Request) {
	cid, cidProvided := mux.Vars(req)["cid"]
	if !cidProvided {
		respond(rw, http.StatusBadRequest, ErrorResponse{"No customer id given"})
		return
	}

	customerId, _ := strconv.Atoi(cid)
	cart, err := cartservice.New(customerId)
	if err != nil {
		respond(rw, http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	respond(rw, http.StatusOK, cart)
}

func CartAddItem(rw http.ResponseWriter, req *http.Request) {
	getVars := mux.Vars(req)
	uuid, uuidProvided := getVars["uuid"]
	productId, productIdProvided := getVars["productId"]

	if !uuidProvided || !productIdProvided {
		respond(rw, http.StatusBadRequest, ErrorResponse{"Invalid parameters"})
		return
	}

	cart, cartFound := cartservice.Get(&uuid)
	if !cartFound {
		respond(rw, http.StatusNotFound, ErrorResponse{"Cart not found"})
		return
	}

	item, itemFound := itemservice.Get(&productId)
	if !itemFound {
		respond(rw, http.StatusNotFound, ErrorResponse{"Item not found"})
		return
	}

	cart.AddItem(*item)

	if config.IsVeryVerbose() {
		fmt.Printf("Item %s added to cart %s\n", item.ProductId, cart.Uuid)
	}

	respond(rw, http.StatusOK, Response{"Ok"})
}

func CartRemoveItem(rw http.ResponseWriter, req *http.Request) {
	getVars := mux.Vars(req)
	uuid, uuidProvided := getVars["uuid"]
	productId, productIdProvided := getVars["productId"]

	if !uuidProvided || !productIdProvided {
		respond(rw, http.StatusBadRequest, ErrorResponse{"Invalid parameters"})
		return
	}

	cart, cartFound := cartservice.Get(&uuid)
	if !cartFound {
		respond(rw, http.StatusNotFound, ErrorResponse{"Cart not found"})
		return
	}

	item, itemFound := itemservice.Get(&productId)
	if !itemFound {
		respond(rw, http.StatusNotFound, ErrorResponse{"Item not found"})
		return
	}

	if err := cart.RemoveItem(*item); err != nil {
		respond(rw, http.StatusConflict, ErrorResponse{err.Error()})
		return
	}

	if config.IsVeryVerbose() {
		fmt.Printf("Item %s removed from cart %s\n", item.ProductId, cart.Uuid)
	}

	respond(rw, http.StatusOK, Response{"Ok"})
}
