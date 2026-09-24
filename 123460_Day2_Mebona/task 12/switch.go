package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\nCalculation")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Print("Choose an operation: ")

		var op int
		fmt.Scan(&op)

		var x, y float64
		fmt.Print("Enter two numbers: ")
		fmt.Scan(&x, &y)

		switch op {
		case 1:
			fmt.Printf("Result: %.2f\n", x+y)
		case 2:
			fmt.Printf("Result: %.2f\n", x-y)
		case 3:
			fmt.Printf("Result: %.2f\n", x*y)
		case 4:
			if y == 0 {
				fmt.Println("Error: Cannot divide by zero.")
			} else {
				fmt.Printf("Result: %.2f\n", x/y)
			}
		}
	}
}