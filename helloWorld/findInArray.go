package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"Sandy", "St.George", "Provo", "Aaaa"}
	for i, v := range str {
		if v == "Sandy" {
			fmt.Println(i, v)
		}
	}
	sortedList := sort.StringSlice(str)
	sortedList.Sort()
	fmt.Println(sortedList)
	for i, v := range str {
		if v == "Sandy" {
			fmt.Println(i, v)
		}
	}
	index := sortedList.Search("Sandy")
	fmt.Println(index)
}
