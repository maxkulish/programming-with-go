package main

import "fmt"

func main() {

	var a, v, s, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter velocity: ")
	fmt.Scan(&v)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, s)

	fmt.Println(fn(t))
}

// GenDisplaceFn takes three float64 arguments:
// acceleration a, initial velocity v, and initial displacement s
// GenDisplaceFn() should return a function
// which computes displacement as a function of time
func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 1/2*a*t*t + v*t + s
	}
}
