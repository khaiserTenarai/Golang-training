package main
// Use RWMutex when reading happens much more frequently than writing.
import (
	"fmt"
	"sync"
)

func main() {
	/*
		sync.RWMutex is useful when we have:

		Many readers
		Few writers

		RLock()  -> allows multiple readers
		RUnlock() -> releases read lock

		Lock()   -> allows one writer
		Unlock() -> releases write lock
	*/

	var mutex sync.RWMutex
	var name = "Ganesh"

	var wg sync.WaitGroup

	// Reader
	wg.Add(1)
	go func() {
		defer wg.Done()

		mutex.RLock()
		fmt.Println("Reading name:", name)
		mutex.RUnlock()
	}()

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()

		mutex.Lock()
		name = "Ganesh Reddy"
		mutex.Unlock()

		fmt.Println("Name updated")
	}()

	wg.Wait()

	fmt.Println("Final name:", name)
}