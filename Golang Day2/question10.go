package main

import "fmt"

type Employee struct {
    ID   int
    Name string
}

func main() {
    
    var db []Employee

    db = append(db, Employee{ID: 1, Name: "Alice"})
    db = append(db, Employee{ID: 2, Name: "Bob"})
    fmt.Println("After Create:", db)

   
    fmt.Println("\nReading Data:")
    for _, emp := range db {
        fmt.Println(emp.ID, "-", emp.Name)
    }

    for i := range db {
        if db[i].ID == 2 {
            db[i].Name = "Robert" 
        }
    }
    fmt.Println("\nAfter Update:", db)

    for i := range db {
        if db[i].ID == 1 {
            db = append(db[:i], db[i+1:]...)
            break 
        }
    }
    fmt.Println("After Delete:", db)
}