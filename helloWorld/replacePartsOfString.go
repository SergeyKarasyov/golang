package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "Hello, World"
	helloMars := strings.Replace(helloWorld, "World", "Mars", -1)
	fmt.Println(helloMars)
}
