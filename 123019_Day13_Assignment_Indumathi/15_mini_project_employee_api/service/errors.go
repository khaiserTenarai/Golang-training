package service

import "errors"

var (
	ErrValidation      = errors.New("invalid employee")
	ErrFeatureDisabled = errors.New("this feature is turned off")
)
