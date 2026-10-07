package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sm.Store(n, fmt.Sprintf("value-%d", n))
		}(i)
	}

	wg.Wait()

	val, ok := sm.Load(2)
	if ok {
		fmt.Printf("Load: %v\n", val)
	}

	actual, loaded := sm.LoadOrStore(2, "new-value-2")
	fmt.Printf("LoadOrStore (existing): actual=%v, loaded=%v\n", actual, loaded)

	actual, loaded = sm.LoadOrStore(5, "value-5")
	fmt.Printf("LoadOrStore (new): actual=%v, loaded=%v\n", actual, loaded)

	sm.Delete(0)

	fmt.Println("Map contents:")
	sm.Range(func(key, value any) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}