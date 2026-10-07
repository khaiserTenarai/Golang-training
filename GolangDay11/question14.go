package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Application struct {
	DB *sql.DB
}

func (app *Application) getEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var emp Employee
	query := `SELECT id, name, department FROM employees WHERE id = $1`
	err := app.DB.QueryRow(query, idStr).Scan(&emp.ID, &emp.Name, &emp.Department)

	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	app := &Application{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /employee/{id}", app.getEmployeeHandler)

	http.ListenAndServe(":8080", mux)
}