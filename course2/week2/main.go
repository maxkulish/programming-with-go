package main

import "fmt"

func main() {
	funcVar := incFn
	fmt.Println(funcVar(3))

	fmt.Println(applyFunc(funcVar, 4))
	fmt.Println(applyFunc(incFn, 5))
	fmt.Println(applyFunc(decFn, 10))

	f = test

	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}

func applyFunc(afunc func(int) int, val int) int {
	return afunc(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x + 1 }

var f func(string) int

func test(x string) int {
	return len(x)
}

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
