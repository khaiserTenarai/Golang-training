package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// Method with pointer receiver
func (emp *Employee) UpdateSalary(newSalary float64) {
	emp.Salary = newSalary
}

func main() {

	// Create Employee object
	employee := Employee{
		ID:       101,
		Name:     "Muneera",
		Position: "Software Developer",
		Salary:   20000,
	}

	fmt.Println("Before Salary Update:")
	fmt.Println("Employee Name  :", employee.Name)
	fmt.Println("Employee Salary :", employee.Salary)

	// Calling method with pointer receiver
	employee.UpdateSalary(25000)

	fmt.Println("\nAfter Salary Update:")
	fmt.Println("Employee Name  :", employee.Name)
	fmt.Println("Employee Salary :", employee.Salary)
}
