package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance(name string) { fmt.Printf("%s balance: %d\n", name, balance) }

func Withdraw(amount int) { balance = balance - amount }

func main() {

	var wg sync.WaitGroup

	// Alice operations with the account
	// Alice add 200 to the account, but we don't know what the result would be
	// Balance could be 0, 200 or 300
	// The reason is - we can't guaranty which goroutine will add money first
	// and which will withdraw money
	wg.Add(1)
	go func() {
		defer wg.Done()

		Deposit(200)
		Balance("Alice")
	}()

	// Max operations
	// Max withdraws 100 print balance and after that add 200 to the account
	// Balance could be 0, 100, 300, 500 because we can't predict goroutines order
	wg.Add(1)
	go func() {
		defer wg.Done()
		Withdraw(100)
		Balance("Max")
		Deposit(200)
		Balance("Max")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(200)
		Balance("Max")
	}()

	// Total balance will be 0 if we don't wait goroutines to finish
	// or 500 if we would wait for all goroutines
	// time.Sleep(100 * time.Millisecond)
	wg.Wait()
	Balance("Total")
}
