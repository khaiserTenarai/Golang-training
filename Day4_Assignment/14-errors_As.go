package main
import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func validateEmployee(name string) error {

	if name == "" {
		return ValidationError{
			Message: "Name cannot be empty",
		}
	}

	return nil
}

func main() {

	err := validateEmployee("")

	if err != nil {

		var validationError ValidationError

		if errors.As(err, &validationError) {

			fmt.Println("Validation Error:")
			fmt.Println(validationError.Message)
		}
	}
}