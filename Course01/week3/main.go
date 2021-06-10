package main

import "fmt"

// Currency a special type for currencies signs
type Currency int

const (
	USD Currency = iota // Amarican dollar sign
	EUR
	RUB
	GBP
	RMB
)

func main() {

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RUB: "₽", RMB: "¥"}

	fmt.Println(RMB, symbol[RMB])
}
