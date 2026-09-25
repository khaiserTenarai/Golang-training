package main

import "fmt"

const company = "Autoapp"

func main() {
	name, age, isManager := "Sachin", 23, false

	fmt.Println("Company:", company)
	fmt.Println("Employee:", name)
	fmt.Println("Age:", age)
	fmt.Println("Manager:", isManager)
}