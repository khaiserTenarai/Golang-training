package utility

import (
 "errors"
 "strings"
 "student-management/model"
)

func ValidateStudent(student model.Student) error {
 if strings.TrimSpace(student.Name) == "" { return errors.New("name is required") }
 if strings.TrimSpace(student.Grade) == "" { return errors.New("grade is required") }
 return nil
}
