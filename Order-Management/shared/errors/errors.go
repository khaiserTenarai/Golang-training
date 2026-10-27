package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInvalidInput = errors.New("invalid input")
	ErrConflict     = errors.New("resource conflict")
)

// AppError adds context to an application error.
type AppError struct {
	Operation string
	Err       error
}

// Error returns the formatted error message.
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Operation, e.Err)
}

// Unwrap exposes the underlying error for errors.Is and errors.As.
func (e *AppError) Unwrap() error {
	return e.Err
}

// Wrap adds an operation name to an existing error.
func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return &AppError{
		Operation: operation,
		Err:       err,
	}
}
