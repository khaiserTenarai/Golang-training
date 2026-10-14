package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-api-day13/repository"
	"employee-api-day13/service"
)

type EmployeeController struct {
	service *service.EmployeeService
}

func NewEmployeeController(s *service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

// Register uses Go's standard net/http ServeMux patterns (Go 1.22+), no Gin.
func (c *EmployeeController) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", c.health)
	mux.HandleFunc("GET /employees", c.list)
	mux.HandleFunc("POST /employees", c.create)
	mux.HandleFunc("GET /employees/{id}", c.get)
	mux.HandleFunc("PUT /employees/{id}", c.update)
	mux.HandleFunc("DELETE /employees/{id}", c.delete)
}

func (c *EmployeeController) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (c *EmployeeController) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Department string `json:"department"`
	}
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	// Task 14: feature flag is checked in service; return a clear response.
	employee, err := c.service.Create(r.Context(), toInput(input.Name, input.Email, input.Department))
	if err != nil {
		if errors.Is(err, service.ErrFeatureOff) {
			writeError(w, http.StatusServiceUnavailable, err.Error())
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			writeError(w, http.StatusConflict, "email already exists")
			return
		}
		writeError(w, http.StatusInternalServerError, "could not create employee")
		return
	}
	writeJSON(w, http.StatusCreated, employee)
}

func (c *EmployeeController) list(w http.ResponseWriter, r *http.Request) {
	employees, err := c.service.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list employees")
		return
	}
	writeJSON(w, http.StatusOK, employees)
}

func (c *EmployeeController) get(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	employee, err := c.service.GetByID(r.Context(), id)
	if errors.Is(err, repository.ErrNotFound) {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not get employee")
		return
	}
	writeJSON(w, http.StatusOK, employee)
}

func (c *EmployeeController) update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	var input struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Department string `json:"department"`
	}
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	employee, err := c.service.Update(r.Context(), id, toInput(input.Name, input.Email, input.Department))
	if errors.Is(err, repository.ErrNotFound) {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	if errors.Is(err, service.ErrInvalidInput) {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			writeError(w, http.StatusConflict, "email already exists")
			return
		}
		writeError(w, http.StatusInternalServerError, "could not update employee")
		return
	}
	writeJSON(w, http.StatusOK, employee)
}

func (c *EmployeeController) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	err := c.service.Delete(r.Context(), id)
	if errors.Is(err, repository.ErrNotFound) {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not delete employee")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		writeError(w, http.StatusBadRequest, "id must be a positive integer")
		return 0, false
	}
	return id, true
}
