package main

import (
	"fmt"
	"sync"
)

func main() {
	// Task 8: Several goroutines safely increase one shared counter.
	var counter int
	var mutex sync.Mutex
	var group sync.WaitGroup

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("Try: go run -race ./assignment_task8")
}
