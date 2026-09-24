package main

import "fmt"

func main() {
	employees := map[int]string{
		101: "Pallavi",
		102: "Rahul",
		103: "Anita",
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