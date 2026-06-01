package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	greet := func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}

	greet("Nagendra")

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(2, 3))
}
