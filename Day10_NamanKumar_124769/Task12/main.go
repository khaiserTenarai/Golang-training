package main

import (
	"context"
	"fmt"
	"time"
)

func slowOp(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if err := slowOp(ctx); err != nil {
		fmt.Println("failed:", err) // context deadline exceeded
		return
	}
	fmt.Println("success")
}
