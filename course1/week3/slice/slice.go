package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	fmt.Println("Enter itegers in each line\nType 'X' to quit.")

	slc := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "x" || input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("You entered not a number: <%s>. Error: %v", input, err)
		}

		slc = append(slc, num)
		sort.Ints(slc)
		fmt.Println(slc)
	}

	fmt.Printf("Bye! Your sorted integers: %v", slc)

}
