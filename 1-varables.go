package main

import "fmt"

func main() {

	var name string = "Shashank"
	var age int = 22

	city := "Bangalore"
	const college = "VTU"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)

	fmt.Println("College:", college)

	age = 23
	fmt.Println("Updated Age:", age)

	// A constant cannot be changed
	// college = "Another College" // This would cause an error
}
