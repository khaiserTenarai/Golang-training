// 11. CLI Employee Search
//
// Maintain employee data in memory.
// Accept employee ID from command line.
// Display employee information.

package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Dept   string
	Salary float64
}

func main() {
	employees := []Employee{
		{ID: 101, Name: "Anil", Dept: "HR", Salary: 45000},
		{ID: 102, Name: "Ravi", Dept: "Engineering", Salary: 60000},
		{ID: 103, Name: "Priya", Dept: "Finance", Salary: 52000},
	}

	searchID := 102

	for _, emp := range employees {
		if emp.ID == searchID {
			fmt.Println("Found Employee:")
			fmt.Println("Name:", emp.Name)
			fmt.Println("Dept:", emp.Dept)
			fmt.Println("Salary:", emp.Salary)

			return
		}
	}

	fmt.Println("No employee found with ID", searchID)
}
