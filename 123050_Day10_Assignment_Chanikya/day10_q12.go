package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()

	err := slowOperation(ctx)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Success")
}
