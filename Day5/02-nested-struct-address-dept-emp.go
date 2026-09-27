package main

import "fmt"

type Address struct {
    City string
    State string
    Pincode int
}

type Department struct {
    Name string
	Loc string
}

type Employee struct {
    ID int
    Name string
    Age int
    Salary float64
    Email string
    Phone string
    Address Address
    Department Department
}

func main() {

    emp := Employee{
        ID: 101,
        Name: "Ganesh",
        Age: 22,
        Salary: 40000,
        Email: "ganesh@gmail.com",
        Phone: "9876543210",

        Address: Address{
            City: "Bengaluru",
            State: "Karnataka",
            Pincode: 560100,
        },

        Department: Department{
            Name: "IT",
            Loc  : "Bangalore",
        },
    }

    fmt.Println(emp)
    fmt.Println(emp.Address.City)
    fmt.Println(emp.Department.Name)
}

/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\02-nested-struct-address-dept-emp.go
{101 Ganesh 22 40000 ganesh@gmail.com 9876543210 {Bengaluru Karnataka 560100} {IT Bangalore}}
Bengaluru
IT
*/