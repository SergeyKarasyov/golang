package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://golang.org"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	html, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic(err)
	}
	fmt.Println(string(html))
}
