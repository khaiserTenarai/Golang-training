package main

import (
	"fmt"
	"sync"
	"time"
)

func employeeTask(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine started")

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutine completed")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("Creating goroutine")

	go employeeTask(&wg)

	fmt.Println("Main function is running")

	wg.Wait()

	fmt.Println("Main function completed")
}
