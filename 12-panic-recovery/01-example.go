package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		fmt.Println("[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("app panic", err)
			return
		}

	}()

	divisor := 0
	q, r := devide(100, divisor)
	fmt.Println(q, r)

}

func devide(x, y int) (q, r int) {
	defer func() {
		fmt.Println("[divide] deferred..")

	}()
	fmt.Println("[deveide] calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	q = x / y
	fmt.Println("[devide] Caluculating Remainder")
	r = x % y
	return
}
