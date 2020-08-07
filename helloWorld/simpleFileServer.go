package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("/tmp/images/")))
	http.ListenAndServe(":5051", nil)
}
