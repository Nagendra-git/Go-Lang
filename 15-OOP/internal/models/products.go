package models

type Products struct {
	list []*Product
}

func NewProducts() *Products {
	return &Products{
		list: make([]*Product, 0),
	}
}

func (products *Products) AddNew(product *Product) {
	products.list = append(products.list, product)
}

func (products Products) GetCount() int{
	return len(products.list)
}