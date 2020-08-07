package main

import "fmt"

func main() {
	items1 := []int{3, 4}
	items2 := []int{5, 6}
	result := append(items1, items2...)
	fmt.Println(result)
}
