package main

import "fmt"

func main() {

	sayHi()
	greet("Nagendra")
	greetFullName("Nagendra", "Burusu")
	msg := getGreet("Nagendra")
	fmt.Println(msg)
	fmt.Println(add(100, 200))
	fmt.Println(devide(100, 7))
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d, remainder = %d\n", q, r)

}

func sayHi() {
	fmt.Println("Hi!")
}

func greet(userName string) {
	fmt.Println("Hello", userName)
}

func greetFullName(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

func getGreet(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

func add(a, b int) int {
	return a + b
}

func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
func devide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
