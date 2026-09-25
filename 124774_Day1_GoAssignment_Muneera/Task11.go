package main

import (
	"fmt"
	"os"
)

func main() {

	// Employee data
	id := [5]string{"201", "202", "203", "204", "205"}

	name := [5]string{"Rahul", "Priya", "Arjun", "Sneha", "Kiran"}

	salary := [5]int{35000, 45000, 55000, 40000, 65000}

	// Get employee ID from command line
	if len(os.Args) < 2 {
		fmt.Println("Please enter an Employee ID")
		return
	}

	searchID := os.Args[1]

	// Search employee
	for i := 0; i < len(id); i++ {

		if searchID == id[i] {

			fmt.Println("Employee ID:", id[i])
			fmt.Println("Employee Name:", name[i])
			fmt.Println("Employee Salary:", salary[i])

			return
		}
	}

	fmt.Println("Employee Not Found")
}
