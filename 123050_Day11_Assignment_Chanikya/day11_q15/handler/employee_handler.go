package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-rest-api/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeHandler struct {
	DB *pgxpool.Pool
}

func NewEmployeeHandler(db *pgxpool.Pool) *EmployeeHandler {
	return &EmployeeHandler{DB: db}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func validEmployee(e model.Employee) bool {
	return strings.TrimSpace(e.Name) != "" &&
		len(e.Name) <= 100 &&
		strings.TrimSpace(e.Department) != "" &&
		len(e.Department) <= 100 &&
		e.Salary >= 0 &&
		e.Salary < 100000000
}

// GET /employees
// GET /employees?department=IT
func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	department := strings.TrimSpace(r.URL.Query().Get("department"))

	var rows pgx.Rows
	var err error

	if department != "" {
		rows, err = h.DB.Query(ctx,
			`SELECT id, name, department, salary
			 FROM employees_14 WHERE department = $1 ORDER BY id`,
			department)
	} else {
		rows, err = h.DB.Query(ctx,
			`SELECT id, name, department, salary
			 FROM employees_14 ORDER BY id`)
	}

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to retrieve employees")
		return
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)

	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Department, &e.Salary); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to read employee data")
			return
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to read employees")
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

// GET /employees/{id}
func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var e model.Employee
	err := h.DB.QueryRow(r.Context(),
		`SELECT id, name, department, salary
		 FROM employees_14 WHERE id = $1`, id).
		Scan(&e.ID, &e.Name, &e.Department, &e.Salary)

	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "Employee not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to retrieve employee")
		return
	}

	writeJSON(w, http.StatusOK, e)
}

// POST /employees
func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var e model.Employee

	if err := decodeEmployee(w, r, &e); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !validEmployee(e) {
		writeError(w, http.StatusBadRequest,
			"Name and department are required; salary must be non-negative")
		return
	}

	err := h.DB.QueryRow(r.Context(),
		`INSERT INTO employees_14 (name, department, salary)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		strings.TrimSpace(e.Name),
		strings.TrimSpace(e.Department),
		e.Salary,
	).Scan(&e.ID)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create employee")
		return
	}

	writeJSON(w, http.StatusCreated, e)
}

// PUT /employees/{id}
func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var e model.Employee
	if err := decodeEmployee(w, r, &e); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !validEmployee(e) {
		writeError(w, http.StatusBadRequest,
			"Name and department are required; salary must be non-negative")
		return
	}

	tag, err := h.DB.Exec(r.Context(),
		`UPDATE employees_14
		 SET name = $1, department = $2, salary = $3
		 WHERE id = $4`,
		strings.TrimSpace(e.Name),
		strings.TrimSpace(e.Department),
		e.Salary,
		id,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update employee")
		return
	}

	if tag.RowsAffected() == 0 {
		writeError(w, http.StatusNotFound, "Employee not found")
		return
	}

	e.ID = id
	writeJSON(w, http.StatusOK, e)
}

// DELETE /employees/{id}
func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	tag, err := h.DB.Exec(r.Context(),
		`DELETE FROM employees_14 WHERE id = $1`, id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to delete employee")
		return
	}

	if tag.RowsAffected() == 0 {
		writeError(w, http.StatusNotFound, "Employee not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "Invalid employee ID")
		return 0, false
	}
	return id, true
}

func decodeEmployee(w http.ResponseWriter, r *http.Request, e *model.Employee) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(e); err != nil {
		return errors.New("invalid JSON request body")
	}

	// Reject a second JSON value or any trailing non-whitespace content.
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, context.Canceled) &&
		err == nil {
		return errors.New("request must contain only one JSON object")
	} else if err != nil && err.Error() != "EOF" {
		return errors.New("invalid JSON request body")
	}

	return nil
}
