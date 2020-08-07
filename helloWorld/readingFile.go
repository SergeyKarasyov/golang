package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contentBytes, err := ioutil.ReadFile("/tmp/log.txt")
	if err == nil {
		var contentStr string = string(contentBytes)
		fmt.Println(contentStr)
	}
}
