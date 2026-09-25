package employee

import (
	"fmt"
	"my-employee-app/pkg/formatter"
)

// Employee defines the structure of our employee data
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// Display prints the employee's details cleanly to the console
func Display(emp Employee) {
	header := formatter.FormatHeader("Employee Profile")
	fmt.Println(header)
	fmt.Printf("ID:       %d\n", emp.ID)
	fmt.Printf("Name:     %s\n", emp.Name)
	fmt.Printf("Position: %s\n", emp.Position)
	fmt.Printf("Salary:   $%.2f\n\n", emp.Salary)
}
