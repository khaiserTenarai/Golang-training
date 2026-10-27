package main

import "errors"

func CheckAge(age int) error {
	if age < 18 {
		return errors.New("must be at least 18")
	}
	return nil
}
