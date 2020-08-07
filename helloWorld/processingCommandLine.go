package main

import (
	"fmt"
	"os"
)

func main() {
	realArgs := os.Args[1:]
	fmt.Println(realArgs)

	if len(realArgs) == 0 {
		fmt.Println("no arguments")
		return
	}

	if realArgs[0] == "a" {
		writeHelloWorld()
	} else if realArgs[0] == "b" {
		writeHelloMars()
	} else {
		fmt.Println("please pass a valid argument")
	}
}

func writeHelloWorld() {
	fmt.Println("Hello World")
}

func writeHelloMars() {
	fmt.Println("Hello Mars")
}
