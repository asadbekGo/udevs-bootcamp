package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Input number:")
	fmt.Scan(&number)

	if number % 2 != 0 {
		fmt.Println("Odd number")
	} else if number % 2 == 0 {
		fmt.Println("Even number")
	}
}