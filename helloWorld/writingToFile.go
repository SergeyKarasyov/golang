package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	hello := "/tmp/HelloWorld"
	err := ioutil.WriteFile(hello, []byte(hello), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
