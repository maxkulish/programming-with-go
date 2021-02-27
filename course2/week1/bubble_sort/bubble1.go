package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var numbers []int

	for i := 0; i < 10; i++ {
		fmt.Printf("enter an integer:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		number, _ := strconv.Atoi(scanner.Text())

		numbers = append(numbers, number)
	}

	bubbleSort(numbers)
	fmt.Println(numbers)
}

func bubbleSort(numbers []int) {
	for i := len(numbers); i > 0; i-- {

		for j := range numbers {
			if j != (len(numbers) - 1) {
				swap(numbers, j)
			}
		}
	}
}

func swap(numbers []int, i int) {
	current := numbers[i]
	next := numbers[i+1]
	if current > next {
		numbers[i] = next
		numbers[i+1] = current
	}
}
