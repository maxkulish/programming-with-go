package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var inputSome string

	fmt.Println("Please enter floating point: ")
	fmt.Scan(&inputSome)

	inputFloat, err := strconv.ParseFloat(inputSome, 64)
	if err != nil {
		fmt.Printf("You entered not a float number: %s", inputSome)
		return
	}

	outNum := math.RoundToEven(inputFloat)

	fmt.Printf("You entered: %s\nTruncated number: %.0f\n", inputSome, outNum)
}
