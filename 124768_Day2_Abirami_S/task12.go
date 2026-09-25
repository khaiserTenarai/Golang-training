package main

import "fmt"

func main() {
	var choice int
	var num1, num2 int
	var continueChoice string
	for {
		fmt.Println("\nCalculator")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")

		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Enter 1st Number: ")
		fmt.Scan(&num1)

		fmt.Println("Enter 2nd Number: ")
		fmt.Scan(&num2)

		switch choice {
		case 1:
			fmt.Println("Sum: ", num1+num2)
		case 2:
			fmt.Println("Difference: ", num1-num2)
		case 3:
			fmt.Println("Product: ", num1*num2)
		case 4:
			if num2 != 0 {
				fmt.Println("Quotient: ", num1/num2)
			} else {
				fmt.Println("Can't divide a number by zero")
			}
		default:
			fmt.Println("Envalid choice")
		}
		fmt.Println("Do you want to continue? yes/no? ")
		fmt.Scan(&continueChoice)
		if continueChoice != "yes" {
			break
		}
	}
}
