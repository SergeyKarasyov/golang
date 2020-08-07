package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("/tmp/")
	if err != nil {
		panic(nil)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
