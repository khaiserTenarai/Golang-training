package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 1. Repository Layer
type Repository struct{}

func (r *Repository) GetByID(id int) Employee {
	return Employee{ID: id, Name: "Alice"}
}

// 2. Service Layer
type Service struct {
	repo *Repository
}

func (s *Service) GetEmployee(id int) Employee {
	return s.repo.GetByID(id)
}

// 3. Handler Layer
type Handler struct {
	svc *Service
}

func (h *Handler) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	emp := h.svc.GetEmployee(1)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func main() {
	repo := &Repository{}
	svc := &Service{repo: repo}
	h := &Handler{svc: svc}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /employee", h.GetEmployeeHandler)
	http.ListenAndServe(":8080", mux)
}