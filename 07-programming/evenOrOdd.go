package main

import "fmt"

func main() {

	fmt.Println("Enter the number :")
	var no int
	fmt.Scanln(&no)
	if no%2 == 0 {
		fmt.Println("Given number is even number")
	} else {
		fmt.Println("Given number is odd number")
	}
}
