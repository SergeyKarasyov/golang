package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := "2"
	valueInt, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		fmt.Print("Error happened!")
	} else {
		if valueInt == 2 {
			fmt.Println("Success")
		}
	}
	fnumber := "2.2222"
	valueFloat, errF := strconv.ParseFloat(fnumber, 64)
	if errF != nil {
		fmt.Print("Error happened!")
	} else {
		if valueFloat == 2.2222 {
			fmt.Println("Success")
		}
	}
}
