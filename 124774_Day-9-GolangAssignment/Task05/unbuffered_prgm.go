package main

import "fmt"

func main() {

	// Create an unbuffered channel
	message := make(chan string)

	// Start a new goroutine
	go func() {

		// Print a message before sending data
		fmt.Println("Employee: Preparing message")

		// Send a message through the channel
		message <- "Salary calculation completed"

		fmt.Println("Employee: Message sent")
	}()
	fmt.Println("Manager: Waiting for message")

	// Receive the message from the channel
	result := <-message

	// Print the received message
	fmt.Println("Manager received:", result)
}
