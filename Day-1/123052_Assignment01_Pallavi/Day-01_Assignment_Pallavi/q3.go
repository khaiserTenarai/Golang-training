package main

import "fmt"

func main() {

	var number1 float64
	var number2 float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&number1)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&number2)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	switch operator {

	case "+":
		fmt.Println("Result:", number1+number2)

	case "-":
		fmt.Println("Result:", number1-number2)

	case "*":
		fmt.Println("Result:", number1*number2)

	case "/":
		if number2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}

		fmt.Println("Result:", number1/number2)

	default:
		fmt.Println("Invalid operator")
	}
}