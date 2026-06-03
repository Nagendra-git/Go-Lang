package main

import "fmt"

func main() {

	fmt.Println("Enter the number :")
	var no int
	fmt.Scanln(&no)
	res := isPrime(no)
	if res {
		fmt.Println("Given number is a prime number")
	} else {
		fmt.Println("Given number is not a prime number")
	}
}

func isPrime(no int) bool {
	for i := 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}

	}
	return true
}
