package main

import "fmt"

type MyError struct {
	ShortMessage    string
	DetailedMessage string
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func doSomething() error {
	//Doing something here...
	return &MyError{ShortMessage: "Wohoo something happened!", DetailedMessage: "File cannot found!"}
}

func main() {
	err := doSomething()
	fmt.Print(err)
}
