package handlers

import (
	"configurable-employee-api/config"
	"configurable-employee-api/service"
	"encoding/json"
	"net/http"
)

type EmployeeHandler struct {
	svc intService
	cfg *config.Config
}

type intService interface {
	ListEmployees(ctx interface{})
}

type Handler struct {
	svc service.EmployeeService
	cfg *config.Config
}

func NewHandler(svc service.EmployeeService, cfg *config.Config) *Handler {
	return &Handler{svc: svc, cfg: cfg}
}

func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	employees, err := h.svc.ListEmployees(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"environment": h.cfg.AppEnv,
		"data":        employees,
	}

	// Feature flag check
	if h.cfg.EnableNewFeature {
		response["feature_new_analytics"] = "ENABLED: Advanced employee metrics active"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}