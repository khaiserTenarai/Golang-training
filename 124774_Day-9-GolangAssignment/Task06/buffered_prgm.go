package main

import "fmt"

func main() {

	// Create a buffered channel with a capacity of 3
	message := make(chan string, 3)

	// Send the first message into the channel
	message <- "Employee 1 completed"

	// Send the second message into the channel
	message <- "Employee 2 completed"

	// Send the third message into the channel
	message <- "Employee 3 completed"

	// Rfirst message
	fmt.Println(<-message)

	// R second message
	fmt.Println(<-message)

	// third message
	fmt.Println(<-message)
}
