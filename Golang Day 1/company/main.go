package main

import (
    "fmt"
    "company/employee" 
    "company/utils"    
)

func main() {
    fmt.Println("--- Company System ---")

    emp := employee.Employee{
        Name:   "Alice",
        Salary: 3000,
    }

    employee.ShowDetails(emp)

   
    finalPay := utils.CalculateBonus(emp.Salary)

    fmt.Println("Final Pay (with bonus): $", finalPay)
}