package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=Day7 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB connection:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database is unreachable:", err)
	}

	repo := NewUserRepository(db)
	service := NewUserService(repo)
	controller := NewUserController(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", controller.RegisterHandler)
	mux.HandleFunc("/users/active", controller.GetActiveUsersHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}