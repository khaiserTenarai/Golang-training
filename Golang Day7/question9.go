package main

import (
	"errors"
	"strings"
)

type Employee struct {
	Name string
	Age  int
	Role string
}

func ValidateEmployee(e Employee) error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if e.Age < 18 || e.Age > 65 {
		return errors.New("age must be between 18 and 65")
	}
	if strings.TrimSpace(e.Role) == "" {
		return errors.New("role cannot be empty")
	}
	return nil
}