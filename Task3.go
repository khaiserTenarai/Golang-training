package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Println("***CLI Calculator***")
	fmt.Print("Enter first number: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Invalid input: please enter a number.")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	if _, err := fmt.Scan(&operator); err != nil {
		fmt.Println("Invalid input: please enter an operator.")
		return
	}

	fmt.Print("Enter second number: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Invalid input: please enter a number.")
		return
	}

	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Error: division by zero is not allowed.")
			return
		}
		result = a / b
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
