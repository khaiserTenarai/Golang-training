// Tasks 11, 12: Context
// Task 11 — Context cancellation
// Task 12 — Timeout handling
// Run: go run main.go

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// ============================================================
// Task 11: Implement Context Cancellation
// ============================================================

func longRunningTask(ctx context.Context, id int, results chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  [Worker %d] Cancelled: %v\n", id, ctx.Err())
			results <- fmt.Sprintf("Worker %d: cancelled", id)
			return
		default:
			// Simulate work
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("  [Worker %d] Working...\n", id)
		}
	}
}

func demoCancellation() {
	fmt.Println("========== TASK 11: Context Cancellation ==========")

	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan string, 3)

	// Launch 3 workers
	for i := 1; i <= 3; i++ {
		go longRunningTask(ctx, i, results)
	}

	// Let them work for 800ms then cancel
	time.Sleep(800 * time.Millisecond)
	fmt.Println("  [Main] Calling cancel()...")
	cancel() // All workers receive ctx.Done()

	// Collect cancellation confirmations
	for i := 0; i < 3; i++ {
		fmt.Printf("  [Main] %s\n", <-results)
	}
	fmt.Println()
}

// ============================================================
// Task 12: Implement Timeout Handling
// ============================================================

func fetchFromAPI(ctx context.Context, apiName string) (string, error) {
	// Simulate variable response time
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond

	select {
	case <-time.After(delay):
		return fmt.Sprintf("%s responded in %v", apiName, delay), nil
	case <-ctx.Done():
		return "", fmt.Errorf("%s: %w", apiName, ctx.Err())
	}
}

func demoTimeout() {
	fmt.Println("========== TASK 12: Timeout Handling ==========")

	// Timeout after 250ms
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	apis := []string{"UserService", "OrderService", "PaymentService", "NotificationService"}
	results := make(chan string, len(apis))

	for _, api := range apis {
		go func(name string) {
			result, err := fetchFromAPI(ctx, name)
			if err != nil {
				results <- fmt.Sprintf("  [TIMEOUT] %s", err)
			} else {
				results <- fmt.Sprintf("  [OK] %s", result)
			}
		}(api)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-results)
	}

	// Deadline-based timeout
	fmt.Println("\n  --- Deadline-Based Timeout ---")
	deadline := time.Now().Add(300 * time.Millisecond)
	ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  Slow operation completed")
	case <-ctx2.Done():
		fmt.Printf("  Deadline exceeded: %v\n", ctx2.Err())
	}
	fmt.Println()
}

func main() {
	demoCancellation()
	demoTimeout()
}
