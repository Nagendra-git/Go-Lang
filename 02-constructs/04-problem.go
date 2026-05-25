package main

import "fmt"

func main() {

	var no, iterCount int
	fmt.Println("Enter the no : ")
	fmt.Scanln(&no)

	for ; no != 1; iterCount++ {
		if no%2 == 0 {
			no = no / 2
		} else {
			no = (3 * no) + 1
		}
	}
	fmt.Println(iterCount)
}
