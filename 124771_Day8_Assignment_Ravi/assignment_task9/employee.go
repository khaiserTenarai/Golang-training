package main

import "errors"

type Employee struct {
	Name string
	Age  int
}

func (e Employee) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	}
	if e.Age <= 0 {
		return errors.New("age must be greater than zero")
	}
	return nil
}
