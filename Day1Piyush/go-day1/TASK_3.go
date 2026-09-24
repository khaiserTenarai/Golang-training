package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Println("Enter first number: ")
	fmt.Scanf("%f",&num1)

	fmt.Println("Enter operator (+, -, *, /): ")
	fmt.Scanf("%f",&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanf("%f",&num2)

	if operator == "+" {
		fmt.Println("Result:", num1+num2)
	} else if operator == "-" {
		fmt.Println("Result:", num1-num2)
	} else if operator == "*" {
		fmt.Println("Result:", num1*num2)
	} else if operator == "/" {
		fmt.Println("Result:", num1/num2)
	} else {
		fmt.Println("Invalid operator")
	}

}
