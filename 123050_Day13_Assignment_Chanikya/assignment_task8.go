package main

import "fmt"

// Task 8: Role-Based Authorization Interceptor (AuthZ)
// Asserts access rights given identity claims.
type AuthZInterceptor struct{}

func (z *AuthZInterceptor) Authorize(role string, requiredRole string) bool {
	return role == requiredRole
}

func main() {
	interceptor := &AuthZInterceptor{}
	allowed := interceptor.Authorize("admin", "admin")
	fmt.Printf("[Task 8] AuthZ permission determination outcome: %t\n", allowed)
}
