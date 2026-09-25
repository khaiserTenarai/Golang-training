package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	fmt.Print("Enter first number: ")
	if _, err := fmt.Scan(&num1); err != nil {
		fmt.Println("Invalid number!")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Enter second number: ")
	if _, err := fmt.Scan(&num2); err != nil {
		fmt.Println("Invalid number!")
		return
	}

	switch op {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero.")
			return
		}
		fmt.Printf("Result: %.2f\n", num1/num2)
	default:
		fmt.Println("Unknown operator:", op)
	}
}