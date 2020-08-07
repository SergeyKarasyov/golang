package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	nosync()
	withsync()
}

func nosync() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Hello world")
		}()
	}

	time.Sleep(1 * time.Second)
}

func withsync() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello world")
			wg.Done()
		}()
	}

	wg.Wait()
}
