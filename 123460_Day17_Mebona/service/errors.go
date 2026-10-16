package service

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail      = errors.New("email already exists")
	ErrDuplicateDepartment = errors.New("department name already exists")
	ErrInvalidCredentials  = errors.New("invalid username or password")
)

// ValidationError is returned when the user's input breaks a business rule.
type ValidationError struct{ Message string }

func (e *ValidationError) Error() string { return e.Message }
func invalid(msg string) error           { return &ValidationError{msg} }

func IsNotFound(err error) bool { return errors.Is(err, gorm.ErrRecordNotFound) }
func IsValidation(err error) bool {
	var v *ValidationError
	return errors.As(err, &v)
}
func IsDuplicate(err error) bool {
	return errors.Is(err, ErrDuplicateEmail) || errors.Is(err, ErrDuplicateDepartment)
}
