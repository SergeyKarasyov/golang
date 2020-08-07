package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("Recovered in function", r)
		}
	}()
	doPanik()
}

func doPanik() {
	panic("WriteOperationNOK")
}
