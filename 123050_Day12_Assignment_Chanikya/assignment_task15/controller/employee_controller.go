package controller

import (
	"encoding/json"
	"employee_api/model"
	"employee_api/service"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	svc *service.EmployeeService
}

func NewEmployeeController(svc *service.EmployeeService) *EmployeeController {
	return &EmployeeController{svc: svc}
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Bad JSON format input request payload", http.StatusBadRequest)
		return
	}

	created, err := c.svc.AddEmployee(r.Context(), emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

func (c *EmployeeController) ListEmployees(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	dept := q.Get("department")
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))

	list, err := c.svc.SearchAndList(r.Context(), dept, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(list)
}

func (c *EmployeeController) LivenessProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func (c *EmployeeController) ReadinessProbe(w http.ResponseWriter, r *http.Format) {
	// Add real systems inspection logic if necessary
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("READY"))
}
