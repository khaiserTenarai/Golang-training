// 14. Perform a code review and identify 10 issues.
//
// This file is deliberately written with real, realistic problems in it.
// CODE_REVIEW.md (in this same folder) lists exactly 10 issues found in
// it, with the reasoning and a suggested fix for each. This file is left
// exactly as originally written (flawed) on purpose - the review is the
// answer, not this file.

package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	Name   string
	Salary float64
}

var Emp_list []*Employee

func AddEmployee(name string, salaryText string) {
	salary, _ := strconv.ParseFloat(salaryText, 64)

	emp := &Employee{Name: name, Salary: salary}
	Emp_list = append(Emp_list, emp)

	tax := emp.Salary * 0.1
	fmt.Println("Added", emp.Name, "tax is", tax)
}

func PrintAllNames() string {
	result := ""
	for _, emp := range Emp_list {
		result = result + emp.Name + ", "
	}
	return result
}

func GetEmployee(index int) *Employee {
	return Emp_list[index]
}

func PrintSalary(index int) {
	emp := GetEmployee(index)
	fmt.Println(emp.Salary)
}

func main() {
	AddEmployee("Anita", "45000")
	AddEmployee("Ravi", "abc")
	fmt.Println(PrintAllNames())
	PrintSalary(5)
}
