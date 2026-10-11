package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

const schema = `
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    department VARCHAR(100) NOT NULL,
    salary NUMERIC(12, 2) NOT NULL
)`

type Server struct {
	db *sql.DB
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func connect(ctx context.Context, url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := s.db.PingContext(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "degraded", "database": "down"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "database": "up"})
}

func (s *Server) collection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.list(w, r)
	case http.MethodPost:
		s.create(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) list(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.QueryContext(r.Context(),
		`SELECT id, name, email, department, salary FROM employees ORDER BY id`)
	if err != nil {
		log.Println("query error:", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			log.Println("scan error:", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}
		employees = append(employees, e)
	}
	if err := rows.Err(); err != nil {
		log.Println("rows error:", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if strings.TrimSpace(e.Name) == "" || strings.TrimSpace(e.Email) == "" {
		writeError(w, http.StatusUnprocessableEntity, "name and email are required")
		return
	}

	err := s.db.QueryRowContext(r.Context(),
		`INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`,
		e.Name, e.Email, e.Department, e.Salary).Scan(&e.ID)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "email already exists")
			return
		}
		log.Println("insert error:", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusCreated, e)
}

func (s *Server) item(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil || id < 1 {
		writeError(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	var e Employee
	err = s.db.QueryRowContext(r.Context(),
		`SELECT id, name, email, department, salary FROM employees WHERE id = $1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	if err != nil {
		log.Println("query error:", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, e)
}

func main() {
	url := getenv("DATABASE_URL", "postgres://postgres:password@localhost:5432/employeedb?sslmode=disable")

	ctx := context.Background()
	db, err := connect(ctx, url)
	if err != nil {
		log.Fatalf("could not connect to PostgreSQL: %v", err)
	}
	defer db.Close()
	log.Println("Connected to PostgreSQL")

	if _, err := db.ExecContext(ctx, schema); err != nil {
		log.Fatalf("could not create schema: %v", err)
	}

	server := &Server{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", server.health)
	mux.HandleFunc("/employee", server.collection)
	mux.HandleFunc("/employee/", server.item)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
