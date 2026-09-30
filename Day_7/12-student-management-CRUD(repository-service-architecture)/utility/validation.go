package utility
import (
	"errors"
	"strings"

	"student_app/model"
)

var (
	InvalidIDError =
		errors.New("invalid student ID")

	InvalidNameError =
		errors.New("invalid student name")

	InvalidAgeError =
		errors.New("invalid student age")

	InvalidEmailError =
		errors.New("invalid email")
)

func ValidateID(id int) error {

	if id <= 0 {
		return InvalidIDError
	}

	return nil
}

func ValidateStudent(
	student model.Student,
) error {

	if strings.TrimSpace(student.Name) == "" {
		return InvalidNameError
	}

	if student.Age <= 0 {
		return InvalidAgeError
	}

	return nil
}

func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)

	if len(email) <= 3 ||
		len(email) >= 100 {
		return InvalidEmailError
	}

	if !strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {
		return InvalidEmailError
	}

	return nil
}
