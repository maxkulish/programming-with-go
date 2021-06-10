package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fileName string

	fmt.Print("Enter the name of the text file: ")
	fmt.Scan(&fileName)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("can't read your file: %s. Error: %v\n", fileName, err)
		os.Exit(1)
	}

	type Name struct {
		Fname string
		Lname string
	}

	var result []Name
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		name := strings.Split(line, " ")

		if len(name) > 0 && len(name) < 3 {
			result = append(result, Name{
				Fname: name[0],
				Lname: name[1],
			})
		} else {
			fmt.Printf("something strange in this line: %s\n", line)
			continue
		}

	}

	fmt.Println("=================")
	fmt.Printf("Names from you file: %s\n", fileName)
	for i, n := range result {
		fmt.Printf("%d %s %s\n", i, n.Fname, n.Lname)
	}

}
