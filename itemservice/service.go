package itemservice

import (
	"strconv"

	"github.com/jaswdr/faker"
)

var itemCatalog = make([]*Item, 0)

func GetAll() []*Item {
	return itemCatalog
}

func Get(ean string) *Item {
	for _, item := range itemCatalog {
		if item.Ean == ean {
			return item
		}
	}
	return nil
}

func CatalogAdd(item *Item) {
	itemCatalog = append(itemCatalog, item)
}

func MockItem() *Item {
	car := faker.New().Car()
	price, _ := strconv.ParseFloat(faker.New().Numerify("#####.##"), 32)
	item := Item{faker.Car.Maker(car) + " " + faker.Car.Model(car), faker.New().Numerify("#############"), float32(price)}
	return &item
}
