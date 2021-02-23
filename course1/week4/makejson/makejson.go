package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	user := make(map[string]string)

	var name, address string

	// Read user name
	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Can't read your name: %s", name)
	}

	// Read user address
	fmt.Print("Enter your address: ")
	reader = bufio.NewReader(os.Stdin)
	address, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Can't read your address: %s", address)
	}

	user["name"] = name
	user["address"] = address

	// Marshal map to JSON
	output, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Can't marshal map to JSON. Error: %v", err)
	}

	fmt.Printf("\nYour personal data: \n%s\n", string(output))
}
