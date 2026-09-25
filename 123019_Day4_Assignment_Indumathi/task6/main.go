package main

import "fmt"

func generateID() func() int {
	id := 100

	return func() int {
		id++
		return id
	}
}

func main() {
	getID := generateID()

	fmt.Println(getID())
	fmt.Println(getID())
	fmt.Println(getID())
}