package main

import "fmt"

type NameAge struct {
	Name string
	Age  int
}

func main() {
	var nameAgeSlice []NameAge
	nameAges := map[string]int{
		"A": 0,
		"B": 1,
	}
	for key, value := range nameAges {
		nameAgeSlice = append(nameAgeSlice, NameAge{key, value})
	}
	fmt.Println(nameAgeSlice)

}
