// Package controller contains HTTP request handlers.
package controller

import (
	"employee-management/model"
	"employee-management/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

type EmployeeController struct{ service service.EmployeeService }

// NewEmployeeController creates the controller with dependency injection.
func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

// RegisterRoutes registers employee REST endpoints.
func (c *EmployeeController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/employees", c.handleCollection)
	mux.HandleFunc("/employees/", c.handleByID)
}
func (c *EmployeeController) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		department := r.URL.Query().Get("department")
		name := r.URL.Query().Get("name")
		sortBy := r.URL.Query().Get("sortBy")
		sortOrder := r.URL.Query().Get("sortOrder")
		page := parseInt(r.URL.Query().Get("page"), 1)
		pageSize := parseInt(r.URL.Query().Get("pageSize"), 10)
		resp, err := c.service.GetAll(r.Context(), department, name, sortBy, sortOrder, page, pageSize)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, resp)
	case http.MethodPost:
		var e model.Employee
		if json.NewDecoder(r.Body).Decode(&e) != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON request body")
			return
		}
		if err := c.service.Create(r.Context(), &e); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, e)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
func (c *EmployeeController) handleByID(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.ParseInt(idText, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "employee id must be numeric")
		return
	}
	switch r.Method {
	case http.MethodGet:
		e, err := c.service.GetByID(r.Context(), id)
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "employee not found")
			return
		}
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, e)
	case http.MethodPut:
		var e model.Employee
		if json.NewDecoder(r.Body).Decode(&e) != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON request body")
			return
		}
		u, err := c.service.Update(r.Context(), id, &e)
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "employee not found")
			return
		}
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, u)
	case http.MethodDelete:
		if err := c.service.Delete(r.Context(), id); err != nil {
			if err.Error() == "employee not found" {
				writeError(w, http.StatusNotFound, err.Error())
				return
			}
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
func parseInt(v string, d int) int {
	if v == "" {
		return d
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return d
	}
	return n
}
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
