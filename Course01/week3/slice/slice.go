package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	s := make([]int, 3)

	for i, j := 0, len(s); i < j; i, j = i+1, j+1 {
		var input string
		fmt.Print("Enter an integer: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("error:", err)
		}

		if input == "X" || input == "x" {
			break
		}

		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("you entered not an integer: %s\n", input)
			continue
		}
		s = append(s, n)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

		fmt.Printf("Sorted slice: %v\n", s[3:])
	}

	fmt.Println("=================")
	fmt.Printf("Final slice: %v\n", s[3:])
}
