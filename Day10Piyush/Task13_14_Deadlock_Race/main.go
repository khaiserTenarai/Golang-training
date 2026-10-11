// Tasks 13, 14: Deadlock and Race Condition
// Task 13 — Create a deadlock example and fix it
// Task 14 — Find and fix a race condition using -race
// Run: go run main.go
// Run with race detector: go run -race main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// Task 13: Create a Deadlock Example and Fix It
// ============================================================

// DEADLOCK EXAMPLE (commented out — would freeze the program)
// Two goroutines each hold one lock and wait for the other → deadlock.
func deadlockExample() {
	fmt.Println("--- Deadlock Example (shows the pattern, won't actually deadlock) ---")
	fmt.Println(`
  // This code WOULD deadlock:
  var muA, muB sync.Mutex

  // Goroutine 1: locks A then waits for B
  go func() {
      muA.Lock()
      time.Sleep(100ms)
      muB.Lock()       // Blocked — Goroutine 2 holds muB
      muB.Unlock()
      muA.Unlock()
  }()

  // Goroutine 2: locks B then waits for A
  go func() {
      muB.Lock()
      time.Sleep(100ms)
      muA.Lock()       // Blocked — Goroutine 1 holds muA
      muA.Unlock()
      muB.Unlock()
  }()
  // RESULT: Both goroutines wait forever = DEADLOCK
	`)
}

// FIXED: Always acquire locks in the SAME order (A → B)
func deadlockFixed() {
	fmt.Println("--- Deadlock Fixed: Consistent Lock Ordering ---")

	var muA, muB sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	// Goroutine 1: locks A then B (same order)
	go func() {
		defer wg.Done()
		muA.Lock()
		time.Sleep(50 * time.Millisecond)
		muB.Lock()
		fmt.Println("  [Goroutine 1] Acquired A then B — working")
		muB.Unlock()
		muA.Unlock()
	}()

	// Goroutine 2: also locks A then B (same order — no deadlock!)
	go func() {
		defer wg.Done()
		muA.Lock()
		time.Sleep(50 * time.Millisecond)
		muB.Lock()
		fmt.Println("  [Goroutine 2] Acquired A then B — working")
		muB.Unlock()
		muA.Unlock()
	}()

	wg.Wait()
	fmt.Println("  No deadlock! Consistent ordering prevents it.\n")
}

// Another deadlock: unbuffered channel in same goroutine
func channelDeadlockExample() {
	fmt.Println("--- Channel Deadlock Example ---")
	fmt.Println(`
  // This WOULD deadlock:
  ch := make(chan int)    // unbuffered
  ch <- 42               // blocks forever — no receiver in another goroutine
  fmt.Println(<-ch)      // never reached

  // FIX 1: Use a buffered channel
  ch := make(chan int, 1)
  ch <- 42
  fmt.Println(<-ch)      // works!

  // FIX 2: Use separate goroutine
  ch := make(chan int)
  go func() { ch <- 42 }()
  fmt.Println(<-ch)      // works!
	`)

	// Demonstrate the fix
	ch := make(chan int, 1)
	ch <- 42
	fmt.Printf("  [Fixed] Buffered channel received: %d\n\n", <-ch)
}

func demoDeadlock() {
	fmt.Println("========== TASK 13: Deadlock Example and Fix ==========\n")
	deadlockExample()
	deadlockFixed()
	channelDeadlockExample()
}

// ============================================================
// Task 14: Find and Fix a Race Condition Using -race
// ============================================================

func demoRace() {
	fmt.Println("========== TASK 14: Race Condition and Fix ==========")
	fmt.Println("  Run with: go run -race main.go\n")

	// RACE CONDITION: multiple goroutines write to shared variable
	fmt.Println("--- Unsafe Version (race condition) ---")
	unsafeBalance := 1000
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// READ, MODIFY, WRITE — not atomic = RACE
			temp := unsafeBalance
			temp += 10
			unsafeBalance = temp
		}()
	}
	wg.Wait()
	fmt.Printf("  Unsafe balance (expected 2000): %d (likely wrong)\n\n", unsafeBalance)

	// FIXED: Protect with Mutex
	fmt.Println("--- Fixed Version (with Mutex) ---")
	safeBalance := 1000
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			safeBalance += 10 // Protected by mutex
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("  Safe balance (expected 2000): %d (always correct)\n\n", safeBalance)
}

func main() {
	demoDeadlock()
	demoRace()
}
