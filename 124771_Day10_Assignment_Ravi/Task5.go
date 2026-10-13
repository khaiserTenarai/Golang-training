package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for sync.Map.
)

func main() {
	var employees sync.Map // Creates a concurrent-safe map.

	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(3)             // Registers three goroutines.

	go func() { // Starts first writer.
		defer wg.Done()                // Signals completion.
		employees.Store(101, "Rajesh") // Stores employee 101.
	}()

	go func() { // Starts second writer.
		defer wg.Done()              // Signals completion.
		employees.Store(102, "Amit") // Stores employee 102.
	}()

	go func() { // Starts third writer.
		defer wg.Done()               // Signals completion.
		employees.Store(103, "Pooja") // Stores employee 103.
	}()

	wg.Wait() // Waits for all stores.

	employees.Range(func(key, value any) bool { // Iterates safely over the map.
		fmt.Println(key, value) // Prints each employee.
		return true             // Continues iteration.
	})
}
