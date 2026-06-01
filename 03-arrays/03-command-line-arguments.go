package main

import (
	"fmt"
	"os"
)

func main() {
	var s, arg string

	for i := 0; i < len(os.Args); i++ {
		s += arg + os.Args[i] + ""
	}
	fmt.Println(s)
}
