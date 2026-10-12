package main

import (
	"context" // Imports context for cancellation.
	"fmt"     // Imports fmt for output.
	"time"    // Imports time for delay.
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Stops when cancellation happens.
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // Creates a cancellable context.

	go worker(ctx) // Starts the worker.

	time.Sleep(2 * time.Second) // Lets the worker run for 2 seconds.

	cancel() // Cancels the context.

	time.Sleep(500 * time.Millisecond) // Gives worker time to stop.

	fmt.Println("Main finished")
}
