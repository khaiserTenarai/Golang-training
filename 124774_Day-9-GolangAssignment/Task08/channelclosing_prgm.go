package main

import "fmt"

func main() {

	// Create an unbuffered integer channel
	ch := make(chan int)

	// Start a new goroutine
	go func() {

		// Send values from 1 to 5
		for i := 1; i <= 5; i++ {

			// Send the value into the channel
			ch <- i
		}

		// Close the channel after sending all values
		close(ch)

	}()

	// Continuously receive values from the channel
	for {

		// Receive a value and check whether the channel is open
		value, ok := <-ch

		// If ok is false, the channel is closed
		if !ok {
			fmt.Println("Channel is closed")
			break
		}

		// Print the received value
		fmt.Println("Received:", value)
	}

	// Program completed
	fmt.Println("Channel completed")
}
