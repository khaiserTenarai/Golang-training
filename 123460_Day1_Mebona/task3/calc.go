package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var op string
	
	fmt.Println("Enter num1, num2 and operator:")
	_, err := fmt.Scan(&num1, &num2, &op)

	if err != nil {
		fmt.Println("Invalid input")
		return
	} else {
		switch op {
		case "+":
			fmt.Println("Result:", num1+num2)
		case "-":
			fmt.Println("Result:", num1-num2)
		case "*":
			fmt.Println("Result:", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
			} else {
				fmt.Println("Result:", num1/num2)
			}
		default:
			fmt.Println("Error: Invalid operator")
		}
	}
}