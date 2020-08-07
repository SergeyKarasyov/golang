package main

import (
	"fmt"
	"strconv"
)

func main() {
	isNew := "true"
	isNewBool, err := strconv.ParseBool(isNew)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("not failed")
		if isNewBool {
			fmt.Print("IsNew")
		} else {
			fmt.Println("Not new")
		}
	}
}
