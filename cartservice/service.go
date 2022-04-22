package cartservice

import (
	"encoding/json"
	"errors"
	"fmt"
)

var shoppingCarts = make([]*ShoppingCart, 0)

func GetAll() []*ShoppingCart {
	return shoppingCarts
}

func getCart(cartOwner string) *ShoppingCart {
	for _, cart := range shoppingCarts {
		if *&cart.Owner == cartOwner {
			return cart
		}
	}
	return nil
}

func NewCart(owner string) (any, error) {
	if getCart(owner) != nil {
		return nil, errors.New("Owner already has a cart")
	}
	sc := newShoppingCart(owner)
	shoppingCarts = append(shoppingCarts, sc)

	// Debug
	bytes, _ := json.Marshal(sc)
	fmt.Println("Created: ", string(bytes))

	return sc, nil
}

func Clear() {
	shoppingCarts = make([]*ShoppingCart, 0)
}
