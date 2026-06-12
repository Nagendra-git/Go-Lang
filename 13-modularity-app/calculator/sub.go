package calculator

import "fmt"

func Sub(x, y int) {

	opCount["sub"]++
	fmt.Println("Subraction is:", x-y)
}
