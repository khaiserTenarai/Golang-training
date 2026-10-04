package main

import "fmt"

// sendMessage can only SEND data to the channel
func sendMessage(message chan<- string) {

	// Send a message into the channel
	message <- "Employee processing completed"
}

// receiveMessage can only RECEIVE data from the channel
func receiveMessage(message <-chan string) {

	// Receive the message from the channel
	result := <-message

	fmt.Println("Manager received:", result)
}

func main() {

	// Create a channel
	message := make(chan string)

	// Start a goroutine for sending
	go sendMessage(message)

	// Receive the message
	receiveMessage(message)
}
