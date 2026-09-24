package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ErrDivideByZero is returned when dividing by zero.
var ErrDivideByZero = errors.New("division by zero is not allowed")

// Calculate applies op to a and b.
func Calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*", "x", "X":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivideByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator %q (use + - * x /)", op)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number> <operator> <number>")
		fmt.Println("Example: calculator 10 + 5")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Error: %q is not a valid number\n", os.Args[1])
		os.Exit(1)
	}
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Error: %q is not a valid number\n", os.Args[3])
		os.Exit(1)
	}

	result, err := Calculate(a, os.Args[2], b)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("%g %s %g = %g\n", a, os.Args[2], b, result)
}
