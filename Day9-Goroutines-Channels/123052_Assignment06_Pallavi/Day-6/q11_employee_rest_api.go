package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:Info%402k26@localhost:5432/assignment_db"

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

var db *pgx.Conn

func handleEmployees(w w, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	switch r.Method {
	case "POST":
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := db.QueryRow(ctx, "INSERT INTO employees (name, email, salary) VALUES ($1, $2, $3) RETURNING id", emp.Name, emp.Email, emp.Salary).Scan(&emp.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(emp)

	case "GET":
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var emp Employee
			err := db.QueryRow(ctx, "SELECT id, name, email, salary FROM employees WHERE id=$1", id).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Salary)
			if err != nil {
				http.Error(w, "Employee not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(emp)
			return
		}

		rows, err := db.Query(ctx, "SELECT id, name, email, salary FROM employees")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var e Employee
			_ = rows.Scan(&e.ID, &e.Name, &e.Email, &e.Salary)
			employees = append(employees, e)
		}
		json.NewEncoder(w).Encode(employees)

	case "PUT":
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := db.Exec(ctx, "UPDATE employees SET name=$1, email=$2, salary=$3 WHERE id=$4", emp.Name, emp.Email, emp.Salary, emp.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(emp)

	case "DELETE":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		_, err := db.Exec(ctx, "DELETE FROM employees WHERE id=$1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	var err error
	ctx := context.Background()
	db, err = pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer db.Close(ctx)

	_, _ = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100) UNIQUE,
			salary NUMERIC(10,2)
		);
	`)

	http.HandleFunc("/api/employees", handleEmployees)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}