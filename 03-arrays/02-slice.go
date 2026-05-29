package main

import (
	"fmt"
)

func main() {

	/*odd := [6]int{2, 4, 6, 8, 10, 12}

	var s []int = odd[1:4]

	fmt.Println(s)*/

	/*names := [4]string{
		"John",
		"Jim",
		"Jack",
		"jen",
	}

	fmt.Println(names)
	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)
	slice2[0] = "ZZZ"
	fmt.Println(slice1, slice2)
	fmt.Println(names)*/

	/*s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
		{5, false},
		{6, true},
	}

	fmt.Println(s)*/

	/*slice1 := []int{2, 4, 6, 8, 10, 12, 14}
	printSlice(slice1)
	slice2 := slice1[:0]
	printSlice(slice2)
	slice3 := slice1[:4]
	printSlice(slice3)
	slice4 := slice1[2:]
	printSlice(slice4)*/

	slice := make([]int, 10)
	printSlice("slice", slice)
	slice1 := make([]int, 0, 10)
	printSlice("slice1", slice1)
	slice2 := slice1[:5]
	printSlice("slice2", slice2)
	slice3 := slice2[2:5]
	printSlice("slice3", slice3)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s length=%d capacity=%d %v\n",
		s, len(x), cap(x), x)
}

// func printSlice(s []int) {

// 	fmt.Printf("length =%d capacity=%d %v\n", len(s), cap(s), s)

// }
