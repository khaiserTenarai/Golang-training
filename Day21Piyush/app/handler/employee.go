// Employee CRUD handler — works in-memory (local) or PostgreSQL (RDS)

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

func Respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if h.db != nil {
		rows, err := h.db.Query(`SELECT id,name,email,age,department,salary,created_at FROM employees ORDER BY id`)
		if err != nil {
			Respond(w, 500, models.APIResponse{Success: false, Message: err.Error()})
			return
		}
		defer rows.Close()
		var list []models.Employee
		for rows.Next() {
			var e models.Employee
			rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Department, &e.Salary, &e.CreatedAt)
			list = append(list, e)
		}
		Respond(w, 200, models.APIResponse{Success: true, Data: list})
		return
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	var list []models.Employee
	for _, e := range h.mem {
		list = append(list, e)
	}
	Respond(w, 200, models.APIResponse{Success: true, Data: list})
}

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if h.db != nil {
		var e models.Employee
		err := h.db.QueryRow(`SELECT id,name,email,age,department,salary,created_at FROM employees WHERE id=$1`, id).
			Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Department, &e.Salary, &e.CreatedAt)
		if err != nil {
			Respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
			return
		}
		Respond(w, 200, models.APIResponse{Success: true, Data: e})
		return
	}
	h.mu.RLock()
	e, ok := h.mem[id]
	h.mu.RUnlock()
	if !ok {
		Respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
		return
	}
	Respond(w, 200, models.APIResponse{Success: true, Data: e})
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var e models.Employee
	json.NewDecoder(r.Body).Decode(&e)
	if strings.TrimSpace(e.Name) == "" || !strings.Contains(e.Email, "@") {
		Respond(w, 400, models.APIResponse{Success: false, Message: "name and valid email required"})
		return
	}
	if h.db != nil {
		err := h.db.QueryRow(
			`INSERT INTO employees (name,email,age,department,salary) VALUES ($1,$2,$3,$4,$5) RETURNING id,created_at`,
			e.Name, e.Email, e.Age, e.Department, e.Salary).Scan(&e.ID, &e.CreatedAt)
		if err != nil {
			Respond(w, 500, models.APIResponse{Success: false, Message: err.Error()})
			return
		}
		Respond(w, 201, models.APIResponse{Success: true, Data: e})
		return
	}
	h.mu.Lock()
	e.ID = h.next
	h.next++
	e.CreatedAt = time.Now()
	h.mem[e.ID] = e
	h.mu.Unlock()
	Respond(w, 201, models.APIResponse{Success: true, Data: e})
}

func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if h.db != nil {
		res, _ := h.db.Exec(`DELETE FROM employees WHERE id=$1`, id)
		n, _ := res.RowsAffected()
		if n == 0 {
			Respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
			return
		}
		Respond(w, 200, models.APIResponse{Success: true, Message: "deleted"})
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.mem[id]; !ok {
		Respond(w, 404, models.APIResponse{Success: false, Message: "not found"})
		return
	}
	delete(h.mem, id)
	Respond(w, 200, models.APIResponse{Success: true, Message: "deleted"})
}

// Generate CSV report for S3 upload
func (h *EmployeeHandler) GenerateCSV() string {
	csv := "ID,Name,Email,Age,Department,Salary\n"
	if h.db != nil {
		rows, _ := h.db.Query(`SELECT id,name,email,age,department,salary FROM employees ORDER BY id`)
		if rows != nil {
			defer rows.Close()
			for rows.Next() {
				var e models.Employee
				rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Department, &e.Salary)
				csv += fmt.Sprintf("%d,%s,%s,%d,%s,%.2f\n", e.ID, e.Name, e.Email, e.Age, e.Department, e.Salary)
			}
		}
	} else {
		h.mu.RLock()
		defer h.mu.RUnlock()
		for _, e := range h.mem {
			csv += fmt.Sprintf("%d,%s,%s,%d,%s,%.2f\n", e.ID, e.Name, e.Email, e.Age, e.Department, e.Salary)
		}
	}
	return csv
}
