package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func calculateBonus(salary float64) float64 {
	bonus := salary * 0.10
	return bonus
}

func updateSalary(employee *Employee, increment float64) {
	employee.Salary = employee.Salary + increment
}

func main() {
	employees := []Employee{
		{ID: 101, Name: "Vimal", Salary: 500000},
		{ID: 102, Name: "Aditya", Salary: 100000},
	}

	totalBonus := 0.0

	for _, employee := range employees {
		// BUG 1: Passing value instead of pointer to updateSalary
		// BUG 2: Incorrect loop calculation logic for total bonus
		bonus := calculateBonus(employee.Salary)
		totalBonus = bonus // Overwrites instead of accumulating

		updateSalary(&employee, bonus)
	}

	fmt.Printf("Total Bonus Calculated: %.2f\n", totalBonus)
	fmt.Printf("Employee 101 Updated Salary: %.2f\n", employees[0].Salary)
}

// Start dlv debug assignment_task1.go
// break calculateBonus
// break updateSalary 
// continue
// print employee
// next
// print employee.Salary 

//for bug2-
// break 31
// continue
// print totalBonus
// continue
// print totalBonus

