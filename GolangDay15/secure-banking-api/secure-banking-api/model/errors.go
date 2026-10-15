package model

import "errors"

var (
	ErrNotFound           = errors.New("resource not found")
	ErrForbidden          = errors.New("you do not have access to this resource")
	ErrEmailExists        = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
)
