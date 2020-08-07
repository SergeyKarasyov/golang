package main

import (
	"fmt"
	"strconv"
)

func main() {
	isNew := true
	message := "Purchased item is " + strconv.FormatBool(isNew)
	fmt.Println(message)
	fmt.Printf("%s %v", message, isNew)
}
