// Tasks 4, 5, 6, 7, 8: Channel Fundamentals
// Task 4  — Unbuffered channel
// Task 5  — Buffered channel
// Task 6  — Directional channels
// Task 7  — Channel closing
// Task 8  — Range over a channel
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

// ============================================================
// Task 4: Implement an Unbuffered Channel
// ============================================================

func demoUnbuffered() {
	fmt.Println("========== TASK 4: Unbuffered Channel ==========")

	ch := make(chan string) // unbuffered: sender blocks until receiver is ready

	go func() {
		fmt.Println("  [Sender] Sending message...")
		ch <- "Hello from goroutine!" // blocks until main reads
		fmt.Println("  [Sender] Message sent!")
	}()

	time.Sleep(500 * time.Millisecond) // simulate delay
	msg := <-ch                        // receive unblocks the sender
	fmt.Printf("  [Receiver] Got: %s\n\n", msg)
}

// ============================================================
// Task 5: Implement a Buffered Channel
// ============================================================

func demoBuffered() {
	fmt.Println("========== TASK 5: Buffered Channel ==========")

	ch := make(chan int, 3) // buffer size 3: can hold 3 values without blocking

	// Send 3 values without a receiver (won't block because buffer has space)
	ch <- 10
	ch <- 20
	ch <- 30
	fmt.Printf("  Buffer length: %d, capacity: %d\n", len(ch), cap(ch))

	// Receive all values
	fmt.Println("  Received:", <-ch)
	fmt.Println("  Received:", <-ch)
	fmt.Println("  Received:", <-ch)
	fmt.Printf("  Buffer length after reads: %d\n\n", len(ch))
}

// ============================================================
// Task 6: Demonstrate Directional Channels
// ============================================================

// sendOnly can only send to the channel
func sendOnly(ch chan<- string, msg string) {
	ch <- msg
	fmt.Printf("  [Send-Only] Sent: %s\n", msg)
}

// receiveOnly can only read from the channel
func receiveOnly(ch <-chan string) string {
	msg := <-ch
	fmt.Printf("  [Receive-Only] Got: %s\n", msg)
	return msg
}

func demoDirectional() {
	fmt.Println("========== TASK 6: Directional Channels ==========")

	ch := make(chan string, 1)

	// Pass as send-only
	sendOnly(ch, "Directional message")

	// Pass as receive-only
	receiveOnly(ch)
	fmt.Println()
}

// ============================================================
// Task 7: Demonstrate Channel Closing
// ============================================================

func demoChannelClosing() {
	fmt.Println("========== TASK 7: Channel Closing ==========")

	ch := make(chan int, 5)

	// Producer sends values then closes
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * 10
			fmt.Printf("  [Producer] Sent: %d\n", i*10)
		}
		close(ch) // Signal no more values
		fmt.Println("  [Producer] Channel closed!")
	}()

	time.Sleep(200 * time.Millisecond)

	// Consumer checks if channel is closed using ok pattern
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("  [Consumer] Channel is closed, stopping.")
			break
		}
		fmt.Printf("  [Consumer] Received: %d\n", val)
	}
	fmt.Println()
}

// ============================================================
// Task 8: Use Range Over a Channel
// ============================================================

func demoRange() {
	fmt.Println("========== TASK 8: Range Over Channel ==========")

	ch := make(chan string, 5)

	// Send employee names
	go func() {
		names := []string{"Piyush", "Alice", "Bob", "Charlie", "Diana"}
		for _, name := range names {
			ch <- name
		}
		close(ch) // Must close so range loop ends
	}()

	// range automatically stops when channel is closed
	fmt.Println("  Employees received via range:")
	for name := range ch {
		fmt.Printf("    - %s\n", name)
	}
	fmt.Println()
}

func main() {
	demoUnbuffered()
	demoBuffered()
	demoDirectional()
	demoChannelClosing()
	demoRange()
}
