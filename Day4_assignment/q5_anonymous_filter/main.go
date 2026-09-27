
// Day 4, Q5. Create an anonymous function for salary filtering.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	employees := []Employee{
		{"Ranjitha", 62000},
		{"Meera", 41000},
		{"Ravi", 55000},
		{"Divya", 38000},
	}

	filterAbove := func(list []Employee, threshold float64) []Employee {
		var result []Employee
		for _, e := range list {
			if e.Salary > threshold {
				result = append(result, e)
			}
		}
		return result
	}

	highEarners := filterAbove(employees, 50000)
	fmt.Println("Employees earning above 50000:")
	for _, e := range highEarners {
		fmt.Println(" -", e.Name, e.Salary)
	}
}
