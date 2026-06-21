package main

import (
	"fmt"
)

func main() {

	//1.Poniters Type

	var x int = 20

	var ptr *int = &x
	fmt.Println(*ptr)
	fmt.Printf("x is a type of %T\n", x)
	fmt.Printf("ptr is a type of %T\n", ptr)

	//2.Array Type

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Printf("arr is a Type of %T\n", arr)

	//3.Struct type
	type Student struct {
		Name string
		Age  int
	}
	var std Student = Student{"Nagendra", 24}
	fmt.Println(std)
	//4.Union Type
	var u interface{} = "Nagendra burusu"
	fmt.Println(u)
	u = 24
	fmt.Println(u)
	u = true
	fmt.Println(u)
	//5.Fuction type

	var AddTwoNumbers = func(a, b int) int {
		return a + b
	}

	fmt.Println(AddTwoNumbers(2, 4))
	//6.Slice type
	var s []int = []int{1, 2, 3, 4}
	fmt.Println(s)

	//7. Map type
	var countryCode = map[string]string{
		"USA":    "+1",
		"India":  "+91",
		"China":  "+86",
		"Brazil": "+55",
		"UK":     "+44",
	}
	fmt.Println(countryCode)
	fmt.Println(countryCode["India"])

	//8. channel type
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println(<-ch)
}
