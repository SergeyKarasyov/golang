package main

import "os"

func main() {
	err := os.Remove("/tmp/new.txt")
	if err != nil {
		panic(err)
	}
}
