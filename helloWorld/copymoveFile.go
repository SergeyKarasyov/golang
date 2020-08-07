package main

import (
	"io"
	"os"
)

func main() {
	original, err := os.Open("/tmp/original.txt")
	if err != nil {
		panic(err)
	}
	defer original.Close()

	original_copy, err2 := os.Create("/tmp/copied.txt")
	if err2 != nil {
		panic(err2)
	}
	defer original_copy.Close()

	_, err3 := io.Copy(original, original_copy)
	if err3 != nil {
		panic(err3)
	}
}
