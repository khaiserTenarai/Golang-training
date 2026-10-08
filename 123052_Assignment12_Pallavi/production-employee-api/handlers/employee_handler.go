package handlers

import (
	"encoding/json"
	"net/http"
	"production-employee-api/models"
	"production-employee-api/service"
	"strconv"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(svc service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: svc}
}

func (h *EmployeeHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "UP"}`))
}

func (h *EmployeeHandler) ReadyCheck(w http.ResponseWriter, r *http.Request) {
	if h.service.CheckReadiness() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "READY"}`))
		return
	}
	http.Error(w, `{"status": "NOT_READY"}`, http.StatusServiceUnavailable)
}

func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	params := models.QueryParams{
		Page:       page,
		Limit:      limit,
		Department: r.URL.Query().Get("department"),
		SortBy:     r.URL.Query().Get("sort_by"),
		Order:      r.URL.Query().Get("order"),
	}

	employees, total := h.service.FetchEmployees(params)

	response := map[string]interface{}{
		"data":  employees,
		"total": total,
		"page":  params.Page,
		"limit": params.Limit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, `{"error": "invalid payload"}`, http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateEmployee(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}