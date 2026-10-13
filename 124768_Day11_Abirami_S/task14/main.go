package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("Set DATABASE_URL before running the server")
	}
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	}
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		rows, err := db.QueryContext(r.Context(), "SELECT id, name FROM employees ORDER BY id")
		if err != nil {
			http.Error(w, "Database query failed", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		result := make([]Employee, 0)
		for rows.Next() {
			var e Employee
			if err := rows.Scan(&e.ID, &e.Name); err != nil {
				http.Error(w, "Could not read employee", http.StatusInternalServerError)
				return
			}
			result = append(result, e)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "Database read failed", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("JSON encode failed: %v", err)
		}
	})
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
