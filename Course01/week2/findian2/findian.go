package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	prefix = "i"
	middle = "a"
	suffix = "n"
)

func main() {

	fmt.Print("Enter a string: ")

	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("an error occured: %s\n", err.Error())
	}

	lower := strings.ToLower(strings.TrimSuffix(input, "\n"))
	if strings.HasPrefix(lower, prefix) &&
		strings.Contains(lower, middle) &&
		strings.HasSuffix(lower, suffix) {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
	return
}
