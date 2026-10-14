package health

import (
	"context"
	"employee-management/service"
	"encoding/json"
	"net/http"
)

type Handler struct{ service service.EmployeeService }

func NewHandler(s service.EmployeeService) *Handler { return &Handler{service: s} }

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	write(w, http.StatusOK, map[string]string{"status": "UP"})
}

func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Health(context.Background()); err != nil {
		write(w, http.StatusServiceUnavailable, map[string]string{"status": "NOT_READY", "database": "DOWN"})
		return
	}
	write(w, http.StatusOK, map[string]string{"status": "READY", "database": "UP"})
}
func write(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
