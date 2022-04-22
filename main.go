package main

import (
	"fmt"
	"net/http"
	"patricpuola/gocart/api"

	"github.com/gorilla/mux"
)

const regexCustomerId string = "[0-9]+"
const regexUuid string = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

func registerHandlers(handler *mux.Router) {
	handler.HandleFunc("/", api.Index).Methods("GET")

	handler.HandleFunc("/cart", api.CartIndex).Methods("GET")
	handler.HandleFunc(fmt.Sprintf("/cart/new/{cid:%s}", regexCustomerId), api.CartNew).Methods("GET")
	handler.HandleFunc("/cart/clear", api.CartClear).Methods("GET")
	handler.HandleFunc(fmt.Sprintf("/cart/{uuid:%s}", regexUuid), api.CartIndex).Methods("GET")

	handler.HandleFunc("/item", api.ItemIndex).Methods("GET")
	handler.HandleFunc("/item/new", api.ItemNew).Methods("GET")
	handler.HandleFunc("/item/remove", api.ItemRemove).Methods("GET")
}

func main() {
	fmt.Println("Starting server")
	handler := mux.NewRouter()

	fmt.Println("Registering handlers")
	registerHandlers(handler)

	http.ListenAndServe(":8000", handler)
}
