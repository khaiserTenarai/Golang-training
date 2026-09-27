package main

import "fmt"

func main() {

	var age int = 22
	var salary float64 = 45000.5
	var isEmployee bool = true
	var name string = "Shashank"
	var grade rune = 'A'
	var employeeID uint = 101
	var complexNumber complex128 = 10 + 5i

	fmt.Println("Integer:", age)
	fmt.Println("Float:", salary)
	fmt.Println("Boolean:", isEmployee)
	fmt.Println("String:", name)
	fmt.Println("Rune:", grade)
	fmt.Println("Unsigned Integer:", employeeID)
	fmt.Println("Complex Number:", complexNumber)
}
