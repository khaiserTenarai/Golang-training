package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	if age < 18 {
		return errors.New("age must be 18 or above")
	}
	return nil
}
func main() {
	err := checkAge(17)
	if err != nil {
		fmt.Println(err)
	}
}
