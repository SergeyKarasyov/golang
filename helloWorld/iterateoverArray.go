package main

import "fmt"

func main() {
	numbers := []int{1, 5, 6}
	for index, value := range numbers {
		fmt.Printf("index : %v, value: %v\n", index, value)
	}
}
