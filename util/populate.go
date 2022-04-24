package util

import (
	"fmt"
	"math/rand"
	"patricpuola/gocart/cartservice"
	"patricpuola/gocart/config"
	"patricpuola/gocart/itemservice"
	"strconv"

	"github.com/jaswdr/faker"
)

const populate_count_items = 15
const populate_count_carts = 4

// For trying out the service
// Fill out empty service with carts and items
func Populate() {
	if config.IsVerbose() {
		fmt.Println("Populating items")
	}
	for i := range make([]int, populate_count_items) {
		itemservice.CatalogAdd(MockItem())
		if config.IsVeryVerbose() {
			fmt.Printf("Item %d added\n", i)
		}
	}

	if config.IsVerbose() {
		fmt.Println("Populating carts")
	}

	for i := range make([]int, populate_count_carts) {
		MockCart()
		if config.IsVeryVerbose() {
			fmt.Printf("Cart %d added\n", i)
		}
		// todo: add items to carts
	}
}

func MockCart() *cartservice.ShoppingCart {
	cart, _ := cartservice.New(rand.Int())
	return cart
}

func MockItem() *itemservice.Item {
	car := faker.New().Car()
	price, _ := strconv.ParseFloat(faker.New().Numerify("#####.##"), 32)
	item := itemservice.New(
		faker.Car.Maker(car)+" "+faker.Car.Model(car),
		faker.New().Numerify("#############"),
		float32(price),
	)
	return item
}
