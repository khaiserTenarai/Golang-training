package main

import (
	"context" // Imports context.
	"fmt"     // Imports fmt.
	"time"    // Imports time.
)

func slowOperation(ctx context.Context) error { // Defines an operation with timeout support.
	select { // Waits for completion or timeout.

	case <-time.After(3 * time.Second): // Simulates a three-second operation.
		fmt.Println("Operation completed") // Prints success.
		return nil                         // Returns successful completion.

	case <-ctx.Done(): // Receives timeout signal.
		return ctx.Err() // Returns the timeout error.
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second, // Sets a one-second timeout.
	)
	defer cancel() // Releases context resources.

	err := slowOperation(ctx) // Runs the slow operation.

	if err != nil { // Checks if an error occurred.
		fmt.Println("Operation failed:", err) // Prints the timeout error.
	}
}
