package main

import (
	"fmt"
	"time"
)

func main() {
	// Limiting maximum queued requests using buffer capacity
	bufferSize := 3
	jobQueue := make(chan int, bufferSize)

	// Rapid request generator
	go func() {
		for i := 1; i <= 7; i++ {
			select {
			case jobQueue <- i:
				fmt.Printf("Job %d accepted\n", i)
			default:
				fmt.Printf("Job %d REJECTED due to backpressure!\n", i)
			}
			time.Sleep(50 * time.Millisecond)
		}
		close(jobQueue)
	}()

	// Slow worker
	for job := range jobQueue {
		fmt.Printf("Processing job %d...\n", job)
		time.Sleep(150 * time.Millisecond)
	}
}