package main

import "fmt"

func main() {
	nameAges := map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println(nameAges["B"])
	fmt.Println(nameAges["C"])
	value, exists := nameAges["B"]
	value1, exists1 := nameAges["C"]
	fmt.Println(value, exists)
	fmt.Println(value1, exists1)

}
