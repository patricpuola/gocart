package itemservice

type Item struct {
	Name      string
	ProductId string
	Price     float32
}

func New(name string, productId string, price float32) *Item {
	item := Item{name, productId, price}
	return &item
}
