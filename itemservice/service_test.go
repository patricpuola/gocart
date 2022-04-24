package itemservice_test

import (
	"fmt"
	"patricpuola/gocart/itemservice"
	"testing"
)

type itemDetails struct {
	name      string
	productId string
	price     float32
}

var itemTestCases = []itemDetails{
	{"Barbie doll", "9532958215221", 29.95},
	{"Airsoft gun", "9254752472452", 109.10},
	{"Black umbrella", "9747878615221", 35.00},
	{"Sock puppet", "52263", 5.99},
	{"Chocolate bar", "953636363631", 2.95},
	{"Millenium Falcon", "SW3252352-14", 150529.00},
}

func TestNewItem(t *testing.T) {
	for _, testItem := range itemTestCases {
		item := itemservice.New(testItem.name, testItem.productId, testItem.price)
		if item.Name != testItem.name {
			t.Error(fmt.Sprintf("Created item's name (%s) does not match given (%s)", item.Name, testItem.name))
		}
		if item.ProductId != testItem.productId {
			t.Error(fmt.Sprintf("Created item's productId (%s) does not match given (%s)", item.ProductId, testItem.productId))
		}
		if item.Price != testItem.price {
			t.Error(fmt.Sprintf("Created item's price (%f) does not match given (%f)", item.Price, testItem.price))
		}
	}
}
