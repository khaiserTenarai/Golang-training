package main

import (
	"fmt"  // Imports fmt for printing output.
	"sync" // Imports sync for WaitGroup.
)

// fanIn combines multiple channels into one channel.
func fanIn(ch1 <-chan string, ch2 <-chan string) <-chan string {

	out := make(chan string) // Creates one output channel.

	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(2)             // Registers two input goroutines.

	// Reads data from the first channel.
	go func() {
		defer wg.Done() // Signals that this goroutine is finished.

		for value := range ch1 { // Receives values from channel 1.
			out <- value // Sends the value to the output channel.
		}
	}()

	// Reads data from the second channel.
	go func() {
		defer wg.Done() // Signals that this goroutine is finished.

		for value := range ch2 { // Receives values from channel 2.
			out <- value // Sends the value to the output channel.
		}
	}()

	// Closes the output channel after both inputs are finished.
	go func() {
		wg.Wait()  // Waits for both input goroutines.
		close(out) // Closes the combined output channel.
	}()

	return out // Returns the combined channel.
}

func main() {

	ch1 := make(chan string) // Creates the first input channel.
	ch2 := make(chan string) // Creates the second input channel.

	// Sends data to the first channel.
	go func() {
		ch1 <- "Apple"  // Sends Apple.
		ch1 <- "Banana" // Sends Banana.
		close(ch1)      // Closes the first channel.
	}()

	// Sends data to the second channel.
	go func() {
		ch2 <- "Car"  // Sends Car.
		ch2 <- "Bike" // Sends Bike.
		close(ch2)    // Closes the second channel.
	}()

	result := fanIn(ch1, ch2) // Combines both channels into one.

	// Reads values from the combined channel.
	for value := range result {
		fmt.Println(value) // Prints each received value.
	}

	fmt.Println("Finished") // Prints after all values are received.
}
