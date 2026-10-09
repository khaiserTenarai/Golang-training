package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	// Non-blocking send
	select {
	case ch <- "Hello":
		fmt.Println("Message sent")
	default:
		fmt.Println("Channel is full, send skipped")
	}

	// Non-blocking receive
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available")
	}

	// Try receiving again
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available")
	}
}
