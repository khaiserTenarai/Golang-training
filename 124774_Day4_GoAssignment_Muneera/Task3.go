package main

import "fmt"

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	result := float64(a) / float64(b)
	return result, nil
}

func main() {
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
