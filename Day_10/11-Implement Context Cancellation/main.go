// context.WithCancel() allows us to stop goroutines when we no longer need them.

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Context was cancelled.
			fmt.Println("Worker stopped")
			return

		default:
			fmt.Println("Worker is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	/*
		Problem:
		A goroutine may continue running even when
		the main program no longer needs it.

		Solution:
		Use context cancellation.

		cancel() sends a cancellation signal.
	*/

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	// Allow worker to run for some time.
	time.Sleep(2 * time.Second)

	// Cancel the worker.
	cancel()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Main completed")
}