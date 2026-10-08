package main

import (
	"fmt"
	"sync"
	"time"
)

func employeeTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Started")
	fmt.Println("Running")
	time.Sleep(2 * time.Second)
	fmt.Println("Completed")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Creating goroutine")
	go employeeTask(&wg)
	wg.Wait()
	fmt.Println("Main function completed")
}
