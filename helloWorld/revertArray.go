package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 8, 2, 3, 4, 5, 5, 6}
	fmt.Println(numbers)
	tobeSorted := sort.IntSlice(numbers)
	sort.Sort(tobeSorted)
	fmt.Println(tobeSorted)
	tobeReverted := sort.IntSlice(numbers)
	sort.Sort(sort.Reverse(tobeReverted))
	fmt.Println(tobeReverted)
}
