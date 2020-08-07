package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 2, 3, 4}
	fmt.Println(intSlice)
	uniqueIntSlice := unique(intSlice)
	fmt.Println(uniqueIntSlice)
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	uniqueElements := []int{}
	for _, entry := range intSlice {

		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements
}
