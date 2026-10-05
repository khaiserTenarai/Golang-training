package main

import "errors"

type Employee struct {
	ID   int
	Name string
	Age  int
}

func ValidateEmployee(e Employee) error {
	if e.ID <= 0 {
		return errors.New("invalid ID")
	}
	if e.Name == "" {
		return errors.New("name cannot be empty")
	}
	if e.Age < 18 {
		return errors.New("employee must be at least 18")
	}
	return nil
}