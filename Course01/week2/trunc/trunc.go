package main

import (
	"fmt"
	"math"
)

func main() {
	var userNumber float64
	fmt.Print("Enter float number: ")
	fmt.Scanln(&userNumber)

	fmt.Printf("Your float number: %.4f\n", userNumber)
	fmt.Printf("Truncated number: %.0f\n", math.Trunc(userNumber))
}
