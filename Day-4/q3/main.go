package main

import (
	"errors"
	"fmt"
)

func divide(a float64, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {

	result, err := divide(20, 5)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}