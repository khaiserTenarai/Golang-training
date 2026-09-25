package main

import "fmt"

func main() {
	var name string = "ganesh"
	var age int = 21
	var salary float64 = 50000.50

	const company = "Tenarai"
	const country = "India"
	// company = "Binqle" error

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Company:", company)
	fmt.Println("Country:", country)
}
