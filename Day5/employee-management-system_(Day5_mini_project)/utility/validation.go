package utility
import "strings"

func ValidateName(name string) error {

	if strings.TrimSpace(name) == "" {
		return ErrInvalidName
	}

	return nil
}

func ValidateAge(age int) error {

	if age <= 18 || age > 100 {
		return ErrInvalidAge
	}

	return nil
}

func ValidateSal(salary float64) error {

	if salary <= 0 {
		return ErrInvalidSalary
	}

	return nil
}