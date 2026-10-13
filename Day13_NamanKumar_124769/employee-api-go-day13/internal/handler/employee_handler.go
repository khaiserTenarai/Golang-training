package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-api-go/internal/domain"
	"employee-api-go/internal/repository"
	"employee-api-go/internal/service"
	"employee-api-go/internal/validation"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateEmployeeInput
	if err := decodeJSON(r, &input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return
	}
	if err := validation.ValidateCreateEmployee(input); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	employee, err := h.service.Create(r.Context(), input)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, employee)
}

func (h *EmployeeHandler) List(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	page, err := parseIntDefault(values.Get("page"), 1)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "page must be an integer"})
		return
	}
	limit, err := parseIntDefault(values.Get("limit"), 10)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "limit must be an integer"})
		return
	}
	result, err := h.service.List(r.Context(), domain.EmployeeQuery{
		Page:       page,
		Limit:      limit,
		Department: strings.TrimSpace(values.Get("department")),
		Search:     strings.TrimSpace(values.Get("search")),
		SortBy:     strings.TrimSpace(values.Get("sort_by")),
		Order:      strings.ToLower(strings.TrimSpace(values.Get("order"))),
	})
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	employee, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	var input domain.UpdateEmployeeInput
	if err := decodeJSON(r, &input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return
	}
	if err := validation.ValidateUpdateEmployee(input); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	employee, err := h.service.Update(r.Context(), id, input)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	if err := h.service.Delete(r.Context(), id); err != nil {
		h.writeServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *EmployeeHandler) writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "employee not found"})
	case errors.Is(err, repository.ErrEmailExists):
		writeJSON(w, http.StatusConflict, map[string]string{"error": "email already exists"})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
}

func pathID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func parseIntDefault(value string, fallback int) (int, error) {
	if value == "" {
		return fallback, nil
	}
	return strconv.Atoi(value)
}

func decodeJSON(r *http.Request, target any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(http.MaxBytesReader(nil, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
