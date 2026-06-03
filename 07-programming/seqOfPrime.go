package main

import "fmt"

func main() {

	fmt.Println("Enter the range :")
	var r int
	fmt.Scanln(&r)

	for i := 1; i <= r; i++ {

		isPrimeNo := true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrimeNo = false
			}
		}
		if isPrimeNo {
			fmt.Println(i)
		}

	}

}
