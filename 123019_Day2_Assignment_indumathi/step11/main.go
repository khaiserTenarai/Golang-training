package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	employees := map[int]Employee{
		101: {
			ID:         101,
			Name:       "Indumathi",
			Department: "Data Engineering",
			Salary:     50000,
		},
		102: {
			ID:         102,
			Name:       "Arun",
			Department: "IT",
			Salary:     45000,
		},
	}

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	employee, exists := employees[id]

	if exists {
		fmt.Println("Employee ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Department:", employee.Department)
		fmt.Println("Salary:", employee.Salary)
	} else {
		fmt.Println("Employee not found")
	}
}