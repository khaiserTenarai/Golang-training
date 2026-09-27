package main

import "fmt"

type Department struct {
	Name string
}

type Employee struct {
	Name       string
	Department Department
}

func main() {

	department := Department{
		Name: "IT",
	}

	employee := Employee{
		Name:       "Muneera",
		Department: department,
	}

	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Department:", employee.Department.Name)
}
