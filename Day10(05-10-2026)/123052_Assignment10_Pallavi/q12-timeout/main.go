package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Create a context with a 2-second timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	// Simulate a 5-second operation
	select {

	case <-time.After(5 * time.Second):
		fmt.Println("Operation completed")

	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}
}