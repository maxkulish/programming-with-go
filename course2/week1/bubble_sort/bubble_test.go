package main

import "testing"

func TestBubbleSort(t *testing.T) {
	t.Helper()

	f := func(slc []int, expSlc []int) {

		resSlc := scl[:]
		BubbleSort(resSlc)

		if len(resSlc) != len(expSlc) {
			t.Fatalf("unexpected result; got\n%v\nwant\n%v", slc, expSlc)
		}
	}

	f([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
}
