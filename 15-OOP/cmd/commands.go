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

func NewCommand(inventory IInventory) *Commands{
	return &Commands{
		inventry: inventory,
	}
}

func (commands *Commands) AddProduct() error{
	var userInput string
	fmt.Println("Enter the product info (name, cost, usnits) : ")
	fmt.Scanln(&userInput)
	fields := strings.Split(userInput, ",")
	name := fields[0]
	if cost , err := strconv.ParseFloat(fields[1], 64); err == nil{
		if units, err := strconv.Atoi(fields[2]); err == nil {
			newProduct := commands.inventry.AddProduct(name, cost, units)
			fmt.Println("New Prodcut added :", newProduct)
			return nil
		}
	}
	return ErrorUserInput
}

func (commands *Commands) QueryCount(){
	count := commands.inventry.GetCount()
	fmt.Printf("Products Count : %d\n", count)
}

