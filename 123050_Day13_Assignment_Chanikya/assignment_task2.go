package main

import "fmt"

// Task 2: Handler and Service Boundary Separation Code
// Demonstrates how structs hold dependencies to act as clear boundary layers.
type UserService struct{}

func (s *UserService) GetUser(id string) string {
	return fmt.Sprintf("User details for ID: %s", id)
}

type UserHandler struct {
	Service *UserService
}

func (h *UserHandler) HandleGetUser(id string) {
	result := h.Service.GetUser(id)
	fmt.Printf("[Task 2] Handler boundary output: %s\n", result)
}

func main() {
	svc := &UserService{}
	handler := &UserHandler{Service: svc}
	handler.HandleGetUser("USR-100")
}
