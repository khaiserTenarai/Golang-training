package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&operator)

	
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator.")
		return
	}


	fmt.Printf("Result: %v %s %v = %v\n", num1, operator, num2, result)
}