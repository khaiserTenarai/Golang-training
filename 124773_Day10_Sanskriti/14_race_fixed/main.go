package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			// FIX: protect shared data with a mutex.
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}

// Test this program with:
// go run -race main.go
//
// The mutex prevents a data race.
