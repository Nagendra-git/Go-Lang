package calculator

import "fmt"

func init() {
	fmt.Println("Calculator[add.go] initilized")
}

func Add(x, y int) {

	opCount["add"]++
	fmt.Println("add result :", x+y)

}
