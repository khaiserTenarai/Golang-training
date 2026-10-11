package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"configurable_api/features"
	"configurable_api/models"
	"configurable_api/service"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	svc   service.EmployeeService
	flags *features.FeatureFlags
	maxPg int
}

func NewEmployeeHandler(svc service.EmployeeService, flags *features.FeatureFlags, maxPageSize int) *EmployeeHandler {
	return &EmployeeHandler{svc: svc, flags: flags, maxPg: maxPageSize}
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respond(w, http.StatusBadRequest, models.APIResponse{Success: false, Message: "Invalid JSON"})
		return
	}
	created, err := h.svc.Create(r.Context(), emp)
	if err != nil {
		respond(w, http.StatusBadRequest, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	respond(w, http.StatusCreated, models.APIResponse{Success: true, Message: "Created", Data: created})
}

func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	minSalary, _ := strconv.ParseFloat(q.Get("min_salary"), 64)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > h.maxPg {
		pageSize = h.maxPg
	}

	params := models.QueryParams{Page: page, PageSize: pageSize}

	// Task 14: Feature flags control which features are active
	if h.flags.IsEnabled("filtering") {
		params.Name = q.Get("name")
		params.Department = q.Get("department")
		params.MinSalary = minSalary
	}
	if h.flags.IsEnabled("sorting") {
		params.SortBy = q.Get("sort_by")
		params.SortOrder = q.Get("sort_order")
	}
	if !h.flags.IsEnabled("pagination") {
		params.Page = 1
		params.PageSize = 1000
	}

	list, total, err := h.svc.GetAll(r.Context(), params)
	if err != nil {
		respond(w, http.StatusInternalServerError, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	respond(w, http.StatusOK, map[string]interface{}{
		"success": true, "data": list, "total": total, "page": page, "page_size": pageSize,
	})
}

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	emp, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		respond(w, http.StatusNotFound, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	respond(w, http.StatusOK, models.APIResponse{Success: true, Data: emp})
}

func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp models.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	updated, err := h.svc.Update(r.Context(), id, emp)
	if err != nil {
		respond(w, http.StatusBadRequest, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	respond(w, http.StatusOK, models.APIResponse{Success: true, Message: "Updated", Data: updated})
}

func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.svc.Delete(r.Context(), id); err != nil {
		respond(w, http.StatusNotFound, models.APIResponse{Success: false, Message: err.Error()})
		return
	}
	respond(w, http.StatusOK, models.APIResponse{Success: true, Message: "Deleted"})
}

func (h *EmployeeHandler) Health(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, map[string]interface{}{"status": "healthy"})
}

func (h *EmployeeHandler) Ready(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, map[string]interface{}{"ready": true, "employees": h.svc.Count()})
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
