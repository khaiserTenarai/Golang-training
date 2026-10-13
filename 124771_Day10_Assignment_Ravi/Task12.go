package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(), 2*time.Second,
	)
	defer cancel()

	worker(ctx)
}
