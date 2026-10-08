package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

const connStr = "postgres://postgres:sasi2356@localhost:5432/assignment_db"

// Domain Model
type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// Repository Layer
type CapstoneRepo struct {
	pool *pgxpool.Pool
}

func (r *CapstoneRepo) ListPaginated(ctx context.Context, limit, offset int) ([]Employee, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, department, salary FROM capstone_employees ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []Employee
	for rows.Next() {
		var e Employee
		_ = rows.Scan(&e.ID, &e.Name, &e.Department, &e.Salary)
		emps = append(emps, e)
	}
	return emps, nil
}

func (r *CapstoneRepo) CreateTransactional(ctx context.Context, emp *Employee) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, "INSERT INTO capstone_employees (name, department, salary) VALUES ($1, $2, $3) RETURNING id", emp.Name, emp.Department, emp.Salary).Scan(&emp.ID)
	if err != nil {
		return err
	}

	// Audit log insertion inside same transaction
	_, err = tx.Exec(ctx, "INSERT INTO capstone_audit (action, employee_id) VALUES ('CREATED', $1)", emp.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Controller Handler
type CapstoneHandler struct {
	repo *CapstoneRepo
}

func (h *CapstoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	switch r.Method {
	case "GET":
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 1 {
			page = 1
		}
		limit := 5
		offset := (page - 1) * limit

		emps, err := h.repo.ListPaginated(ctx, limit, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"page":  page,
			"items": emps,
		})

	case "POST":
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.repo.CreateTransactional(ctx, &emp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(emp)
	}
}

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to connect pool: %v\n", err)
	}
	defer pool.Close()

	// Initialize tables
	_, _ = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS capstone_employees (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			department VARCHAR(100),
			salary NUMERIC(10,2)
		);
		CREATE TABLE IF NOT EXISTS capstone_audit (
			id SERIAL PRIMARY KEY,
			action VARCHAR(50),
			employee_id INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	repo := &CapstoneRepo{pool: pool}
	handler := &CapstoneHandler{repo: repo}

	http.Handle("/api/v1/employees", handler)
	fmt.Println("Capstone API Server listening on http://localhost:8080/api/v1/employees")
	log.Fatal(http.ListenAndServe(":8080", nil))
}