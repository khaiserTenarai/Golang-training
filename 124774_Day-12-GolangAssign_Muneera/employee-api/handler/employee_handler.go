package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"day12/model"
	"day12/service"
)

type EmployeeHandler struct {
	Service service.EmployeeService
	DB      *pgxpool.Pool
}

func NewEmployeeHandler(
	s service.EmployeeService,
	db *pgxpool.Pool,
) *EmployeeHandler {
	return &EmployeeHandler{
		Service: s,
		DB:      db,
	}
}

func (h *EmployeeHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/employees", h.collection)
	mux.HandleFunc("/employees/", h.byID)
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("/ready", h.ready)
}

func (h *EmployeeHandler) collection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query()

		page, _ := strconv.Atoi(q.Get("page"))
		pageSize, _ := strconv.Atoi(q.Get("pageSize"))

		result, err := h.Service.GetAll(
			r.Context(),
			q.Get("department"),
			q.Get("name"),
			q.Get("sortBy"),
			q.Get("sortOrder"),
			page,
			pageSize,
		)
		if err != nil {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, result)

	case http.MethodPost:
		var e model.Employee

		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "invalid JSON"})
			return
		}

		if err := h.Service.Create(r.Context(), &e); err != nil {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, e)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *EmployeeHandler) byID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(
		strings.TrimPrefix(r.URL.Path, "/employees/"), 10, 64,
	)
	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "invalid employee id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		e, err := h.Service.GetByID(r.Context(), id)

		if errors.Is(err, pgx.ErrNoRows) {
			writeJSON(w, http.StatusNotFound,
				map[string]string{"error": "employee not found"})
			return
		}
		if err != nil {
			writeJSON(w, http.StatusInternalServerError,
				map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, e)

	case http.MethodPut:
		var e model.Employee

		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "invalid JSON"})
			return
		}

		updated, err := h.Service.Update(r.Context(), id, &e)

		if errors.Is(err, pgx.ErrNoRows) {
			writeJSON(w, http.StatusNotFound,
				map[string]string{"error": "employee not found"})
			return
		}
		if err != nil {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		err := h.Service.Delete(r.Context(), id)

		if err != nil {
			if err.Error() == "employee not found" {
				writeJSON(w, http.StatusNotFound,
					map[string]string{"error": err.Error()})
			} else {
				writeJSON(w, http.StatusInternalServerError,
					map[string]string{"error": err.Error()})
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *EmployeeHandler) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "UP"})
}

func (h *EmployeeHandler) ready(w http.ResponseWriter, r *http.Request) {
	if err := h.DB.Ping(r.Context()); err != nil {
		writeJSON(w, http.StatusServiceUnavailable,
			map[string]string{"status": "NOT_READY"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "READY"})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
