package main

import "fmt"

func main() {

	//First, define outer function
	upperCounter := func(intiatValue int) func() int {
		//define local varibel inside the function
		count := intiatValue
		return func() int {
			count++
			return count
		}
	}

	increment1 := upperCounter(100)

	fmt.Println(increment1())
	fmt.Println(increment1())

	increment2 := upperCounter(200)
	fmt.Println(increment2())
	fmt.Println(increment2())
}
