package main

import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup

	fmt.Println("Starting program")

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from anonymous goroutine")
	}()

	customMessage := "Go Is ready"

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println("The custom message",msg)

	}(customMessage)
	wg.Wait()
	fmt.Println("Program Completed")
}