package utility
import (
	"errors"
	"strings"

	"employee-management/model"
)

var (
	InvalidEmployeeID       = errors.New("invalid employee ID")
	InvalidEmployeeName     = errors.New("invalid employee name")
	InvalidEmployeeAge      = errors.New("invalid employee age")
	InvalidEmployeeEmail    = errors.New("invalid employee email")
	InvalidEmployeeSalary   = errors.New("invalid employee salary")
	InvalidDepartmentID     = errors.New("invalid department ID")
	InvalidDepartmentName   = errors.New("invalid department name")
)

func ValidateEmployee(employee model.Employee) error {

	if employee.ID < 0 {
		return InvalidEmployeeID
	}

	if strings.TrimSpace(employee.Name) == "" {
		return InvalidEmployeeName
	}

	if employee.Age < 18 || employee.Age > 60 {
		return InvalidEmployeeAge
	}

	if !strings.Contains(employee.Email, "@") ||
		!strings.Contains(employee.Email, ".") {
		return InvalidEmployeeEmail
	}

	if employee.Salary <= 0 {
		return InvalidEmployeeSalary
	}

	if employee.Department.ID <= 0 {
		return InvalidDepartmentID
	}

	return nil
}

func ValidateDepartment(
	department model.Department,
) error {

	if strings.TrimSpace(department.Name) == "" {
		return InvalidDepartmentName
	}

	return nil
}
