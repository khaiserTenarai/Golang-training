package main

import "fmt"

func main() {

	// Create an unbuffered integer channel
	ch := make(chan int)

	// Start a new goroutine
	go func() {

		// Send values from 1 to 5
		for i := 1; i <= 5; i++ {

			// Send the current value into the channel
			ch <- i
		}

		// Close the channel after sending all values
		close(ch)

	}()

	// Use range to receive values from the channel
	// The loop continues until the channel is closed
	for value := range ch {

		// Print the received value
		fmt.Println("Received:", value)
	}

	// This executes after the range loop ends
	fmt.Println("All values received")
}
