
package main

import (
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("employee not found")

func getSalary(id int) (float64, error) {
	if id != 1 {
		return 0, fmt.Errorf("getSalary(%d): %w", id, ErrEmployeeNotFound)
	}
	return 60000, nil
}

func main() {
	_, err := getSalary(7)

	if errors.Is(err, ErrEmployeeNotFound) {
		fmt.Println("Confirmed: employee not found for ranjitha's team lookup")
	} else {
		fmt.Println("Some other error occurred:", err)
	}
}
