package main

import "fmt"

type Employee struct {
    ID       int
    Name     string
    Age      int
    Salary   float64
    Email    string
    Phone    string
    City     string
    Position string
}

func main() {
    emp := Employee{
        ID:       101,
        Name:     "Ganesh",
        Age:      22,
        Salary:   40000,
        Email:    "ganesh@gmail.com",
        Phone:    "9876543210",
        City:     "Bengaluru",
        Position: "Software Engineer",
    }

    fmt.Println(emp)
}

/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\01-employee-struct-8Fields.go
{101 Ganesh 22 40000 ganesh@gmail.com 9876543210 Bengaluru Software Engineer}
*/