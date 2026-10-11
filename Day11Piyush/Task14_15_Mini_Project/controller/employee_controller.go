package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task14_15_employee_api/models"
	"task14_15_employee_api/repository"

	"github.com/gorilla/mux"
)

type EmployeeController struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeController(repo *repository.EmployeeRepository) *EmployeeController {
	return &EmployeeController{Repo: repo}
}

// Task 13: Validate
func validate(emp models.Employee) []models.ValidationError {
	var errs []models.ValidationError
	if strings.TrimSpace(emp.Name) == "" {
		errs = append(errs, models.ValidationError{Field: "name", Message: "Name is required"})
	}
	if strings.TrimSpace(emp.Email) == "" || !strings.Contains(emp.Email, "@") {
		errs = append(errs, models.ValidationError{Field: "email", Message: "Valid email is required"})
	}
	if emp.Age < 18 || emp.Age > 65 {
		errs = append(errs, models.ValidationError{Field: "age", Message: "Age must be 18-65"})
	}
	if strings.TrimSpace(emp.Department) == "" {
		errs = append(errs, models.ValidationError{Field: "department", Message: "Department is required"})
	}
	if emp.Salary < 0 {
		errs = append(errs, models.ValidationError{Field: "salary", Message: "Salary cannot be negative"})
	}
	return errs
}

// Task 4, 9: POST (create, accept JSON body)
func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.APIResponse{Success: false, Message: "Invalid JSON"})
		return
	}
	if errs := validate(emp); len(errs) > 0 {
		sendJSON(w, http.StatusUnprocessableEntity, models.APIResponse{Success: false, Message: "Validation failed", Data: errs})
		return
	}
	created, err := c.Repo.Create(emp)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusCreated, models.APIResponse{Success: true, Message: "Employee created", Data: created})
}

// Task 2, 11: GET all with query params
func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	dept := q.Get("department")
	minSalary, _ := strconv.ParseFloat(q.Get("min_salary"), 64)
	maxSalary, _ := strconv.ParseFloat(q.Get("max_salary"), 64)
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	employees, total, err := c.Repo.GetAll(name, dept, minSalary, maxSalary, page, pageSize)
	if err != nil {
		sendJSON(w, http.StatusInternalServerError, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.PaginatedResponse{
		Success: true, Data: employees, Total: total, Page: page, PageSize: pageSize,
	})
}

// Task 3, 10: GET by ID (path parameter)
func (c *EmployeeController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	emp, err := c.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.APIResponse{Success: false, Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusInternalServerError, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.APIResponse{Success: true, Message: "Employee found", Data: emp})
}

// Task 5: PUT (update)
func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest, models.APIResponse{Success: false, Message: "Invalid JSON"})
		return
	}
	if errs := validate(emp); len(errs) > 0 {
		sendJSON(w, http.StatusUnprocessableEntity, models.APIResponse{Success: false, Message: "Validation failed", Data: errs})
		return
	}
	emp.ID = id
	updated, err := c.Repo.Update(emp)
	if err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.APIResponse{Success: false, Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusInternalServerError, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.APIResponse{Success: true, Message: "Employee updated", Data: updated})
}

// Task 6: DELETE
func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := c.Repo.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			sendJSON(w, http.StatusNotFound, models.APIResponse{Success: false, Message: "Employee not found"})
			return
		}
		sendJSON(w, http.StatusInternalServerError, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	sendJSON(w, http.StatusOK, models.APIResponse{Success: true, Message: "Employee deleted"})
}

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
