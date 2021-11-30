package main

import (
	"fmt"
)

func main() {

	var a, b, c, limit int

	fmt.Println("Fibonacci number limit:")
	fmt.Scan(&limit)

	b = 1

	for i := 0; i < limit; i++ {

		fmt.Printf("%d ", a + b)
		c = a + b
		a = b
		b = c
	}
}