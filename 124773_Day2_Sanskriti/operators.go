package main

import "fmt"

func main() {
	age := 22
	salary := 50000

	// Comparison operators
	fmt.Println("Age == 22:", age == 22)
	fmt.Println("Age != 25:", age != 25)
	fmt.Println("Age > 18:", age > 18)
	fmt.Println("Age < 18:", age < 18)
	fmt.Println("Salary >= 50000:", salary >= 50000)
	fmt.Println("Salary <= 30000:", salary <= 30000)

	// Logical operators
	isEmployee := true
	hasID := true
	isManager := false

	fmt.Println("Employee AND ID:", isEmployee && hasID)
	fmt.Println("Employee OR Manager:", isEmployee || isManager)
	fmt.Println("NOT Manager:", !isManager)
}