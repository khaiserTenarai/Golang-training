
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func main() {
	var err error

	db, err = connectDB()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/employees", employeesHandler)
	mux.HandleFunc("/employees/", employeeHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

// GET /employees - List all employees
// POST /employees - Create an employee
func employeesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query(context.Background(),
			`SELECT id, name, email, department, salary
			 FROM employees ORDER BY id`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		employees := []Employee{}

		for rows.Next() {
			var e Employee
			err := rows.Scan(
				&e.ID, &e.Name, &e.Email,
				&e.Department, &e.Salary,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			employees = append(employees, e)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(employees)

	case http.MethodPost:
		var e Employee

		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		e.Name = strings.TrimSpace(e.Name)
		e.Email = strings.TrimSpace(e.Email)
		e.Department = strings.TrimSpace(e.Department)

		if e.Name == "" || e.Email == "" ||
			e.Department == "" || e.Salary <= 0 {
			http.Error(w, "Invalid employee data", http.StatusBadRequest)
			return
		}

		err := db.QueryRow(context.Background(),
			`INSERT INTO employees (name, email, department, salary)
			 VALUES ($1, $2, $3, $4)
			 RETURNING id`,
			e.Name, e.Email, e.Department, e.Salary,
		).Scan(&e.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /employees/{id} - Get one employee
// PUT /employees/{id} - Update an employee
// DELETE /employees/{id} - Delete an employee
func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idText := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idText)

	if err != nil || id <= 0 || strings.Contains(idText, "/") {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		var e Employee

		err := db.QueryRow(context.Background(),
			`SELECT id, name, email, department, salary
			 FROM employees WHERE id=$1`, id,
		).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)

		if err != nil {
			writeDBError(w, err)
			return
		}

		json.NewEncoder(w).Encode(e)

	case http.MethodPut:
		var e Employee

		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		e.Name = strings.TrimSpace(e.Name)
		e.Email = strings.TrimSpace(e.Email)
		e.Department = strings.TrimSpace(e.Department)

		if e.Name == "" || e.Email == "" ||
			e.Department == "" || e.Salary <= 0 {
			http.Error(w, "Invalid employee data", http.StatusBadRequest)
			return
		}

		tag, err := db.Exec(context.Background(),
			`UPDATE employees
			 SET name=$1, email=$2, department=$3, salary=$4
			 WHERE id=$5`,
			e.Name, e.Email, e.Department, e.Salary, id,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if tag.RowsAffected() == 0 {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		e.ID = id
		json.NewEncoder(w).Encode(e)

	case http.MethodDelete:
		tag, err := db.Exec(context.Background(),
			`DELETE FROM employees WHERE id=$1`, id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if tag.RowsAffected() == 0 {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeDBError(w http.ResponseWriter, err error) {
	if err == pgx.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	http.Error(w, "Database error",
		http.StatusInternalServerError)
}