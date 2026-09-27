// 5. Create an anonymous function for salary filtering.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	list := []Employee{
		{Name: "Sam", Salary: 45000},
		{Name: "Alex", Salary: 65000},
		{Name: "Taylor", Salary: 52000},
	}

	filter := func(emps []Employee, min float64) []Employee {
		var result []Employee
		for _, e := range emps {
			if e.Salary >= min {
				result = append(result, e)
			}
		}
		return result
	}

	highEarners := filter(list, 50000)
	fmt.Println("High Earners:", highEarners)
}
