package main

import "fmt"

func employeeIDGenerator() func() int {
	id := 1000

	return func() int {
		id++
		return id
	}
}

func main() {
	generateID := employeeIDGenerator()

	fmt.Println("Employee ID:", generateID())
	fmt.Println("Employee ID:", generateID())
	fmt.Println("Employee ID:", generateID())
	fmt.Println("Employee ID:", generateID())
}
