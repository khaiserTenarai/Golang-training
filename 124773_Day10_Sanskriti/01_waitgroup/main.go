package main

import (
	"fmt"
	"sync"
)

func employeeTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Tell WaitGroup that this goroutine is finished.
	fmt.Println("Employee", id, "is working")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // We are starting 3 goroutines.

	go employeeTask(1, &wg)
	go employeeTask(2, &wg)
	go employeeTask(3, &wg)

	wg.Wait() // Wait until all goroutines finish.

	fmt.Println("All employees completed their work")
}
