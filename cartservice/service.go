package cartservice

import (
	"errors"
	"patricpuola/gocart/config"
)

var shoppingCarts = make([]*ShoppingCart, 0)

func GetAll() []*ShoppingCart {
	return shoppingCarts
}

func Get(uuid *string) *ShoppingCart {
	for _, cart := range shoppingCarts {
		if *&cart.Uuid == *uuid {
			return cart
		}
	}
	return nil
}

func New(customerId int) (*ShoppingCart, error) {
	cart_limit := config.GetInt("cart_limit")
	if cart_limit > 0 && len(shoppingCarts) >= cart_limit {
		return nil, errors.New("Maximum number of carts reached")
	}
	cart := newShoppingCart(customerId)
	shoppingCarts = append(shoppingCarts, cart)

	return cart, nil
}

func Clear() {
	shoppingCarts = make([]*ShoppingCart, 0)
}
