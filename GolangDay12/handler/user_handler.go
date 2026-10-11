package handler

import (
	"api/models"
	"api/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	svc service.EmployeeService
}

func NewEmployeeHandler(svc service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}

func (h *EmployeeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		dept := r.URL.Query().Get("dept")
		sort := r.URL.Query().Get("sort")

		emps := h.svc.GetEmployees(page, limit, dept, sort)
		json.NewEncoder(w).Encode(emps)
	} else if r.Method == http.MethodPost {
		var emp models.Employee
		json.NewDecoder(r.Body).Decode(&emp)
		
		created, err := h.svc.CreateEmployee(emp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}