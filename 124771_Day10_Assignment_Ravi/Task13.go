package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu1, mu2 sync.Mutex

	mu1.Lock()
	mu2.Lock()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1")
		mu2.Unlock()
		mu1.Unlock()
	}()

	mu2.Lock()
	fmt.Println("Main goroutine")
	mu1.Unlock()
	mu2.Unlock()
}
