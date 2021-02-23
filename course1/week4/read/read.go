package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// User type with fields: First Name and Last Name
type User struct {
	FName string
	LName string
}

// How to run
// go run ./read.go
// enter names.txt
func main() {

	var path string
	fmt.Print("Please enter path to your file: ")
	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Printf("Can't read your path: %s", path)
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Can't open your file: %s\n%v", path, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	var usersList []User
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var fname, lname string
		line := scanner.Text()
		if len(line) > 0 {
			names := strings.Split(line, " ")
			fname, lname = names[0], names[len(names)-1]
		}

		usersList = append(usersList, User{
			FName: fname,
			LName: lname,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Can't scan your file: %s. Error: %v", path, err.Error())
	}

	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t", "First Name", "Last Name")
	fmt.Fprintf(w, "\n %s\t%s\t", "----------", "---------")
	for _, user := range usersList {
		fmt.Fprintf(w, "\n %s\t%s\t", user.FName, user.LName)
	}

}
