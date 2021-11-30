package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Input number: ")
	fmt.Scan(&number)

	sum, temp := 0, number

	for temp > 0 {

		remain := temp % 10

		sum = sum * 10 + remain

		temp /= 10
	}

	if sum == number {
		fmt.Println("Palindrom number")
		return
	}

	fmt.Println("Not Palindrome")
}