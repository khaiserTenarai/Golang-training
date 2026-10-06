// context.WithTimeout() automatically cancels the operation when the specified time expires...
package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")

	case <-ctx.Done():
		fmt.Println("Work timed out")
	}
}

func main() {
	/*
		Problem:
		Our operation may take too long.

		Solution:
		Set a timeout using context.WithTimeout().

		Here timeout = 2 seconds.
		But work takes 3 seconds.
		So timeout happens first.
	*/

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	doWork(ctx)

	fmt.Println("Main completed")
}