// 3. CLI Calculator
//    - Create a command-line calculator supporting:
//      - Addition
//      - Subtraction
//      - Multiplication
//      - Division
//    - Handle invalid input.

package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	fmt.Print("Enter the first number:")
	fmt.Scan(&num1)

	fmt.Print("Enter an the operation you should make(+,-,*,/): ")
	fmt.Scan(&op)

	fmt.Print("Enter the second number:")
	fmt.Scan(&num2)

	switch op {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("can't divide by zero.")
		} else {
			fmt.Println("Result:", num1/num2)
		}
	default:
		fmt.Println("can't recognize that oprator")
	}
}
