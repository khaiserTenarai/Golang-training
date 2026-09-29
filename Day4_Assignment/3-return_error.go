package main

import (
	"errors"
	"fmt"
)

func findEmployee(id int) (string, error) {

	if id == 1 {
		return "Vittesh", nil
	}

	return "", errors.New("employee not found")
}

func main() {

	name, err := findEmployee(1)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee:", name)
	}

	name, err = findEmployee(2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee:", name)
	}
}