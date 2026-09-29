// 6. Create a closure for generating employee IDs.

package main

import "fmt"

func newIDGenerator() func() int {
	lastID := 0
	return func() int {
		lastID++
		return lastID
	}
}

func main() {
	nextID := newIDGenerator()

	fmt.Println("Employee 1 ID:", nextID())
	fmt.Println("Employee 2 ID:", nextID())
	fmt.Println("Employee 3 ID:", nextID())

	anotherGenerator := newIDGenerator()
	fmt.Println("New batch, Employee 1 ID:", anotherGenerator())
}
