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
			fmt.Println("Worker: context cancelled, stopping work")
			return
		default:
			fmt.Println("Worker: doing work")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)

	fmt.Println("Main: cancelling context")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: finished")
}