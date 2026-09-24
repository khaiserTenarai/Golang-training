package employee

import "fmt"

type Employee struct {
    Name   string
    Salary int
}

func ShowDetails(emp Employee) {
    fmt.Println("Employee Name:", emp.Name)
    fmt.Println("Base Salary: $", emp.Salary)
}