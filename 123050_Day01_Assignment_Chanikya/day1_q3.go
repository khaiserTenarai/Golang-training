package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operator. ")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	switch operator {
	case "+":
		fmt.Println("Result: ", num1+num2)
	case "-":
		fmt.Println("Result: ", num1-num2)
	case "*":
		fmt.Println("Result: ", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero is not allowed.")
		} else {
			fmt.Println("Result: ", num1/num2)
		}
	}
}
