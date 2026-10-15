package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered channel acting as a queue with a capacity of 3
	queue := make(chan int, 3)

	// FAST PRODUCER
	go func() {
		for i := 1; i <= 6; i++ {
			fmt.Printf("[Fast Producer] Attempting to send item %d...\n", i)
			
			// This will be instant for the first 3 items because of the buffer.
			// From item 4 onwards, it will block (backpressure) until the consumer frees up space.
			queue <- i
			
			fmt.Printf("[Fast Producer] Successfully sent item %d\n", i)
			time.Sleep(50 * time.Millisecond) // Producer works very fast
		}
		close(queue)
		fmt.Println("[Fast Producer] Finished and closed queue.")
	}()

	// SLOW CONSUMER
	// Consuming items from the main thread to demonstrate the delay
	for item := range queue {
		fmt.Printf("[Slow Consumer] Processing item %d...\n", item)
		time.Sleep(400 * time.Millisecond) // Consumer works much slower than producer
		fmt.Printf("[Slow Consumer] Finished item %d\n", item)
	}

	fmt.Println("Program completed...")
}