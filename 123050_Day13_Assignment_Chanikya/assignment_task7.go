package main

import "fmt"

// Task 7: Authentication Token Validator (AuthN)
// Performs core token evaluation layers.
type AuthNValidator struct{}

func (v *AuthNValidator) ValidateToken(token string) (string, error) {
	if token == "secret-valid-token" {
		return "admin-user", nil
	}
	return "", fmt.Errorf("invalid token credentials")
}

func main() {
	validator := &AuthNValidator{}
	user, err := validator.ValidateToken("secret-valid-token")
	fmt.Printf("[Task 7] AuthN verification check. User: %s, Err: %v\n", user, err)
}
