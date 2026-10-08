package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()

	c.counter++

	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {
	var wg sync.WaitGroup
	safeCount := SafeCounter{}

	numGoroutines := 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go safeCount.Increment(&wg)
	}
	wg.Wait()

	fmt.Printf("Final Counter Value: %d\n", safeCount.Value())
}