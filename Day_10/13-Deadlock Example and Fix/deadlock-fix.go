package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		FIX:

		Unlock the mutex before trying to lock it again.
	*/

	var mutex sync.Mutex

	mutex.Lock()

	fmt.Println("First lock")

	mutex.Unlock()

	// Now locking is safe.
	mutex.Lock()

	fmt.Println("Second lock")

	mutex.Unlock()

	fmt.Println("Program completed")
}

// A deadlock happens when goroutines wait for each other forever and none of them can continue.