package cmd

import "inventory-app/internal/models"

type IInventory interface {
	AddProduct(name string, cost float64, units int) *models.Product
	GetCount() int
	GetByIndex(idx int) *models.Product
	GetProducts() []*models.Product
	GetValue() float64
	SortByCost()
}
