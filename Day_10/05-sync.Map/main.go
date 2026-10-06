package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		sync.Map is a concurrent-safe map.

		Normal map is NOT safe when multiple goroutines
		are reading and writing at the same time.

		sync.Map provides:

		Store() -> add/update data
		Load()  -> read data
		Delete() -> remove data
	*/

	var users sync.Map

	var wg sync.WaitGroup

	// Store data
	wg.Add(1)
	go func() {
		defer wg.Done()

		users.Store(1, "Ganesh")
		users.Store(2, "Reddy")
	}()

	wg.Wait()

	// Read data
	value, found := users.Load(1)

	if found {
		fmt.Println("User:", value)
	} else {
		fmt.Println("User not found")
	}

	// Delete data
	users.Delete(2)

	fmt.Println("User 2 deleted")
}