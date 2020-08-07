package main

import (
	"errors"
	"fmt"
	"time"
)

func doSomething() (string, error) {
	return "", errors.New("Something happened.")
}

func main() {
	parsedDate, err := time.Parse("2006", "2018")
	if err != nil {
		fmt.Println("An error occured", err.Error())
	} else {
		fmt.Println(parsedDate)
	}
	parsedDate1, err1 := time.Parse("1", "2018")
	if err1 != nil {
		fmt.Println("An error occured", err1.Error())
	} else {
		fmt.Println(parsedDate1)
	}
	_, err34 := doSomething()
	if err34 != nil {
		fmt.Println(err34)
	}
}
