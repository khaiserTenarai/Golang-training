package utility
import (
	"errors"
	"strings"

	"ems/model"
)

/*
	Custom validation errors.
*/
var (
	InvalidEmployeeIDError =
		errors.New("invalid employee ID")

	InvalidEmployeeNameError =
		errors.New("invalid employee name")

	InvalidEmployeeAgeError =
		errors.New("invalid employee age")

	InvalidEmailError =
		errors.New("invalid employee email")

	InvalidEmployeeSalaryError =
		errors.New("invalid employee salary")
)

/*
	ValidateEmployee validates employee data.
*/
func ValidateEmployee(
	employee model.Employee,
) error {

	if employee.ID <= 0 {
		return InvalidEmployeeIDError
	}

	if strings.TrimSpace(employee.Name) == "" {
		return InvalidEmployeeNameError
	}

	if employee.Age <= 18 || employee.Age >= 60 {
		return InvalidEmployeeAgeError
	}

	if employee.Salary <= 0 {
		return InvalidEmployeeSalaryError
	}

	return nil
}

/*
	ValidateEmployeeID validates employee ID.
*/
func ValidateEmployeeID(id int) error {

	if id <= 0 {
		return InvalidEmployeeIDError
	}

	return nil
}

/*
	ValidateEmail validates employee email.
*/
func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)

	if len(email) <= 3 ||
		len(email) >= 100 ||
		!strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {

		return InvalidEmailError
	}

	return nil
}