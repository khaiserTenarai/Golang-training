package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func (m *SafeMap) Read(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data[key]
}

func (m *SafeMap) Write(key, val string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

func main() {
	m := SafeMap{data: make(map[string]string)}
	var wg sync.WaitGroup

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Write("status", "active")
	}()

	// Multiple Concurrent Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Reader %d read: %s\n", id, m.Read("status"))
		}(i)
	}

	wg.Wait()
}