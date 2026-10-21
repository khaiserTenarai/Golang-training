package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Task 2: Handler and Service Boundary Separation Code
// The Handler manages HTTP parsing, status codes, and JSON responses.
// The Service holds the core business logic independent of the transport layer (HTTP, gRPC, CLI).

type SimpleService struct{}

func (s *SimpleService) GetMessage() string {
	return "Hello from the bounded Service Layer!"
}

type SimpleHandler struct {
	Service *SimpleService
}

func (h *SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := h.Service.GetMessage()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

func main() {
	fmt.Println("Task 2 - Handler & Service Separation definition compiled successfully.")
}
