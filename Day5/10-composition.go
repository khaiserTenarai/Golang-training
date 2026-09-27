package main

import "fmt"

type Address struct {
    City string
    State string
}

type Employee struct {
    Name string
    Address Address
}

func main() {

    emp := Employee{
        Name: "Ganesh",
        Address: Address{
            City: "Bengaluru",
            State: "Karnataka",
        },
    }

    fmt.Println(emp.Name)
    fmt.Println(emp.Address.City)
    fmt.Println(emp.Address.State)
}


/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\10-composition.go
Ganesh
Bengaluru
Karnataka
*/