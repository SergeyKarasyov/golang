package main

import (
	"fmt"
	"time"
)

func main() {

	nameChannel2 := make(chan string)

	go func() {
		names := []string{"channel1", "channel2", "channel3"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
		nameChannel2 <- ""
	}()

	nameChannel := make(chan string)

	go func() {
		names := []string{"oneChannel", "twoChannel", "threeChannel"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			// fmt.Println(name)
			nameChannel <- name
		}
		close(nameChannel)
	}()
	go func() {
		names := []string{"one", "two", "three"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
	}()

	go func() {
		names := []int{1, 2, 3}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
	}()
	time.Sleep(2 * time.Second)
	for data := range nameChannel {
		fmt.Println(data)
	}

	<-nameChannel2

}
