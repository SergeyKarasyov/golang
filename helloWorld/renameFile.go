package main

import "os"

func main() {
	os.Rename("/tmp/old.txt", "/tmp/new.txt")

}
