package utility

import (
	"errors"
	"strings"

	"student-management/model"
)

// ValidateStudent checks student details.
func ValidateStudent(student model.Student) error {
	if strings.TrimSpace(student.Name) == "" {
		return errors.New("name is required")
	}

	if student.Age < 1 || student.Age > 120 {
		return errors.New("age must be between 1 and 120")
	}

	if strings.TrimSpace(student.Grade) == "" {
		return errors.New("grade is required")
	}

	return nil
}
