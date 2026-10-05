package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu sync.RWMutex

	data := "Employee Data"

	var wg sync.WaitGroup

	// Create 5 readers
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func(id int) {
			defer wg.Done()

			// Multiple readers can read at the same time
			mu.RLock()

			fmt.Println("Reader", id, "read:", data)

			time.Sleep(100 * time.Millisecond)

			mu.RUnlock()

		}(i)
	}

	// Create one writer
	wg.Add(1)

	go func() {
		defer wg.Done()

		time.Sleep(50 * time.Millisecond)

		// Writer needs exclusive access
		mu.Lock()

		data = "Updated Employee Data"

		fmt.Println("Writer updated data")

		mu.Unlock()
	}()

	wg.Wait()

	fmt.Println("Finished")
}