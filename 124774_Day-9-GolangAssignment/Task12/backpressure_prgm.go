package main

import (
	"fmt"
	"time"
)

func main() {

	// Create an unbuffered channel
	ch := make(chan int)

	// Start the producer goroutine
	go func() {
		// Produce 5 values
		for i := 1; i <= 5; i++ {
			// Send the value to the channel
			// The producer waits here until the consumer receives it
			fmt.Println("Producer: Sending", i)
			ch <- i

			fmt.Println("Producer: Sent", i)
		}

		// Close the channel after producing all values
		close(ch)
	}()

	// Consumer receives values from the channel
	for value := range ch {
		fmt.Println("Consumer: Received", value)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Processing completed")
}
