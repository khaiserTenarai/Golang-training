package main

import "fmt"

func main() {

	// Create an unbuffered channel
	ch := make(chan string)

	// Start a goroutine to send a message
	// go func() {

	// 	 Send the message into the channel
	// 	ch <- "Hello from employee"

	// }()

	// Use select for a non-blocking receive
	select {

	// Try to receive a message from the channel
	case message := <-ch:
		fmt.Println("Received:", message)

	// If no message is available, do not wait
	default:
		fmt.Println("No message available")
	}

	// Program completed
	fmt.Println("Main completed")
}
