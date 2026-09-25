package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
}

func main() {
	var employees []Employee

	employees = append(employees, Employee{ID: 101, Name: "Aarav", Position: "Developer"})
	employees = append(employees, Employee{ID: 102, Name: "Priya", Position: "Designer"})
	employees = append(employees, Employee{ID: 103, Name: "Rohan", Position: "QA Engineer"})
	display(employees)

	for _, emp := range employees {
		if emp.ID == 102 {
			fmt.Printf("Found: ID: %d | Name: %s | Position: %s\n", emp.ID, emp.Name, emp.Position)
		}
	}

	for i := range employees {
		if employees[i].ID == 102 {
			employees[i].Position = "Lead Designer"
		}
	}
	display(employees)

	deleteID := 101
	for i, emp := range employees {
		if emp.ID == deleteID {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}
	display(employees)
}

func display(list []Employee) {
	for _, emp := range list {
		fmt.Printf("ID: %d | Name: %-10s | Position: %s\n", emp.ID, emp.Name, emp.Position)
	}
	fmt.Println()
}
