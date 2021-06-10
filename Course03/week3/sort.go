package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const parts = 4

	var allNum []int
	fmt.Println("Type min 4 integers separated by space. Press Enter to finish.")

	fmt.Print("Slice: ")
	r := bufio.NewReader(os.Stdin)

	input, _, err := r.ReadLine()
	if err != nil && err != io.EOF {
		fmt.Printf("an error occured: %s\n", err.Error())
	}

	// skip letters and add numbers to the slice
	allNum = normalizeInput(string(input))

	// devide unsorted slice to the parts
	chunk := len(allNum) / parts
	fmt.Println("chunk: ", chunk)

	var end int
	var result []int
	for i := 0; i < parts; i++ {
		var slice []int

		start := end
		end += chunk
		slice = allNum[start:end]

		if i+1 == parts {
			slice = allNum[start:]
		}

		wg.Add(1)
		go sortSlice(slice)

		fmt.Printf("sorted subslice: %v\n", slice)
		result = append(result, slice...)
	}

	wg.Wait()
	wg.Add(1)
	sortSlice(result)
	fmt.Printf("final sorted slice: %v", result)
}

func sortSlice(s []int) {
	defer wg.Done()
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

func normalizeInput(input string) []int {
	var result []int
	for _, n := range strings.Split(input, " ") {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		result = append(result, num)
	}
	return result
}
