package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	first  = "i"
	middle = "a"
	last   = "n"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var inputStr string
	fmt.Print("Please type a string: ")
	for scanner.Scan() {
		inputStr = scanner.Text()
		break
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Can't read your string", err)
	}

	result := MatchString(inputStr, first, middle, last)

	fmt.Println(result)
}

// MatchString check if input text containes 3 strings
// first - string starts with this character
// middle - contains this character
// last - ends with this character
func MatchString(text, first, middle, last string) string {

	lowText := strings.ToLower(text)

	if strings.HasPrefix(lowText, first) &&
		strings.HasSuffix(lowText, last) &&
		strings.Contains(lowText, middle) {
		return "Found!"
	}

	return "Not Found!"
}
