package main

import (
    "fmt"
    "strconv"
)

type Employee struct {
    ID   int
    Name string
}

func main() {
    fmt.Println("=== Employee Management & Reporting System ===")

    // Employee Logic
    strVal := "101"
    empID, err := strconv.Atoi(strVal)
    if err != nil {
        fmt.Println("Error converting Employee ID")
        return
    }
    emp := Employee{ID: empID, Name: "Alice"}
    fmt.Printf("Employee Created: ID=%d, Name=%s\n", emp.ID, emp.Name)

    // Reporting Logic
    reportVal := "50"
    totalReports, _ := strconv.Atoi(reportVal)
    fmt.Printf("Total Reports Processed: %d\n", totalReports)
}
