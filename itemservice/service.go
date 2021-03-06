package itemservice

import (
	"errors"
	"patricpuola/gocart/config"
)

var itemCatalog = make([]*Item, 0)

func GetAll() []*Item {
	return itemCatalog
}

func Get(productId *string) (item *Item, itemFound bool) {
	for _, item := range itemCatalog {
		if item.ProductId == *productId {
			return item, true
		}
	}
	return nil, false
}

func CatalogAdd(item *Item) error {
	item_limit := config.GetInt("item_limit")
	if item_limit > 0 && len(itemCatalog) >= item_limit {
		return errors.New("Maximum number of cataloged items reached")
	}
	itemCatalog = append(itemCatalog, item)
	return nil
}
