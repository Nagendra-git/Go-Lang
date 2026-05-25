package main

import "fmt"

func main() {

	/*switch no := 22; no % 2 {
	case 0:
		fmt.Println("Even Number")
	case 1:
		fmt.Println("Odd Number")
	}*/

	//simulate nested-if in switch

	/*switch no := 21; {
	case no%2 == 0:
		fmt.Println("Even number")
	case no%2 != 0:
		fmt.Println("Odd number")
	default:
		fmt.Println("Neither an even not an odd")
	}*/

	//fall-through(opposite of 'break')

	/*switch no := 4; no {
	case 0:
		fmt.Print("no == 0")
		fallthrough
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("no <= 3")
		fallthrough
	case 4:
		fmt.Println((no <= 4))
		fallthrough
	case 5:
		fmt.Println("no <= 5")
		fallthrough
	case 6:
		fmt.Println("no <= 6")
		// fallthrough
	case 7:
		fmt.Println("no <= 7")

	}*/

	//fallthrough applied

	fmt.Println("fallthrough applied")

	switch plan := "SUPER"; plan {
	case "SUPREME":
		fmt.Println("[SUPREME] : offline download allowed!")
		fallthrough
	case "SUPER":
		fmt.Println("[SUPER] : For a family of 3")
		fallthrough
	case "PRO":
		fmt.Println("[PRO] : Create you own playlist")
		fallthrough
	case "FREE":
		fmt.Println("[FREE] : Listen to music!")
	}
}
