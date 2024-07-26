package main

import (
	"fmt"
)

// This function takes a number and returns its square
func square(num float64) float64 {
	return num * num
}

func main() {
	// We want to find the square of the number 5
	number := 5.0
	result := square(number)
	fmt.Printf("The square of %.2f is %.2f\n", number, result)
}
