package cartservice

import (
	"errors"
	"fmt"
	"patricpuola/gocart/itemservice"
	"time"

	"github.com/google/uuid"
)

type ShoppingCart struct {
	Uuid       string
	CustomerId int
	Contents   []CartRow
	Created    time.Time
}

type CartRow struct {
	Item     itemservice.Item
	Quantity int
}

func newShoppingCart(customerId int) *ShoppingCart {
	sc := ShoppingCart{Uuid: uuid.New().String(), CustomerId: customerId, Contents: make([]CartRow, 0), Created: time.Now()}
	return &sc
}

func newCartRow(item itemservice.Item, quantity int) *CartRow {
	row := CartRow{Item: item, Quantity: quantity}
	return &row
}

func getItemCartRowIndex(cart *ShoppingCart, item itemservice.Item) (idx int, found bool) {
	for index, row := range cart.Contents {
		if row.Item == item {
			idx, found = index, true
			break
		}
	}

	return idx, found
}

func (cart *ShoppingCart) AddItem(item itemservice.Item) {
	if idx, itemfound := getItemCartRowIndex(cart, item); itemfound {
		cart.Contents[idx].Quantity++
	} else {
		cartRow := newCartRow(item, 1)
		cart.Contents = append(cart.Contents, *cartRow)
	}
	fmt.Println(cart)
}

func (cart *ShoppingCart) RemoveItem(item itemservice.Item) error {
	idx, itemFound := getItemCartRowIndex(cart, item)
	if itemFound {
		if cart.Contents[idx].Quantity == 1 {
			cart.Contents[idx] = cart.Contents[len(cart.Contents)-1]
			cart.Contents = cart.Contents[:len(cart.Contents)-1]
		} else {
			cart.Contents[idx].Quantity--
		}
		return nil
	} else {
		return errors.New("Item is not in the cart")
	}
}

func (cart ShoppingCart) Empty() {
	cart.Contents = make([]CartRow, 0)
}
