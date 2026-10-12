package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for sync.Map.
)

func main() {
	var products sync.Map // Creates a concurrent-safe map.

	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(3)             // Registers three goroutines.

	go func() { // Starts first writer.
		defer wg.Done()               // Signals completion.
		products.Store(101, "Laptop") // Stores product.
	}()

	go func() { // Starts second writer.
		defer wg.Done()               // Signals completion.
		products.Store(102, "Mobile") // Stores product.
	}()

	go func() { // Starts third writer.
		defer wg.Done()               // Signals completion.
		products.Store(103, "Camera") // Stores product.
	}()

	wg.Wait() // Waits for all stores.

	products.Range(func(key, value any) bool { // Iterates safely over the map.
		fmt.Println(key, value) // Prints each product.
		return true             // Continues iteration.
	})
}
