package main

import "fmt"

// Deadlock idea:
// A receive from an unbuffered channel blocks until another goroutine sends.
// The fixed example below starts a goroutine that sends the value.

func main() {
	ch := make(chan int)

	// FIX: send data from another goroutine.
	go func() {
		ch <- 100
	}()

	// Receive data.
	value := <-ch

	fmt.Println(value)
}
