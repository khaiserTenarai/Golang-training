package utils

import (
	"fmt"
	"strings"
)

func ValidateName(name string) error{
	name = strings.TrimSpace(name)
	if name == ""{
		return fmt.Errorf("Name Cannot be empty")
	}
	if len(name) < 3{
		return fmt.Errorf("Name must contain atleast 3 characters")
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 18 || age > 100 {
		return fmt.Errorf("Age must be between 18 and 100")
	} 
	return nil
}

func ValidateID(id int) error{
	if id <= 0{
		return fmt.Errorf("ID must be postive Integer")
	}
	return nil
}