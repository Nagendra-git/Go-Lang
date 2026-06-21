package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)
	swap(a, b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

	std := Student{"Nagendra", 24}

	fmt.Println("Student data before function call:\n", std)
	// Call by value
	changeStudentData(std)

	fmt.Println("Student data after function call:\n", std)
}

func changeStudentData(s Student) {
	s.name = "Janu"
	s.age = 26
	fmt.Println("Student Data Inside changeStudentData():\n", s)
}

func swap(x, y int) {
	var temp int
	fmt.Printf("Inside Swap Before swap, value of a : %d\n", x)
	fmt.Printf("Inside Swap Before swap, value of b : %d\n", y)
	temp = x
	x = y
	y = temp
	fmt.Printf("Inside Swap After swap, value of a : %d\n", x)
	fmt.Printf("Inside Swap After swap, value of b : %d\n", y)
}
