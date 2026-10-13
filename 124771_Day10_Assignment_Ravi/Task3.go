package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for RWMutex.
	"time" // Imports time for simulated work.
)

func main() {
	var mu sync.RWMutex     // Creates a read/write mutex.
	data := "Employee Data" // Creates shared data.

	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(5)             // Registers five readers.

	for i := 1; i <= 5; i++ { // Starts five readers.
		go func(id int) { // Creates a reader goroutine.
			defer wg.Done()                          // Signals reader completion.
			mu.RLock()                               // Allows multiple readers at the same time.
			fmt.Println("Reader", id, "read:", data) // Reads shared data.
			time.Sleep(100 * time.Millisecond)       // Simulates reading.
			mu.RUnlock()                             // Releases the read lock.
		}(i) // Passes the reader ID.
	}

	wg.Add(1)   // Registers one writer.
	go func() { // Creates the writer goroutine.
		defer wg.Done()                   // Signals writer completion.
		time.Sleep(50 * time.Millisecond) // Lets readers start first.
		mu.Lock()                         // Gets exclusive write access.
		data = "Updated Employee Data"    // Changes shared data.
		fmt.Println("Writer updated data:   ", data)
		mu.Unlock() // Releases the write lock.
	}()

	wg.Wait()               // Waits for all readers and writer.
	fmt.Println("Finished") // Prints when all operations complete.
}
