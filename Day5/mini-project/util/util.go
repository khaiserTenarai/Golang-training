package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

var idCounter int

func NextID() int {
	idCounter++
	return idCounter
}

func ToJSON(v interface{}) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}
	return string(data), nil
}

func FromJSON(jsonStr string, target interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), target); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}
	return nil
}

func IsValidEmail(email string) bool {
	at := strings.Index(email, "@")
	dot := strings.LastIndex(email, ".")
	return at > 0 && dot > at+1 && dot < len(email)-1
}

func ValidateName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if len(name) < 3 {
		return fmt.Errorf("name must contain at least 3 characters")
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 18 || age > 100 {
		return fmt.Errorf("age must be between 18 and 100")
	}
	return nil
}

func Divider() {
	fmt.Println(strings.Repeat("-", 40))
}
