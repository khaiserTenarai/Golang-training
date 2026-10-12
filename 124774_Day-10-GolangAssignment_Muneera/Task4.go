package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for sync.Once.
)

var once sync.Once // Creates a Once object that allows one-time execution.

func addition() { // Defines the addition function.
	fmt.Println("Addition done successfully") // Prints initialization message.
}

func main() {
	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(5)             // Registers five goroutines.

	for i := 1; i <= 5; i++ { // Starts five goroutines.
		go func(id int) { // Creates a goroutine.
			defer wg.Done()               // Signals completion.
			once.Do(addition)             // Runs addition only once across all goroutines.
			fmt.Println("Goroutine:", id) // Prints the goroutine ID.
		}(i) // Passes the ID.
	}

	wg.Wait() // Waits for all goroutines.
}
