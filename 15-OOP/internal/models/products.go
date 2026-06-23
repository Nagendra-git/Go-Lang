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

func (products *Products) GetCount() int {
	return len(products.list)
}

// sort.Interface interface implementation
func (products *Products) Len() int {
	return len(products.list)
}

func (products *Products) Less(i, j int) bool {
	return products.list[i].Cost < products.list[j].Cost
}

func (products *Products) Swap(i, j int) {
	products.list[i], products.list[j] = products.list[j], products.list[i]
}
