package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) { defer wg.Done(); m.Store(i, i*i) }(i)
	}
	wg.Wait()

	v, ok := m.Load(3)
	fmt.Println(v, ok)

	actual, loaded := m.LoadOrStore(10, "new")
	fmt.Println(actual, loaded)

	m.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
	m.Delete(3)
}
