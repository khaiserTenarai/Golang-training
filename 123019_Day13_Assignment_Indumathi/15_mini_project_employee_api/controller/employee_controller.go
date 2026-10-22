package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"employeeapi/dao"
	"employeeapi/model"
	"employeeapi/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func respondError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, dao.ErrNotFound), errors.Is(err, service.ErrFeatureDisabled):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, service.ErrValidation):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		writeError(w, http.StatusServiceUnavailable, "request was cancelled or timed out")
	default:
		writeError(w, http.StatusInternalServerError, "something went wrong")
	}
}

func readID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return 0, false
	}
	return id, true
}

func (c *EmployeeController) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		writeError(w, http.StatusBadRequest, "request body is not valid JSON")
		return
	}

	saved, err := c.service.AddEmployee(r.Context(), employee)
	if err != nil {
		respondError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, saved)
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.service.GetAllEmployees(r.Context())
	if err != nil {
		respondError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, employees)
}

func (c *EmployeeController) GetOne(w http.ResponseWriter, r *http.Request) {
	id, ok := readID(w, r)
	if !ok {
		return
	}

	employee, err := c.service.GetEmployee(r.Context(), id)
	if err != nil {
		respondError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, employee)
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := readID(w, r)
	if !ok {
		return
	}

	if err := c.service.DeleteEmployee(r.Context(), id); err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *EmployeeController) Report(w http.ResponseWriter, r *http.Request) {
	report, err := c.service.SalaryReport(r.Context())
	if err != nil {
		respondError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report)
}
