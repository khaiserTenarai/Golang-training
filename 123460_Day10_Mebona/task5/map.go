package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	// Store key-value pairs
	sm.Store("golang", "1.22")
	sm.Store("python", "3.12")

	// Load
	if val, ok := sm.Load("golang"); ok {
		fmt.Println("Loaded golang:", val)
	}

	// LoadOrStore
	actual, loaded := sm.LoadOrStore("rust", "1.75")
	fmt.Printf("Loaded: %v, Value: %v\n", loaded, actual)

	// Iterate over map elements
	sm.Range(func(key, value any) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}