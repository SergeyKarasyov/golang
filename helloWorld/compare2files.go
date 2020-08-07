package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	one, err := ioutil.ReadFile("/tmp/data1.txt")
	if err != nil {
		panic(err)
	}
	two, err2 := ioutil.ReadFile("/tmp/data2.txt")
	if err2 != nil {
		panic(err2)
	}
	same := bytes.Equal(one, two)
	fmt.Println(same)
}
