package main

import "fmt"

func main() {
	no := 100

	var noPtr *int

	//address from a value
	noPtr = &no

	//dereferencing - value form an address

	var x int
	x = *noPtr
	fmt.Println(x)

	// in other words
	fmt.Println(no == *(&no))

}
