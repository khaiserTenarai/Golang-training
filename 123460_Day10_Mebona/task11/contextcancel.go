package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Cancelling context...")
	cancel()

	time.Sleep(100 * time.Millisecond)
}