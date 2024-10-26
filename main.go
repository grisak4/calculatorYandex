package main

import (
	"calc/calculator"
	"fmt"
)

func main() {
	someMath := "2 + 2 * 4"
	fmt.Println(calculator.Calc(someMath))
}
