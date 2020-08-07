package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log_file, err := os.Create("/tmp/log.file")
	if err != nil {
		fmt.Println("An error occured")
	}
	defer log_file.Close()
	log.SetOutput(log_file)
	log.Println("1Doing some logging here...")
	//fatalln is followed with os.exit
	log.Fatalln("2Fatal:Doing some logging here...")
	log.Println("3Doing some logging here...")

}
