package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Salary float64
}

func calculateSalary(employee Employee) {
	finalSalary := employee.Salary + 5000

	fmt.Printf("%s final salary: %.2f\n", employee.Name, finalSalary)

	time.Sleep(500 * time.Millisecond)
}

func main() {
	employees := []Employee{
		{Name: "max", Salary: 50000},
		{Name: "Rahul", Salary: 60000},
		{Name: "Suresh", Salary: 70000},
	}

	for _, employee := range employees {
		go calculateSalary(employee)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("All calculations completed")
}
