package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"employeeapi/model"
	"employeeapi/repository"
	"employeeapi/service"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	var validationErr *service.ValidationError

	switch {
	case errors.As(err, &validationErr):
		writeJSON(w, http.StatusUnprocessableEntity, map[string]interface{}{"errors": validationErr.Problems})
	case errors.Is(err, repository.ErrNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "employee not found"})
	case errors.Is(err, repository.ErrDuplicateEmail):
		writeJSON(w, http.StatusConflict, map[string]string{"error": "an employee with that email already exists"})
	default:
		log.Println("unexpected error:", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
	}
}

func (h *EmployeeHandler) Employees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (h *EmployeeHandler) EmployeeByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a number"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		e, err := h.service.Get(id)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, e)

	case http.MethodPut:
		var e model.Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}
		updated, err := h.service.Update(id, e)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			writeError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (h *EmployeeHandler) list(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	filter := model.ListFilter{
		Department: q.Get("department"),
		Search:     q.Get("search"),
		SortBy:     q.Get("sort_by"),
		Order:      q.Get("order"),
	}
	filter.MinSalary, _ = strconv.ParseFloat(q.Get("min_salary"), 64)
	filter.MaxSalary, _ = strconv.ParseFloat(q.Get("max_salary"), 64)
	filter.Page, _ = strconv.Atoi(q.Get("page"))
	filter.Limit, _ = strconv.Atoi(q.Get("limit"))

	result, err := h.service.List(filter)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *EmployeeHandler) create(w http.ResponseWriter, r *http.Request) {
	var e model.Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return
	}
	created, err := h.service.Create(e)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}
