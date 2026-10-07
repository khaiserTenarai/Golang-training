package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

func (s *Store) Read(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[key]
}

func (s *Store) Write(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func main() {
	store := Store{
		data: make(map[string]string),
	}

	store.Write("name", "Go")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Printf("Reader %d: %s\n", id, store.Read("name"))
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
}
