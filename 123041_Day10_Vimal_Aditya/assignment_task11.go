package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doWork(ctx)

	time.Sleep(350 * time.Millisecond)
	fmt.Println("Triggering cancellation...")
	cancel()

	time.Sleep(100 * time.Millisecond)
}