package main

import "fmt"

func main() {
	helloWorldByte := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100}
	fmt.Println(string(helloWorldByte))
	fmt.Println([]byte(string(helloWorldByte)))
}
