package model

import "errors"

var (
	ErrNotFound        = errors.New("employee not found")
	ErrConflict        = errors.New("conflict")
	ErrFeatureDisabled = errors.New("feature disabled")
)

// ValidationError lists the invalid fields, e.g. {"email": "must be a valid email address"}.
type ValidationError struct {
	Fields map[string]string
}

func (e *ValidationError) Error() string { return "validation failed" }
