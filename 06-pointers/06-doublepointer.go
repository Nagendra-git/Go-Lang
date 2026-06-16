package main

import "fmt"

func main() {
	var v int = 100
	var pt1 *int = &v
	var pt2 **int = &pt1

	fmt.Println("The Value of Variable V is = ", v)
	fmt.Println("Address of variable V is = ", &v)
	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The value of pt2 is = ", pt2)

	// Dereferencing the
	// pointer to pointer
	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)

	// double pointer will give the value of variable V
	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
}
