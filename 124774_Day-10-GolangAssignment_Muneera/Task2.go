package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for Mutex.
)

func main() {
	var mu sync.Mutex // Creates a mutex protecting the shared counter.
	counter := 0      // Creates shared data.

	var wg sync.WaitGroup // Creates a WaitGroup to wait for workers.
	wg.Add(100)           // Registers 100 goroutines.

	for i := 0; i < 100; i++ { // Starts 100 goroutines.
		go func() { // Creates one goroutine.
			defer wg.Done() // Signals completion.
			mu.Lock()       // Locks the shared resource.
			counter++       // Safely changes the shared counter.
			mu.Unlock()     // Releases the lock.
		}()
	}

	wg.Wait()                              // Waits for all increments to finish.
	fmt.Println("Final counter:", counter) // Prints the correct counter value.
}
