package main

import "fmt"

func main() {

	/*i := 1

	start:
		fmt.Println(i)
		i++

		if i <= 3 {
			goto start
		}
		fmt.Println("Loop finished")*/

	for m := 0; m < 4; m++ {

		for n := 0; n < 4; n++ {

			if m == 2 && n == 2 {
				goto EndLoop
			}
			fmt.Printf("m : %d, n = %d\n", m, n)
		}

	}
EndLoop:
	fmt.Println("Exited the loop")
}
