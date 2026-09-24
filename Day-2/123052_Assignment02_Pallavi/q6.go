package main

import "fmt"

func main() {
	age := 25
	salary := 50000

	fmt.Println("Age > 18:", age > 18)
	fmt.Println("Salary > 40000:", salary > 40000)

	if age >= 18 && salary >= 40000 {
		fmt.Println("Employee is eligible")
	}

	if age < 18 || salary < 30000 {
		fmt.Println("Employee is not eligible")
	}
}