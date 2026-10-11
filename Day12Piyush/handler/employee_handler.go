// Task 2: Handler Layer — HTTP concerns only (parse request, send response).
// Task 4: Depends on service INTERFACE (dependency injection).
// Handlers call service methods and never touch the repository directly.

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"production_api/middleware"
	"production_api/models"
	"production_api/repository"
	"production_api/service"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	service service.EmployeeService // depends on interface (Task 4)
}

// NewEmployeeHandler injects service dependency (Task 4)
func NewEmployeeHandler(svc service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: svc}
}

// ---- POST /api/employees (Tasks 2, 4, 9) ----

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, r, http.StatusBadRequest, models.APIResponse{
			Success: false, Message: "Invalid JSON body",
		})
		return
	}

	created, validationErrs, err := h.service.Create(emp)
	if validationErrs != nil {
		sendJSON(w, r, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false, Message: "Validation failed", Data: validationErrs,
		})
		return
	}
	if err != nil {
		sendJSON(w, r, http.StatusConflict, models.APIResponse{
			Success: false, Message: err.Error(),
		})
		return
	}

	sendJSON(w, r, http.StatusCreated, models.APIResponse{
		Success: true, Message: "Employee created", Data: created,
	})
}

// ---- GET /api/employees (Tasks 2, 10, 11, 12) ----

func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	minSalary, _ := strconv.ParseFloat(q.Get("min_salary"), 64)
	maxSalary, _ := strconv.ParseFloat(q.Get("max_salary"), 64)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	params := models.QueryParams{
		Name:       q.Get("name"),       // Task 11: filter
		Department: q.Get("department"), // Task 11: filter
		MinSalary:  minSalary,           // Task 11: filter
		MaxSalary:  maxSalary,           // Task 11: filter
		SortBy:     q.Get("sort_by"),    // Task 12: sort
		SortOrder:  q.Get("sort_order"), // Task 12: sort
		Page:       page,                // Task 10: pagination
		PageSize:   pageSize,            // Task 10: pagination
	}

	employees, total, err := h.service.GetAll(params)
	if err != nil {
		sendJSON(w, r, http.StatusInternalServerError, models.APIResponse{
			Success: false, Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", middleware.GetRequestID(r))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.PaginatedResponse{
		Success:    true,
		Data:       employees,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: repository.TotalPages(total, pageSize),
		RequestID:  middleware.GetRequestID(r),
	})
}

// ---- GET /api/employees/{id} (Task 2) ----

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendJSON(w, r, http.StatusBadRequest, models.APIResponse{
			Success: false, Message: "Invalid employee ID",
		})
		return
	}

	emp, err := h.service.GetByID(id)
	if err != nil {
		sendJSON(w, r, http.StatusNotFound, models.APIResponse{
			Success: false, Message: err.Error(),
		})
		return
	}

	sendJSON(w, r, http.StatusOK, models.APIResponse{
		Success: true, Message: "Employee found", Data: emp,
	})
}

// ---- PUT /api/employees/{id} ----

func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, r, http.StatusBadRequest, models.APIResponse{
			Success: false, Message: "Invalid JSON body",
		})
		return
	}

	updated, validationErrs, err := h.service.Update(id, emp)
	if validationErrs != nil {
		sendJSON(w, r, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false, Message: "Validation failed", Data: validationErrs,
		})
		return
	}
	if err != nil {
		sendJSON(w, r, http.StatusNotFound, models.APIResponse{
			Success: false, Message: err.Error(),
		})
		return
	}

	sendJSON(w, r, http.StatusOK, models.APIResponse{
		Success: true, Message: "Employee updated", Data: updated,
	})
}

// ---- DELETE /api/employees/{id} ----

func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.service.Delete(id); err != nil {
		sendJSON(w, r, http.StatusNotFound, models.APIResponse{
			Success: false, Message: err.Error(),
		})
		return
	}
	sendJSON(w, r, http.StatusOK, models.APIResponse{
		Success: true, Message: "Employee deleted",
	})
}

// ---- Tasks 13, 14: Health and Ready Endpoints ----

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, r, http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "employee-api",
		"version":   "1.0.0",
	})
}

func ReadyHandler(svc service.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the service can respond (e.g., DB connectivity in production)
		count := svc.Count()
		sendJSON(w, r, http.StatusOK, map[string]interface{}{
			"ready":          true,
			"employee_count": count,
			"timestamp":      time.Now().Format(time.RFC3339),
		})
	}
}

// ---- Helper ----

func sendJSON(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", middleware.GetRequestID(r))
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
