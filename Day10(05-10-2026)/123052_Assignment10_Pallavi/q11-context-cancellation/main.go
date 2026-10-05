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

	// Allow worker to run
	time.Sleep(2 * time.Second)

	// Cancel the worker
	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Finished")
}