package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Input number:")
	fmt.Scan(&number)

	if number % 3 == 0 && number % 5 == 0 {
		fmt.Println("FIZZBUZZ")
	} else if number % 3 == 0 {
		fmt.Println("FIZZ")
	} else if number % 5 == 0 {
		fmt.Println("BUZZ")
	}
	
	fmt.Println("The End")
}