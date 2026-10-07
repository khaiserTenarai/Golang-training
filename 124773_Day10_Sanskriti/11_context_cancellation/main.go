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
			fmt.Println("Worker cancelled")
			return
		default:
			fmt.Println("Worker is working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	// Let the worker work for 2 seconds.
	time.Sleep(2 * time.Second)

	// Send cancellation signal.
	cancel()

	// Give the worker time to stop.
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Main finished")
}
