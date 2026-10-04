package main

import "fmt"

func main() {
	employees := map[int]string{
		1001: "Sasikumar",
		1002: "Vijay",
		1003: "Ajith",
	}

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	name, exists := employees[id]

	if exists {
		fmt.Println("Employee Name:", name)
	} else {
		fmt.Println("Employee not found")
	}
}