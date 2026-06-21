package main

import (
	"fmt"
	"math"
)

func main() {

	//declare a fuction valraible
	getSquareRoot := func(x float64) float64 {

		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	//Assigning a function to a variable

	sum := AddTwoNumbers

	result := sum(100, 200)
	fmt.Println(result)

	// passing fuction as argument
	result1 := calculation(2, 5, multiplayNumbers)
	fmt.Println(result1)

	// Return fuction as a value

	multiplayByTwo := calculate(2)

	result2 := multiplayByTwo(20)

	fmt.Println(result2)

}

func calculate(fact int) func(int) int{
	return func(i int) int {
		return fact*i
	}
}

func AddTwoNumbers(x, y int) int {
	return x + y
}

func calculation(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func multiplayNumbers(x, y int) int {
	return x * y
}
