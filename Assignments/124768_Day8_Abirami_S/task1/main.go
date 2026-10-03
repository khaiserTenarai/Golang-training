package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func calculateSalary(salary float64) float64 {
	bonus := salary * 0.10
	return salary + float64(bonus)
}
func main() {
	employee := Employee{
		ID:     101,
		Name:   "Tom",
		Salary: 70000,
	}
	finalSalary := calculateSalary(employee.Salary)
	fmt.Println("Employee ID: ", employee.ID)
	fmt.Println("Employee Name: ", employee.Name)
	fmt.Println("Basic Salary: ", employee.Salary)
	fmt.Println("Final Salary: ", finalSalary)
}
