package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("/tmp/hello")
	if err != nil {
		fmt.Println(err)
	}
	err1 := os.Remove("/tmp/hello")
	if err1 != nil {
		fmt.Println(err1)
	}
}
