package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {

	var a int = 100
	var b int = 200

	fmt.Println("[main] A before Swap: ", a)
	fmt.Println("[main] B before swap: ", b)
	swapByRef(&a, &b)
	fmt.Println("[main] A after swap: ", a)
	fmt.Println("[main] b after swap: ", b)

	std := Student{"Nagendra", 24}
	fmt.Println("[main] Before change Student :", std)
	changeStudentDataRef(&std)
	fmt.Println("[main] After change :", std)
}

func changeStudentDataRef(s *Student) {
	s.name = "Janu"
	s.age = 24
}

func swapByRef(x, y *int) {
	fmt.Println("[swapByRef] A before swap: ", *x)
	fmt.Println("[swapByRef] B before swap: ", *y)
	var temp int

	temp = *x
	*x = *y
	*y = temp
	fmt.Println("[swapByRef] A after swap: ", *x)
	fmt.Println("[swapByRef] B after swap: ", *y)

}
