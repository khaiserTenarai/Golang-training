package main

import "fmt"

func generateEmployeeID() func() int {

	id := 100

	return func() int {
		id++
		return id
	}
}

func main() {

	generateID := generateEmployeeID()

	fmt.Println("Employee ID:", generateID())
	fmt.Println("Employee ID:", generateID())
	fmt.Println("Employee ID:", generateID())
	fmt.Println("Employee ID:", generateID())
}