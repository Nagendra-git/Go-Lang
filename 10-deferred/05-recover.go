package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd ", r)
		}

	}()

	fmt.Println("Before panic")
	panic("Something is wrong!")

	fmt.Println("Code won't execute")
}
