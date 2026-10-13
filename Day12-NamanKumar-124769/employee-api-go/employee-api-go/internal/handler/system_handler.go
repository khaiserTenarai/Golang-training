package handler

import (
	"net/http"

	"employee-api-go/internal/repository"
)

type SystemHandler struct {
	repo repository.EmployeeRepository
}

func NewSystemHandler(repo repository.EmployeeRepository) *SystemHandler {
	return &SystemHandler{repo: repo}
}

func (h *SystemHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *SystemHandler) Ready(w http.ResponseWriter, r *http.Request) {
	if !h.repo.Ready() {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "not_ready"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}
