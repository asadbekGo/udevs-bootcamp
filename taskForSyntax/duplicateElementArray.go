package main

import (
	"fmt"
)

func main() {

	var check bool

	numbers := []int { 3, 4, 2, 5, 0, 8, 7, 4, 4, 21, 9 }

	for n, i := range numbers {

		for j := n; j < len(numbers) - 1; j++ {

			if i == numbers[j + 1] {
				check = true
				break
			}
		}
	}

	if check {
		fmt.Println("Duplicate array")
		return
	}

	fmt.Println("Not Duplicate array")
}