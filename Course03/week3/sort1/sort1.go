package sort1

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func main() {
	parts := 4

	var allNum []int
	fmt.Println("Type integers separated by space. Press Enter to finish.")

	for i := 1; i <= parts; i++ {
		wg.Add(1)
		fmt.Printf("Slice %d: ", i)

		s := input([]int{}, nil)

		go sortSlice(s)

		wg.Wait()
		fmt.Printf("sorted slice %d: %v\n", i, s)
		allNum = append(allNum, s...)
	}

	wg.Add(1)
	sortSlice(allNum)
	fmt.Printf("final sorted slice: %v", allNum)

}

func sortSlice(s []int) {
	defer wg.Done()
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

func input(x []int, err error) []int {

	if err != nil {
		return x
	}

	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}

	return input(x, err)
}
