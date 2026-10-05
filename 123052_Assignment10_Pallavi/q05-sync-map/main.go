package main

import (
	"fmt"
	"sync"
)

func main() {

	var employees sync.Map

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		employees.Store(101, "Pallavi")
	}()

	go func() {
		defer wg.Done()

		employees.Store(102, "Rahul")
	}()

	go func() {
		defer wg.Done()

		employees.Store(103, "Anita")
	}()

	wg.Wait()

	// Read all values
	employees.Range(func(key, value interface{}) bool {
		fmt.Println("Employee ID:", key, "Name:", value)
		return true
	})
}