package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "hello World"
	helloWorldCapitalized := strings.Title(helloWorld)
	fmt.Println(helloWorldCapitalized)
	helloWorldAllCapitalized := strings.ToUpper(helloWorld)
	fmt.Println(helloWorldAllCapitalized)
}
