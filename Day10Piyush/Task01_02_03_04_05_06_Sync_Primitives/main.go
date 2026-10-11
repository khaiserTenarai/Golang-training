// Tasks 1–6: Sync Primitives
// Task 1 — sync.WaitGroup
// Task 2 — sync.Mutex
// Task 3 — sync.RWMutex
// Task 4 — sync.Once
// Task 5 — sync.Map
// Task 6 — Atomic counter
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================
// Task 1: Use sync.WaitGroup
// ============================================================

func demoWaitGroup() {
	fmt.Println("========== TASK 1: sync.WaitGroup ==========")

	var wg sync.WaitGroup

	tasks := []string{"Fetch Data", "Process Report", "Send Email", "Generate PDF", "Clean Cache"}

	for _, task := range tasks {
		wg.Add(1)
		go func(t string) {
			defer wg.Done() // Signals completion when goroutine finishes
			fmt.Printf("  [Start] %s\n", t)
			time.Sleep(time.Duration(100+len(t)*10) * time.Millisecond)
			fmt.Printf("  [Done]  %s\n", t)
		}(task)
	}

	wg.Wait() // Blocks until all goroutines call Done()
	fmt.Println("  All tasks completed!")
}

// ============================================================
// Task 2: Protect Shared Data with sync.Mutex
// ============================================================

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func demoMutex() {
	fmt.Println("========== TASK 2: sync.Mutex ==========")

	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	fmt.Printf("  Final counter (expected 1000): %d\n\n", counter.Value())
}

// ============================================================
// Task 3: Implement sync.RWMutex
// ============================================================

type EmployeeCache struct {
	mu   sync.RWMutex
	data map[int]string
}

func NewEmployeeCache() *EmployeeCache {
	return &EmployeeCache{data: make(map[int]string)}
}

func (c *EmployeeCache) Write(id int, name string) {
	c.mu.Lock() // Exclusive lock for writing
	defer c.mu.Unlock()
	c.data[id] = name
	fmt.Printf("  [Write] ID=%d, Name=%s\n", id, name)
}

func (c *EmployeeCache) Read(id int) string {
	c.mu.RLock() // Shared lock — multiple readers allowed
	defer c.mu.RUnlock()
	return c.data[id]
}

func demoRWMutex() {
	fmt.Println("========== TASK 3: sync.RWMutex ==========")

	cache := NewEmployeeCache()
	var wg sync.WaitGroup

	// Writers
	names := map[int]string{1: "Piyush", 2: "Alice", 3: "Bob"}
	for id, name := range names {
		wg.Add(1)
		go func(i int, n string) {
			defer wg.Done()
			cache.Write(i, n)
		}(id, name)
	}
	wg.Wait()

	// Concurrent readers (many readers at once — no blocking between them)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(readID int) {
			defer wg.Done()
			id := (readID % 3) + 1
			name := cache.Read(id)
			fmt.Printf("  [Read]  ID=%d → %s\n", id, name)
		}(i)
	}
	wg.Wait()
	fmt.Println()
}

// ============================================================
// Task 4: Demonstrate sync.Once
// ============================================================

type DBConnection struct {
	Host string
}

var (
	dbInstance *DBConnection
	dbOnce     sync.Once
)

func GetDB() *DBConnection {
	dbOnce.Do(func() {
		fmt.Println("  [Once] Initializing database connection...")
		time.Sleep(100 * time.Millisecond)
		dbInstance = &DBConnection{Host: "localhost:5432"}
		fmt.Println("  [Once] Connection established!")
	})
	return dbInstance
}

func demoOnce() {
	fmt.Println("========== TASK 4: sync.Once ==========")

	var wg sync.WaitGroup
	// 5 goroutines all try to get DB — only the first one initializes
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			db := GetDB()
			fmt.Printf("  [Goroutine %d] Got DB: %s\n", id, db.Host)
		}(i)
	}
	wg.Wait()
	fmt.Println()
}

// ============================================================
// Task 5: Demonstrate sync.Map
// ============================================================

func demoSyncMap() {
	fmt.Println("========== TASK 5: sync.Map ==========")

	var empMap sync.Map
	var wg sync.WaitGroup

	// Concurrent writes — safe without external locking
	employees := map[int]string{1: "Piyush", 2: "Alice", 3: "Bob", 4: "Charlie", 5: "Diana"}
	for id, name := range employees {
		wg.Add(1)
		go func(i int, n string) {
			defer wg.Done()
			empMap.Store(i, n)
			fmt.Printf("  [Store] ID=%d, Name=%s\n", i, n)
		}(id, name)
	}
	wg.Wait()

	// Concurrent reads
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val, ok := empMap.Load(id)
			if ok {
				fmt.Printf("  [Load]  ID=%d → %s\n", id, val)
			}
		}(i)
	}
	wg.Wait()

	// Range over all entries
	fmt.Println("  [Range] All entries:")
	empMap.Range(func(key, value interface{}) bool {
		fmt.Printf("    %d: %s\n", key, value)
		return true
	})
	fmt.Println()
}

// ============================================================
// Task 6: Implement an Atomic Counter
// ============================================================

func demoAtomic() {
	fmt.Println("========== TASK 6: Atomic Counter ==========")

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // Lock-free atomic increment
		}()
	}
	wg.Wait()

	fmt.Printf("  Atomic counter (expected 1000): %d\n", atomic.LoadInt64(&counter))

	// Demonstrate CompareAndSwap
	old := atomic.LoadInt64(&counter)
	swapped := atomic.CompareAndSwapInt64(&counter, old, 0)
	fmt.Printf("  CAS reset to 0: swapped=%v, new value=%d\n\n", swapped, atomic.LoadInt64(&counter))
}

func main() {
	demoWaitGroup()
	demoMutex()
	demoRWMutex()
	demoOnce()
	demoSyncMap()
	demoAtomic()
}
