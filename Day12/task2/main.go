package main

import (
	"encoding/json"
	"net/http"
)

type Service interface {
	CalculateBonus(salary float64) float64
}

type employeeService struct{}

func (s *employeeService) CalculateBonus(salary float64) float64 {
	return salary * 0.10 // 10% bonus logic inside Service
}

type Handler struct {
	svc Service
}

func (h *Handler) BonusHandler(w http.ResponseWriter, r *http.Request) {
	// Handler only manages HTTP protocols and delegates business logic
	bonus := h.svc.CalculateBonus(50000)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"bonus": bonus})
}

func main() {
	svc := &employeeService{}
	h := &Handler{svc: svc}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /bonus", h.BonusHandler)
	http.ListenAndServe(":8080", mux)
}