package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n3. CLI Calculator")
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scanln(&num1)
	if err1 != nil {
		fmt.Println("Error: Invalid number input.")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scanln(&num2)
	if err2 != nil {
		fmt.Println("Error: Invalid number input.")
		return
	}

	if operator == "+" {
		fmt.Println("Result:", num1+num2)
	} else if operator == "-" {
		fmt.Println("Result:", num1-num2)
	} else if operator == "*" {
		fmt.Println("Result:", num1*num2)
	} else if operator == "/" {
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			fmt.Println("Result:", num1/num2)
		}
	} else {
		fmt.Println("Error: Invalid operator. Use +, -, *, or /.")
	}

}