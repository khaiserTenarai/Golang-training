package model

import "fmt"

type Address struct {
    City    string
    State   string
    Pincode string
}

type Employee struct {
    ID      int64
    Name    string
    Email   string
    Age     int
    Salary  float64
    Address Address
    DepartmentID *int
    DepartmentName string
}

// Method with value receiver.
func (e Employee) Display() {
    fmt.Println("----------------------------------------")
    fmt.Println("ID      :", e.ID)
    fmt.Println("Name    :", e.Name)
    fmt.Println("Email   :", e.Email)
    fmt.Println("Age     :", e.Age)
    fmt.Printf("Salary  : %.2f\n", e.Salary)
    fmt.Println("City    :", e.Address.City)
    fmt.Println("State   :", e.Address.State)
    fmt.Println("Pincode :", e.Address.Pincode)
    fmt.Println("----------------------------------------")
}

// Method with pointer receiver.
func (e *Employee) UpdateSalary(salary float64) {
    e.Salary = salary
}
