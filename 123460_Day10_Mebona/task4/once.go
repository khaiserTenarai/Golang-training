package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance string
)

func initialize() {
	fmt.Println("Initializing expensive resource...")
	instance = "Resource Loaded"
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(initialize)
			fmt.Printf("Goroutine %d got: %s\n", id, instance)
		}(i)
	}

	wg.Wait()
}