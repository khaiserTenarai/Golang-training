package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-api/model"
	"employee-api/repository"
	"employee-api/service"
)

type EmployeeHandler struct {
	Service service.EmployeeService
}

func NewEmployeeHandler(s service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{Service: s}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *EmployeeHandler) Collection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.List(w, r)
	case http.MethodPost:
		var e model.Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		created, err := h.Service.Create(r.Context(), e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeJSON(w, http.StatusCreated, created)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *EmployeeHandler) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page := 1
	limit := 10

	if raw := q.Get("page"); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil || value < 1 {
			http.Error(w, "page must be a positive integer", http.StatusBadRequest)
			return
		}
		page = value
	}
	if raw := q.Get("limit"); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil || value < 1 || value > 100 {
			http.Error(w, "limit must be between 1 and 100", http.StatusBadRequest)
			return
		}
		limit = value
	}

	employees, err := h.Service.List(r.Context(),
		strings.TrimSpace(q.Get("department")),
		strings.ToLower(q.Get("sort")),
		strings.ToLower(q.Get("order")),
		limit, (page-1)*limit)
	if err != nil {
		http.Error(w, "Could not retrieve employees", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"page": page, "limit": limit, "employees": employees,
	})
}

func (h *EmployeeHandler) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		e, err := h.Service.Get(r.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrEmployeeNotFound) {
				http.Error(w, "Employee not found", http.StatusNotFound)
			} else {
				http.Error(w, "Could not retrieve employee", http.StatusInternalServerError)
			}
			return
		}
		writeJSON(w, http.StatusOK, e)

	case http.MethodPut:
		var e model.Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		updated, err := h.Service.Update(r.Context(), id, e)
		if err != nil {
			if errors.Is(err, repository.ErrEmployeeNotFound) {
				http.Error(w, "Employee not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		err := h.Service.Delete(r.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrEmployeeNotFound) {
				http.Error(w, "Employee not found", http.StatusNotFound)
			} else {
				http.Error(w, "Could not delete employee", http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
