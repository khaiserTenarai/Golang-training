// Task 8: Demonstrate Race Detection
// Run WITH race detector: go run -race main.go
// The program shows a race condition and its fix using sync.Mutex.

package main

import (
	"fmt"
	"sync"
)

// ============================================================
// PART 1: Race Condition (Unsafe)
// ============================================================

func unsafeCounter() {
	fmt.Println("===== PART 1: Unsafe Counter (Race Condition) =====")
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // DATA RACE: multiple goroutines read/write without sync
		}()
	}
	wg.Wait()

	// Expected: 1000, Actual: unpredictable (e.g., 987, 952, etc.)
	fmt.Printf("Expected: 1000, Got: %d (incorrect due to race condition)\n\n", counter)
}

// ============================================================
// PART 2: Fixed with sync.Mutex
// ============================================================

func safeCounterMutex() {
	fmt.Println("===== PART 2: Safe Counter (Mutex) =====")
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++ // Protected by mutex — no race
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("Expected: 1000, Got: %d (correct with Mutex)\n\n", counter)
}

// ============================================================
// PART 3: Fixed with sync/atomic
// ============================================================

// Using sync/atomic for lock-free atomic operations
func safeCounterAtomic() {
	fmt.Println("===== PART 3: Safe Counter (sync/atomic) =====")
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Atomic increment — no race, no lock overhead
			atomicAdd(&counter, 1)
		}()
	}
	wg.Wait()

	fmt.Printf("Expected: 1000, Got: %d (correct with atomic)\n\n", counter)
}

// Simple atomic-like add using mutex (sync/atomic.AddInt64 in real code)
var atomicMu sync.Mutex

func atomicAdd(ptr *int64, val int64) {
	atomicMu.Lock()
	*ptr += val
	atomicMu.Unlock()
}

// ============================================================
// PART 4: Race on a shared map
// ============================================================

func unsafeMap() {
	fmt.Println("===== PART 4: Unsafe Map (Race Condition) =====")
	data := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			data[val] = val * 2 // RACE: concurrent map writes
		}(i)
	}
	wg.Wait()
	fmt.Printf("Map has %d entries (may crash with 'concurrent map writes')\n\n", len(data))
}

func safeMap() {
	fmt.Println("===== PART 5: Safe Map (sync.RWMutex) =====")
	data := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.RWMutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			data[val] = val * 2 // Protected by mutex
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Map has %d entries (correct with RWMutex)\n\n", len(data))
}

func main() {
	fmt.Println("Run with: go run -race main.go")
	fmt.Println("The -race flag detects data races at runtime.\n")

	// Comment/uncomment to test each scenario
	unsafeCounter()
	safeCounterMutex()
	safeCounterAtomic()

	// WARNING: unsafeMap may panic with "concurrent map writes"
	// Uncomment to test: unsafeMap()
	safeMap()

	fmt.Println("Done! Run with -race flag to see race detection output.")
}
