package main

import (
	"flag"
	"fmt"
	"net/http"
	"patricpuola/gocart/api"
	"patricpuola/gocart/config"
	"patricpuola/gocart/util"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

const regexCustomerId string = "[0-9]+"
const regexUuid string = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

var Populate *bool

func loadFlags() {
	verbose := flag.Bool("v", false, "Verbose server output")
	very_verbose := flag.Bool("vv", false, "Very verbose server output")
	Populate = flag.Bool("populate", false, "Populate service with carts and items on startup")
	flag.Parse()

	if *very_verbose {
		config.SetVerbosity(config.VERY_VERBOSE)
	} else if *verbose {
		config.SetVerbosity(config.VERBOSE)
	}
}

func loadConfig() error {
	viper.SetConfigFile(".env")

	viper.SetDefault("cart_limit", 0)
	viper.SetDefault("item_limit", 0)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	context := viper.GetString("ENV")
	viper.SetConfigFile(".env." + context)
	viper.SetConfigType("env")
	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println(err)
	}

	viper.Unmarshal(&config.Config)
	if config.IsVerbose() {
		fmt.Println("*** GoCart SERVER ENVIRONMENT ***")
		util.PrettyPrintMap(config.Config, " ")
		fmt.Println("***                  ***")
		fmt.Println()
	}
	return nil
}

func postSetup() {
	if *Populate {
		util.Populate()
	}
}

func registerHandlers(handler *mux.Router) {
	handler.StrictSlash(true)

	handler.HandleFunc("/", api.Index).Methods("GET")

	handler.HandleFunc("/cart/", api.CartIndex).Methods("GET")
	handler.HandleFunc(fmt.Sprintf("/cart/new/{cid:%s}", regexCustomerId), api.CartNew).Methods("GET")
	handler.HandleFunc("/cart/clear/", api.CartClear).Methods("GET")
	handler.HandleFunc(fmt.Sprintf("/cart/{uuid:%s}", regexUuid), api.CartIndex).Methods("GET")
	handler.HandleFunc(fmt.Sprintf("/cart/{uuid:%s}/item/add/{productId}", regexUuid), api.CartAddItem).Methods("GET")
	handler.HandleFunc(fmt.Sprintf("/cart/{uuid:%s}/item/remove/{productId}", regexUuid), api.CartRemoveItem).Methods("GET")

	handler.HandleFunc("/item/", api.ItemIndex).Methods("GET")
	handler.HandleFunc("/item/new/", api.ItemNew).Methods("GET")
	handler.HandleFunc("/item/remove/", api.ItemRemove).Methods("GET")
}

func main() {
	loadFlags()
	loadConfig()
	postSetup()

	fmt.Println("Starting server")
	handler := mux.NewRouter()

	fmt.Println("Registering handlers")
	registerHandlers(handler)

	fmt.Println("Ready")
	port := config.GetString("port")
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
