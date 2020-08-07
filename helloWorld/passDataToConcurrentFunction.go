package main

import "fmt"

func main() {
	noncached()
	cached()
}

func cached() {
	nameChannel := make(chan string, 5)
	done := make(chan string)

	go func() {
		names := []string{"first", "second", "third", "fourth", "fifth", "sixth"}
		for _, name := range names {
			fmt.Println("cached FUNC1: " + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	go func() {
		for name := range nameChannel {
			fmt.Println("cached FUNC2: " + name)
		}
		done <- ""
	}()
	<-done
}

func noncached() {
	nameChannel := make(chan string)
	done := make(chan string)

	go func() {
		names := []string{"first", "second", "third"}
		for _, name := range names {
			fmt.Println("FUNC1: " + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	go func() {
		for name := range nameChannel {
			fmt.Println("FUNC2: " + name)
		}
		done <- ""
	}()
	<-done
}
