package main

import "fmt"

func employeeID() func() int {
	id := 999

	return func() int {
		id++
		return id
	}
}

func main() {
	nextID := employeeID()

	fmt.Println(nextID())
	fmt.Println(nextID())
}
