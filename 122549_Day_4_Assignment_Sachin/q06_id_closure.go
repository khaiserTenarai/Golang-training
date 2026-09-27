// 6. Create a closure for generating employee IDs.

package main

import "fmt"

func idGenerator(startID int) func() int {
	current := startID
	return func() int {
		id := current
		current++
		return id
	}
}

func main() {
	nextID := idGenerator(1001)

	fmt.Println("ID:", nextID())
	fmt.Println("ID:", nextID())
	fmt.Println("ID:", nextID())
}
