package main

import "fmt"

func main() {
	messageChannel := make(chan string)

	// Non-blocking receive
	select {
	case message := <-messageChannel:
		fmt.Println("Received:", message)

	default:
		fmt.Println("No message available")
	}

	// Non-blocking send
	select {
	case messageChannel <- "Employee completed":
		fmt.Println("Message sent")

	default:
		fmt.Println("Receiver is not ready")
	}
}
