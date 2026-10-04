package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Basic anonymous goroutine defined and invoked inline
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from an anonymous goroutine!")
	}()

	// Anonymous goroutine receiving parameters
	wg.Add(1)
	go func(name string) {
		defer wg.Done()
		fmt.Printf("Hello, %s!\n", name)
	}("Alice")

	wg.Wait()
}