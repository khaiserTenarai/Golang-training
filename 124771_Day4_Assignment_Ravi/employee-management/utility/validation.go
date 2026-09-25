package utility

import (
	"regexp"

	"employee-management/model"
)

// ValidationError represents a validation failure.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// ValidateName validates employee name.
func ValidateName(name string) error {
	name = CleanString(name)

	if name == "" {
		return &ValidationError{
			Field:   "Name",
			Message: "name cannot be empty",
		}
	}

	if len(name) < 2 {
		return &ValidationError{
			Field:   "Name",
			Message: "name must contain at least 2 characters",
		}
	}

	return nil
}

// ValidateEmail validates employee email.
func ValidateEmail(email string) error {
	email = CleanString(email)

	if email == "" {
		return &ValidationError{
			Field:   "Email",
			Message: "email cannot be empty",
		}
	}

	emailRegex := regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	if !emailRegex.MatchString(email) {
		return &ValidationError{
			Field:   "Email",
			Message: "invalid email format",
		}
	}

	return nil
}

// ValidateAge validates employee age.
func ValidateAge(age int) error {
	if age < 18 || age > 100 {
		return &ValidationError{
			Field:   "Age",
			Message: "age must be between 18 and 100",
		}
	}

	return nil
}

// ValidateSalary validates employee salary.
func ValidateSalary(salary float64) error {
	if salary < 0 {
		return &ValidationError{
			Field:   "Salary",
			Message: "salary cannot be negative",
		}
	}

	return nil
}

// ValidateEmployee validates all employee fields.
func ValidateEmployee(employee model.Employee) error {
	if err := ValidateName(employee.Name); err != nil {
		return err
	}

	if err := ValidateEmail(employee.Email); err != nil {
		return err
	}

	if err := ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := ValidateSalary(employee.Salary); err != nil {
		return err
	}

	return nil
}
