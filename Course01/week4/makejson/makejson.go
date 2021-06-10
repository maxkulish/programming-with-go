package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	input := map[string]string{
		"name":    "",
		"address": "",
	}

	var name string
	fmt.Print("Enter your name: ")
	r := bufio.NewReader(os.Stdin)
	name, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("an error occured: %s\n", err.Error())
	}

	var address string
	fmt.Print("Enter your address: ")
	address, err = r.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("an error occured: %s\n", err.Error())
	}

	if name != "" && address != "" {
		input["name"] = strings.TrimSuffix(name, "\n")
		input["address"] = strings.TrimSuffix(address, "\n")
	}

	result, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Printf("can't marhsal json: %v", err)
	}

	fmt.Println("=============")
	fmt.Println(string(result))
}
