package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}
}

func main() {
	// Set timeout to 2 seconds.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	slowOperation(ctx)

	fmt.Println("Main finished")
}
