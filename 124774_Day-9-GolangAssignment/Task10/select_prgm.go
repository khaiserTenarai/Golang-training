//select is used to wait for multiple channel operations at the same time.
//  When one or more channel operations are ready, select chooses one ready case and executes it.

package main

import "fmt"

func main() {

	// Create two channels
	employee := make(chan string)
	manager := make(chan string)

	// Start a goroutine for the employee channel
	go func() {

		// Send a message to the employee channel
		employee <- "Employee task completed"
	}()

	// Start a goroutine for the manager channel
	go func() {

		// Send a message to the manager channel
		manager <- "Manager task completed"
	}()

	// Use select to wait for a channel operation
	select {

	// Receive a message from the employee channel
	case message := <-employee:
		fmt.Println("Received:", message)

	// Receive a message from the manager channel
	case message := <-manager:
		fmt.Println("Received:", message)
	}
}
