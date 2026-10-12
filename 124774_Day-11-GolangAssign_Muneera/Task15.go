package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Employee15 struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var db *sql.DB

// GET all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT id, name, department, salary FROM employees")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee15

	for rows.Next() {
		var employee Employee15

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Department,
			&employee.Salary,
		)

		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		employees = append(employees, employee)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// GET employee by ID
func getEmployee(w http.ResponseWriter, r *http.Request) {

	idString := strings.TrimPrefix(r.URL.Path, "/employee/")
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var employee Employee15

	err = db.QueryRow(
		"SELECT id, name, department, salary FROM employees WHERE id=$1",
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Department,
		&employee.Salary,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

// POST create employee
func createEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Employee15

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		"INSERT INTO employees (name, department, salary) VALUES ($1, $2, $3) RETURNING id",
		employee.Name,
		employee.Department,
		employee.Salary,
	).Scan(&employee.ID)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

// PUT update employee
func updateEmployee(w http.ResponseWriter, r *http.Request) {

	idString := strings.TrimPrefix(r.URL.Path, "/employee/")
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var employee Employee15

	err = json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"UPDATE employees SET name=$1, department=$2, salary=$3 WHERE id=$4",
		employee.Name,
		employee.Department,
		employee.Salary,
		id,
	)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	count, _ := result.RowsAffected()

	if count == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	employee.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

// DELETE employee
func deleteEmployee(w http.ResponseWriter, r *http.Request) {

	idString := strings.TrimPrefix(r.URL.Path, "/employee/")
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"DELETE FROM employees WHERE id=$1",
		id,
	)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	count, _ := result.RowsAffected()

	if count == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Employee deleted successfully")
}

func main() {

	connStr := "host=localhost port=5432 user=postgres password=Info%40131 dbname=go_training sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL:", err)
	}

	fmt.Println("Connected to PostgreSQL")

	// GET all and POST
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			getEmployees(w, r)
			return
		}

		if r.Method == http.MethodPost {
			createEmployee(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// GET, PUT and DELETE by ID
	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			getEmployee(w, r)
			return
		}

		if r.Method == http.MethodPut {
			updateEmployee(w, r)
			return
		}

		if r.Method == http.MethodDelete {
			deleteEmployee(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server running at http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
