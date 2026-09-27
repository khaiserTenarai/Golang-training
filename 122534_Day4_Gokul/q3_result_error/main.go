// Day 4, Q3. Create functions returning (result, error).

package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func calculateBonus(salary float64, years int) (float64, error) {
	if years < 0 {
		return 0, errors.New("years of experience cannot be negative")
	}
	return salary * float64(years) * 0.02, nil
}

func main() {
	result, err := divide(100, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("100 / 5 =", result)
	}

	_, err = divide(50, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	bonus, err := calculateBonus(50000, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Bonus for Gokul:", bonus)
	}
}
