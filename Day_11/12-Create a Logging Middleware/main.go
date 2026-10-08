package main

import (
	"fmt"
	"log"
	"net/http"

	"employee-management-go/middleware"

	"github.com/gorilla/mux"
)

// Employee handler.
func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Employee list")
}

func main() {

	// Create router.
	router := mux.NewRouter()

	// Create employee route.
	router.HandleFunc("/employees", getEmployees).Methods("GET")

	// Add logging middleware.
	router.Use(middleware.LoggingMiddleware)

	// Start server.
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
