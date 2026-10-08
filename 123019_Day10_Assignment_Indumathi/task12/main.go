package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) (string, error) {
	select {
	case <-time.After(500 * time.Millisecond):
		return "Success!", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	res, err := slowOperation(ctx)
	if err != nil {
		fmt.Println("Operation failed:", err)
		return
	}
	fmt.Println("Operation result:", res)
}