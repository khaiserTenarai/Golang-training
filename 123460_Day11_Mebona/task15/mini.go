package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := "postgres://postgres:pgadmin@localhost:5432/employee_db?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		position TEXT NOT NULL,
		salary NUMERIC NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Connected to PostgreSQL & table verified.")
}

func validateEmployee(emp Employee) string {
	if emp.Name == "" {
		return "Name is required"
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(emp.Email) {
		return "Invalid email format"
	}
	if emp.Position == "" {
		return "Position is required"
	}
	if emp.Salary <= 0 {
		return "Salary must be greater than zero"
	}
	return ""
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s - Duration: %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Supports query parameters: /employee?position=Developer
		positionFilter := r.URL.Query().Get("position")
		var rows *sql.Rows
		var err error

		if positionFilter != "" {
			rows, err = db.Query("SELECT id, name, email, position, salary FROM employees WHERE position = $1", positionFilter)
		} else {
			rows, err = db.Query("SELECT id, name, email, position, salary FROM employees")
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		employees := []Employee{}
		for rows.Next() {
			var e Employee
			if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.Salary); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			employees = append(employees, e)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employees)

	case http.MethodPost:
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if valErr := validateEmployee(emp); valErr != "" {
			http.Error(w, valErr, http.StatusBadRequest)
			return
		}

		query := "INSERT INTO employees (name, email, position, salary) VALUES ($1, $2, $3, $4) RETURNING id"
		err := db.QueryRow(query, emp.Name, emp.Email, emp.Position, emp.Salary).Scan(&emp.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(emp)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func employeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		var e Employee
		query := "SELECT id, name, email, position, salary FROM employees WHERE id = $1"
		err := db.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.Salary)
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(e)

	case http.MethodPut:
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if valErr := validateEmployee(emp); valErr != "" {
			http.Error(w, valErr, http.StatusBadRequest)
			return
		}

		query := "UPDATE employees SET name = $1, email = $2, position = $3, salary = $4 WHERE id = $5"
		res, err := db.Exec(query, emp.Name, emp.Email, emp.Position, emp.Salary, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		emp.ID = id
		json.NewEncoder(w).Encode(emp)

	case http.MethodDelete:
		query := "DELETE FROM employees WHERE id = $1"
		res, err := db.Exec(query, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	initDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employeeHandler)
	mux.HandleFunc("/employee/", employeeDetailHandler)

	fmt.Println("Mini Project Server running on :8080...")
	http.ListenAndServe(":8080", loggingMiddleware(mux))
}