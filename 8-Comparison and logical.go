package main

import "fmt"

func main() {
	age := 25
	salary := 30000

	// Comparison operators
	fmt.Println("Age is 25:", age == 25)
	fmt.Println("Age is greater than 18:", age > 18)
	fmt.Println("Salary is less than 50000:", salary < 50000)

	// Logical operators
	fmt.Println("Age > 18 AND Salary > 20000:", age > 18 && salary > 20000)
	fmt.Println("Age < 18 OR Salary > 20000:", age < 18 || salary > 20000)
	fmt.Println("Age is NOT 25:", !(age == 25))
}
