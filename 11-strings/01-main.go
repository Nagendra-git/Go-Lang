package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var x string = "Nagendra"

	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	fmt.Println(len(x))

	// Print Ascii value of A

	fmt.Println("Ascii value of N is ", "N"[0])

	//To convert the string into upper case
	fmt.Println(strings.ToUpper(x))

	// To convert the string into lower case

	fmt.Println(strings.ToLower(x))
	fmt.Println(strings.HasPrefix(x, "Na"))
	fmt.Println(strings.HasSuffix(x, "ra"))

	var arr = []string{"a", "b", "c", "d"}

	fmt.Println(strings.Join(arr, "*"))

	fmt.Println(strings.Repeat(x, 3))

	fmt.Println(strings.Contains(x, "n"))
	fmt.Println(strings.Index(x, "n"))
	fmt.Println(strings.Count(x, "a"))
}
