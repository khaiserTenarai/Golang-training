package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter.value)
}