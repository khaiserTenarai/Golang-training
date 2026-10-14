package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"employee-rest-api/model"
	"employee-rest-api/repository"
	"employee-rest-api/service"
)

type EmployeeController struct {
	service *service.EmployeeService
}

func NewEmployeeController(s *service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var e model.Employee
	if err := decodeJSON(w, r, &e); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := c.service.Create(r.Context(), &e); err != nil {
		c.writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, e)
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.service.GetAll(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not fetch employees"})
		return
	}
	writeJSON(w, http.StatusOK, employees)
}

func (c *EmployeeController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	e, err := c.service.GetByID(r.Context(), id)
	if err != nil {
		c.writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, e)
}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	var e model.Employee
	if err := decodeJSON(w, r, &e); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	updated, err := c.service.Update(r.Context(), id, &e)
	if err != nil {
		c.writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	if err := c.service.Delete(r.Context(), id); err != nil {
		c.writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return errors.New("invalid JSON body: " + err.Error())
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("request body must contain a single JSON object")
	}
	return nil
}

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	raw := strings.TrimPrefix(r.URL.Path, "/api/employees/")
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func (c *EmployeeController) writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.Is(err, repository.ErrEmailExists):
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if status != http.StatusNoContent {
		_ = json.NewEncoder(w).Encode(value)
	}
}
