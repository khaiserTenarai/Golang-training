package main

import (
	"errors"
	"fmt"
)

func validateMarks(marks int) (string, error) {
	if marks < 90 {
		return "", errors.New("Not Eligible for Scholorship")
	}
	return "Eligible for Scholorship", nil
}
func main() {
	result, err := validateMarks(89)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
