package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println("Anonymous Goroutine executed:", msg)
	}("Hello Go!")

	wg.Wait()
}