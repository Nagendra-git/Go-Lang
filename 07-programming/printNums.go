package main

import "fmt"

func main() {

	fmt.Println("Enter the range:")
	var no int
	fmt.Scanln(&no)
	fmt.Println("Forword")
	// Print forword numbers
	for i := 0; i <= no; i++ {
		fmt.Println(i)
	}

	fmt.Println("Backword")
	//print backword numbers
	for i := no; i >= 0; i-- {
		fmt.Println(i)
	}
}
