package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for RWMutex and WaitGroup.
)

func main() {
	var mu sync.RWMutex // Creates a read/write mutex.

	name := "John" // Creates shared data.

	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(3)             // Registers three goroutines.

	go func() { // Creates the first reader.
		defer wg.Done() // Marks the reader as finished.

		mu.RLock()                     // Gets read lock.
		fmt.Println("Reader 1:", name) // Reads the data.
		mu.RUnlock()                   // Releases read lock.
	}()

	go func() { // Creates the second reader.
		defer wg.Done() // Marks the reader as finished.

		mu.RLock()                     // Gets read lock.
		fmt.Println("Reader 2:", name) // Reads the data.
		mu.RUnlock()                   // Releases read lock.
	}()

	go func() { // Creates the writer.
		defer wg.Done() // Marks the writer as finished.

		mu.Lock()                   // Gets write lock.
		name = "David"              // Changes the data.
		fmt.Println("Name updated") // Prints update message.
		mu.Unlock()                 // Releases write lock.
	}()

	wg.Wait() // Waits for all goroutines.

	fmt.Println("Finished") // Prints when all work is complete.
}
