package cartservice

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

func New(customerId int) *ShoppingCart {
	cart := newShoppingCart(customerId)
	shoppingCarts = append(shoppingCarts, cart)

	return cart
}

func Clear() {
	shoppingCarts = make([]*ShoppingCart, 0)
}
