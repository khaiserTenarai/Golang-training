package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"employee-management-go/model"
	"employee-management-go/service"

	"github.com/gorilla/mux"
)

// EmployeeController handles employee APIs.
type EmployeeController struct {
	Service service.EmployeeService
}

// NewEmployeeController creates a controller.
func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		Service: service,
	}
}

// CreateEmployee handles POST /employees.
func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee model.Employee

	// Read JSON body.
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// Create employee.
	employee, err = c.Service.Create(employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(employee)
}

// GetEmployees handles GET /employees.
func (c *EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request) {

	employees := c.Service.GetAll()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

// GetEmployee handles GET /employees/{id}.
func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "invalid employee ID", http.StatusBadRequest)
		return
	}

	employee, err := c.Service.GetByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employee)
}

// DeleteEmployee handles DELETE /employees/{id}.
func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "invalid employee ID", http.StatusBadRequest)
		return
	}

	err = c.Service.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
