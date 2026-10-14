package service

import "errors"

var ErrDuplicateEmail = errors.New("an employee with this email already exists")
