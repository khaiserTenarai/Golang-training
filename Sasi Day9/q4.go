package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from a basic anonymous goroutine!")
	}()

	employees := []string{"Sasi", "Anand", "Siva"}

	for _, name := range employees {
		wg.Add(1)
		go func(empName string) {
			defer wg.Done()
			fmt.Println("Processing employee:", empName)
		}(name)
	}

	resultChan := make(chan string, 1)
	wg.Add(1)
	go func(ch chan<- string) {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		ch <- "Task completed by channel anonymous goroutine"
	}(resultChan)

	wg.Wait()
	close(resultChan)

	fmt.Println(<-resultChan)
	fmt.Println("All anonymous goroutines finished successfully.")
}