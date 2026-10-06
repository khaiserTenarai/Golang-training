package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		DEADLOCK:

		Goroutine locks the mutex and then waits for
		the same mutex again.

		The mutex can never be unlocked.
	*/

	var mutex sync.Mutex

	mutex.Lock()

	fmt.Println("First lock")

	// ❌ Trying to lock the same mutex again.
	mutex.Lock()

	fmt.Println("This line will never execute")
}