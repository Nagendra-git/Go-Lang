package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Commands struct {
	inventry IInventory
}

var ErrorUserInput error = errors.New("user Input error")

func NewCommand(inventory IInventory) *Commands {
	return &Commands{
		inventry: inventory,
	}
}

func (commands *Commands) AddProduct() error {
	var userInput string
	fmt.Println("Enter the product info (name, cost, usnits) : ")
	fmt.Scanln(&userInput)
	fields := strings.Split(userInput, ",")
	name := fields[0]
	if cost, err := strconv.ParseFloat(fields[1], 64); err == nil {
		if units, err := strconv.Atoi(fields[2]); err == nil {
			newProduct := commands.inventry.AddProduct(name, cost, units)
			fmt.Println("New Prodcut added :", newProduct)
			return nil
		}
	}
	return ErrorUserInput
}

func (commands *Commands) QueryCount() {
	count := commands.inventry.GetCount()
	fmt.Printf("Products Count : %d\n", count)
}

func (commands *Commands) QuaryByIndex() error {
	var idx int
	fmt.Println("Enter the index :")
	if _, err := fmt.Scanln(&idx); err != nil {
		return ErrorUserInput
	}
	if product := commands.inventry.GetByIndex(idx); product != nil {
		fmt.Println(product)
		return nil
	}
	fmt.Println("Prduct not found for the given index", idx)
	return nil
}

func (commands *Commands) QueryProducts() {
	if products := commands.inventry.GetProducts(); len(products) != 0 {
		for _, p := range products {
			fmt.Println(p)
		}
		return
	}
	fmt.Println("Noo products Found!")
}

func (commands *Commands) QueryInventoryValue() {
	invValue := commands.inventry.GetValue()
	fmt.Printf("Inventory value : %02.f\n", invValue)
}

func (commands *Commands) SortProductsByCost() {
	commands.inventry.SortByCost()
	commands.QueryProducts()
}
