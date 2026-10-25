package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-management/middleware"
	"employee-management/model"
	"employee-management/service"

	"gorm.io/gorm"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

// RegisterRoutes registers protected employee routes.
// The caller wraps this handler with authentication and authorization middleware.
func (c *EmployeeController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/employees", c.collection)
	mux.HandleFunc("/employees/", c.byID)
}

func (c *EmployeeController) collection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query()

		page, err := queryInt(q.Get("page"), 1)
		if err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid page"})
			return
		}
		pageSize, err := queryInt(q.Get("pageSize"), 10)
		if err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid pageSize"})
			return
		}

		result, err := c.service.GetAll(
			model.EmployeeFilter{
				Name: q.Get("name"), Department: q.Get("department"),
			},
			model.EmployeeSort{
				Field: q.Get("sortBy"), Direction: q.Get("sortOrder"),
			},
			page, pageSize,
		)
		if err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, result)

	case http.MethodPost:
		if middleware.RoleFromContext(r.Context()) != "admin" {
			writeJSON(w, 403, map[string]string{"error": "admin role required"})
			return
		}

		var e model.Employee
		if err := decodeEmployee(w, r, &e); err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid employee JSON"})
			return
		}
		if err := c.service.Create(&e); err != nil {
			writeServiceError(w, err)
			return
		}
		writeJSON(w, 201, e)

	default:
		w.Header().Set("Allow", "GET, POST")
		writeJSON(w, 405, map[string]string{"error": "method not allowed"})
	}
}

func (c *EmployeeController) byID(w http.ResponseWriter, r *http.Request) {
	rawID := strings.TrimPrefix(r.URL.Path, "/employees/")
	id64, err := strconv.ParseUint(rawID, 10, 32)
	if err != nil || id64 == 0 {
		writeJSON(w, 400, map[string]string{"error": "invalid employee ID"})
		return
	}
	id := uint(id64)

	switch r.Method {
	case http.MethodGet:
		e, err := c.service.GetByID(id)
		if err != nil {
			writeServiceError(w, err)
			return
		}
		writeJSON(w, 200, e)

	case http.MethodPut:
		if middleware.RoleFromContext(r.Context()) != "admin" {
			writeJSON(w, 403, map[string]string{"error": "admin role required"})
			return
		}

		var e model.Employee
		if err := decodeEmployee(w, r, &e); err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid employee JSON"})
			return
		}
		updated, err := c.service.Update(id, &e)
		if err != nil {
			writeServiceError(w, err)
			return
		}
		writeJSON(w, 200, updated)

	case http.MethodDelete:
		if middleware.RoleFromContext(r.Context()) != "admin" {
			writeJSON(w, 403, map[string]string{"error": "admin role required"})
			return
		}
		if err := c.service.Delete(id); err != nil {
			writeServiceError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		writeJSON(w, 405, map[string]string{"error": "method not allowed"})
	}
}

func decodeEmployee(w http.ResponseWriter, r *http.Request, e *model.Employee) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(e)
}

func queryInt(value string, fallback int) (int, error) {
	if value == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound),
		err.Error() == "employee not found":
		writeJSON(w, 404, map[string]string{"error": "employee not found"})
	case service.IsDuplicate(err):
		writeJSON(w, 409, map[string]string{"error": "email already exists"})
	default:
		// Validation errors are client errors; unexpected database failures
		// should be logged and returned as a generic server error in production.
		msg := err.Error()
		if strings.Contains(msg, "required") ||
			strings.Contains(msg, "must ") ||
			strings.Contains(msg, "valid") ||
			strings.Contains(msg, "unsupported") {
			writeJSON(w, 400, map[string]string{"error": msg})
			return
		}
		writeJSON(w, 500, map[string]string{"error": "internal server error"})
	}
}
