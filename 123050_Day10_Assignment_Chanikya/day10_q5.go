package main

import (
	"fmt"
	"sync"
)

func main() {
	var users sync.Map

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			users.Store(id, fmt.Sprintf("User-%d", id))
		}(i)
	}

	wg.Wait()

	users.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
