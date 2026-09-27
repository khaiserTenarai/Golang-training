package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
}

// Method with value receiver
func (emp Employee) ShowInfo() {
	fmt.Println("Employee ID      :", emp.ID)
	fmt.Println("Employee Name    :", emp.Name)
	fmt.Println("Employee Position:", emp.Position)
}

func main() {

	// Create Employee object
	employee := Employee{
		ID:       101,
		Name:     "Muneera",
		Position: "Software Developer",
	}

	// Calling ShowInfo() method
	employee.ShowInfo()
}
