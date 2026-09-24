package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// readNumber repeatedly prompts the user until a valid float is entered.
func readNumber(prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err == nil {
			return num
		}
		fmt.Println("Invalid number! Please enter a valid numerical value.")
	}
}

// readOperator repeatedly prompts until a valid operator (+, -, *, /) is selected.
func readOperator() string {
	validOperators := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	for {
		fmt.Print("Enter operator (+, -, *, /): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading operator. Please try again.")
			continue
		}

		op := strings.TrimSpace(input)
		if validOperators[op] {
			return op
		}
		fmt.Println("Invalid operator! Please choose one of: +, -, *, /")
	}
}

func main() {

	num1 := readNumber("Enter first number: ")
	operator := readOperator()

	var num2 float64
	for {
		num2 = readNumber("Enter second number: ")
		if operator == "/" && num2 == 0 {
			fmt.Println("Error: Division by zero is undefined! Please enter a non-zero second number.")
			continue
		}
		break
	}

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	fmt.Printf("Result: %g %s %g = %g\n", num1, operator, num2, result)

}
