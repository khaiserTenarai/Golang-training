
package main

import "fmt"

// Employee represents an employee.
type Employee struct {
	Name string
	ID   int
}

// ShowEmployee displays employee information.
func ShowEmployee(employee Employee) {
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Employee ID:", employee.ID)
}

func main() {

	employee := Employee{
		Name: "Swathi",
		ID:   101,
	}

	ShowEmployee(employee)
}


