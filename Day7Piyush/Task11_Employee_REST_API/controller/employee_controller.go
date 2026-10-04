package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"task11_employee_rest_api/models"
	"task11_employee_rest_api/repository"

	"github.com/gorilla/mux"
)

type EmployeeController struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeController(repo *repository.EmployeeRepository) *EmployeeController {
	return &EmployeeController{Repo: repo}
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	created, err := c.Repo.Create(emp)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to create employee: " + err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.Response{Message: "Employee created", Data: created})
}

func (c *EmployeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := c.Repo.GetAll()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to fetch employees"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employees retrieved", Data: employees})
}

func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	emp, err := c.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Error fetching employee"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee found", Data: emp})
}

func (c *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.Response{Message: "Invalid request body"})
		return
	}
	emp.ID = id
	updated, err := c.Repo.Update(emp)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to update employee"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee updated", Data: updated})
}

func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.Repo.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.Response{Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to delete employee"})
		return
	}
	sendJSON(w, http.StatusOK, models.Response{Message: "Employee deleted"})
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
