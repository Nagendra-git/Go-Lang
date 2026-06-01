package main

import (
	"fmt"
	"log"
)

func main() {

	/*add(10, 20)
	subtract(100, 200)*/

	/*logAdd(10, 20)
	logSub(100, 200)*/

	logOperation(add, 10, 20)
	logOperation(subtract, 100, 200)

}

func logOperation(op func(int, int), x, y int) {
	log.Println("Invocation started")
	op(x, y)
	log.Println("Invocation completed")

}

func logAdd(x, y int) {
	log.Println("Invocation started")
	add(x, y)
	log.Println("Invocation completed")
}

func logSub(x, y int) {
	log.Println("Invocation started")
	subtract(x, y)
	log.Println("invocation Completed")
}

func add(x, y int) {
	fmt.Println("Sum of x, y is :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtraction of x, y is ", y-x)
}
