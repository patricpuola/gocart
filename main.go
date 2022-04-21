package main

import (
	"fmt"
	"net/http"
	"patricpuola/gocart/api"
	"patricpuola/gocart/cart"

	"github.com/gorilla/mux"
)

func registerHandlers(handler *mux.Router) {
	handler.HandleFunc("/", api.Index)
	handler.HandleFunc("/new", api.New)
	handler.HandleFunc("/add", api.Add)
	handler.HandleFunc("/remove", api.Remove)
	handler.HandleFunc("/clear", api.Clear)
}

func main() {
	fmt.Println("Starting server...")
	handler := mux.NewRouter()

	cart.Clear()
	fmt.Println("Carts cleared")

	fmt.Println("Registering handlers")
	registerHandlers(handler)

	http.ListenAndServe(":8000", handler)
}
