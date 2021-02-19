package main

import "testing"

func TestMatchString(t *testing.T) {
	t.Helper()

	f := func(text, first, middle, last, expResult string) {

		result := MatchString(text, middle, last)

	}

	f("ian", "i", "a", "n", "Found!")
	f("Ian", "i", "a", "n", "Found!")
	f("iuiygaygn", "i", "a", "n", "Found!")
	f("I d skd a efju N", "i", "a", "n", "Found!")
	f("ihhhhhn", "i", "a", "n", "Not Found!")
	f("ina", "i", "a", "n", "Not Found!")
	f("xian", "i", "a", "n", "Not Found!")
}
