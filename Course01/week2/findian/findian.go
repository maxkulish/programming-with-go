package main

import (
	"fmt"
	"strings"
)

func main() {

	var userString string

	fmt.Print("Enter a string: ")
	n, err := fmt.Scanf("%v", &userString)
	if err != nil {
		fmt.Printf("n: %d, %+v\n, error: %v", n, userString, err)
	}

	userString = strings.ToLower(userString)

	if strings.HasPrefix(userString, "i") &&
		strings.Contains(userString, "a") &&
		strings.HasSuffix(userString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
