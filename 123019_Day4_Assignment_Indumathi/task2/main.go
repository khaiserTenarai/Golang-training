package main

import "fmt"

func employeeDetails() (string, int) {
	name := "Indu"
	age := 22

	return name, age
}

func main() {
	name, age := employeeDetails()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}