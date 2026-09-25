package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("CLI Calculator")
	fmt.Println("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("Addition: ", num1+num2)
	fmt.Println("Subtraction: ", num1-num2)
	fmt.Println("Multiplication: ", num1*num2)

	if num2 != 0 {
		fmt.Println("Division: ", num1/num2)
	} else {
		fmt.Println("Can't Divide a number by zero")
	}
}
