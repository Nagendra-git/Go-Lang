package main

import "fmt"

func main() {

	/*var a int = 1

	for a < 10 {
		fmt.Println("Value of a is ", a, "\n")
		a++

		if a > 5 {
			break
		}

	}*/

	/*for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			if a == 2 && b == 2 {
				break
			}
			fmt.Println("a: ", a, "b: ", b, "\n")
		}
	}*/

	//Infinite loop
	/*i := 1

	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++

	}*/

	//Switch case break statement

	/*days := 4

	switch days {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")
		break
	case 3:
		fmt.Println("Tuesday")
		break
	case 4:
		fmt.Println("Wednesday")
		break
	case 5:
		fmt.Println("Thursday")
		break
	case 6:
		fmt.Println("Friday")
		break
	case 7:
		fmt.Println("Saturday")
		break
	default:
		fmt.Println("Not Valid day number")
	}*/

	//Label Name break

Outer:
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i == 2 && j == 2 {
				break Outer
			}
			fmt.Println(i, j)
		}
	}
}
