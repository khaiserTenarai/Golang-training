package main

import "fmt"

func main() {
	// Variable declaration using var
	var name string = "Ravi"
	var age int = 25

	// Variable declaration with type inference
	var city = "Mumbai"

	country := "India"

	// Boolean variable
	isTrue := true
	// Constant declaration
	const Pi = 3.14159

	// Display constants
	fmt.Println("Pi:", Pi)

	// Display variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Is Student:", isTrue)

	// Changing a variable
	age = 26
	fmt.Println("Updated Age:", age)

}
