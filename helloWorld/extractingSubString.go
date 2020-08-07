package main

import (
	"fmt"
)

func main() {
	greetings := "\t Hello, world and Mars"
	nomarsstring := greetings[0:15]
	fmt.Printf("%d '%s' %d '%s' \n", len(greetings), greetings, len(nomarsstring), nomarsstring)
}
