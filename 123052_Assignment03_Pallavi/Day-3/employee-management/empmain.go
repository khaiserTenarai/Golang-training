package main

import "fmt"

func main() {
    employee := Employee{
        ID:     101,
        Name:   "Rajesh",
        Salary: 75000,
    }

    fmt.Println("Employee Management Application")
    fmt.Println("Employee ID:", employee.ID)
    fmt.Println("Employee Name:", employee.Name)
    fmt.Println("Employee Salary:", employee.Salary)
}