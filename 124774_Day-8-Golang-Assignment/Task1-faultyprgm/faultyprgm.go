package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func calculateBonus(salary float64) float64 {
	// BUG: bonus should be 10% of salary
	bonus := salary * 10
	return bonus
}

func updateSalary(employee *Employee, increment float64) {
	// BUG: salary should increase by increment
	employee.Salary = employee.Salary - increment
}

func main() {

	employees := []Employee{
		{
			ID:     101,
			Name:   "Muneera",
			Salary: 20000,
		},
	}

	fmt.Println("Employee Management Application")
	fmt.Println("================================")

	for _, employee := range employees {

		fmt.Println("Employee ID:", employee.ID)
		fmt.Println("Employee Name:", employee.Name)
		fmt.Println("Employee Salary:", employee.Salary)

		bonus := calculateBonus(employee.Salary)

		fmt.Println("Bonus:", bonus)

		updateSalary(&employee, 5000)

		fmt.Println("Updated Salary:", employee.Salary)

		fmt.Println("--------------------------------")
	}
}
