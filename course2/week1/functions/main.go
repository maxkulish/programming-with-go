package main

import (
	"fmt"
	"strings"
)

func arrayInc(x *[3]int) int {
	(*x)[0] = (*x)[0] + 1
	return (*x)[0]
}

func sliceInc(sli []int) int {
	sli[0] = sli[0] + 1
	return sli[0]
}

func square(n int) int { return n * n }

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	a := [3]int{1, 2, 3}
	arrayInc(&a)
	fmt.Println(a)

	b := []int{1, 2, 3}
	sliceInc(b)
	fmt.Println(b)

	f := square
	fmt.Println(f(3))

	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))

	fsquare := squares()
	fmt.Println(fsquare())
	fmt.Println(fsquare())
	fmt.Println(fsquare())
	fmt.Println(fsquare())

	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
