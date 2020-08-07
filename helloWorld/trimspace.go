package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "\t Hello, world"
	trim_greetings := strings.TrimSpace(greetings)
	fmt.Printf("%d '%s' %d '%s' \n", len(greetings), greetings, len(trim_greetings), trim_greetings)
}
