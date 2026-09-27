/*
Day 4, Q6. Create a closure for generating employee IDs.
*/
package main

import "fmt"

// idGenerator returns a function that remembers its own counter between
// calls - that's the closure part, the returned func "closes over" nextID.
func idGenerator(start int) func() int {
	nextID := start
	return func() int {
		id := nextID
		nextID++
		return id
	}
}

func main() {
	generateID := idGenerator(1001)

	fmt.Println("Gokul's ID:", generateID())
	fmt.Println("Meera's ID:", generateID())
	fmt.Println("Ravi's ID :", generateID())

	// a second, independent generator starting from a different number
	internGen := idGenerator(5000)
	fmt.Println("Intern ID:", internGen())
}
