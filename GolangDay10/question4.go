package main

import (
	"fmt"
	"sync"
)

var (
	once   sync.Once
	config map[string]string
)

func loadConfig() {
	fmt.Println("Loading configuration...")
	config = map[string]string{
		"environment": "production",
		"version":     "1.0.0",
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	once.Do(loadConfig)

	fmt.Printf("Worker %d using config: %s\n", id, config["environment"])
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}