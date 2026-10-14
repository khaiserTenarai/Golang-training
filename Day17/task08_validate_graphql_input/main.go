package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

func ValidateAndCreate(salary float64) error {
	if salary <= 0 {
		return fmt.Errorf("input validation failed: salary must be > 0, got %f", salary)
	}
	return nil
}

func main() {
	err := ValidateAndCreate(-100)
	fmt.Printf("GraphQL Input Validation check: %v
", err)
}
