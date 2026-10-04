package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func calculateEmployee(employee Employee) {
	bonus := employee.Salary * 0.10
	totalSalary := employee.Salary + bonus

	fmt.Printf(
		"Employee: %s | Salary: %.2f | Bonus: %.2f | Total Salary: %.2f\n",
		employee.Name,
		employee.Salary,
		bonus,
		totalSalary,
	)

	time.Sleep(1 * time.Second)
}

func main() {
	employees := []Employee{
		{ID: 101, Name: "Indu", Salary: 30000},
		{ID: 102, Name: "Rahul", Salary: 40000},
		{ID: 103, Name: "Priya", Salary: 50000},
		{ID: 104, Name: "Arun", Salary: 60000},
	}

	for _, employee := range employees {
		go calculateEmployee(employee)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("All employee calculations completed.")
}