package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"day13_project/model"
	"day13_project/service"
)

type EmployeeController struct {
	svc *service.EmployeeService
}

func NewEmployeeController(svc *service.EmployeeService) *EmployeeController {
	return &EmployeeController{svc: svc}
}

func (c *EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	
	list, err := c.svc.ListEmployees(r.Context(), true, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Malformed request body payload data context", http.StatusBadRequest)
		return
	}

	if err := c.svc.AddEmployee(r.Context(), emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Employee record operational mapping updated successfully"))
}

func (c *EmployeeController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (c *EmployeeController) ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("READY"))
}
