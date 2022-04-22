package cartservice

import "github.com/google/uuid"

type ShoppingCart struct {
	Uuid     string
	Owner    string
	Contents []CartRow
}

type CartRow struct {
	Item     Item
	Quantity int
}

type Item struct {
	Name  string
	Ean   string
	Price float32
}

func newShoppingCart(owner string) *ShoppingCart {
	sc := ShoppingCart{Uuid: uuid.New().String(), Owner: owner, Contents: make([]CartRow, 0)}
	return &sc
}

func newCartRow(item Item, quantity int) *CartRow {
	row := CartRow{Item: item, Quantity: quantity}
	return &row
}

func getItemCartRowIndex(cart *ShoppingCart, item Item) int {
	for index, row := range cart.Contents {
		if row.Item == item {
			return index
		}
	}

	return -1
}

func AddItem(cartOwner string, item Item) {
	cart := getCart(cartOwner)
	existingRowIdx := getItemCartRowIndex(cart, item)
	if existingRowIdx < 0 {
		*&cart.Contents[existingRowIdx].Quantity++
	} else {
		cartRow := newCartRow(item, 1)
		*&cart.Contents = append(*&getCart(cartOwner).Contents, *cartRow)
	}
}

func RemoveItem(cartOwner string, item Item) {
	cart := getCart(cartOwner)
	existingRowIdx := getItemCartRowIndex(cart, item)
	if existingRowIdx < 0 {
		if cart.Contents[existingRowIdx].Quantity == 1 {
			cart.Contents[existingRowIdx] = cart.Contents[len(cart.Contents)-1]
			cart.Contents = cart.Contents[:len(cart.Contents)-1]
		} else {
			cart.Contents[existingRowIdx].Quantity--
		}
	} else {
		// item does not exist
	}
}
