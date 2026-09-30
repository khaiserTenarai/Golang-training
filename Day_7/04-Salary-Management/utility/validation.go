package utility
import (
	"errors"
	"strings"

	"salary-management/model"
)

var (
	InvalidEmployeeIDError =
		errors.New("invalid employee ID")

	InvalidEmployeeNameError =
		errors.New("invalid employee name")

	InvalidEmployeeAgeError =
		errors.New("invalid employee age")

	InvalidEmailError =
		errors.New("invalid employee email")

	InvalidSalaryError =
		errors.New("invalid salary")
)

func ValidateEmployee(
	employee model.Employee,
) error {

	if strings.TrimSpace(employee.Name) == "" {
		return InvalidEmployeeNameError
	}

	if employee.Age <= 18 || employee.Age >= 60 {
		return InvalidEmployeeAgeError
	}

	if employee.Salary < 0 {
		return InvalidSalaryError
	}

	return nil
}

func ValidateEmployeeID(
	id int,
) error {

	if id <= 0 {
		return InvalidEmployeeIDError
	}

	return nil
}

func ValidateEmail(
	email string,
) error {

	email = strings.TrimSpace(email)

	if email == "" {
		return InvalidEmailError
	}

	if len(email) <= 3 {
		return InvalidEmailError
	}

	if !strings.Contains(email, "@") {
		return InvalidEmailError
	}

	if !strings.Contains(email, ".") {
		return InvalidEmailError
	}

	return nil
}

func ValidateSalary(
	salary float64,
) error {

	if salary < 0 {
		return InvalidSalaryError
	}

	return nil
}
