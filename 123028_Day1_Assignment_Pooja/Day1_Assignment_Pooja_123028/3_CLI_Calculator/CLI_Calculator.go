package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid first number")
		return
	}
	operator := os.Args[2]
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid second number")
		return
	}

	switch operator {

	case "+":
		fmt.Println("Result:", x+y)

	case "-":
		fmt.Println("Result:", x-y)

	case "*":
		fmt.Println("Result:", x*y)

	case "/":
		if y == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Println("Result:", x/y)
		}

	default:
		fmt.Println("Invalid operator")
	}
}
