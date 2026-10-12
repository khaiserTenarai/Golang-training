package main

import (
	"fmt"  // Imports fmt.
	"sync" // Imports sync.
)

func main() {
	counter := 0 // Shared variable.

	var mu sync.Mutex // Creates a mutex.

	var wg sync.WaitGroup // Creates WaitGroup.
	wg.Add(2)             // Registers two goroutines.

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			mu.Lock()   // Locks before changing counter.
			counter++   // Safely updates counter.
			mu.Unlock() // Unlocks after updating.
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			mu.Lock()   // Locks before changing counter.
			counter++   // Safely updates counter.
			mu.Unlock() // Unlocks after updating.
		}
	}()

	wg.Wait() // Waits for both goroutines.

	fmt.Println("Counter:", counter)
}
