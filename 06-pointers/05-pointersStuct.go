package main

import "fmt"

//Define a struct
type Person struct {
	name string
	age  int
}

func main() {
	//Creating an instance of the struct
	p1 := Person{name: "A", age: 25}
	fmt.Println("Original struct", p1)

	// Using the & operator

	//creating a pointer to the struct
	personPointer := &p1
	//Accessing fields using the pointer
	fmt.Println("Name ", personPointer.name) //Automatically dereferences

	fmt.Println("Age: ", personPointer.age)

	//Modifying struct values using the pointer

	personPointer.age = 26

	fmt.Println("Updated struct :", p1)

	//Using the new function

	personPointer1 := new(Person)

	personPointer1.name = "Nagendra"
	personPointer1.age = 24
	fmt.Println("Struct created with new: ", *personPointer1)
}
