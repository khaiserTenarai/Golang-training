package main

import "fmt"

func main() {
	name := "Sanskriti"
	age := 22

	fmt.Println("Employee Name:", name)
	fmt.Println("Employee Age:", age)

	if age >= 18 {
		fmt.Println("Employee is eligible")
	} else {
		fmt.Println("Employee is not eligible")
	}
}