// 14. Demonstrate errors.As.

package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	MinAllowed float64
	Actual     float64
}

func (e *SalaryError) Error() string {
	return fmt.Sprintf("salary %.2f is below min limit %.2f", e.Actual, e.MinAllowed)
}

func checkSalary(amount float64) error {
	if amount < 30000 {
		return &SalaryError{MinAllowed: 30000, Actual: amount}
	}
	return nil
}

func main() {
	err := checkSalary(25000)
	if err != nil {
		var sErr *SalaryError
		if errors.As(err, &sErr) {
			fmt.Println("Salary error details:")
			fmt.Println("Minimum required:", sErr.MinAllowed)
			fmt.Println("Actual given:", sErr.Actual)
		}
	}
}
