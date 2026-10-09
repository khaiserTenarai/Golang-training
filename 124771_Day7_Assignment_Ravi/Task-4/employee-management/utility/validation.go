package utility

import "fmt"

func ValidateEmployeeID(
	employeeID int64,
) error {

	if employeeID <= 0 {

		return fmt.Errorf(
			"employee ID must be greater than 0",
		)
	}

	return nil
}

func ValidateSalary(
	salary float64,
) error {

	if salary < 0 {

		return fmt.Errorf(
			"salary cannot be negative",
		)
	}

	return nil
}
