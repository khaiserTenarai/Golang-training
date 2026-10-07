package main

import (
	"fmt"
	"sync"
)

func main() {
	var employees sync.Map

	// Store data.
	employees.Store(1, "Sanskriti")
	employees.Store(2, "Rahul")
	employees.Store(3, "Priya")

	// Load data.
	value, found := employees.Load(1)

	if found {
		fmt.Println("Employee:", value)
	}

	// Display all employees.
	employees.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
