package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func main() {
	employees := []Employee{
		{"Indu", 50000},
		{"Arun", 40000},
		{"Ravi", 60000},
	}

	filterSalary := func(salary int) bool {
		return salary >= 50000
	}

	for _, employee := range employees {
		if filterSalary(employee.salary) {
			fmt.Println(employee.name, employee.salary)
		}
	}
}