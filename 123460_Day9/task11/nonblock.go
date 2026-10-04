package main

import (
	"fmt"
)

func main() {
	// Using buffered channels with a capacity of 1 so they can hold 
	// one value immediately without blocking.
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	// Non-blocking send 1 (Pre-filling buffer so it succeeds instead of falling back to default)
	messages <- "Hello from buffered channel!"

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message available (non-blocking receive)")
	}

	// Non-blocking send
	select {
	case signals <- true:
		fmt.Println("Sent signal successfully")
	default:
		fmt.Println("Channel not ready for send (non-blocking send)")
	}

	// Non-blocking receive for the signal we just sent
	select {
	case sig := <-signals:
		fmt.Printf("Received signal: %v\n", sig)
	default:
		fmt.Println("No signal available")
	}
}