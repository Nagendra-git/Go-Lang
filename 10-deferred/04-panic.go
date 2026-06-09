package main

import "fmt"

func main() {

	fmt.Println(" Before Panic")
	panic("Something went wrong")
	fmt.Println("After panic")
}
