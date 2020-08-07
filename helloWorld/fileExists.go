package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("/tmp/log.txt"); err == nil {
		fmt.Println("/tmp/log.txt exists")
	}
	if _, err := os.Stat("/tmp/log.txt"); os.IsNotExist(err) {
		fmt.Println("/tmp/log.txt doesn't exist")
	} else {
		fmt.Println("/tmp/log.txt exists")
	}

}
