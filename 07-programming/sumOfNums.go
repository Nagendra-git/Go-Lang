package main

import "fmt"

func main() {

	fmt.Println("Enter the range of the number :")
	var no int
	sum := 0
	fmt.Scanln(&no)
	for i := 1; i < no; i++ {
		sum += i
	}
	fmt.Println("Sum of the number is :", sum)

}
