package main

import (
	"log"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func logEmployeeCreated(employee Employee) {
	log.Printf(
		"event=employee_created employee_id=%d employee_name=%s salary=%.2f",
		employee.ID,
		employee.Name,
		employee.Salary,
	)
}

func logEmployeeUpdated(employee Employee) {
	log.Printf(
		"event=employee_updated employee_id=%d employee_name=%s salary=%.2f",
		employee.ID,
		employee.Name,
		employee.Salary,
	)
}

func main() {

	employee := Employee{
		ID:     1,
		Name:   "ram",
		Salary: 50000,
	}

	logEmployeeCreated(employee)

	employee.Salary = 55000

	logEmployeeUpdated(employee)
}
