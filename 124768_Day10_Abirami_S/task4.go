package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialization is done")
}
func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Worker", id, " started")
	once.Do(initialize)
	fmt.Println("Worker", id, " completed")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go worker(&wg, i)
	}
	wg.Wait()
	fmt.Println("All workers completed")
}
