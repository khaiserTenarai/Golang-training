package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Read(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Write(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func main() {
	cache := Cache{data: make(map[string]string)}
	var wg sync.WaitGroup

	cache.Write("status", "active")

	// Multiple concurrent readers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val := cache.Read("status")
			fmt.Printf("Reader %d saw: %s\n", id, val)
		}(i)
	}

	// Concurrent writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		cache.Write("status", "updated")
		fmt.Println("Writer updated cache")
	}()

	wg.Wait()
}