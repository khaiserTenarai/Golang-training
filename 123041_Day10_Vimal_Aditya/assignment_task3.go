package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *SharedCache) Read(key string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.RLock()
	val := c.data[key]
	fmt.Printf("Reader %d -> Key: %s | Value: %s\n", id, key, val)
	time.Sleep(50 * time.Millisecond)
	c.mu.RUnlock()                
}

func (c *SharedCache) Write(key, val string, wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()
	c.data[key] = val
	fmt.Printf("[WRITER] Updated key '%s' to '%s'\n", key, val)
	time.Sleep(100 * time.Millisecond)
	c.mu.Unlock()            
}

func main() {
	var wg sync.WaitGroup
	cache := SharedCache{
		data: map[string]string{
			"config": "default_v1",
		},
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go cache.Read("config", i, &wg)
	}

	wg.Add(1)
	go cache.Write("config", "updated_v2", &wg)

	for i := 5; i <= 7; i++ {
		wg.Add(1)
		go cache.Read("config", i, &wg)
	}

	wg.Wait()
}