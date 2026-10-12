package main

import (
	"log"
	"net/http"
	"time"

	"employee-rest-api/handler"
	"employee-rest-api/middleware"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file: ", err)
	}

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	employeeHandler := handler.NewEmployeeHandler(db)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /employees", employeeHandler.GetEmployees)
	mux.HandleFunc("GET /employees/{id}", employeeHandler.GetEmployee)
	mux.HandleFunc("POST /employees", employeeHandler.CreateEmployee)
	mux.HandleFunc("PUT /employees/{id}", employeeHandler.UpdateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", employeeHandler.DeleteEmployee)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           middleware.Logging(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Employee REST API running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
