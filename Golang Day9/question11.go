package main

import "fmt"

func main() {
	
	messages := make(chan string, 1)

	fmt.Println("Attempting to read from an empty channel...")
	select {
	case msg := <-messages:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("[Default] No messages available right now. Moving on!")
	}

	fmt.Println("\nAttempting to send a message...")
	select {
	case messages <- "First Message":
		fmt.Println("Successfully sent the first message!")
	default:
		fmt.Println("[Default] Channel is full. Cannot send.")
	}

	fmt.Println("\nAttempting to send a second message...")
	select {
	case messages <- "Second Message":
		fmt.Println("Successfully sent the second message!")
	default:
		fmt.Println("[Default] Channel is full! Aborting send and moving on!")
	}
}