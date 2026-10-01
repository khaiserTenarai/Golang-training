package main

import "fmt"

func main() {
	messages := make(chan string)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received (non-blocking)")
	}

	// Non-blocking send
	select {
	case messages <- "Hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("Could not send message (channel busy/unbuffered)")
	}
}