package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	// Store values concurrently
	keys := []string{"carrot", "tomato", "cherry"}
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			m.Store(k, len(k))
		}(key)
	}
	wg.Wait()

	// Load or Store
	actual, loaded := m.LoadOrStore("carrot", 999)
	fmt.Printf("LoadOrStore 'carrot': val=%v, loaded=%v\n", actual, loaded)

	// Range iteration
	fmt.Println("Map contents:")
	m.Range(func(key, value any) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})
}