package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func (c *Cache) Get(k string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[k]
	return v, ok
}
func (c *Cache) Set(k string, v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func main() {
	c := &Cache{data: map[string]int{}}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			c.Set(fmt.Sprint("k", i), i)

		}(i)
		go func(i int) {
			defer wg.Done()
			c.Get(fmt.Sprint("k", i))
		}(i)
	}
	wg.Wait()
	fmt.Println(len(c.data))
}
