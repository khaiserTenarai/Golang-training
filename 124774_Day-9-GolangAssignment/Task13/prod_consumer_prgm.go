package main

import "fmt"

// Producer function
func producer(ch chan<- int) {

	// Produce numbers from 1 to 5
	for i := 1; i <= 5; i++ {

		// Send the number to the channel
		ch <- i

		fmt.Println("Producer: Sent", i)
	}

	// Close the channel after producing all values
	close(ch)
}

// Consumer function
func consumer(ch <-chan int) {

	// Receive values from the channel
	for value := range ch {

		// Process the received value
		fmt.Println("Consumer: Received", value)
	}
}

func main() {

	// Create an unbuffered channel
	ch := make(chan int)

	// Start the producer goroutine
	go producer(ch)

	// Start the consumer
	consumer(ch)

	// Program completed
	fmt.Println("Producer-Consumer completed")
}
