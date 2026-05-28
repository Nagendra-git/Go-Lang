package main

import (
	"fmt"
)

func main() {

	/*var num [5]int

	num[0] = 10
	num[1] = 20
	num[2] = 30
	num[3] = 40
	num[4] = 50
	fmt.Println("Array elements")
	for m := 0; m < len(num); m++ {
		fmt.Println(num[m])
	}*/

	/*num := [...]int{100, 200, 300, 400, 500}

	l := len(num)
	fmt.Println("The array elements are ", num)
	fmt.Println("The length of the array is ", l)
	*/

	/*n := [5]int{100, 200, 300, 400, 500}
	fmt.Println("The element at index 0 is ", n[0])
	fmt.Println("The element at index 2 is ", n[2])
	fmt.Println("The element at index 4 is ", n[4])
	l := len(n)
	fmt.Println("Lenth of the array is ", l)*/

	/*var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("The matrix elements are ")

	for m := 0; m < 2; m++ {
		for n := 0; n < 3; n++ {
			fmt.Print(matrix[m][n], " ")
		}
		fmt.Println()
	}*/

	num := [5]int{100, 200, 300, 400, 500}

	sum := 0
	l := len(num)
	for i := 0; i < l; i++ {
		sum += num[i]
	}
	fmt.Println("Sum of an array is ", sum)
	fmt.Println("Average of an array is ", sum/l)

}
