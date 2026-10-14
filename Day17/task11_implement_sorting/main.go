package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	emps := []Employee{
		{"Bob", 60000},
		{"Alice", 90000},
		{"Charlie", 75000},
	}

	// Sort by Salary Descending
	sort.Slice(emps, func(i, j int) bool {
		return emps[i].Salary > emps[j].Salary
	})

	fmt.Printf("Sorted Employees (Salary Desc): %+v
", emps)
}
