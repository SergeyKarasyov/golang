package main

import "fmt"

func main() {
	map1 := map[string]int{
		"A": 1,
		"b": 3,
	}
	map2 := map[string]int{
		"c": 1,
		"D": 8,
	}
	fmt.Println(map1)
	fmt.Println(map2)
	for key, value := range map2 {

		map1[key] = value
	}
	fmt.Println(map1)

}
