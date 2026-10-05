package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		// Send data from another goroutine
		ch <- "Employee Data"
	}()

	// Receive data
	message := <-ch

	fmt.Println(message)

	wg.Wait()

	fmt.Println("Deadlock fixed")
}	