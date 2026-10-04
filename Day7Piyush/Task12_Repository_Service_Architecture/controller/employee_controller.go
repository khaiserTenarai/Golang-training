package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"task12_repository_service_architecture/models"
	"task12_repository_service_architecture/service"

	"github.com/gorilla/mux"
)

// EmployeeController depends on service interface (not concrete type)
type EmployeeController struct {
	service service.EmployeeService
}

// NewEmployeeController injects the service dependency via interface
func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: svc}
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	created, err := c.service.CreateEmployee(emp)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "Employee created", Data: created})
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.service.GetAllEmployees()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to fetch employees"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Success", Data: employees})
}

func (c *EmployeeController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	emp, err := c.service.GetEmployeeByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Success", Data: emp})
}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	emp.ID = id
	updated, err := c.service.UpdateEmployee(emp)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee updated", Data: updated})
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.service.DeleteEmployee(id); err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee deleted"})
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
