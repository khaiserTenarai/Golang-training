package main

import "fmt"
// 2. Major Go Primitive Data Types

func main() {
	var age int = 30
	var userID uint64 = 9876543210

	var salary float64 = 85400.50
	var rating float32 = 4.8

	var isManager bool = true

	var employeeName string = "Alice Johnson"



	fmt.Printf("Name: %s, Age: %d, UserID: %d\n", employeeName, age, userID)
	fmt.Printf("Salary: $%.2f, Rating: %.1f, Manager: %t\n", salary, rating, isManager)

}

