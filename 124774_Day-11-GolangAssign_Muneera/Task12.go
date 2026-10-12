package main

import (
	"fmt"      // Used for printing messages
	"log"      // Used for printing log messages
	"net/http" // Provides HTTP server functionality
	"time"     // Used to calculate request duration
)

// loggingMiddleware is a middleware that logs information
// about each HTTP request.
func loggingMiddleware(next http.Handler) http.Handler {

	// Return a new HTTP handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Record the time when the request starts
		start := time.Now()

		// Pass the request to the next handler
		next.ServeHTTP(w, r)

		// Log the HTTP method, URL path, and time taken
		// after the handler has finished processing the request.
		log.Printf("method=%s path=%s duration=%s", r.Method, r.URL.Path, time.Since(start))
	})
}

// employeeHandler handles requests for employees.
func employeeHandler(w http.ResponseWriter, r *http.Request) {

	// Send a response back to the client
	fmt.Fprintln(w, "Employee handler executed")
}

func main() {

	// Wrap the employee handler with the logging middleware.
	// The middleware will run before and after employeeHandler.
	handler := loggingMiddleware(http.HandlerFunc(employeeHandler))

	// Register the handler for the /employees URL path.
	http.Handle("/employees", handler)
	fmt.Println("Server running at http://localhost:2020/employees")

	// Start the HTTP server on port 2020.
	http.ListenAndServe(":2020", nil)
}
