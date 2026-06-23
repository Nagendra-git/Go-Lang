package models

type Inventory struct {
	products *Products
}

func NewInventory() *Inventory {
	products := NewProducts()
	return &Inventory{
		products: products,
	}
}

func (inventory *Inventory) AddProduct(name string, cost float64, uints int) *Product {
	newProduct := NewProduct(len(inventory.products.list)+1, name, cost, uints)
	inventory.products.AddNew(newProduct)
	return newProduct
}

func (inventory *Inventory) GetProducts() []*Product {
	return inventory.products.list
}

func (inventory *Inventory) GetCount() int {
	return inventory.products.GetCount()
}

func (inventory *Inventory) GetByIndex(idx int) *Product {
	return inventory.products.list[idx]
}


