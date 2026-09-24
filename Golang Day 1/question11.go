package main

import "fmt"

type Employee struct {
    Name       string
    Department string
}

func main() {
    
    database := map[string]Employee{
        "101": {Name: "Lakshmi", Department: "Engineering"},
        "102": {Name: "Esther", Department: "HR"},
        "103": {Name: "Ayush", Department: "Design"},
    }

    var searchID string

    fmt.Println("Employee Search System ")
    
    fmt.Print("Please enter an Employee ID (e.g., 101): ")
    fmt.Scan(&searchID)

    emp, exists := database[searchID]

    if exists {
        fmt.Println("\nEmployee Found!")
        fmt.Println("ID        :", searchID)
        fmt.Println("Name      :", emp.Name)
        fmt.Println("Department:", emp.Department)
    } else {
        fmt.Println("\n No employee found with ID:", searchID)
    }
}