// package main

// import (
// 	"fmt"  // Imports fmt.
// 	"sync" // Imports sync.
// )

// func main() {
// 	var mu sync.Mutex // Creates a mutex.

// 	mu.Lock() // Locks the mutex.

// 	fmt.Println("Mutex locked")

// 	mu.Lock() // Tries to lock the same mutex again.
// 	// Deadlock happens because the mutex is already locked.
// }

package main

import (
	"fmt"  // Imports fmt.
	"sync" // Imports sync.
)

func main() {
	var mu sync.Mutex // Creates a mutex.

	mu.Lock() // Locks the mutex.

	fmt.Println("Mutex locked")

	mu.Unlock() // Unlocks the mutex before locking again.

	mu.Lock() // Locks the mutex again safely.

	fmt.Println("Mutex locked again")

	mu.Unlock() // Releases the mutex.
}
