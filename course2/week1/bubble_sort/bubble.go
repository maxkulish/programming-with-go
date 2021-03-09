package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter itegers in each line\nType 'x' to quit.")

	slc := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "x" || input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("You entered not a number: <%s>\n", input)
			fmt.Print("Type 'x' to quit or next number: ")
		}
		slc = append(slc, num)
	}

	fmt.Printf("Unsorted: %v\n", slc)

	BubbleSort(slc)
	fmt.Printf("Sorted: %v\n", slc)
}

// BubbleSort uses bubble sorting algorithm to sort integers in the slice
func BubbleSort(slc []int) {
	for range slc {
		for i := 0; i < len(slc)-1; i++ {
			if slc[i] > slc[i+1] {
				Swap(i, slc)
			}
		}
	}
}

// Swap takes slice and move one element with index i to 1 position ahead
func Swap(i int, slc []int) {
	slc[i], slc[i+1] = slc[i+1], slc[i]
}
