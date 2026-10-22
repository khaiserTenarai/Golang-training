// Day 20: Go Employee API — runs locally (in-memory) or on AWS (EC2 + RDS + S3)
// Task 5: Deploy on EC2 | Task 6: Configure as service | Task 12: Connect to RDS
//
// Local:  go run main.go                              (in-memory, port 8080)
// With DB: DATABASE_URL="postgres://..." go run main.go (PostgreSQL/RDS)

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"employee-api/config"
	"employee-api/handler"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	var db *sql.DB
	if cfg.UseDB {
		var err error
		db, err = sql.Open("postgres", cfg.DatabaseURL)
		if err != nil {
			log.Fatal("DB connection failed:", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatal("DB ping failed:", err)
		}
		// Create table if not exists
		db.Exec(`CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			age INT,
			department VARCHAR(50),
			salary DECIMAL(10,2),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)
		fmt.Println("[DB] Connected to PostgreSQL (RDS)")
		defer db.Close()
	} else {
		fmt.Println("[DB] Running with in-memory store (no DATABASE_URL set)")
	}

	empHandler := handler.NewEmployeeHandler(db)
	s3Handler := handler.NewS3Handler(cfg)

	r := mux.NewRouter()

	// Health check (Task 8: behind security group, Task 13: CloudWatch target)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		status := "healthy"
		dbStatus := "in-memory"
		if db != nil {
			if err := db.Ping(); err != nil {
				status = "degraded"
				dbStatus = "disconnected"
			} else {
				dbStatus = "connected"
			}
		}
		handler.respond(w, 200, map[string]interface{}{
			"service": "employee-api", "status": status,
			"database": dbStatus, "timestamp": time.Now().Format(time.RFC3339),
		})
	}).Methods("GET")

	// Employee CRUD
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/employees", empHandler.GetAll).Methods("GET")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.GetByID).Methods("GET")
	api.HandleFunc("/employees", empHandler.Create).Methods("POST")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Update).Methods("PUT")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Delete).Methods("DELETE")
	api.HandleFunc("/employees/report", empHandler.GenerateReport).Methods("GET")

	// S3 operations (Tasks 9, 10)
	api.HandleFunc("/s3/upload-report", s3Handler.UploadReport).Methods("POST")

	fmt.Printf("\n[Server] Employee API running on http://0.0.0.0:%s\n", cfg.Port)
	fmt.Println("[Server] Endpoints:")
	fmt.Println("  GET    /health")
	fmt.Println("  GET    /api/employees")
	fmt.Println("  GET    /api/employees/{id}")
	fmt.Println("  POST   /api/employees")
	fmt.Println("  PUT    /api/employees/{id}")
	fmt.Println("  DELETE /api/employees/{id}")
	fmt.Println("  GET    /api/employees/report  (CSV download)")
	fmt.Println("  POST   /api/s3/upload-report  (upload to S3)")

	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
