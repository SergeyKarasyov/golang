package main

import (
	"fmt"
	"time"
)

func main() {
	selector()
}

func selector() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hola de channel1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Hola de channel2"
	}()
	var result string
	select {
	case result = <-channel1:
		fmt.Println(result)
	case result = <-channel2:
		fmt.Println(result)
	}
}
