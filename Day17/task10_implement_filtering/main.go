package main

import "fmt"

type Employee struct {
	Name string
	Dept string
}

func FilterByDept(list []Employee, dept string) []Employee {
	var filtered []Employee
	for _, e := range list {
		if e.Dept == dept {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func main() {
	emps := []Employee{{"Alice", "IT"}, {"Bob", "HR"}, {"Charlie", "IT"}}
	result := FilterByDept(emps, "IT")
	fmt.Printf("Filtered Employees (IT): %+v
", result)
}
