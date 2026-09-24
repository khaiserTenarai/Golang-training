package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println("=== Employee Reporting System ===")

    strVal := "50"
    totalReports, err := strconv.Atoi(strVal)
    if err != nil {
        fmt.Println("Error reading report count")
        return
    }

    fmt.Printf("Total Reports Processed: %d\n", totalReports)
}