package main

import "fmt"

func main() {
	ch := make(chan int) // Unbuffered channel

	// DEADLOCK: Sending without a receiver in a separate goroutine
	// ch <- 42
	// fmt.Println(<-ch)

	// FIX: Use a goroutine to send or read asynchronously
	go func() {
		ch <- 42
	}()

	fmt.Println("Received:", <-ch)
}