package main

import (
	calc "Users/nagen/GoProjects/Go-Lang/13-modularity-app/calculator"
	"Users/nagen/GoProjects/Go-Lang/13-modularity-app/calculator/utils"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Red("%q app started\n", appName)

	calc.Add(100, 200)
	calc.Sub(200, 100)

	calc.Add(300, 200)
	calc.Sub(300, 100)

	fmt.Println("OpCount = ", calc.OpCount())

	fmt.Println("IsEven(21) : ", utils.IsEven(21))
	fmt.Println("IsOdd(21) : ", utils.IsOdd(21))
}
