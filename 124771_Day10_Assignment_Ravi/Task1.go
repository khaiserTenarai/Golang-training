package main

import (
	"fmt"
	"sync" // Imports sync for sync.Once.
)

var once sync.Once // Creates a Once object that allows one-time execution.

func initialize() { // Defines the initialization function.
	fmt.Println("Application initialized once") // Prints initialization message.
}

func main() {
	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(5)

	for i := 1; i <= 5; i++ { // Starts five goroutines.
		go func(id int) { // Creates a goroutine.
			defer wg.Done()               // Signals completion.
			once.Do(initialize)           // Runs initialize only once across all goroutines.
			fmt.Println("Goroutine:", id) // Prints the goroutine ID.
		}(i) // Passes the ID.
	}

	wg.Wait() // Waits for all goroutines.
}
