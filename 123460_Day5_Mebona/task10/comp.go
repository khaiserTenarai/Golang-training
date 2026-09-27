package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func (e Employee) DisplayEmployee() {
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
}

type Department struct {
	Name string
}

func (d Department) DisplayDepartment() {
	fmt.Println("Department:", d.Name)
}

// Composition
type EmployeeDetails struct {
	Employee
	Department
}

func main() {
	details := EmployeeDetails{
		Employee: Employee{
			ID:   101,
			Name: "John",
		},
		Department: Department{
			Name: "IT",
		},
	}

	details.DisplayEmployee()
	details.DisplayDepartment()
}