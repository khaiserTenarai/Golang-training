package utility

import(
	"fmt"
	"errors"
	"strings"
)

//custom validation error
type ValidationError struct{
	Field string
	Message string
}

func (e ValidationError) Error() string{
	return fmt.Sprintf("%s: %s",e.Field, e.Message)
}

//sentinel errors

var ErrEmployeeNotFound = errors.New("Employee not found")
var ErrDuplicateEmployee = errors.New("Employee already exists")

func ValidateName(name string) error{
	if strings.TrimSpace(name) == ""{
		return ValidationError{
			Field: "Name",
			Message: "name cannot be empty",
		}
	}
	return nil
}

func ValidateEmail(email string) error{
	email = CleanString(email)

	if email==""{
		return ValidationError{
			Field: "Email",
			Message: "email cannot be empty",
		}
	}

	if !strings.Contains(email,"@"){
		return ValidationError{
			Field: "Email",
			Message: "invalid email",
		}
	}
	return nil
	
}

func ValidateAge(age int) error{
	if age < 18 || age > 60{
		return ValidationError{
			Field: "Age",
			Message: "Age is less than 18",
		}
	}
	return nil
}

func ValidateSalary(salary float64) error{
	if salary<=0{
		return ValidationError{
			Field: "Salary",
			Message: "Salary must be greater than 0",
		}
	}
	return nil
}

