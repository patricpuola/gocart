package api

import (
	"net/http"
	"patricpuola/gocart/itemservice"
	"patricpuola/gocart/util"
)

func ItemIndex(rw http.ResponseWriter, req *http.Request) {
	respond(rw, http.StatusOK, itemservice.GetAll())
}

func ItemNew(rw http.ResponseWriter, req *http.Request) {
	item := util.MockItem()
	err := itemservice.CatalogAdd(item)
	if err != nil {
		respond(rw, http.StatusConflict, ErrorResponse{err.Error()})
		return
	}
	respond(rw, http.StatusOK, item)
}

func ItemRemove(rw http.ResponseWriter, req *http.Request) {
	respond(rw, http.StatusOK, Response{"OK"})
}
