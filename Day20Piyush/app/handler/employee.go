// Tasks 5, 6, 12: Employee CRUD — works with in-memory OR PostgreSQL (RDS).
// If DATABASE_URL is set → uses PostgreSQL. Otherwise → in-memory store.

package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"employee-api/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type EmployeeHandler struct {
	db   *sql.DB
	mu   sync.RWMutex
	mem  map[int]models.Employee
	next int
}

func NewEmployeeHandler(db *sql.DB) *EmployeeHandler {
	h := &EmployeeHandler{db: db, mem: make(map[int]models.Employee), next: 1}
	if db == nil {
		// Seed in-memory data
		for _, e := range []models.Employee{
			{Name: "Piyush", Email: "piyush@test.com", Age: 25, Department: "Engineering", Salary: 60000},
			{Name: "Alice", Email: "alice@test.com", Age: 30, Department: "Design", Salary: 55000},
			{Name: "Bob", Email: "bob@test.com", Age: 28, Department: "Marketing", Salary: 45000},
		} {
			e.ID = h.next
			e.CreatedAt = time.Now()
			h.mem[e.ID] = e
			h.next++
		}
	}
	return h
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// GET /api/employees
func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if h.db != nil {
		rows, err := h.db.Query(`SELECT id, name, email, age, department, salary, created_at FROM employees ORDER BY id`)
		if err != nil {
			respond(w, 500, models.APIResponse{Success: false, Message: err.Error()})
			return
		}
		defer rows.Close()
		var list []models.Employee
		for rows.Next() {
			var e models.Employee
			rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Department, &e.Salary, &e.CreatedAt)
			list = append(list, e)
		}
		respond(w, 200, models.APIResponse{Success: true, Data: list})
		return
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	var list []models.Employee
	for _, e := range h.mem {
		list = append(list, e)
	}
	respond(w, 200, models.APIResponse{Success: true, Data: list})
}

// GET /api/employees/{id}
func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if h.db != nil {
		var e models.Employee
		err := h.db.QueryRow(`SELECT id, name, email, age, department, salary, created_at FROM employees WHERE id=$1`, id).
			Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Department, &e.Salary, &e.CreatedAt)
		if err != nil {
			respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
			return
		}
		respond(w, 200, models.APIResponse{Success: true, Data: e})
		return
	}
	h.mu.RLock()
	e, ok := h.mem[id]
	h.mu.RUnlock()
	if !ok {
		respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
		return
	}
	respond(w, 200, models.APIResponse{Success: true, Data: e})
}

// POST /api/employees
func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var e models.Employee
	json.NewDecoder(r.Body).Decode(&e)
	if strings.TrimSpace(e.Name) == "" || !strings.Contains(e.Email, "@") {
		respond(w, 400, models.APIResponse{Success: false, Message: "name and valid email required"})
		return
	}
	if h.db != nil {
		err := h.db.QueryRow(
			`INSERT INTO employees (name, email, age, department, salary) VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at`,
			e.Name, e.Email, e.Age, e.Department, e.Salary).Scan(&e.ID, &e.CreatedAt)
		if err != nil {
			respond(w, 500, models.APIResponse{Success: false, Message: err.Error()})
			return
		}
		respond(w, 201, models.APIResponse{Success: true, Data: e})
		return
	}
	h.mu.Lock()
	e.ID = h.next
	h.next++
	e.CreatedAt = time.Now()
	h.mem[e.ID] = e
	h.mu.Unlock()
	respond(w, 201, models.APIResponse{Success: true, Data: e})
}

// PUT /api/employees/{id}
func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var e models.Employee
	json.NewDecoder(r.Body).Decode(&e)
	if h.db != nil {
		err := h.db.QueryRow(
			`UPDATE employees SET name=$1, email=$2, age=$3, department=$4, salary=$5 WHERE id=$6 RETURNING id, created_at`,
			e.Name, e.Email, e.Age, e.Department, e.Salary, id).Scan(&e.ID, &e.CreatedAt)
		if err != nil {
			respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
			return
		}
		respond(w, 200, models.APIResponse{Success: true, Data: e})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	old, ok := h.mem[id]
	if !ok {
		respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
		return
	}
	e.ID = id
	e.CreatedAt = old.CreatedAt
	h.mem[id] = e
	respond(w, 200, models.APIResponse{Success: true, Data: e})
}

// DELETE /api/employees/{id}
func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if h.db != nil {
		res, err := h.db.Exec(`DELETE FROM employees WHERE id=$1`, id)
		if err != nil {
			respond(w, 500, models.APIResponse{Success: false, Message: err.Error()})
			return
		}
		n, _ := res.RowsAffected()
		if n == 0 {
			respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
			return
		}
		respond(w, 200, models.APIResponse{Success: true, Message: "deleted"})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.mem[id]; !ok {
		respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
		return
	}
	delete(h.mem, id)
	respond(w, 200, models.APIResponse{Success: true, Message: "deleted"})
}

// GET /api/employees/report — generates CSV report (Task 10: upload to S3)
func (h *EmployeeHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
	var list []models.Employee
	if h.db != nil {
		rows, _ := h.db.Query(`SELECT id, name, email, age, department, salary, created_at FROM employees ORDER BY id`)
		if rows != nil {
			defer rows.Close()
			for rows.Next() {
				var e models.Employee
				rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Department, &e.Salary, &e.CreatedAt)
				list = append(list, e)
			}
		}
	} else {
		h.mu.RLock()
		for _, e := range h.mem {
			list = append(list, e)
		}
		h.mu.RUnlock()
	}

	csv := "ID,Name,Email,Age,Department,Salary,CreatedAt\n"
	for _, e := range list {
		csv += fmt.Sprintf("%d,%s,%s,%d,%s,%.2f,%s\n",
			e.ID, e.Name, e.Email, e.Age, e.Department, e.Salary, e.CreatedAt.Format("2006-01-02"))
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=employee_report.csv")
	w.Write([]byte(csv))
}
