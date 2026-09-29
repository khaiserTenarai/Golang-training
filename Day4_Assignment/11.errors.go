package main
import "fmt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func validateEmployee(name string, salary float64) error {

	if name == "" {
		return ValidationError{
			Message: "Employee name cannot be empty",
		}
	}

	if salary <= 0 {
		return ValidationError{
			Message: "Salary must be greater than zero",
		}
	}

	return nil
}

func main() {

	err := validateEmployee("", 30000)

	if err != nil {
		fmt.Println("Error:", err)
	}
}