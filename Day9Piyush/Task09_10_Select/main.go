// Tasks 9, 10: Select Statement and Non-Blocking Operations
// Task 9  — Use select
// Task 10 — Implement non-blocking channel operations
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

// ============================================================
// Task 9: Use Select
// ============================================================

func demoSelect() {
	fmt.Println("========== TASK 9: Select Statement ==========")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "Result from Service A"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Result from Service B"
	}()

	// Select waits on multiple channels — whichever is ready first wins
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("  [Select] ch1:", msg)
		case msg := <-ch2:
			fmt.Println("  [Select] ch2:", msg)
		}
	}

	// Select with timeout
	fmt.Println("\n  --- Select with Timeout ---")
	ch3 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second) // slow response
		ch3 <- "Slow result"
	}()

	select {
	case msg := <-ch3:
		fmt.Println("  Got:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  [Timeout] No response within 500ms, moving on.")
	}

	// Select with multiple cases for fan-in
	fmt.Println("\n  --- Fan-In with Select ---")
	emails := make(chan string, 3)
	sms := make(chan string, 3)

	go func() { emails <- "email: piyush@test.com" }()
	go func() { sms <- "sms: +91-9999999999" }()
	go func() { emails <- "email: alice@test.com" }()

	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case e := <-emails:
			fmt.Println("  [Fan-In]", e)
		case s := <-sms:
			fmt.Println("  [Fan-In]", s)
		}
	}
	fmt.Println()
}

// ============================================================
// Task 10: Non-Blocking Channel Operations
// ============================================================

func demoNonBlocking() {
	fmt.Println("========== TASK 10: Non-Blocking Channel Operations ==========")

	ch := make(chan int, 1)

	// Non-blocking send
	select {
	case ch <- 42:
		fmt.Println("  [Non-Block Send] Sent 42 to channel")
	default:
		fmt.Println("  [Non-Block Send] Channel full, skipping send")
	}

	// Non-blocking receive
	select {
	case val := <-ch:
		fmt.Printf("  [Non-Block Recv] Received: %d\n", val)
	default:
		fmt.Println("  [Non-Block Recv] No data available")
	}

	// Non-blocking receive on empty channel
	select {
	case val := <-ch:
		fmt.Printf("  [Non-Block Recv] Got: %d\n", val)
	default:
		fmt.Println("  [Non-Block Recv] Channel empty, nothing to read")
	}

	// Non-blocking multi-channel
	fmt.Println("\n  --- Non-Blocking Multi-Channel ---")
	chA := make(chan string, 1)
	chB := make(chan string, 1)
	chA <- "Message from A"

	select {
	case msg := <-chA:
		fmt.Println("  Got from A:", msg)
	case msg := <-chB:
		fmt.Println("  Got from B:", msg)
	default:
		fmt.Println("  Neither channel had data")
	}
	fmt.Println()
}

func main() {
	demoSelect()
	demoNonBlocking()
}
