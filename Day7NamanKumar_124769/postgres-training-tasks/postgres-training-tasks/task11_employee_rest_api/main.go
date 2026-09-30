package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

var db *pgxpool.Pool

func main() {
	connString := "postgres://postgres:password@localhost:5432/employeeapidb"
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	db = pool
	defer db.Close()

	http.HandleFunc("/employees", employeesHandler)
	http.HandleFunc("/employees/", employeeByIDHandler)

	http.ListenAndServe(":8080", nil)
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listEmployees(w, r)
	case http.MethodPost:
		createEmployee(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func employeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getEmployee(w, r, id)
	case http.MethodPut:
		updateEmployee(w, r, id)
	case http.MethodDelete:
		deleteEmployee(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func listEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(r.Context(), `SELECT id, name, email, salary FROM employees ORDER BY id`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Salary); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.QueryRow(r.Context(), `INSERT INTO employees (name, email, salary) VALUES ($1, $2, $3) RETURNING id`,
		e.Name, e.Email, e.Salary).Scan(&e.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func getEmployee(w http.ResponseWriter, r *http.Request, id int) {
	var e Employee
	err := db.QueryRow(r.Context(), `SELECT id, name, email, salary FROM employees WHERE id = $1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Salary)
	if errors.Is(err, pgx.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func updateEmployee(w http.ResponseWriter, r *http.Request, id int) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	e.ID = id

	result, err := db.Exec(r.Context(), `UPDATE employees SET name = $1, email = $2, salary = $3 WHERE id = $4`,
		e.Name, e.Email, e.Salary, e.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request, id int) {
	result, err := db.Exec(r.Context(), `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
