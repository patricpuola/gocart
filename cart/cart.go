package cart

import "encoding/json"

var savedShoppingCarts = make([]*shoppingCart, 0)

type shoppingCart struct {
	owner    string
	contents []cartRow
}

type cartRow struct {
	item     Item
	quantity int
}

type Item struct {
	name  string
	ean   string
	price float32
}

func newShoppingCart(owner string) *shoppingCart {
	sc := shoppingCart{owner: owner, contents: make([]cartRow, 0)}
	return &sc
}

func newCartRow(item Item, quantity int) *cartRow {
	row := cartRow{item: item, quantity: quantity}
	return &row
}

func getCart(cartOwner string) *shoppingCart {
	for _, cart := range savedShoppingCarts {
		if *&cart.owner == cartOwner {
			return cart
		}
	}
	return nil
}

func getItemCartRowIndex(cart *shoppingCart, item Item) int {
	for index, row := range cart.contents {
		if row.item == item {
			return index
		}
	}

	return -1
}

func NewCart(owner string) {
	sc := newShoppingCart(owner)
	savedShoppingCarts = append(savedShoppingCarts, sc)
}

func AddItem(cartOwner string, item Item) {
	cart := getCart(cartOwner)
	existingRowIdx := getItemCartRowIndex(cart, item)
	if existingRowIdx < 0 {
		*&cart.contents[existingRowIdx].quantity++
	} else {
		cartRow := newCartRow(item, 1)
		*&cart.contents = append(*&getCart(cartOwner).contents, *cartRow)
	}
}

func RemoveItem(cartOwner string, item Item) {
	cart := getCart(cartOwner)
	existingRowIdx := getItemCartRowIndex(cart, item)
	if existingRowIdx < 0 {
		if cart.contents[existingRowIdx].quantity == 1 {
			cart.contents[existingRowIdx] = cart.contents[len(cart.contents)-1]
			cart.contents = cart.contents[:len(cart.contents)-1]
		} else {
			cart.contents[existingRowIdx].quantity--
		}
	} else {
		// item does not exist
	}
}

func ListJson(cartOwner string) string {
	cart := getCart(cartOwner)
	if cart == nil {
		return ""
	}
	bytes, _ := json.Marshal(cart)
	return string(bytes)
}

func Clear() {
	savedShoppingCarts = make([]*shoppingCart, 0)
}
