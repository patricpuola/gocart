package util

import (
	"fmt"
	"math/rand"
	"patricpuola/gocart/cartservice"
	"patricpuola/gocart/itemservice"
	"strconv"

	"github.com/jaswdr/faker"
)

const populate_count_items = 15
const populate_count_carts = 4

// For trying out the service
// Fill out empty service with carts and items
func Populate() {

	PrintVerbose("Populating items")

	for i := range make([]int, populate_count_items) {
		itemservice.CatalogAdd(MockItem())
		PrintVeryVerbose(fmt.Sprintf("Item %d added\n", i))
	}

	PrintVerbose("Populating carts")

	for i := range make([]int, populate_count_carts) {
		MockCart()
		PrintVeryVerbose(fmt.Sprintf("Cart %d added\n", i))
		// todo: add items to carts
	}
}

func MockCart() *cartservice.ShoppingCart {
	cart, _ := cartservice.New(rand.Int())
	return cart
}

func MockItem() *itemservice.Item {
	fkr := faker.NewWithSeed(rand.NewSource(rand.Int63()))
	car := fkr.Car()
	price, _ := strconv.ParseFloat(fkr.Numerify("#####.##"), 32)
	item := itemservice.New(
		car.Maker()+" "+car.Model(),
		strconv.Itoa(fkr.RandomNumber(13)),
		float32(price),
	)
	return item
}
