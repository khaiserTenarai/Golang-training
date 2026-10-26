package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-management/model"
	"employee-management/service"

	"github.com/jackc/pgx/v5"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

func (c *EmployeeController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/employees", c.handleCollection)
	mux.HandleFunc("/employees/", c.handleByID)
}

func (c *EmployeeController) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		sortBy := r.URL.Query().Get("sort")
		employees, err := c.service.GetAll(r.Context(), sortBy)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, employees)
	case http.MethodPost:
		var emp model.Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON request body")
			return
		}
		if err := c.service.Create(r.Context(), &emp); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, emp)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c *EmployeeController) handleByID(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/employees/")
	if idText == "" {
		writeError(w, http.StatusBadRequest, "employee id is required")
		return
	}

	id, err := strconv.ParseInt(idText, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "employee id must be numeric")
		return
	}

	switch r.Method {
	case http.MethodGet:
		emp, err := c.service.GetByID(r.Context(), id)
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "employee not found")
			return
		}
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, emp)
	case http.MethodPut:
		var emp model.Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON request body")
			return
		}
		updated, err := c.service.Update(r.Context(), id, &emp)
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "employee not found")
			return
		}
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		err := c.service.Delete(r.Context(), id)
		if err != nil && err.Error() == "employee not found" {
			writeError(w, http.StatusNotFound, "employee not found")
			return
		}
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}