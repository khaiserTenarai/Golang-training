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
			fmt.Println("Worker is working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main function finished")
}
