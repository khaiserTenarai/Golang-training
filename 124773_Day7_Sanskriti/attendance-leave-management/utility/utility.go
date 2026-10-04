package utility

import (
	"errors"
	"strings"
)

var (
	ErrEmployeeNotFound   = errors.New("employee not found")
	ErrAlreadyCheckedIn   = errors.New("employee already checked in today")
	ErrNotCheckedIn       = errors.New("employee has not checked in today")
	ErrAlreadyCheckedOut  = errors.New("employee already checked out today")
	ErrLeaveNotFound      = errors.New("leave application not found")
	ErrInvalidLeaveStatus = errors.New("leave is already processed")
)

func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" || !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	return nil
}

func CleanString(value string) string {
	return strings.TrimSpace(value)
}
