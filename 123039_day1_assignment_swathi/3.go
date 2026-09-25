
package main

import "fmt"

func main() {

	var num1, num2 float64
	var operator string
	var count int
	var err error

	fmt.Println("===== CLI CALCULATOR =====")

	fmt.Print("Enter first number: ")
	count, err = fmt.Scan(&num1)

	if err != nil || count != 1 {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	count, err = fmt.Scan(&operator)

	if err != nil || count != 1 {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Enter second number: ")
	count, err = fmt.Scan(&num2)

	if err != nil || count != 1 {
		fmt.Println("Invalid input")
		return
	}

	switch operator {

	case "+":
		fmt.Println("Result:", num1+num2)

	case "-":
		fmt.Println("Result:", num1-num2)

	case "*":
		fmt.Println("Result:", num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}

		fmt.Println("Result:", num1/num2)

	default:
		fmt.Println("Invalid operator")
	}
}

