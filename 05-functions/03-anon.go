/*
anonymus functions
	- function with no name
	- have to be immediately invoked
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!", userName)
	}("Nagendra")
}
