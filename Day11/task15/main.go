package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// Model definition
type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// Request validation extension
func (e *Employee) Validate() error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(e.Email, "@") {
		return errors.New("invalid email domain format")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("department is required")
	}
	if e.Salary <= 0 {
		return errors.New("salary must be a positive value")
	}
	return nil
}

// App environment dependency injection container
type App struct {
	DB *sql.DB
}

// Logging middleware wrapper
func (a *App) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s - Execution Duration: %v", r.Method, r.URL.String(), time.Since(start))
	})
}

// JSON Helper method
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// Error Helper method
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

// 1. GET /employees?department=X
func (a *App) getEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	deptQuery := r.URL.Query().Get("department")

	var rows *sql.Rows
	var err error

	if deptQuery != "" {
		query := `SELECT id, name, email, department, salary FROM employees WHERE LOWER(department) = LOWER($1)`
		rows, err = a.DB.Query(query, deptQuery)
	} else {
		query := `SELECT id, name, email, department, salary FROM employees`
		rows, err = a.DB.Query(query)
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to query database")
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Salary); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to parse records")
			return
		}
		employees = append(employees, emp)
	}

	respondWithJSON(w, http.StatusOK, employees)
}

// 2. GET /employees/{id}
func (a *App) getEmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID parameter")
		return
	}

	var emp Employee
	query := `SELECT id, name, email, department, salary FROM employees WHERE id = $1`
	err = a.DB.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Salary)

	if errors.Is(err, sql.ErrNoRows) {
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	} else if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}

	respondWithJSON(w, http.StatusOK, emp)
}

// 3. POST /employees
func (a *App) createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}

	if err := emp.Validate(); err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	query := `INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`
	err := a.DB.QueryRow(query, emp.Name, emp.Email, emp.Department, emp.Salary).Scan(&emp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create record")
		return
	}

	respondWithJSON(w, http.StatusCreated, emp)
}

// 4. PUT /employees/{id}
func (a *App) updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Malformed JSON input")
		return
	}

	if err := emp.Validate(); err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	query := `UPDATE employees SET name = $1, email = $2, department = $3, salary = $4 WHERE id = $5`
	result, err := a.DB.Exec(query, emp.Name, emp.Email, emp.Department, emp.Salary, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database update failure")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Target employee record does not exist")
		return
	}

	emp.ID = id
	respondWithJSON(w, http.StatusOK, emp)
}

// 5. DELETE /employees/{id}
func (a *App) deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	query := `DELETE FROM employees WHERE id = $1`
	result, err := a.DB.Exec(query, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database execution error")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

// Database Auto-Initialization
func initDB(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		department VARCHAR(50) NOT NULL,
		salary NUMERIC(10, 2) NOT NULL
	);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=company sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection initialization failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("PostgreSQL server unreachable: %v", err)
	}

	initDB(db)
	app := &App{DB: db}

	// Router Setup using Go 1.22+ ServeMux routing syntax
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", app.getEmployeesHandler)
	mux.HandleFunc("GET /employees/{id}", app.getEmployeeDetailHandler)
	mux.HandleFunc("POST /employees", app.createEmployeeHandler)
	mux.HandleFunc("PUT /employees/{id}", app.updateEmployeeHandler)
	mux.HandleFunc("DELETE /employees/{id}", app.deleteEmployeeHandler)

	handler := app.loggingMiddleware(mux)

	fmt.Println("PostgreSQL API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}