package main

import "fmt"

func main() {

    sales := [12]int{1000, 1200, 1500, 1300, 1600, 2000, 2100, 1900, 1800, 2200, 2500, 3000}

    total := 0

    for _, monthSale := range sales {
        total = total + monthSale 
    }

    average := float64(total) / 12.0

    fmt.Println("--- Yearly Sales Report ---")
    fmt.Println("All Months   :", sales)
    fmt.Println("Total Sales  : $", total)
    
    fmt.Printf("Average Sales: $%.2f\n", average) 
}